// Package emitrust emits Rust protocol definitions from the canonical v2
// manifest. The manifest remains the sole wire-schema artifact.
package emitrust

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"protocolgen/internal/manifest"
)

type definition struct {
	Name         string
	Kind         manifest.NodeKind
	Fields       []rustField
	Underlying   string
	Control      string
	Variants     []manifest.Variant
	Union        []rustVariant
	BitLength    uint64
	Tuple        bool
	WrapperCodec string
}

type rustVariant struct {
	Name         string
	Payload      string
	Fields       []rustField
	Discriminant int64
	SourceName   string
}

type rustField struct {
	Name string
	Type string
	Docs []string
}

type packetInfo struct {
	packet manifest.Packet
	name   string
	fields []rustFieldInfo
	size   int
}

type generator struct {
	definitions map[string]definition
	identities  map[string]string
	used        map[string]bool
	usesNbt     bool
	usesUUID    bool
	usesGlam    bool
	usesBytes   bool
	standalone  map[string]bool
}

func prepare(m manifest.Manifest) (*generator, []packetInfo, error) {
	if err := manifest.Validate(m); err != nil {
		return nil, nil, err
	}
	g := &generator{definitions: map[string]definition{}, identities: map[string]string{}, used: map[string]bool{}, standalone: standaloneRustStructs(m)}
	packets := append([]manifest.Packet(nil), m.Packets...)
	sort.Slice(packets, func(i, j int) bool { return packets[i].ID < packets[j].ID })
	infos := make([]packetInfo, 0, len(packets))
	for _, packet := range packets {
		name := g.unique(packetTypeName(packet.Name))
		used := map[string]bool{}
		fields := make([]rustFieldInfo, 0, len(packet.Fields))
		for _, field := range packet.Fields {
			if err := ensureCodecSymmetric(field); err != nil {
				return nil, nil, fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
			fieldName := uniqueField(fieldName(field.Name), used)
			typ, err := g.rustType(field.Encode, name+typeName(field.Name))
			if err != nil {
				return nil, nil, fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
			fields = append(fields, rustFieldInfo{name: fieldName, typ: typ, docs: fieldDocs(field)})
		}
		infos = append(infos, packetInfo{packet: packet, name: name, fields: fields, size: g.estimatePacketSize(fields)})
	}
	return g, infos, nil
}

func ensureCodecSymmetric(field manifest.Field) error {
	if field.Decode != nil || field.Symmetry != manifest.Symmetric {
		return fmt.Errorf("asymmetric encode/decode layouts require separate codec methods")
	}
	return ensureNodeCodecSymmetric(field.Encode)
}

func ensureNodeCodecSymmetric(node manifest.Node) error {
	for _, field := range node.Fields {
		if err := ensureCodecSymmetric(field); err != nil {
			return fmt.Errorf("nested field %s: %w", field.Name, err)
		}
	}
	for _, variant := range node.Variants {
		if variant.Decode != nil {
			return fmt.Errorf("union variant %s has an asymmetric decode layout", variant.Name)
		}
		if err := ensureNodeCodecSymmetric(variant.Encode); err != nil {
			return err
		}
	}
	for _, child := range []*manifest.Node{node.Prefix, node.Element, node.Value, node.Key, node.Control, node.Default} {
		if child != nil {
			if err := ensureNodeCodecSymmetric(*child); err != nil {
				return err
			}
		}
	}
	for _, child := range node.Elements {
		if err := ensureNodeCodecSymmetric(child); err != nil {
			return err
		}
	}
	for _, oneCase := range node.Cases {
		if len(oneCase.Decode) != 0 {
			return fmt.Errorf("conditional case %s has an asymmetric decode layout", oneCase.Value)
		}
		for _, child := range oneCase.Encode {
			if err := ensureNodeCodecSymmetric(child); err != nil {
				return err
			}
		}
	}
	return nil
}

// standaloneRustStructs returns named structs that are used outside the root
// payload of a union variant and therefore retain their own public type.
func standaloneRustStructs(m manifest.Manifest) map[string]bool {
	standalone := map[string]bool{}
	var walk func(manifest.Node, bool)
	walk = func(node manifest.Node, directVariant bool) {
		if node.Kind == manifest.KindStruct && node.TypeID != "" && !directVariant {
			standalone[node.TypeID] = true
		}
		for _, child := range []*manifest.Node{node.Prefix, node.Element, node.Value, node.Key, node.Control, node.Default} {
			if child != nil {
				walk(*child, false)
			}
		}
		for _, child := range node.Elements {
			walk(child, false)
		}
		for _, field := range node.Fields {
			walk(field.Encode, false)
			if field.Decode != nil {
				walk(*field.Decode, false)
			}
		}
		for _, variant := range node.Variants {
			walk(variant.Encode, true)
			if variant.Decode != nil {
				walk(*variant.Decode, true)
			}
		}
		for _, oneCase := range node.Cases {
			for _, child := range oneCase.Encode {
				walk(child, false)
			}
			for _, child := range oneCase.Decode {
				walk(child, false)
			}
		}
	}
	for _, packet := range m.Packets {
		for _, field := range packet.Fields {
			walk(field.Encode, false)
			if field.Decode != nil {
				walk(*field.Decode, false)
			}
		}
	}
	return standalone
}

func GenerateFiles(m manifest.Manifest) (map[string]string, error) {
	g, infos, err := prepare(m)
	if err != nil {
		return nil, err
	}
	definitions := g.sortedDefinitions()
	files := map[string]string{
		"Cargo.toml":     emitCargo(m, g),
		"src/lib.rs":     emitLib(m),
		"src/enums.rs":   emitRustEnums(definitions),
		"src/types.rs":   emitRustTypes(definitions),
		"src/wire.rs":    emitWire(g.usesNbt),
		"src/packets.rs": emitRustPackets(infos),
	}
	return files, nil
}

func (g *generator) sortedDefinitions() []definition {
	definitions := make([]definition, 0, len(g.definitions))
	for _, item := range g.definitions {
		definitions = append(definitions, item)
	}
	sort.Slice(definitions, func(i, j int) bool { return definitions[i].Name < definitions[j].Name })
	return definitions
}

func emitLib(m manifest.Manifest) string {
	return fmt.Sprintf(`// Code generated from canonical protocol manifest v2. DO NOT EDIT.

pub const PROTOCOL_VERSION: i32 = %d;

pub mod enums;
pub mod types;
pub mod packets;
pub mod wire;
`, m.Target.ProtocolVersion)
}

func emitRustEnums(definitions []definition) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	for _, item := range definitions {
		if item.Kind == manifest.KindEnum {
			emitRustEnum(&b, item)
		}
	}
	return strings.TrimSpace(b.String()) + "\n"
}

func emitRustTypes(definitions []definition) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	b.WriteString("use crate::enums::*;\n\n")
	b.WriteString("use crate::wire;\n\n")
	for _, item := range definitions {
		switch item.Kind {
		case manifest.KindStruct:
			if item.Tuple {
				derive := "Clone, Copy, Debug, Default, PartialEq"
				if item.WrapperCodec != "" && item.Fields[0].Type != "f32" && item.Fields[0].Type != "f64" {
					derive += ", Eq, Hash"
				}
				fmt.Fprintf(&b, "#[derive(%s)]\npub struct %s(pub %s);\n\n", derive, item.Name, item.Fields[0].Type)
				if item.WrapperCodec != "" {
					fmt.Fprintf(&b, "impl wire::WireCodec for %s {\n    fn encode<W: std::io::Write>(&self, writer: &mut W) -> std::io::Result<()> {\n        <wire::%s as wire::WireCodec>::encode(&wire::%s(self.0), writer)\n    }\n\n    fn decode<R: std::io::Read>(reader: &mut R) -> std::io::Result<Self> {\n        <wire::%s as wire::WireCodec>::decode(reader).map(|value| Self(value.0))\n    }\n}\n\n", item.Name, item.WrapperCodec, item.WrapperCodec, item.WrapperCodec)
				}
				continue
			}
			fmt.Fprintf(&b, "#[derive(Clone, Debug, Default, PartialEq)]\npub struct %s {\n", item.Name)
			for _, field := range item.Fields {
				for _, doc := range field.Docs {
					fmt.Fprintf(&b, "    %s\n", doc)
				}
				fmt.Fprintf(&b, "    pub %s: %s,\n", field.Name, field.Type)
			}
			b.WriteString("}\n\n")
		case manifest.KindUnion:
			deriveDefault := unionCanDeriveDefault(item)
			derive := "Clone, Debug, PartialEq"
			if deriveDefault {
				derive += ", Default"
			}
			fmt.Fprintf(&b, "#[derive(%s)]\npub enum %s {\n", derive, item.Name)
			for index, variant := range item.Union {
				if isPlaceholderVariantName(variant.SourceName) {
					fmt.Fprintf(&b, "    /// Naming overlay required: source placeholder `%s`.\n", variant.SourceName)
				}
				if deriveDefault && index == 0 {
					b.WriteString("    #[default]\n")
				}
				if len(variant.Fields) != 0 {
					fmt.Fprintf(&b, "    %s {\n", variant.Name)
					for _, field := range variant.Fields {
						for _, doc := range field.Docs {
							fmt.Fprintf(&b, "        %s\n", doc)
						}
						fmt.Fprintf(&b, "        %s: %s,\n", field.Name, field.Type)
					}
					b.WriteString("    },\n")
				} else if variant.Payload == "" {
					fmt.Fprintf(&b, "    %s,\n", variant.Name)
				} else {
					fmt.Fprintf(&b, "    %s(%s),\n", variant.Name, variant.Payload)
				}
			}
			b.WriteString("}\n\n")
			fmt.Fprintf(&b, "impl %s {\n    pub fn discriminant(&self) -> %s {\n        match self {\n", item.Name, item.Control)
			for _, variant := range item.Union {
				if len(variant.Fields) != 0 {
					fmt.Fprintf(&b, "            Self::%s { .. } => %d,\n", variant.Name, variant.Discriminant)
				} else if variant.Payload != "" {
					fmt.Fprintf(&b, "            Self::%s(..) => %d,\n", variant.Name, variant.Discriminant)
				} else {
					fmt.Fprintf(&b, "            Self::%s => %d,\n", variant.Name, variant.Discriminant)
				}
			}
			b.WriteString("        }\n    }\n}\n\n")
			if len(item.Union) > 0 && !deriveDefault {
				first := item.Union[0]
				fmt.Fprintf(&b, "impl Default for %s {\n    fn default() -> Self {\n", item.Name)
				if len(first.Fields) != 0 {
					fmt.Fprintf(&b, "        Self::%s {\n", first.Name)
					for _, field := range first.Fields {
						fmt.Fprintf(&b, "            %s: Default::default(),\n", field.Name)
					}
					b.WriteString("        }\n")
				} else if first.Payload != "" {
					fmt.Fprintf(&b, "        Self::%s(Default::default())\n", first.Name)
				} else {
					fmt.Fprintf(&b, "        Self::%s\n", first.Name)
				}
				b.WriteString("    }\n}\n\n")
			}
		case manifest.KindBitset:
			fmt.Fprintf(&b, "/// Stores the %d-bit value used by the wire bitset encoding.\n", item.BitLength)
			fmt.Fprintf(&b, "#[derive(Clone, Debug, Default, PartialEq, Eq)]\npub struct %s(pub [u64; %d]);\n\n", item.Name, (item.BitLength+63)/64)
		}
	}
	return strings.TrimSpace(b.String()) + "\n"
}

func emitCargo(m manifest.Manifest, g *generator) string {
	var b strings.Builder
	b.WriteString("# Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	crateName := "bedrock-protocol-" + strings.NewReplacer(".", "-", "_", "-").Replace(m.Target.MinecraftVersion)
	fmt.Fprintf(&b, "[package]\nname = \"%s\"\nversion = \"0.1.0\"\nedition = \"2024\"\npublish = false\n", crateName)
	if g.usesUUID || g.usesGlam || g.usesBytes {
		b.WriteString("\n[dependencies]\n")
		if g.usesBytes {
			b.WriteString("bytes = \"1\"\n")
		}
		if g.usesGlam {
			b.WriteString("glam = \"0.30\"\n")
		}
		if g.usesUUID {
			b.WriteString("uuid = \"1\"\n")
		}
	}
	return b.String()
}

func emitWire(usesNbt bool) string {
	source := `// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use std::io::{self, Read, Write};

pub trait WireCodec: Sized {
    fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()>;
    fn decode<R: Read>(reader: &mut R) -> io::Result<Self>;
}

macro_rules! fixed_codec {
    ($name:ident, $inner:ty, $size:expr, $write:ident, $read:ident) => {
        #[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
        pub struct $name(pub $inner);

        impl WireCodec for $name {
            fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
                writer.write_all(&self.0.$write())
            }

            fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
                let mut bytes = [0u8; $size];
                reader.read_exact(&mut bytes)?;
                Ok(Self(<$inner>::$read(bytes)))
            }
        }
    };
}

macro_rules! fixed_float_codec {
    ($name:ident, $inner:ty, $size:expr, $write:ident, $read:ident) => {
        #[derive(Clone, Copy, Debug, Default, PartialEq)]
        pub struct $name(pub $inner);

        impl WireCodec for $name {
            fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
                writer.write_all(&self.0.$write())
            }

            fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
                let mut bytes = [0u8; $size];
                reader.read_exact(&mut bytes)?;
                Ok(Self(<$inner>::$read(bytes)))
            }
        }
    };
}

fixed_codec!(I8, i8, 1, to_ne_bytes, from_ne_bytes);
fixed_codec!(U8, u8, 1, to_ne_bytes, from_ne_bytes);
fixed_codec!(I16LE, i16, 2, to_le_bytes, from_le_bytes);
fixed_codec!(I16BE, i16, 2, to_be_bytes, from_be_bytes);
fixed_codec!(U16LE, u16, 2, to_le_bytes, from_le_bytes);
fixed_codec!(U16BE, u16, 2, to_be_bytes, from_be_bytes);
fixed_codec!(I32LE, i32, 4, to_le_bytes, from_le_bytes);
fixed_codec!(I32BE, i32, 4, to_be_bytes, from_be_bytes);
fixed_codec!(U32LE, u32, 4, to_le_bytes, from_le_bytes);
fixed_codec!(U32BE, u32, 4, to_be_bytes, from_be_bytes);
fixed_codec!(I64LE, i64, 8, to_le_bytes, from_le_bytes);
fixed_codec!(I64BE, i64, 8, to_be_bytes, from_be_bytes);
fixed_codec!(U64LE, u64, 8, to_le_bytes, from_le_bytes);
fixed_codec!(U64BE, u64, 8, to_be_bytes, from_be_bytes);
fixed_float_codec!(F32LE, f32, 4, to_le_bytes, from_le_bytes);
fixed_float_codec!(F32BE, f32, 4, to_be_bytes, from_be_bytes);
fixed_float_codec!(F64LE, f64, 8, to_le_bytes, from_le_bytes);
fixed_float_codec!(F64BE, f64, 8, to_be_bytes, from_be_bytes);

fn write_var_u64<W: Write>(writer: &mut W, mut value: u64) -> io::Result<()> {
    while value >= 0x80 {
        writer.write_all(&[(value as u8) | 0x80])?;
        value >>= 7;
    }
    writer.write_all(&[value as u8])
}

fn read_var_u64<R: Read>(reader: &mut R) -> io::Result<u64> {
    let mut value = 0u64;
    for shift in (0..64).step_by(7) {
        let mut byte = [0u8; 1];
        reader.read_exact(&mut byte)?;
        value |= u64::from(byte[0] & 0x7f) << shift;
        if byte[0] & 0x80 == 0 {
            return Ok(value);
        }
    }
    Err(io::Error::new(io::ErrorKind::InvalidData, "varint too large"))
}

macro_rules! var_codec {
    ($name:ident, $inner:ty) => {
        #[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
        pub struct $name(pub $inner);

        impl WireCodec for $name {
            fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
                write_var_u64(writer, self.0 as u64)
            }

            fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
                Ok(Self(read_var_u64(reader)? as $inner))
            }
        }
    };
}

var_codec!(VarInt, i32);
var_codec!(VarUInt, u32);
var_codec!(VarLong, i64);
var_codec!(VarULong, u64);

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ZigZag32(pub i32);

impl WireCodec for ZigZag32 {
    fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
        write_var_u64(writer, ((self.0 << 1) ^ (self.0 >> 31)) as u32 as u64)
    }

    fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
        let raw = read_var_u64(reader)? as u32;
        Ok(Self(((raw >> 1) as i32) ^ -((raw & 1) as i32)))
    }
}

#[derive(Clone, Copy, Debug, Default, PartialEq, Eq, Hash)]
pub struct ZigZag64(pub i64);

impl WireCodec for ZigZag64 {
    fn encode<W: Write>(&self, writer: &mut W) -> io::Result<()> {
        write_var_u64(writer, ((self.0 << 1) ^ (self.0 >> 63)) as u64)
    }

    fn decode<R: Read>(reader: &mut R) -> io::Result<Self> {
        let raw = read_var_u64(reader)?;
        Ok(Self(((raw >> 1) as i64) ^ -((raw & 1) as i64)))
    }
}
`
	if usesNbt {
		source += `
#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct NetworkNbt(pub bytes::Bytes);

#[derive(Clone, Debug, Default, PartialEq, Eq)]
pub struct PersistentNbt(pub bytes::Bytes);
`
	}
	return source
}

func emitRustPackets(infos []packetInfo) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	b.WriteString("use crate::enums::*;\nuse crate::types::*;\n")
	b.WriteString("use crate::wire;\n\n")
	for _, info := range infos {
		emitRustPacketDefinition(&b, info)
	}
	emitPacketRegistry(&b, infos)
	return strings.TrimSpace(b.String()) + "\n"
}

func emitRustPacket(info packetInfo) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	b.WriteString("use crate::enums::*;\nuse crate::types::*;\nuse crate::wire;\n\n")
	emitRustPacketDefinition(&b, info)
	return b.String()
}

func emitRustPacketDefinition(b *strings.Builder, info packetInfo) {
	fmt.Fprintf(b, "#[derive(Clone, Debug, Default, PartialEq)]\npub struct %s {\n", info.name)
	for _, field := range info.fields {
		for _, doc := range field.docs {
			fmt.Fprintf(b, "    %s\n", doc)
		}
		fmt.Fprintf(b, "    pub %s: %s,\n", field.name, field.typ)
	}
	fmt.Fprintf(b, "}\n\nimpl %s {\n    pub const ID: u32 = %d;\n}\n", info.name, info.packet.ID)
}

func emitPacketRegistry(b *strings.Builder, infos []packetInfo) {
	b.WriteString("\n#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]\n#[repr(u32)]\npub enum PacketId {\n")
	for _, info := range infos {
		fmt.Fprintf(b, "    %s = %d,\n", info.name, info.packet.ID)
	}
	b.WriteString("}\n\nimpl PacketId {\n    pub fn from_raw(raw: u32) -> Option<Self> {\n        match raw {\n")
	for _, info := range infos {
		fmt.Fprintf(b, "            %d => Some(Self::%s),\n", info.packet.ID, info.name)
	}
	b.WriteString("            _ => None,\n        }\n    }\n}\n\n")
	b.WriteString("#[derive(Clone, Debug, PartialEq)]\npub enum Packet {\n")
	for _, info := range infos {
		if packetNeedsBox(info) {
			fmt.Fprintf(b, "    %s(Box<%s>),\n", info.name, info.name)
		} else {
			fmt.Fprintf(b, "    %s(%s),\n", info.name, info.name)
		}
	}
	b.WriteString("}\n\n")
	for _, info := range infos {
		fmt.Fprintf(b, "impl From<%s> for Packet {\n    fn from(value: %s) -> Self {\n", info.name, info.name)
		if packetNeedsBox(info) {
			fmt.Fprintf(b, "        Self::%s(Box::new(value))\n", info.name)
		} else {
			fmt.Fprintf(b, "        Self::%s(value)\n", info.name)
		}
		b.WriteString("    }\n}\n\n")
	}
}

func packetNeedsBox(info packetInfo) bool {
	return len(info.fields) >= 8
}

func fieldDocs(field manifest.Field) []string {
	if field.Encode.Kind == manifest.KindOptional {
		return []string{"/// Wire presence: optional value is preceded by a presence marker."}
	}
	return nil
}

func moduleFileName(value string) string {
	name := fieldName(value)
	return strings.TrimPrefix(name, "r#")
}

func rustModuleName(value string) string {
	if rustKeywords[value] || rustUnrawableKeywords[value] {
		return "r#" + value
	}
	return value
}

type rustFieldInfo struct {
	name string
	typ  string
	docs []string
}

func (g *generator) rustType(node manifest.Node, hint string) (string, error) {
	if typ, matched, err := g.nativeRustType(node); matched || err != nil {
		return typ, err
	}
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return "", fmt.Errorf("primitive has no shape")
		}
		return primitiveRustType(node.Primitive.Code)
	case manifest.KindString:
		return "String", nil
	case manifest.KindBytes:
		g.usesBytes = true
		return "bytes::Bytes", nil
	case manifest.KindBitset:
		name := fmt.Sprintf("Bitset%d", node.Length)
		if _, ok := g.definitions[name]; !ok {
			g.definitions[name] = definition{Name: name, Kind: manifest.KindBitset, BitLength: node.Length}
		}
		return name, nil
	case manifest.KindArray:
		if node.Element == nil {
			return "", fmt.Errorf("array has no element")
		}
		element, err := g.rustType(*node.Element, hint+"Item")
		return "Vec<" + element + ">", err
	case manifest.KindFixedArray:
		if node.Element == nil {
			return "", fmt.Errorf("fixed array has no element")
		}
		element, err := g.rustType(*node.Element, hint+"Item")
		return fmt.Sprintf("[%s; %d]", element, node.Length), err
	case manifest.KindSequence:
		return "", fmt.Errorf("sequence nodes require a target-specific representation")
	case manifest.KindOptional:
		if node.Value == nil {
			return "", fmt.Errorf("optional has no value")
		}
		valueNode := *node.Value
		// Cereal uses an always-present outer optional around some optional
		// values. Retain both markers in the manifest, but expose only the
		// meaningful inner state, matching the Go and gophertunnel APIs.
		if valueNode.Kind == manifest.KindOptional {
			if valueNode.Value == nil {
				return "", fmt.Errorf("nested optional has no value")
			}
			valueNode = *valueNode.Value
		}
		value, err := g.rustType(valueNode, hint+"Value")
		return "Option<" + value + ">", err
	case manifest.KindStruct:
		return g.registerStruct(node, hint)
	case manifest.KindMap:
		if node.Key == nil || node.Value == nil {
			return "", fmt.Errorf("map has no key/value")
		}
		key, err := g.rustType(*node.Key, hint+"Key")
		if err != nil {
			return "", err
		}
		value, err := g.rustType(*node.Value, hint+"Value")
		if err != nil {
			return "", err
		}
		return "Vec<(" + key + ", " + value + ")>", nil
	case manifest.KindUnion:
		name := g.registerIdentity(node, hint+"Union")
		if _, ok := g.definitions[name]; !ok {
			control, err := unionControlType(node)
			if err != nil {
				return "", err
			}
			g.definitions[name] = definition{Name: name, Kind: manifest.KindUnion, Control: control}
			variants := make([]rustVariant, 0, len(node.Variants))
			used := map[string]bool{}
			for _, variant := range node.Variants {
				variantName := strings.TrimSuffix(rustPascalName(shortTypeName(variant.Name)), "Payload")
				variantName = uniqueTypeVariant(variantName, used)
				payload := ""
				var fields []rustField
				if !emptyPayload(variant.Encode) {
					var err error
					if g.inlineRustVariant(variant.Encode) {
						fields, err = g.rustFieldsForUnion(variant.Encode, name+variantName, node.Control)
					} else {
						payload, err = g.rustType(variant.Encode, name+variantName)
					}
					if err != nil {
						return "", err
					}
					if payload != "" && g.largeRustType(payload) {
						payload = "Box<" + payload + ">"
					}
				}
				variants = append(variants, rustVariant{Name: variantName, Payload: payload, Fields: fields, Discriminant: variant.Value, SourceName: variant.Name})
			}
			item := g.definitions[name]
			item.Union = variants
			g.definitions[name] = item
		}
		return name, nil
	case manifest.KindEnum:
		if node.Primitive == nil {
			return "", fmt.Errorf("enum has no underlying primitive")
		}
		underlying, err := primitiveRustRawType(node.Primitive.Code)
		if err != nil {
			return "", err
		}
		name := g.registerIdentity(node, hint+"Enum")
		if _, ok := g.definitions[name]; !ok {
			g.definitions[name] = definition{Name: name, Kind: manifest.KindEnum, Underlying: underlying, Variants: append([]manifest.Variant(nil), node.Variants...)}
		}
		return name, nil
	case manifest.KindReserved, manifest.KindIgnored:
		return "", fmt.Errorf("%s nodes require explicit write/discard codec semantics", node.Kind)
	case manifest.KindRecursive:
		name, ok := g.identities[node.Target]
		if !ok {
			return "", fmt.Errorf("recursive target %q is not a registered named type", node.Target)
		}
		return name, nil
	case manifest.KindVoid:
		return "()", nil
	case manifest.KindOpaque, manifest.KindUnresolved:
		return "", fmt.Errorf("reachable %s node cannot be emitted: %s", node.Kind, node.Reason)
	default:
		return "", fmt.Errorf("unsupported node kind %q", node.Kind)
	}
}

func (g *generator) nativeRustType(node manifest.Node) (string, bool, error) {
	if node.Kind == manifest.KindPrimitive && node.Primitive != nil {
		switch node.Primitive.Code {
		case "uuid":
			g.usesUUID = true
			return "uuid::Uuid", true, nil
		case "nbt_le":
			if !manifest.ValidNBTEncoding(node.Encoding) {
				return "", true, fmt.Errorf("NBT node has invalid encoding %q", node.Encoding)
			}
			g.usesNbt = true
			g.usesBytes = true
			if node.Encoding == string(manifest.NBTNetwork) {
				return "wire::NetworkNbt", true, nil
			}
			return "wire::PersistentNbt", true, nil
		}
	}
	if node.Kind != manifest.KindStruct {
		return "", false, nil
	}
	switch node.TypeID {
	case "Vec2":
		if !isPrimitiveStruct(node, "f32le", "f32le") {
			return "", true, fmt.Errorf("native Vec2 mapping requires exactly two f32le fields")
		}
		g.usesGlam = true
		return "glam::Vec2", true, nil
	case "Vec3":
		if !isPrimitiveStruct(node, "f32le", "f32le", "f32le") {
			return "", true, fmt.Errorf("native Vec3 mapping requires exactly three f32le fields")
		}
		g.usesGlam = true
		return "glam::Vec3", true, nil
	default:
		return "", false, nil
	}
}

func isPrimitiveStruct(node manifest.Node, codes ...string) bool {
	if len(node.Fields) != len(codes) {
		return false
	}
	for index, code := range codes {
		field := node.Fields[index].Encode
		if field.Kind != manifest.KindPrimitive || field.Primitive == nil || field.Primitive.Code != code {
			return false
		}
	}
	return true
}

func (g *generator) registerStruct(node manifest.Node, hint string) (string, error) {
	name := g.registerIdentity(node, hint+"Struct")
	if _, ok := g.definitions[name]; ok {
		return name, nil
	}
	if raw, codec, ok := singleFieldWrapper(node); ok {
		g.definitions[name] = definition{
			Name:         name,
			Kind:         manifest.KindStruct,
			Fields:       []rustField{{Name: "value", Type: raw}},
			Tuple:        true,
			WrapperCodec: codec,
		}
		return name, nil
	}
	g.definitions[name] = definition{Name: name, Kind: manifest.KindStruct}
	fields, err := g.rustFields(node, name)
	if err != nil {
		return "", err
	}
	definition := g.definitions[name]
	definition.Fields = fields
	g.definitions[name] = definition
	return name, nil
}

func (g *generator) rustFields(node manifest.Node, parentName string) ([]rustField, error) {
	return g.rustFieldsForUnion(node, parentName, nil)
}

func (g *generator) rustFieldsForUnion(node manifest.Node, parentName string, control *manifest.Node) ([]rustField, error) {
	used := map[string]bool{}
	fields := make([]rustField, 0, len(node.Fields))
	for _, field := range node.Fields {
		if control != nil && isUnionDiscriminantField(field, *control) {
			continue
		}
		fieldName := uniqueField(fieldName(field.Name), used)
		typ, err := g.rustType(field.Encode, parentName+typeName(field.Name))
		if err != nil {
			return nil, err
		}
		fields = append(fields, rustField{Name: fieldName, Type: typ, Docs: fieldDocs(field)})
	}
	if control != nil {
		g.boxLargeUnionFields(fields)
	}
	return fields, nil
}

func singleFieldWrapper(node manifest.Node) (string, string, bool) {
	if len(node.Fields) != 1 {
		return "", "", false
	}
	field := node.Fields[0].Encode
	if field.Kind != manifest.KindPrimitive || field.Primitive == nil || field.Primitive.Code == "bool" {
		return "", "", false
	}
	raw, err := primitiveRustRawType(field.Primitive.Code)
	if err != nil {
		return "", "", false
	}
	return raw, wireCodecType(field.Primitive.Code), true
}

func wireCodecType(code string) string {
	codeToType := map[string]string{
		"i8":         "I8",
		"u8":         "U8",
		"i16le":      "I16LE",
		"i16be":      "I16BE",
		"u16le":      "U16LE",
		"u16be":      "U16BE",
		"i32le":      "I32LE",
		"i32be":      "I32BE",
		"u32le":      "U32LE",
		"u32be":      "U32BE",
		"i64le":      "I64LE",
		"i64be":      "I64BE",
		"u64le":      "U64LE",
		"u64be":      "U64BE",
		"var_i32":    "VarInt",
		"var_u32":    "VarUInt",
		"var_i64":    "VarLong",
		"var_u64":    "VarULong",
		"f32le":      "F32LE",
		"f32be":      "F32BE",
		"f64le":      "F64LE",
		"f64be":      "F64BE",
		"zigzag_i32": "ZigZag32",
		"zigzag_i64": "ZigZag64",
	}
	return codeToType[code]
}

func (g *generator) boxLargeUnionFields(fields []rustField) {
	for index := range fields {
		if g.largeRustType(fields[index].Type) {
			fields[index].Type = "Box<" + fields[index].Type + ">"
		}
	}
}

func (g *generator) largeRustType(typ string) bool {
	if strings.HasPrefix(typ, "Box<") {
		return false
	}
	return g.estimateRustType(typ, map[string]bool{}) > 128
}

func (g *generator) estimatePacketSize(fields []rustFieldInfo) int {
	size := 0
	for _, field := range fields {
		size += g.estimateRustType(field.typ, map[string]bool{})
	}
	return size
}

func (g *generator) estimateRustType(typ string, visiting map[string]bool) int {
	if strings.HasPrefix(typ, "Box<") || strings.HasPrefix(typ, "Vec<") || strings.HasPrefix(typ, "Option<") || typ == "String" || typ == "bytes::Bytes" {
		return 24
	}
	if strings.HasPrefix(typ, "wire::") {
		return 8
	}
	switch typ {
	case "bool", "i8", "u8":
		return 1
	case "i16", "u16":
		return 2
	case "i32", "u32", "f32":
		return 4
	case "i64", "u64", "f64":
		return 8
	case "uuid::Uuid":
		return 16
	case "glam::Vec2":
		return 8
	case "glam::Vec3":
		return 12
	}
	item, ok := g.definitions[typ]
	if !ok || visiting[typ] {
		return 8
	}
	visiting[typ] = true
	size := 0
	for _, field := range item.Fields {
		size += g.estimateRustType(field.Type, visiting)
	}
	delete(visiting, typ)
	return size
}

func isUnionDiscriminantField(field manifest.Field, control manifest.Node) bool {
	name := fieldName(field.Name)
	name = strings.TrimPrefix(name, "r#")
	name = strings.TrimSuffix(name, "_")
	if name != "type" || control.Primitive == nil {
		return false
	}
	if field.Encode.Primitive == nil {
		return false
	}
	return field.Encode.Primitive.Code == control.Primitive.Code
}

func (g *generator) inlineRustVariant(node manifest.Node) bool {
	if node.Kind != manifest.KindStruct || len(node.Fields) == 0 || node.TypeID == "Vec2" || node.TypeID == "Vec3" {
		return false
	}
	return node.TypeID == "" || !g.standalone[node.TypeID]
}

func (g *generator) registerIdentity(node manifest.Node, hint string) string {
	key := node.TypeID
	inferred := inferredTypeName(node)
	if key == "" {
		identityHint := inferred
		if identityHint == "" {
			identityHint = hint
		} else {
			identityHint += "/" + unionIdentity(node)
		}
		key = fmt.Sprintf("%s/%s/%s", node.Kind, node.Semantic, identityHint)
	}
	if name, ok := g.identities[key]; ok {
		return name
	}
	base := hint
	if node.TypeID != "" {
		base = node.TypeID
	} else if node.Semantic != "" {
		base = node.Semantic
	} else if inferred != "" {
		base = inferred
	}
	candidate := publicTypeName(base)
	if g.used[candidate] {
		// Packet names own their unqualified name. Give a nested definition a
		// semantic role instead of leaking a numeric collision suffix into the API.
		switch node.Kind {
		case manifest.KindEnum:
			candidate += "Type"
		case manifest.KindUnion:
			hinted := strings.TrimSuffix(publicTypeName(hint), "Union")
			if hinted != "" && !g.used[hinted] {
				candidate = hinted
			} else {
				candidate += "Data"
			}
		default:
			candidate += "Data"
		}
	}
	name := g.unique(candidate)
	g.identities[key] = name
	return name
}

func unionIdentity(node manifest.Node) string {
	type variantIdentity struct {
		Value  int64  `json:"value"`
		Name   string `json:"name"`
		TypeID string `json:"type_id,omitempty"`
	}
	identity := struct {
		Control  string            `json:"control"`
		Variants []variantIdentity `json:"variants"`
	}{}
	if node.Control != nil && node.Control.Primitive != nil {
		identity.Control = node.Control.Primitive.Code
	}
	for _, variant := range node.Variants {
		identity.Variants = append(identity.Variants, variantIdentity{Value: variant.Value, Name: variant.Name, TypeID: variant.Encode.TypeID})
	}
	encoded, _ := json.Marshal(identity)
	return string(encoded)
}

func inferredTypeName(node manifest.Node) string {
	if node.Kind != manifest.KindUnion || len(node.Variants) == 0 {
		return ""
	}
	prefix := ""
	for _, variant := range node.Variants {
		qualified := variant.Name
		if !strings.Contains(qualified, "::") && variant.Encode.TypeID != "" {
			qualified = variant.Encode.TypeID
		}
		position := strings.LastIndex(qualified, "::")
		if position < 0 {
			return ""
		}
		candidate := qualified[:position]
		if prefix == "" {
			prefix = candidate
		} else if prefix != candidate {
			return ""
		}
	}
	return prefix
}

func shortTypeName(name string) string {
	if position := strings.LastIndex(name, "::"); position >= 0 {
		return name[position+2:]
	}
	return name
}

func emptyPayload(node manifest.Node) bool {
	return node.Kind == manifest.KindVoid || node.Kind == manifest.KindStruct && len(node.Fields) == 0
}

func uniqueTypeVariant(name string, used map[string]bool) string {
	if !used[name] {
		used[name] = true
		return name
	}
	for index := 2; ; index++ {
		candidate := fmt.Sprintf("%s%d", name, index)
		if !used[candidate] {
			used[candidate] = true
			return candidate
		}
	}
}

func (g *generator) unique(base string) string {
	if !g.used[base] {
		g.used[base] = true
		return base
	}
	for index := 2; ; index++ {
		candidate := fmt.Sprintf("%s%d", base, index)
		if !g.used[candidate] {
			g.used[candidate] = true
			return candidate
		}
	}
}

func primitiveRustType(code string) (string, error) {
	switch code {
	case "bool":
		return "bool", nil
	case "i8":
		return "wire::I8", nil
	case "u8":
		return "wire::U8", nil
	case "i16le":
		return "wire::I16LE", nil
	case "i16be":
		return "wire::I16BE", nil
	case "u16le":
		return "wire::U16LE", nil
	case "u16be":
		return "wire::U16BE", nil
	case "i32le":
		return "wire::I32LE", nil
	case "i32be":
		return "wire::I32BE", nil
	case "u32le":
		return "wire::U32LE", nil
	case "u32be":
		return "wire::U32BE", nil
	case "i64le":
		return "wire::I64LE", nil
	case "i64be":
		return "wire::I64BE", nil
	case "u64le":
		return "wire::U64LE", nil
	case "u64be":
		return "wire::U64BE", nil
	case "var_i32":
		return "wire::VarInt", nil
	case "var_u32":
		return "wire::VarUInt", nil
	case "var_i64":
		return "wire::VarLong", nil
	case "var_u64":
		return "wire::VarULong", nil
	case "zigzag_i32":
		return "wire::ZigZag32", nil
	case "zigzag_i64":
		return "wire::ZigZag64", nil
	case "f32le":
		return "wire::F32LE", nil
	case "f32be":
		return "wire::F32BE", nil
	case "f64le":
		return "wire::F64LE", nil
	case "f64be":
		return "wire::F64BE", nil
	default:
		return "", fmt.Errorf("unsupported primitive code %q", code)
	}
}

func primitiveRustRawType(code string) (string, error) {
	switch code {
	case "bool":
		return "bool", nil
	case "i8":
		return "i8", nil
	case "u8":
		return "u8", nil
	case "i16le", "i16be":
		return "i16", nil
	case "u16le", "u16be":
		return "u16", nil
	case "i32le", "i32be", "var_i32", "zigzag_i32":
		return "i32", nil
	case "u32le", "u32be", "var_u32":
		return "u32", nil
	case "i64le", "i64be", "var_i64", "zigzag_i64":
		return "i64", nil
	case "u64le", "u64be", "var_u64":
		return "u64", nil
	case "f32le", "f32be":
		return "f32", nil
	case "f64le", "f64be":
		return "f64", nil
	default:
		return "", fmt.Errorf("unsupported primitive code %q", code)
	}
}

func unionControlType(node manifest.Node) (string, error) {
	if node.Control == nil || node.Control.Primitive == nil {
		return "", fmt.Errorf("union has no primitive control")
	}
	return primitiveRustRawType(node.Control.Primitive.Code)
}

func emitRustEnum(b *strings.Builder, item definition) {
	derive := "Clone, Copy, Debug, PartialEq, Eq, Hash"
	if len(item.Variants) > 0 {
		derive += ", Default"
	}
	fmt.Fprintf(b, "#[derive(%s)]\npub enum %s {\n", derive, item.Name)
	used := map[string]bool{}
	variantNames := make([]string, 0, len(item.Variants))
	for index, variant := range item.Variants {
		name := uniqueTypeVariant(enumVariantName(item.Name, variant.Name), used)
		variantNames = append(variantNames, name)
		if index == 0 {
			b.WriteString("    #[default]\n")
		}
		if isPlaceholderVariantName(variant.Name) {
			fmt.Fprintf(b, "    /// Naming overlay required: source placeholder `%s`.\n", variant.Name)
		}
		fmt.Fprintf(b, "    %s,\n", name)
	}
	unknownName := uniqueTypeVariant("Unknown", used)
	fmt.Fprintf(b, "    %s(%s),\n", unknownName, item.Underlying)
	b.WriteString("}\n\n")
	fmt.Fprintf(b, "impl From<%s> for %s {\n    fn from(value: %s) -> Self {\n        match value {\n", item.Underlying, item.Name, item.Underlying)
	for index, variant := range item.Variants {
		fmt.Fprintf(b, "            %d => Self::%s,\n", variant.Value, variantNames[index])
	}
	fmt.Fprintf(b, "            value => Self::%s(value),\n        }\n    }\n}\n\n", unknownName)
	fmt.Fprintf(b, "impl %s {\n    pub fn to_raw(self) -> %s {\n        match self {\n", item.Name, item.Underlying)
	for index, variant := range item.Variants {
		fmt.Fprintf(b, "            Self::%s => %d,\n", variantNames[index], variant.Value)
	}
	fmt.Fprintf(b, "            Self::%s(value) => value,\n        }\n    }\n}\n\n", unknownName)
	fmt.Fprintf(b, "impl From<%s> for %s {\n    fn from(value: %s) -> Self {\n        value.to_raw()\n    }\n}\n\n", item.Name, item.Underlying, item.Name)
	if len(variantNames) == 0 {
		fmt.Fprintf(b, "impl Default for %s {\n    fn default() -> Self {\n        Self::%s(0)\n    }\n}\n\n", item.Name, unknownName)
	}
}

func enumVariantName(enumName, value string) string {
	name := rustPascalName(value)
	base := enumName
	for _, suffix := range []string{"PacketType", "Settings", "Type", "Mode", "Status", "Action", "Event"} {
		base = strings.TrimSuffix(base, suffix)
	}
	if base != "" && strings.HasSuffix(name, base) && len(name) > len(base) {
		name = strings.TrimSuffix(name, base)
	}
	name = strings.ReplaceAll(name, "Joincode", "JoinCode")
	name = strings.ReplaceAll(name, "Molang", "MoLang")
	if name == "" {
		name = "Value"
	}
	if name == "Self" {
		name = "SelfValue"
	}
	return name
}

func rustPascalName(value string) string {
	parts := strings.FieldsFunc(value, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
	var b strings.Builder
	for _, part := range parts {
		for _, word := range camelWords(part) {
			runes := []rune(strings.ToLower(word))
			if len(runes) == 0 {
				continue
			}
			runes[0] = unicode.ToUpper(runes[0])
			b.WriteString(string(runes))
		}
	}
	if b.Len() == 0 {
		return "Value"
	}
	return strings.NewReplacer(
		"Fishhook", "FishHook",
		"Fishpos", "FishPos",
		"Hooktime", "HookTime",
		"Tntcart", "TntCart",
		"Joincode", "JoinCode",
	).Replace(b.String())
}

func camelWords(value string) []string {
	runes := []rune(value)
	if len(runes) == 0 {
		return nil
	}
	start := 0
	words := make([]string, 0, 4)
	for index := 1; index < len(runes); index++ {
		previous, current := runes[index-1], runes[index]
		nextLower := index+1 < len(runes) && unicode.IsLower(runes[index+1])
		boundary := unicode.IsUpper(current) && (unicode.IsLower(previous) || unicode.IsDigit(previous) || unicode.IsUpper(previous) && nextLower)
		if boundary {
			words = append(words, string(runes[start:index]))
			start = index
		}
	}
	return append(words, string(runes[start:]))
}

func packetTypeName(value string) string {
	name := typeName(value)
	name = strings.TrimSuffix(name, "Packet")
	if name == "" {
		return "Packet"
	}
	return name
}

func publicTypeName(value string) string {
	value = strings.TrimSuffix(value, ".json#")
	value = strings.TrimSuffix(value, ".json")
	value = strings.TrimPrefix(value, "enums/")
	value = stripSharedTypeVersion(value)
	value = strings.ReplaceAll(value, "Packet::", "::")
	value = strings.ReplaceAll(value, "PacketPayload::", "::")
	value = strings.ReplaceAll(value, "PacketPayload", "")
	value = collapseRedundantGroup(value)
	if strings.HasSuffix(value, "PayloadUnion") {
		value = strings.TrimSuffix(value, "PayloadUnion") + "Value"
	}
	value = strings.TrimSuffix(value, "Union")
	value = strings.TrimSuffix(value, "Payload")
	return strings.ReplaceAll(typeName(value), "Molang", "MoLang")
}

func collapseRedundantGroup(value string) string {
	parts := strings.Split(value, "::")
	if len(parts) == 2 && parts[0] == parts[1]+"Group" {
		return parts[1]
	}
	return value
}

func stripSharedTypeVersion(value string) string {
	const prefix = "SharedTypes::"
	if !strings.HasPrefix(value, prefix) {
		return value
	}
	rest := strings.TrimPrefix(value, prefix)
	if separator := strings.Index(rest, "::"); separator >= 0 && strings.HasPrefix(rest[:separator], "v") {
		return rest[separator+2:]
	}
	return rest
}

func typeName(value string) string {
	var b strings.Builder
	upper := true
	for _, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if upper {
				r = unicode.ToUpper(r)
				upper = false
			}
			b.WriteRune(r)
		} else {
			upper = true
		}
	}
	result := b.String()
	if result == "" {
		return "Generated"
	}
	if unicode.IsDigit([]rune(result)[0]) {
		return "Generated" + result
	}
	return result
}

func fieldName(value string) string {
	value = strings.NewReplacer("'s", "", "'S", "", "’s", "", "’S", "").Replace(value)
	runes := []rune(value)
	var b strings.Builder
	boundary := false
	for index, r := range runes {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			boundary = b.Len() != 0
			continue
		}
		upper := unicode.IsUpper(r)
		previousLowerOrDigit := index > 0 && (unicode.IsLower(runes[index-1]) || unicode.IsDigit(runes[index-1]))
		nextLower := index+1 < len(runes) && unicode.IsLower(runes[index+1])
		if b.Len() != 0 && (boundary || upper && (previousLowerOrDigit || nextLower)) {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
		boundary = false
	}
	name := b.String()
	if name == "" {
		return "field"
	}
	name = strings.NewReplacer("no_pv_m", "no_pvm", "no_mv_p", "no_mvp").Replace(name)
	if rustUnrawableKeywords[name] || rustKeywords[name] {
		return name + "_"
	}
	return name
}

func isPlaceholderVariantName(name string) bool {
	if !strings.HasPrefix(name, "Empty") || len(name) == len("Empty") {
		return false
	}
	for _, r := range name[len("Empty"):] {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func unionCanDeriveDefault(item definition) bool {
	return len(item.Union) > 0 && len(item.Union[0].Fields) == 0 && item.Union[0].Payload == ""
}

var rustUnrawableKeywords = map[string]bool{"crate": true, "self": true, "super": true}

var rustKeywords = map[string]bool{
	"Self": true, "abstract": true, "as": true, "async": true, "await": true,
	"become": true, "box": true, "break": true, "const": true, "continue": true,
	"crate": true, "do": true, "dyn": true, "else": true, "enum": true,
	"extern": true, "false": true, "final": true, "fn": true, "for": true,
	"gen": true, "if": true, "impl": true, "in": true, "let": true,
	"loop": true, "macro": true, "match": true, "mod": true, "move": true,
	"mut": true, "override": true, "priv": true, "pub": true, "ref": true,
	"return": true, "self": true, "static": true, "struct": true, "super": true,
	"trait": true, "true": true, "try": true, "type": true, "typeof": true,
	"union": true, "unsafe": true, "unsized": true, "use": true, "virtual": true,
	"where": true, "while": true, "yield": true,
}

func uniqueField(base string, used map[string]bool) string {
	if !used[base] {
		used[base] = true
		return base
	}
	for index := 2; ; index++ {
		candidate := fmt.Sprintf("%s%d", base, index)
		if !used[candidate] {
			used[candidate] = true
			return candidate
		}
	}
}
