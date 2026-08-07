// Package emitgo emits a small standalone Go view from the canonical v2
// manifest. It receives no source documents and has no profile mapping table.
package emitgo

import (
	"encoding/json"
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
	Union      []string
	Implements []string
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
		name := g.unique(packetTypeName(packet.Name))
		packetNames[packet.ID] = name
		for _, field := range packet.Fields {
			if _, err := g.goType(field.Encode, name+exportName(field.Name)); err != nil {
				return nil, fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
		}
	}
	files, err := g.emitFiles(packageName, packets, packetNames)
	if err != nil {
		return nil, err
	}
	files["wire.go"] = emitWire(packageName)
	return files, nil
}

func (g *generator) goType(node manifest.Node, hint string) (string, error) {
	if typ, matched, err := nativeGoType(node); matched || err != nil {
		return typ, err
	}
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
	case manifest.KindBitset:
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
		return "[]OrderedEntry[" + key + ", " + value + "]", nil
	case manifest.KindUnion:
		name := g.registerIdentity(node, hint+"Union")
		if _, exists := g.definitions[name]; !exists {
			g.definitions[name] = typeDefinition{Name: name, Kind: manifest.KindUnion}
			members := make([]string, 0, len(node.Variants))
			for _, variant := range node.Variants {
				member, err := g.registerUnionMember(name, variant)
				if err != nil {
					return "", err
				}
				members = append(members, member)
			}
			definition := g.definitions[name]
			definition.Union = members
			g.definitions[name] = definition
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

// nativeGoType maps canonical wire semantics to established Go ecosystem types.
func nativeGoType(node manifest.Node) (string, bool, error) {
	if node.Kind == manifest.KindPrimitive && node.Primitive != nil {
		switch node.Primitive.Code {
		case "uuid":
			return "uuid.UUID", true, nil
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
		return "mgl32.Vec2", true, nil
	case "Vec3":
		if !isPrimitiveStruct(node, "f32le", "f32le", "f32le") {
			return "", true, fmt.Errorf("native Vec3 mapping requires exactly three f32le fields")
		}
		return "mgl32.Vec3", true, nil
	case "mce::Color":
		if !isPrimitiveStruct(node, "i32le") {
			return "", true, fmt.Errorf("native colour mapping requires exactly one i32le field")
		}
		return "color.RGBA", true, nil
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

func (g *generator) registerUnionMember(union string, variant manifest.Variant) (string, error) {
	member, err := g.goType(variant.Encode, union+exportName(shortTypeName(variant.Name)))
	if err != nil {
		return "", err
	}
	if variant.Encode.Kind == manifest.KindVoid {
		member = g.unique(union + exportName(shortTypeName(variant.Name)))
		g.definitions[member] = typeDefinition{Name: member, Kind: manifest.KindStruct}
	}
	definition, ok := g.definitions[member]
	if !ok || definition.Kind != manifest.KindStruct {
		wrapper := g.unique(union + exportName(shortTypeName(variant.Name)))
		g.definitions[wrapper] = typeDefinition{Name: wrapper, Kind: manifest.KindStruct, Fields: []typedField{{Name: "Value", Type: member}}, Implements: []string{union}}
		return wrapper, nil
	}
	if !containsString(definition.Implements, union) {
		definition.Implements = append(definition.Implements, union)
	}
	g.definitions[member] = definition
	return member, nil
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
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
	inferred := inferredTypeName(node)
	if key == "" {
		identityHint := inferred
		if identityHint == "" {
			identityHint = hint
		} else {
			identityHint += "/" + unionIdentity(node)
		}
		key = fmt.Sprintf("%q/%q/%q", node.Kind, node.Semantic, identityHint)
	}
	if name, ok := g.identity[key]; ok {
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
	if g.usedNames[candidate] {
		switch node.Kind {
		case manifest.KindEnum:
			candidate += "Type"
		case manifest.KindUnion:
			hinted := strings.TrimSuffix(publicTypeName(hint), "Union")
			if hinted != "" && !g.usedNames[hinted] {
				candidate = hinted
			} else {
				candidate += "Data"
			}
		default:
			candidate += "Data"
		}
	}
	name := g.unique(candidate)
	g.identity[key] = name
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

func (g *generator) emitFiles(packageName string, packets []manifest.Packet, packetNames map[uint32]string) (map[string]string, error) {
	definitions := make([]typeDefinition, 0, len(g.definitions))
	for _, definition := range g.definitions {
		definitions = append(definitions, definition)
	}
	sort.Slice(definitions, func(i, j int) bool { return definitions[i].Name < definitions[j].Name })

	files := map[string]string{}
	typesSource, err := emitTypeDefinitions(packageName, definitions)
	if err != nil {
		return nil, err
	}
	files["types.go"] = typesSource
	enumsSource, err := emitEnumDefinitions(packageName, definitions)
	if err != nil {
		return nil, err
	}
	files["enums.go"] = enumsSource
	usedFiles := map[string]bool{}
	for _, packet := range packets {
		packetName := packetNames[packet.ID]
		base := snakeName(packetName) + ".go"
		name := uniqueFileName(base, packet.ID, usedFiles)
		source, err := g.emitPacket(packageName, packet, packetName)
		if err != nil {
			return nil, err
		}
		files[name] = source
	}
	return files, nil
}

func emitTypeDefinitions(packageName string, definitions []typeDefinition) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	writeGoImports(&b, goImportsForDefinitions(definitions))
	b.WriteString("// OrderedEntry preserves the source order and duplicate keys of a wire map.\n")
	b.WriteString("type OrderedEntry[K, V any] struct {\n\tKey K\n\tValue V\n}\n\n")
	for _, definition := range definitions {
		switch definition.Kind {
		case manifest.KindStruct:
			fmt.Fprintf(&b, "type %s struct {\n", definition.Name)
			for _, field := range definition.Fields {
				fmt.Fprintf(&b, "\t%s %s\n", field.Name, field.Type)
			}
			b.WriteString("}\n\n")
			for _, union := range definition.Implements {
				fmt.Fprintf(&b, "func (%s) is%s() {}\n\n", definition.Name, union)
			}
		case manifest.KindUnion:
			fmt.Fprintf(&b, "type %s interface {\n\tis%s()\n}\n\n", definition.Name, definition.Name)
		}
	}
	return formatGoSource(b.String())
}

func emitEnumDefinitions(packageName string, definitions []typeDefinition) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	for _, definition := range definitions {
		if definition.Kind != manifest.KindEnum {
			continue
		}
		fmt.Fprintf(&b, "type %s %s\n\nconst (\n", definition.Name, definition.Underlying)
		for _, variant := range definition.Variants {
			fmt.Fprintf(&b, "\t%s%s %s = %d\n", definition.Name, exportName(variant.Name), definition.Name, variant.Value)
		}
		b.WriteString(")\n\n")
	}
	return formatGoSource(b.String())
}

type packetField struct {
	field manifest.Field
	name  string
	typ   string
}

func (g *generator) emitPacket(packageName string, packet manifest.Packet, packetName string) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	used := map[string]bool{}
	fields := make([]packetField, 0, len(packet.Fields))
	for _, field := range packet.Fields {
		name := uniqueFieldName(exportName(field.Name), used)
		typ, err := g.goType(field.Encode, packetName+name)
		if err != nil {
			return "", err
		}
		fields = append(fields, packetField{field: field, name: name, typ: typ})
	}
	imports := goImportsForFields(fields)
	if len(packet.Fields) != 0 {
		imports = append(imports, "fmt")
	}
	writeGoImports(&b, imports)
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
	return formatGoSource(b.String())
}

func goImportsForDefinitions(definitions []typeDefinition) []string {
	var types []string
	for _, definition := range definitions {
		for _, field := range definition.Fields {
			types = append(types, field.Type)
		}
	}
	return goImportsForTypes(types)
}

func goImportsForFields(fields []packetField) []string {
	types := make([]string, 0, len(fields))
	for _, field := range fields {
		types = append(types, field.typ)
	}
	return goImportsForTypes(types)
}

func goImportsForTypes(types []string) []string {
	used := map[string]bool{}
	for _, typ := range types {
		if strings.Contains(typ, "color.") {
			used["image/color"] = true
		}
		if strings.Contains(typ, "mgl32.") {
			used["github.com/go-gl/mathgl/mgl32"] = true
		}
		if strings.Contains(typ, "uuid.") {
			used["github.com/google/uuid"] = true
		}
	}
	imports := make([]string, 0, len(used))
	for path := range used {
		imports = append(imports, path)
	}
	sort.Strings(imports)
	return imports
}

func writeGoImports(b *strings.Builder, imports []string) {
	if len(imports) == 0 {
		return
	}
	sort.Strings(imports)
	if len(imports) == 1 {
		fmt.Fprintf(b, "import %q\n\n", imports[0])
		return
	}
	var standard, external []string
	for _, path := range imports {
		first := strings.SplitN(path, "/", 2)[0]
		if strings.Contains(first, ".") {
			external = append(external, path)
		} else {
			standard = append(standard, path)
		}
	}
	b.WriteString("import (\n")
	for _, path := range standard {
		fmt.Fprintf(b, "\t%q\n", path)
	}
	if len(standard) != 0 && len(external) != 0 {
		b.WriteByte('\n')
	}
	for _, path := range external {
		fmt.Fprintf(b, "\t%q\n", path)
	}
	b.WriteString(")\n\n")
}

func formatGoSource(source string) (string, error) {
	formatted, err := format.Source([]byte(source))
	if err != nil {
		return source, fmt.Errorf("format generated Go: %w", err)
	}
	return string(formatted), nil
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

func packetTypeName(value string) string {
	name := exportName(value)
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
	if strings.HasSuffix(value, "PayloadUnion") {
		value = strings.TrimSuffix(value, "PayloadUnion") + "Value"
	}
	value = strings.TrimSuffix(value, "Union")
	value = strings.TrimSuffix(value, "Payload")
	return strings.ReplaceAll(exportName(value), "Molang", "MoLang")
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

func snakeName(value string) string {
	var b strings.Builder
	runes := []rune(value)
	for index, r := range runes {
		if unicode.IsUpper(r) && index > 0 {
			previousLower := unicode.IsLower(runes[index-1]) || unicode.IsDigit(runes[index-1])
			nextLower := index+1 < len(runes) && unicode.IsLower(runes[index+1])
			if previousLower || nextLower {
				b.WriteByte('_')
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

func uniqueFileName(base string, packetID uint32, used map[string]bool) string {
	if !used[base] {
		used[base] = true
		return base
	}
	name := strings.TrimSuffix(base, ".go") + fmt.Sprintf("_%d.go", packetID)
	used[name] = true
	return name
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
