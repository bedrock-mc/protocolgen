// Package emitrust emits a compact Rust consumer surface from the canonical
// v2 manifest. The emitted shape snapshots are serialized from manifest nodes;
// no Rust source is reverse-lowered.
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
}

type rustField struct {
	Name string
	Type string
}

type generator struct {
	definitions map[string]definition
	identities  map[string]string
	used        map[string]bool
}

func Generate(m manifest.Manifest) (string, error) {
	if err := manifest.Validate(m); err != nil {
		return "", err
	}
	g := &generator{definitions: map[string]definition{}, identities: map[string]string{}, used: map[string]bool{}}
	packets := append([]manifest.Packet(nil), m.Packets...)
	sort.Slice(packets, func(i, j int) bool { return packets[i].ID < packets[j].ID })
	type packetInfo struct {
		packet manifest.Packet
		name   string
		fields []rustFieldInfo
	}
	infos := make([]packetInfo, 0, len(packets))
	for _, packet := range packets {
		name := g.unique(typeName(packet.Name))
		used := map[string]bool{}
		fields := make([]rustFieldInfo, 0, len(packet.Fields))
		for _, field := range packet.Fields {
			fieldName := uniqueField(fieldName(field.Name), used)
			typ, err := g.rustType(field.Encode, name+typeName(field.Name))
			if err != nil {
				return "", fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
			shape, err := json.Marshal(field.Encode)
			if err != nil {
				return "", err
			}
			decodeShape := shape
			if field.Decode != nil {
				decodeShape, err = json.Marshal(*field.Decode)
				if err != nil {
					return "", err
				}
			}
			fields = append(fields, rustFieldInfo{field: field, name: fieldName, typ: typ, shape: string(shape), decodeShape: string(decodeShape)})
		}
		infos = append(infos, packetInfo{packet: packet, name: name, fields: fields})
	}

	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\n")
	b.WriteString("#![allow(dead_code)]\n\n")
	b.WriteString("pub trait WireEncoder {\n    fn field(&mut self, path: &'static str, shape: &'static str);\n}\n")
	b.WriteString("pub trait WireDecoder {\n    fn field(&mut self, path: &'static str, shape: &'static str);\n}\n\n")
	definitions := make([]definition, 0, len(g.definitions))
	for _, item := range g.definitions {
		definitions = append(definitions, item)
	}
	sort.Slice(definitions, func(i, j int) bool { return definitions[i].Name < definitions[j].Name })
	for _, item := range definitions {
		switch item.Kind {
		case manifest.KindStruct:
			fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq)]\npub struct %s {\n", item.Name)
			for _, field := range item.Fields {
				fmt.Fprintf(&b, "    pub %s: %s,\n", field.Name, field.Type)
			}
			b.WriteString("}\n\n")
		case manifest.KindUnion:
			fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq)]\npub struct %s {\n    pub tag: i64,\n    pub value: Vec<u8>,\n}\n\n", item.Name)
		case manifest.KindEnum:
			fmt.Fprintf(&b, "#[derive(Clone, Copy, Debug, PartialEq, Eq)]\n#[repr(transparent)]\npub struct %s(pub %s);\n", item.Name, item.Underlying)
			for _, variant := range item.Variants {
				fmt.Fprintf(&b, "pub const %s_%s: %s = %s(%d);\n", strings.ToUpper(item.Name), strings.ToUpper(typeName(variant.Name)), item.Name, item.Name, variant.Value)
			}
			b.WriteString("\n")
		}
	}
	for _, info := range infos {
		fmt.Fprintf(&b, "#[derive(Clone, Debug, PartialEq)]\npub struct %s {\n", info.name)
		for _, field := range info.fields {
			fmt.Fprintf(&b, "    pub %s: %s,\n", field.name, field.typ)
		}
		b.WriteString("}\n\n")
		for _, field := range info.fields {
			constName := shapeConst(info.name, field.name)
			fmt.Fprintf(&b, "pub const %s: &str = %s;\n", constName, rustRawString(field.shape))
			if field.field.Decode != nil {
				fmt.Fprintf(&b, "pub const %s_DECODE: &str = %s;\n", constName, rustRawString(field.decodeShape))
			}
		}
		b.WriteString("\n")
		fmt.Fprintf(&b, "impl %s {\n    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {\n", info.name)
		for _, field := range info.fields {
			fmt.Fprintf(&b, "        encoder.field(%s, %s);\n", rustString(info.packet.Name+"."+field.field.Name), shapeConst(info.name, field.name))
		}
		b.WriteString("    }\n    pub fn decode<D: WireDecoder>(decoder: &mut D) {\n")
		for _, field := range info.fields {
			decodeConst := shapeConst(info.name, field.name)
			if field.field.Decode != nil {
				decodeConst += "_DECODE"
			}
			fmt.Fprintf(&b, "        decoder.field(%s, %s);\n", rustString(info.packet.Name+"."+field.field.Name), decodeConst)
		}
		b.WriteString("    }\n}\n\n")
	}
	return strings.TrimSpace(b.String()) + "\n", nil
}

type rustFieldInfo struct {
	field       manifest.Field
	name        string
	typ         string
	shape       string
	decodeShape string
}

func (g *generator) rustType(node manifest.Node, hint string) (string, error) {
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return "", fmt.Errorf("primitive has no shape")
		}
		return primitiveRustType(node.Primitive.Code)
	case manifest.KindString:
		return "String", nil
	case manifest.KindBytes:
		return "Vec<u8>", nil
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
		return "Vec<Vec<u8>>", nil
	case manifest.KindOptional:
		if node.Value == nil {
			return "", fmt.Errorf("optional has no value")
		}
		value, err := g.rustType(*node.Value, hint+"Value")
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
		entry := g.registerIdentity(node, hint+"Entry")
		if _, ok := g.definitions[entry]; !ok {
			g.definitions[entry] = definition{Name: entry, Kind: manifest.KindStruct, Fields: []rustField{{Name: "key", Type: key}, {Name: "value", Type: value}}}
		}
		return "Vec<" + entry + ">", nil
	case manifest.KindUnion:
		name := g.registerIdentity(node, hint+"Union")
		if _, ok := g.definitions[name]; !ok {
			g.definitions[name] = definition{Name: name, Kind: manifest.KindUnion}
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
		if node.Element == nil {
			return "", fmt.Errorf("compatibility node has no element")
		}
		return g.rustType(*node.Element, hint)
	case manifest.KindRecursive:
		return "Vec<u8>", nil
	case manifest.KindVoid:
		return "()", nil
	case manifest.KindOpaque, manifest.KindUnresolved:
		return "", fmt.Errorf("reachable %s node cannot be emitted: %s", node.Kind, node.Reason)
	default:
		return "", fmt.Errorf("unsupported node kind %q", node.Kind)
	}
}

func (g *generator) registerStruct(node manifest.Node, hint string) (string, error) {
	name := g.registerIdentity(node, hint+"Struct")
	if _, ok := g.definitions[name]; ok {
		return name, nil
	}
	g.definitions[name] = definition{Name: name, Kind: manifest.KindStruct}
	used := map[string]bool{}
	fields := make([]rustField, 0, len(node.Fields))
	for _, field := range node.Fields {
		name := uniqueField(fieldName(field.Name), used)
		typ, err := g.rustType(field.Encode, name+typeName(field.Name))
		if err != nil {
			return "", err
		}
		fields = append(fields, rustField{Name: name, Type: typ})
	}
	definition := g.definitions[name]
	definition.Fields = fields
	g.definitions[name] = definition
	return name, nil
}

func (g *generator) registerIdentity(node manifest.Node, hint string) string {
	key := node.TypeID
	if key == "" {
		key = fmt.Sprintf("%s/%s/%s", node.Kind, node.Semantic, hint)
	}
	if name, ok := g.identities[key]; ok {
		return name
	}
	name := g.unique(typeName(hint))
	g.identities[key] = name
	return name
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
	case "uuid":
		return "[u8; 16]", nil
	case "nbt_le":
		return "Vec<u8>", nil
	default:
		return "", fmt.Errorf("unsupported primitive code %q", code)
	}
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
	name := typeName(value)
	if name == "" {
		return "field"
	}
	runes := []rune(name)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
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

func shapeConst(packet, field string) string {
	return strings.ToUpper(typeName(packet)) + "_" + strings.ToUpper(fieldName(field)) + "_SHAPE"
}

func rustString(value string) string {
	return "\"" + strings.NewReplacer("\\", "\\\\", "\"", "\\\"", "\n", "\\n", "\r", "\\r").Replace(value) + "\""
}

func rustRawString(value string) string {
	delimiter := "#"
	for strings.Contains(value, delimiter+"\"") {
		delimiter += "#"
	}
	return "r" + delimiter + "\"" + value + "\"" + delimiter
}
