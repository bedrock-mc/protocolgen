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
	Name       string
	Kind       manifest.NodeKind
	Fields     []rustField
	Underlying string
	Variants   []manifest.Variant
	Union      []rustVariant
	BitLength  uint64
}

type rustVariant struct {
	Name    string
	Payload string
	Fields  []rustField
}

type rustField struct {
	Name string
	Type string
}

type packetInfo struct {
	packet manifest.Packet
	name   string
	fields []rustFieldInfo
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
			fields = append(fields, rustFieldInfo{name: fieldName, typ: typ})
		}
		infos = append(infos, packetInfo{packet: packet, name: name, fields: fields})
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
		"Cargo.toml":   emitCargo(m, g),
		"src/lib.rs":   emitLib(infos),
		"src/enums.rs": emitRustEnums(definitions),
		"src/types.rs": emitRustTypes(definitions, g.usesNbt),
	}
	usedFiles := map[string]bool{}
	var modules strings.Builder
	modules.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	for _, info := range infos {
		base := moduleFileName(info.name)
		fileName := base + ".rs"
		moduleName := rustModuleName(base)
		if usedFiles[fileName] {
			base += fmt.Sprintf("_%d", info.packet.ID)
			fileName = base + ".rs"
			moduleName = rustModuleName(base)
		}
		usedFiles[fileName] = true
		fmt.Fprintf(&modules, "mod %s;\npub use %s::*;\n", moduleName, moduleName)
		files["src/packets/"+fileName] = emitRustPacket(info)
	}
	files["src/packets/mod.rs"] = strings.TrimSpace(modules.String()) + "\n"
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

func emitLib(_ []packetInfo) string {
	return `// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#![allow(dead_code)]

mod enums;
pub use enums::*;
mod types;
pub use types::*;
mod packets;
pub use packets::*;
`
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

func emitRustTypes(definitions []definition, usesNbt bool) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	b.WriteString("use crate::enums::*;\n\n")
	if usesNbt {
		b.WriteString("#[derive(Clone, Debug, Default, PartialEq, Eq)]\npub struct Nbt(pub Vec<u8>);\n\n")
	}
	for _, item := range definitions {
		switch item.Kind {
		case manifest.KindStruct:
			fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq)]\npub struct %s {\n", item.Name)
			for _, field := range item.Fields {
				fmt.Fprintf(&b, "    pub %s: %s,\n", field.Name, field.Type)
			}
			b.WriteString("}\n\n")
		case manifest.KindUnion:
			fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq)]\npub enum %s {\n", item.Name)
			for _, variant := range item.Union {
				if len(variant.Fields) != 0 {
					fmt.Fprintf(&b, "    %s {\n", variant.Name)
					for _, field := range variant.Fields {
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
		case manifest.KindBitset:
			fmt.Fprintf(&b, "/// Stores the %d-bit value used by the wire bitset encoding.\n", item.BitLength)
			fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq, Eq)]\npub struct %s(pub [u64; %d]);\n\n", item.Name, (item.BitLength+63)/64)
		}
	}
	return strings.TrimSpace(b.String()) + "\n"
}

func emitCargo(m manifest.Manifest, g *generator) string {
	var b strings.Builder
	b.WriteString("# Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	fmt.Fprintf(&b, "[package]\nname = \"bedrock-protocol-%d\"\nversion = \"0.0.0\"\nedition = \"2021\"\npublish = false\n", m.Target.ProtocolVersion)
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

func emitRustPacket(info packetInfo) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	b.WriteString("#[allow(unused_imports)]\nuse crate::*;\n\n")
	fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq)]\npub struct %s {\n", info.name)
	for _, field := range info.fields {
		fmt.Fprintf(&b, "    pub %s: %s,\n", field.name, field.typ)
	}
	fmt.Fprintf(&b, "}\n\nimpl %s {\n    pub const ID: u32 = %d;\n}\n", info.name, info.packet.ID)
	return b.String()
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
			g.definitions[name] = definition{Name: name, Kind: manifest.KindUnion}
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
				}
				variants = append(variants, rustVariant{Name: variantName, Payload: payload, Fields: fields})
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
		underlying, err := primitiveRustType(node.Primitive.Code)
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
			g.usesNbt = true
			return "Nbt", true, nil
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
		fields = append(fields, rustField{Name: fieldName, Type: typ})
	}
	return fields, nil
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

func emitRustEnum(b *strings.Builder, item definition) {
	fmt.Fprintf(b, "#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash)]\npub enum %s {\n", item.Name)
	used := map[string]bool{}
	variantNames := make([]string, 0, len(item.Variants))
	for _, variant := range item.Variants {
		name := uniqueTypeVariant(enumVariantName(item.Name, variant.Name), used)
		variantNames = append(variantNames, name)
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
	fmt.Fprintf(b, "impl Default for %s {\n    fn default() -> Self {\n", item.Name)
	if len(variantNames) > 0 {
		fmt.Fprintf(b, "        Self::%s\n", variantNames[0])
	} else {
		fmt.Fprintf(b, "        Self::%s(0)\n", unknownName)
	}
	b.WriteString("    }\n}\n\n")
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
	if rustUnrawableKeywords[name] {
		return name + "_"
	}
	if rustKeywords[name] {
		return "r#" + name
	}
	return name
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
