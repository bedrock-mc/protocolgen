// Package emitgo emits a small standalone Go view from the canonical v2
// manifest. It receives no source documents and has no profile mapping table.
package emitgo

import (
	"fmt"
	"go/format"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"protocolgen/internal/manifest"
)

type typeDefinition struct {
	Name       string
	Kind       manifest.NodeKind
	Fields     []typedField
	EntryKey   string
	EntryValue string
	Underlying string
	Variants   []manifest.Variant
}

type typedField struct {
	Name string
	Type string
}

type generator struct {
	definitions map[string]typeDefinition
	identity    map[string]string
	usedNames   map[string]bool
}

func Generate(m manifest.Manifest, packageName string) (map[string]string, error) {
	if err := manifest.Validate(m); err != nil {
		return nil, err
	}
	if !validPackageName(packageName) {
		return nil, fmt.Errorf("invalid generated package name %q", packageName)
	}
	g := &generator{definitions: map[string]typeDefinition{}, identity: map[string]string{}, usedNames: map[string]bool{}}
	packets := append([]manifest.Packet(nil), m.Packets...)
	sort.Slice(packets, func(i, j int) bool { return packets[i].ID < packets[j].ID })
	packetNames := map[uint32]string{}
	for _, packet := range packets {
		name := g.unique(exportName(packet.Name))
		packetNames[packet.ID] = name
		for _, field := range packet.Fields {
			if _, err := g.goType(field.Encode, name+exportName(field.Name)); err != nil {
				return nil, fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
		}
	}
	packetsSource, err := g.emitPackets(packageName, packets, packetNames)
	if err != nil {
		return nil, err
	}
	wireSource := emitWire(packageName)
	return map[string]string{"wire.go": wireSource, "packets.go": packetsSource}, nil
}

func (g *generator) goType(node manifest.Node, hint string) (string, error) {
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return "", fmt.Errorf("primitive has no shape")
		}
		return primitiveGoType(node.Primitive.Code)
	case manifest.KindString:
		return "string", nil
	case manifest.KindBytes:
		return "[]byte", nil
	case manifest.KindArray:
		if node.Element == nil {
			return "", fmt.Errorf("array has no element")
		}
		element, err := g.goType(*node.Element, hint+"Item")
		if err != nil {
			return "", err
		}
		return "[]" + element, nil
	case manifest.KindFixedArray:
		if node.Element == nil {
			return "", fmt.Errorf("fixed array has no element")
		}
		element, err := g.goType(*node.Element, hint+"Item")
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("[%d]%s", node.Length, element), nil
	case manifest.KindSequence:
		if len(node.Elements) == 0 {
			return "[]any", nil
		}
		return "[]any", nil
	case manifest.KindOptional:
		if node.Value == nil {
			return "", fmt.Errorf("optional has no value")
		}
		value, err := g.goType(*node.Value, hint+"Value")
		if err != nil {
			return "", err
		}
		if strings.HasPrefix(value, "[]") || strings.HasPrefix(value, "map[") || strings.HasPrefix(value, "*") {
			return "*" + value, nil
		}
		return "*" + value, nil
	case manifest.KindStruct:
		return g.registerStruct(node, hint)
	case manifest.KindMap:
		if node.Key == nil || node.Value == nil {
			return "", fmt.Errorf("map has no key/value")
		}
		key, err := g.goType(*node.Key, hint+"Key")
		if err != nil {
			return "", err
		}
		value, err := g.goType(*node.Value, hint+"Value")
		if err != nil {
			return "", err
		}
		entryName := g.registerIdentity(node, hint+"Entry")
		if _, exists := g.definitions[entryName]; !exists {
			g.definitions[entryName] = typeDefinition{Name: entryName, Kind: manifest.KindStruct, Fields: []typedField{{Name: "Key", Type: key}, {Name: "Value", Type: value}}}
		}
		return "[]" + entryName, nil
	case manifest.KindUnion:
		name := g.registerIdentity(node, hint+"Union")
		if _, exists := g.definitions[name]; !exists {
			g.definitions[name] = typeDefinition{Name: name, Kind: manifest.KindUnion}
		}
		return name, nil
	case manifest.KindEnum:
		if node.Primitive == nil {
			return "", fmt.Errorf("enum has no underlying primitive")
		}
		underlying, err := primitiveGoType(node.Primitive.Code)
		if err != nil {
			return "", err
		}
		name := g.registerIdentity(node, hint+"Enum")
		if _, exists := g.definitions[name]; !exists {
			g.definitions[name] = typeDefinition{Name: name, Kind: manifest.KindEnum, Underlying: underlying, Variants: append([]manifest.Variant(nil), node.Variants...)}
		}
		return name, nil
	case manifest.KindReserved, manifest.KindIgnored:
		if node.Element == nil {
			return "", fmt.Errorf("compatibility node has no preserved element")
		}
		return g.goType(*node.Element, hint)
	case manifest.KindRecursive:
		// Recursive nodes retain their identity in Shape; a generic value keeps the
		// generated API compilable until a profile chooses a recursive view.
		return "any", nil
	case manifest.KindVoid:
		return "struct{}", nil
	case manifest.KindOpaque, manifest.KindUnresolved:
		return "", fmt.Errorf("reachable %s node cannot be emitted: %s", node.Kind, node.Reason)
	default:
		return "", fmt.Errorf("unsupported node kind %q", node.Kind)
	}
}

func (g *generator) registerStruct(node manifest.Node, hint string) (string, error) {
	name := g.registerIdentity(node, hint+"Struct")
	if _, exists := g.definitions[name]; exists {
		return name, nil
	}
	g.definitions[name] = typeDefinition{Name: name, Kind: manifest.KindStruct}
	used := map[string]bool{}
	var fields []typedField
	for _, field := range node.Fields {
		fieldName := uniqueFieldName(exportName(field.Name), used)
		fieldType, err := g.goType(field.Encode, name+fieldName)
		if err != nil {
			return "", err
		}
		fields = append(fields, typedField{Name: fieldName, Type: fieldType})
	}
	definition := g.definitions[name]
	definition.Fields = fields
	g.definitions[name] = definition
	return name, nil
}

func (g *generator) registerIdentity(node manifest.Node, hint string) string {
	key := node.TypeID
	if key == "" {
		key = fmt.Sprintf("%q/%q/%q", node.Kind, node.Semantic, hint)
	}
	if name, ok := g.identity[key]; ok {
		return name
	}
	name := g.unique(exportName(hint))
	g.identity[key] = name
	return name
}

func (g *generator) unique(base string) string {
	if base == "" {
		base = "GeneratedType"
	}
	if !g.usedNames[base] {
		g.usedNames[base] = true
		return base
	}
	for index := 2; ; index++ {
		candidate := fmt.Sprintf("%s%d", base, index)
		if !g.usedNames[candidate] {
			g.usedNames[candidate] = true
			return candidate
		}
	}
}

func (g *generator) emitPackets(packageName string, packets []manifest.Packet, packetNames map[uint32]string) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	b.WriteString("import \"fmt\"\n\n")
	definitions := make([]typeDefinition, 0, len(g.definitions))
	for _, definition := range g.definitions {
		definitions = append(definitions, definition)
	}
	sort.Slice(definitions, func(i, j int) bool { return definitions[i].Name < definitions[j].Name })
	for _, definition := range definitions {
		switch definition.Kind {
		case manifest.KindStruct:
			fmt.Fprintf(&b, "type %s struct {\n", definition.Name)
			for _, field := range definition.Fields {
				fmt.Fprintf(&b, "\t%s %s\n", field.Name, field.Type)
			}
			b.WriteString("}\n\n")
		case manifest.KindUnion:
			fmt.Fprintf(&b, "type %s struct {\n\tTag int64\n\tValue any\n}\n\n", definition.Name)
		case manifest.KindEnum:
			fmt.Fprintf(&b, "type %s %s\n\nconst (\n", definition.Name, definition.Underlying)
			for _, variant := range definition.Variants {
				fmt.Fprintf(&b, "\t%s%s %s = %d\n", definition.Name, exportName(variant.Name), definition.Name, variant.Value)
			}
			b.WriteString(")\n\n")
		}
	}
	for _, packet := range packets {
		packetName := packetNames[packet.ID]
		used := map[string]bool{}
		type packetField struct {
			field manifest.Field
			name  string
			typ   string
		}
		fields := make([]packetField, 0, len(packet.Fields))
		for _, field := range packet.Fields {
			name := uniqueFieldName(exportName(field.Name), used)
			typ, err := g.goType(field.Encode, packetName+name)
			if err != nil {
				return "", err
			}
			fields = append(fields, packetField{field: field, name: name, typ: typ})
		}
		fmt.Fprintf(&b, "type %s struct {\n", packetName)
		for _, field := range fields {
			fmt.Fprintf(&b, "\t%s %s\n", field.name, field.typ)
		}
		b.WriteString("}\n\n")
		fmt.Fprintf(&b, "func (p *%s) Encode(w Encoder) error {\n", packetName)
		for _, field := range fields {
			shape, err := shapeExpr(field.field.Encode)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(&b, "\tif err := w.Write(%s, %s, p.%s); err != nil { return err }\n", strconv.Quote(packet.Name+"."+field.field.Name), shape, field.name)
		}
		b.WriteString("\treturn nil\n}\n\n")
		fmt.Fprintf(&b, "func Decode%s(r Decoder) (%s, error) {\n\tvar p %s\n", packetName, packetName, packetName)
		for _, field := range fields {
			decodeNode := field.field.Encode
			if field.field.Decode != nil {
				decodeNode = *field.field.Decode
			}
			shape, err := shapeExpr(decodeNode)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(&b, "\t{\n\t\traw, err := r.Read(%s, %s)\n\t\tif err != nil { return p, err }\n\t\tvalue, ok := raw.(%s)\n\t\tif !ok { return p, fmt.Errorf(\"field %s has unexpected decoded type %%T\", raw) }\n\t\tp.%s = value\n\t}\n", strconv.Quote(packet.Name+"."+field.field.Name), shape, field.typ, packet.Name+"."+field.field.Name, field.name)
		}
		b.WriteString("\treturn p, nil\n}\n\n")
	}
	formatted, err := format.Source([]byte(b.String()))
	if err != nil {
		return b.String(), fmt.Errorf("format generated Go: %w", err)
	}
	return string(formatted), nil
}

func shapeExpr(node manifest.Node) (string, error) {
	parts := []string{fmt.Sprintf("Kind: %s", strconv.Quote(string(node.Kind)))}
	if node.Semantic != "" {
		parts = append(parts, "Semantic: "+strconv.Quote(node.Semantic))
	}
	if node.TypeID != "" {
		parts = append(parts, "TypeID: "+strconv.Quote(node.TypeID))
	}
	if node.Primitive != nil {
		parts = append(parts, "PrimitiveCode: "+strconv.Quote(node.Primitive.Code))
	}
	if node.Encoding != "" {
		parts = append(parts, "Encoding: "+strconv.Quote(node.Encoding))
	}
	if node.Representation != "" {
		parts = append(parts, "Representation: "+strconv.Quote(node.Representation))
	}
	if node.Length != 0 {
		parts = append(parts, fmt.Sprintf("Length: %d", node.Length))
	}
	if node.Prefix != nil {
		prefix, err := shapeExpr(*node.Prefix)
		if err != nil {
			return "", err
		}
		parts = append(parts, "Prefix: &"+prefix)
	}
	if node.Element != nil {
		element, err := shapeExpr(*node.Element)
		if err != nil {
			return "", err
		}
		parts = append(parts, "Element: &"+element)
	}
	if node.Value != nil {
		value, err := shapeExpr(*node.Value)
		if err != nil {
			return "", err
		}
		parts = append(parts, "Value: &"+value)
	}
	if len(node.Elements) > 0 {
		elements := make([]string, 0, len(node.Elements))
		for _, element := range node.Elements {
			shape, err := shapeExpr(element)
			if err != nil {
				return "", err
			}
			elements = append(elements, shape)
		}
		parts = append(parts, "Elements: []Shape{"+strings.Join(elements, ", ")+"}")
	}
	if node.Key != nil {
		key, err := shapeExpr(*node.Key)
		if err != nil {
			return "", err
		}
		parts = append(parts, "Key: &"+key)
	}
	if node.Control != nil {
		control, err := shapeExpr(*node.Control)
		if err != nil {
			return "", err
		}
		parts = append(parts, "Control: &"+control)
	}
	if node.Default != nil {
		defaultShape, err := shapeExpr(*node.Default)
		if err != nil {
			return "", err
		}
		parts = append(parts, "Default: &"+defaultShape)
	}
	if node.Target != "" {
		parts = append(parts, "Target: "+strconv.Quote(node.Target))
	}
	if node.CompareTo != "" {
		parts = append(parts, "CompareTo: "+strconv.Quote(node.CompareTo))
	}
	if len(node.Fields) > 0 {
		fields := make([]string, 0, len(node.Fields))
		for _, field := range node.Fields {
			fieldShape, err := shapeExpr(field.Encode)
			if err != nil {
				return "", err
			}
			fields = append(fields, fmt.Sprintf("{Ordinal: %d, Name: %s, Shape: %s}", field.Ordinal, strconv.Quote(field.Name), fieldShape))
		}
		parts = append(parts, "Fields: []ShapeField{"+strings.Join(fields, ", ")+"}")
	}
	if len(node.Variants) > 0 {
		variants := make([]string, 0, len(node.Variants))
		for _, variant := range node.Variants {
			variantShape, err := shapeExpr(variant.Encode)
			if err != nil {
				return "", err
			}
			variants = append(variants, fmt.Sprintf("{Value: %d, Name: %s, Shape: %s}", variant.Value, strconv.Quote(variant.Name), variantShape))
		}
		parts = append(parts, "Variants: []ShapeVariant{"+strings.Join(variants, ", ")+"}")
	}
	if len(node.Cases) > 0 {
		cases := make([]string, 0, len(node.Cases))
		for _, oneCase := range node.Cases {
			shapes := make([]string, 0, len(oneCase.Encode))
			for _, child := range oneCase.Encode {
				childShape, err := shapeExpr(child)
				if err != nil {
					return "", err
				}
				shapes = append(shapes, childShape)
			}
			cases = append(cases, fmt.Sprintf("{Value: %s, Shapes: []Shape{%s}}", strconv.Quote(oneCase.Value), strings.Join(shapes, ", ")))
		}
		parts = append(parts, "Cases: []ShapeCase{"+strings.Join(cases, ", ")+"}")
	}
	return "Shape{" + strings.Join(parts, ", ") + "}", nil
}

func emitWire(packageName string) string {
	return fmt.Sprintf(`// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package %s

// Shape is the frozen wire vocabulary rendered from the canonical manifest.
// A profile may implement Encoder/Decoder, but it cannot alter Shape.
type Shape struct {
	Kind string
	Semantic string
	TypeID string
	PrimitiveCode string
	Encoding string
	Representation string
	Prefix *Shape
	Element *Shape
	Length uint64
	Value *Shape
	Elements []Shape
	Key *Shape
	Fields []ShapeField
	Variants []ShapeVariant
	Control *Shape
	CompareTo string
	Cases []ShapeCase
	Default *Shape
	Target string
}

type ShapeField struct { Ordinal int; Name string; Shape Shape }
type ShapeVariant struct { Value int64; Name string; Shape Shape }
type ShapeCase struct { Value string; Shapes []Shape }

type Encoder interface { Write(path string, shape Shape, value any) error }
type Decoder interface { Read(path string, shape Shape) (any, error) }
`, packageName)
}

func primitiveGoType(code string) (string, error) {
	switch code {
	case "bool":
		return "bool", nil
	case "i8":
		return "int8", nil
	case "u8":
		return "uint8", nil
	case "i16le", "i16be":
		return "int16", nil
	case "u16le", "u16be":
		return "uint16", nil
	case "i32le", "i32be", "var_i32", "zigzag_i32":
		return "int32", nil
	case "u32le", "u32be", "var_u32":
		return "uint32", nil
	case "i64le", "i64be", "var_i64", "zigzag_i64":
		return "int64", nil
	case "u64le", "u64be", "var_u64":
		return "uint64", nil
	case "f32le", "f32be":
		return "float32", nil
	case "f64le", "f64be":
		return "float64", nil
	case "uuid":
		return "[16]byte", nil
	case "nbt_le":
		return "[]byte", nil
	default:
		return "", fmt.Errorf("unsupported primitive code %q", code)
	}
}

func validPackageName(name string) bool {
	if name == "" {
		return false
	}
	for index, r := range name {
		if !(r == '_' || unicode.IsLetter(r) || (index > 0 && unicode.IsDigit(r))) {
			return false
		}
	}
	return true
}

func exportName(value string) string {
	var b strings.Builder
	upperNext := true
	for _, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if upperNext {
				r = unicode.ToUpper(r)
				upperNext = false
			}
			b.WriteRune(r)
		} else {
			upperNext = true
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

func uniqueFieldName(base string, used map[string]bool) string {
	if base == "" {
		base = "Field"
	}
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
