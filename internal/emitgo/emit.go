// Package emitgo emits a small standalone Go view from the canonical v2
// manifest. It receives no source documents and has no profile mapping table.
package emitgo

import (
	"encoding/json"
	"fmt"
	"go/format"
	"sort"
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
	Union      []goUnionMember
	Implements []string
	BitLength  uint64
}

type typedField struct {
	Name string
	Type string
	Node manifest.Node
}

type goUnionMember struct {
	Name  string
	Value int64
	Node  manifest.Node
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
			if err := ensureCodecSymmetric(field); err != nil {
				return nil, fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
			if _, err := g.goType(field.Encode, name+exportName(field.Name)); err != nil {
				return nil, fmt.Errorf("packet %s field %s: %w", packet.Name, field.Name, err)
			}
		}
	}
	files, err := g.emitFiles(packageName, packets, packetNames)
	if err != nil {
		return nil, err
	}
	return files, nil
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
		name := fmt.Sprintf("Bitset%d", node.Length)
		if _, exists := g.definitions[name]; !exists {
			g.definitions[name] = typeDefinition{Name: name, Kind: manifest.KindBitset, BitLength: node.Length}
		}
		return name, nil
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
		return "", fmt.Errorf("sequence nodes require a target-specific representation")
	case manifest.KindOptional:
		if node.Value == nil {
			return "", fmt.Errorf("optional has no value")
		}
		valueNode := *node.Value
		// Cereal double optionals carry an always-present outer marker. Keep both
		// nodes in the manifest for codec generation, but expose only the
		// meaningful inner presence state, matching gophertunnel's Optional[T].
		if valueNode.Kind == manifest.KindOptional {
			if valueNode.Value == nil {
				return "", fmt.Errorf("nested optional has no value")
			}
			valueNode = *valueNode.Value
		}
		value, err := g.goType(valueNode, hint+"Value")
		if err != nil {
			return "", err
		}
		return "Optional[" + value + "]", nil
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
			if node.Control == nil || node.Control.Primitive == nil {
				return "", fmt.Errorf("union has no primitive discriminator")
			}
			g.definitions[name] = typeDefinition{Name: name, Kind: manifest.KindUnion, Underlying: node.Control.Primitive.Code}
			members := make([]goUnionMember, 0, len(node.Variants))
			usedMembers := map[string]bool{}
			for _, variant := range node.Variants {
				member, err := g.registerUnionMember(name, variant)
				if err != nil {
					return "", err
				}
				if usedMembers[member] {
					wrapper := g.unique(name + exportName(shortTypeName(variant.Name)))
					g.definitions[wrapper] = typeDefinition{
						Name:       wrapper,
						Kind:       manifest.KindStruct,
						Fields:     []typedField{{Name: "Value", Type: member, Node: variant.Encode}},
						Implements: []string{name},
					}
					member = wrapper
				}
				usedMembers[member] = true
				members = append(members, goUnionMember{Name: member, Value: variant.Value, Node: variant.Encode})
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
		return "", fmt.Errorf("%s nodes require explicit write/discard codec semantics", node.Kind)
	case manifest.KindRecursive:
		name, ok := g.identity[node.Target]
		if !ok {
			return "", fmt.Errorf("recursive target %q is not a registered named type", node.Target)
		}
		return name, nil
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
		g.definitions[wrapper] = typeDefinition{Name: wrapper, Kind: manifest.KindStruct, Fields: []typedField{{Name: "Value", Type: member, Node: variant.Encode}}, Implements: []string{union}}
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
		fields = append(fields, typedField{Name: fieldName, Type: fieldType, Node: field.Encode})
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
	files["codec.go"] = emitCodecRuntime(packageName)
	marshalSource, err := g.emitSharedMarshalers(packageName, definitions)
	if err != nil {
		return nil, err
	}
	files["marshal.go"] = marshalSource
	enumsSource, err := emitEnumDefinitions(packageName, definitions)
	if err != nil {
		return nil, err
	}
	files["enums.go"] = enumsSource
	files["ids.go"] = emitPacketIDs(packageName, packets, packetNames)
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
	b.WriteString("// Optional holds a value that may be absent from the wire.\n")
	b.WriteString("type Optional[T any] struct {\n\tset bool\n\tval T\n}\n\n")
	b.WriteString("// Option creates a present Optional containing value.\n")
	b.WriteString("func Option[T any](value T) Optional[T] {\n\treturn Optional[T]{set: true, val: value}\n}\n\n")
	b.WriteString("// Value returns the optional value and whether it is present.\n")
	b.WriteString("func (o Optional[T]) Value() (T, bool) {\n\treturn o.val, o.set\n}\n\n")
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
		case manifest.KindBitset:
			fmt.Fprintf(&b, "// %s stores the %d-bit value used by the wire bitset encoding.\n", definition.Name, definition.BitLength)
			fmt.Fprintf(&b, "type %s [%d]uint64\n\n", definition.Name, (definition.BitLength+63)/64)
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
	name string
	typ  string
	node manifest.Node
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
		fields = append(fields, packetField{name: name, typ: typ, node: field.Encode})
	}
	imports := goImportsForFields(fields)
	writeGoImports(&b, imports)
	fmt.Fprintf(&b, "type %s struct {\n", packetName)
	for _, field := range fields {
		fmt.Fprintf(&b, "\t%s %s\n", field.name, field.typ)
	}
	b.WriteString("}\n\n")
	fmt.Fprintf(&b, "// Marshal reads or writes %s using its canonical wire layout.\n", packetName)
	fmt.Fprintf(&b, "func (x *%s) Marshal(io IO) {\n", packetName)
	emitter := marshalEmitter{g: g}
	for _, field := range fields {
		if err := emitter.node(&b, field.node, "x."+field.name, packetName+field.name, "\t"); err != nil {
			return "", fmt.Errorf("packet %s field %s marshal: %w", packet.Name, field.name, err)
		}
	}
	b.WriteString("}\n")
	return formatGoSource(b.String())
}

func emitPacketIDs(packageName string, packets []manifest.Packet, packetNames map[uint32]string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\nconst (\n", packageName)
	for _, packet := range packets {
		fmt.Fprintf(&b, "\tID%s uint32 = %d\n", packetNames[packet.ID], packet.ID)
	}
	b.WriteString(")\n")
	return b.String()
}

func emitCodecRuntime(packageName string) string {
	return fmt.Sprintf(`// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package %s

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// IO is the minimal symmetric wire interface used by generated Marshal methods.
// Reading reports whether calls populate values. InvalidValue must stop the
// current codec operation, typically by panicking or recording a terminal error.
// String and Bytes use a varuint32 byte-length prefix. UUID uses Bedrock's
// little-endian 64-bit halves, NBT consumes exactly one little-endian tag, and
// Bitset uses seven payload bits per continuation byte.
type IO interface {
	Reading() bool
	InvalidValue(value any, context string)

	Bool(*bool)
	Int8(*int8)
	Uint8(*uint8)
	Int16(*int16)
	Uint16(*uint16)
	BEInt16(*int16)
	BEUint16(*uint16)
	Int32(*int32)
	Uint32(*uint32)
	BEInt32(*int32)
	BEUint32(*uint32)
	Int64(*int64)
	Uint64(*uint64)
	BEInt64(*int64)
	BEUint64(*uint64)
	Float32(*float32)
	Float64(*float64)
	BEFloat32(*float32)
	BEFloat64(*float64)
	Varint32(*int32)
	Varuint32(*uint32)
	Varint64(*int64)
	Varuint64(*uint64)
	SignedVarint32(*int32)
	SignedVarint64(*int64)

	String(*string)
	Bytes(*[]byte)
	NBT(*[]byte)
	UUID(*uuid.UUID)
	Vec2(*mgl32.Vec2)
	Vec3(*mgl32.Vec3)
	RGBA(*color.RGBA)
	Bitset(words []uint64, bits uint64)
}
`, packageName)
}

func (g *generator) emitSharedMarshalers(packageName string, definitions []typeDefinition) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	writeGoImports(&b, goImportsForDefinitions(definitions))
	emitter := marshalEmitter{g: g}
	for _, definition := range definitions {
		switch definition.Kind {
		case manifest.KindStruct:
			fmt.Fprintf(&b, "// Marshal reads or writes %s using its canonical wire layout.\n", definition.Name)
			fmt.Fprintf(&b, "func (x *%s) Marshal(io IO) {\n", definition.Name)
			for _, field := range definition.Fields {
				if err := emitter.node(&b, field.Node, "x."+field.Name, definition.Name+field.Name, "\t"); err != nil {
					return "", fmt.Errorf("type %s field %s marshal: %w", definition.Name, field.Name, err)
				}
			}
			b.WriteString("}\n\n")
		case manifest.KindUnion:
			if err := emitter.union(&b, definition); err != nil {
				return "", err
			}
		}
	}
	return formatGoSource(b.String())
}

type marshalEmitter struct {
	g       *generator
	counter int
}

func (e *marshalEmitter) temporary(prefix string) string {
	e.counter++
	return fmt.Sprintf("%s%d", prefix, e.counter)
}

func (e *marshalEmitter) node(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	if native, matched, err := nativeGoType(node); matched || err != nil {
		if err != nil {
			return err
		}
		switch native {
		case "uuid.UUID":
			fmt.Fprintf(b, "%sio.UUID(&%s)\n", indent, expression)
		case "mgl32.Vec2":
			fmt.Fprintf(b, "%sio.Vec2(&%s)\n", indent, expression)
		case "mgl32.Vec3":
			fmt.Fprintf(b, "%sio.Vec3(&%s)\n", indent, expression)
		case "color.RGBA":
			fmt.Fprintf(b, "%sio.RGBA(&%s)\n", indent, expression)
		default:
			return fmt.Errorf("native type %s has no codec operation", native)
		}
		return nil
	}
	switch node.Kind {
	case manifest.KindVoid:
		return nil
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return fmt.Errorf("primitive has no shape")
		}
		if node.Primitive.Code == "nbt_le" {
			fmt.Fprintf(b, "%sio.NBT(&%s)\n", indent, expression)
			return nil
		}
		method, err := primitiveIOMethod(node.Primitive.Code)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%sio.%s(&%s)\n", indent, method, expression)
		return nil
	case manifest.KindString:
		if !varuint32Prefix(node) {
			return fmt.Errorf("string has unsupported length prefix")
		}
		fmt.Fprintf(b, "%sio.String(&%s)\n", indent, expression)
		return nil
	case manifest.KindBytes:
		if !varuint32Prefix(node) {
			return fmt.Errorf("bytes have unsupported length prefix")
		}
		fmt.Fprintf(b, "%sio.Bytes(&%s)\n", indent, expression)
		return nil
	case manifest.KindBitset:
		fmt.Fprintf(b, "%sio.Bitset(%s[:], %d)\n", indent, expression, node.Length)
		return nil
	case manifest.KindStruct:
		fmt.Fprintf(b, "%s%s.Marshal(io)\n", indent, expression)
		return nil
	case manifest.KindRecursive:
		fmt.Fprintf(b, "%smarshal%s(io, &%s)\n", indent, e.g.identity[node.Target], expression)
		return nil
	case manifest.KindEnum:
		return e.enum(b, node, expression, hint, indent)
	case manifest.KindOptional:
		if node.Value == nil {
			return fmt.Errorf("optional has no value")
		}
		value := *node.Value
		if value.Kind == manifest.KindOptional {
			if value.Value == nil {
				return fmt.Errorf("nested optional has no value")
			}
			outer := e.temporary("outer")
			fmt.Fprintf(b, "%s%s := true\n%sio.Bool(&%s)\n", indent, outer, indent, outer)
			fmt.Fprintf(b, "%sif %s {\n", indent, outer)
			if err := e.optional(b, *value.Value, expression, hint+"Value", indent+"\t"); err != nil {
				return err
			}
			fmt.Fprintf(b, "%s} else {\n%s\t%s = Optional[%s]{}\n%s}\n", indent, indent, expression, mustGoType(e.g, *value.Value, hint+"Value"), indent)
			return nil
		}
		return e.optional(b, value, expression, hint+"Value", indent)
	case manifest.KindArray:
		if node.Element == nil || node.Prefix == nil {
			return fmt.Errorf("array has no element or prefix")
		}
		return e.collection(b, *node.Prefix, *node.Element, expression, hint+"Item", indent)
	case manifest.KindFixedArray:
		if node.Element == nil {
			return fmt.Errorf("fixed array has no element")
		}
		index := e.temporary("index")
		fmt.Fprintf(b, "%sfor %s := range %s {\n", indent, index, expression)
		if err := e.node(b, *node.Element, expression+"["+index+"]", hint+"Item", indent+"\t"); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil
	case manifest.KindMap:
		if node.Key == nil || node.Value == nil || node.Prefix == nil {
			return fmt.Errorf("map has no key, value, or prefix")
		}
		return e.mapEntries(b, node, expression, hint, indent)
	case manifest.KindUnion:
		name, err := e.g.goType(node, hint)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%smarshal%s(io, &%s)\n", indent, name, expression)
		return nil
	case manifest.KindReserved, manifest.KindIgnored:
		return fmt.Errorf("%s nodes require explicit write/discard codec semantics", node.Kind)
	case manifest.KindSequence, manifest.KindConditional:
		return fmt.Errorf("%s nodes do not yet have a generated codec", node.Kind)
	case manifest.KindOpaque, manifest.KindUnresolved:
		return fmt.Errorf("%s node blocks codec generation: %s", node.Kind, node.Reason)
	default:
		return fmt.Errorf("unsupported node kind %q", node.Kind)
	}
}

func (e *marshalEmitter) optional(b *strings.Builder, value manifest.Node, expression, hint, indent string) error {
	fmt.Fprintf(b, "%sio.Bool(&%s.set)\n%sif %s.set {\n", indent, expression, indent, expression)
	if err := e.node(b, value, expression+".val", hint, indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s} else if io.Reading() {\n%s\tvar zero %s\n%s\t%s.val = zero\n%s}\n", indent, indent, mustGoType(e.g, value, hint), indent, expression, indent)
	return nil
}

func (e *marshalEmitter) collection(b *strings.Builder, prefix, element manifest.Node, expression, hint, indent string) error {
	countType, err := e.g.goType(prefix, hint+"Count")
	if err != nil {
		return err
	}
	count := e.temporary("count")
	if maximum := integerTypeMaximum(countType); maximum != "" {
		fmt.Fprintf(b, "%sif !io.Reading() && uint64(len(%s)) > %s { io.InvalidValue(len(%s), \"collection length overflows %s\"); return }\n", indent, expression, maximum, expression, countType)
	}
	fmt.Fprintf(b, "%s%s := %s(len(%s))\n", indent, count, countType, expression)
	if err := e.node(b, prefix, count, hint+"Count", indent); err != nil {
		return err
	}
	fmt.Fprintf(b, "%sif io.Reading() {\n", indent)
	if strings.HasPrefix(countType, "int") {
		fmt.Fprintf(b, "%s\tif %s < 0 { io.InvalidValue(%s, \"negative collection length\"); return }\n", indent, count, count)
	}
	fmt.Fprintf(b, "%s\tif uint64(%s) > uint64(^uint(0)>>1) { io.InvalidValue(%s, \"collection length overflows int\"); return }\n", indent, count, count)
	fmt.Fprintf(b, "%s\t%s = make(%s, int(%s))\n%s}\n", indent, expression, mustGoType(e.g, manifest.Array(prefix, element), strings.TrimSuffix(hint, "Item")), count, indent)
	index := e.temporary("index")
	fmt.Fprintf(b, "%sfor %s := range %s {\n", indent, index, expression)
	if err := e.node(b, element, expression+"["+index+"]", hint, indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s}\n", indent)
	return nil
}

func (e *marshalEmitter) mapEntries(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	countType, err := e.g.goType(*node.Prefix, hint+"Count")
	if err != nil {
		return err
	}
	count := e.temporary("count")
	if maximum := integerTypeMaximum(countType); maximum != "" {
		fmt.Fprintf(b, "%sif !io.Reading() && uint64(len(%s)) > %s { io.InvalidValue(len(%s), \"map length overflows %s\"); return }\n", indent, expression, maximum, expression, countType)
	}
	fmt.Fprintf(b, "%s%s := %s(len(%s))\n", indent, count, countType, expression)
	if err := e.node(b, *node.Prefix, count, hint+"Count", indent); err != nil {
		return err
	}
	mapType, err := e.g.goType(node, hint)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sif io.Reading() {\n", indent)
	if strings.HasPrefix(countType, "int") {
		fmt.Fprintf(b, "%s\tif %s < 0 { io.InvalidValue(%s, \"negative map length\"); return }\n", indent, count, count)
	}
	fmt.Fprintf(b, "%s\tif uint64(%s) > uint64(^uint(0)>>1) { io.InvalidValue(%s, \"map length overflows int\"); return }\n", indent, count, count)
	fmt.Fprintf(b, "%s\t%s = make(%s, int(%s))\n%s}\n", indent, expression, mapType, count, indent)
	index := e.temporary("index")
	fmt.Fprintf(b, "%sfor %s := range %s {\n", indent, index, expression)
	if err := e.node(b, *node.Key, expression+"["+index+"].Key", hint+"Key", indent+"\t"); err != nil {
		return err
	}
	if err := e.node(b, *node.Value, expression+"["+index+"].Value", hint+"Value", indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s}\n", indent)
	return nil
}

func (e *marshalEmitter) enum(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	if node.Primitive == nil {
		return fmt.Errorf("enum has no primitive")
	}
	if len(node.Variants) == 0 {
		return fmt.Errorf("enum has no variants")
	}
	underlying, err := primitiveGoType(node.Primitive.Code)
	if err != nil {
		return err
	}
	wire := e.temporary("enumValue")
	fmt.Fprintf(b, "%s%s := %s(%s)\n", indent, wire, underlying, expression)
	method, err := primitiveIOMethod(node.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sio.%s(&%s)\n%s%s = %s(%s)\n", indent, method, wire, indent, expression, mustGoType(e.g, node, hint), wire)
	fmt.Fprintf(b, "%sswitch int64(%s) {\n%scase ", indent, wire, indent)
	for index, variant := range node.Variants {
		if index != 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(b, "%d", variant.Value)
	}
	fmt.Fprintf(b, ":\n%sdefault:\n%s\tio.InvalidValue(%s, \"unknown enum value\")\n%s}\n", indent, indent, wire, indent)
	return nil
}

func integerTypeMaximum(typ string) string {
	switch typ {
	case "uint8":
		return "uint64(^uint8(0))"
	case "uint16":
		return "uint64(^uint16(0))"
	case "uint32":
		return "uint64(^uint32(0))"
	case "uint64":
		return "^uint64(0)"
	case "int8":
		return "uint64(^uint8(0) >> 1)"
	case "int16":
		return "uint64(^uint16(0) >> 1)"
	case "int32":
		return "uint64(^uint32(0) >> 1)"
	case "int64":
		return "uint64(^uint64(0) >> 1)"
	default:
		return ""
	}
}

func (e *marshalEmitter) union(b *strings.Builder, definition typeDefinition) error {
	fmt.Fprintf(b, "func marshal%s(io IO, x *%s) {\n", definition.Name, definition.Name)
	if len(definition.Union) == 0 {
		b.WriteString("\tio.InvalidValue(nil, \"union has no variants\")\n}\n\n")
		return nil
	}
	// Every union member is a generated struct (or a wrapper around a scalar),
	// so payload dispatch can remain fully typed in both directions.
	b.WriteString("\tif io.Reading() {\n")
	if definition.Underlying == "" {
		return fmt.Errorf("union %s has no discriminator primitive", definition.Name)
	}
	controlNode := manifest.Primitive(definition.Underlying)
	tagType, err := primitiveGoType(definition.Underlying)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "\t\tvar tag %s\n", tagType)
	if err := e.node(b, controlNode, "tag", definition.Name+"Tag", "\t\t"); err != nil {
		return err
	}
	b.WriteString("\t\tswitch int64(tag) {\n")
	for _, member := range definition.Union {
		fmt.Fprintf(b, "\t\tcase %d:\n\t\t\tvar value %s\n\t\t\tvalue.Marshal(io)\n\t\t\t*x = value\n", member.Value, member.Name)
	}
	b.WriteString("\t\tdefault:\n\t\t\tio.InvalidValue(tag, \"unknown union tag\")\n\t\t}\n\t\treturn\n\t}\n\tswitch value := (*x).(type) {\n")
	for _, member := range definition.Union {
		fmt.Fprintf(b, "\tcase %s:\n\t\ttag := %s(%d)\n", member.Name, tagType, member.Value)
		if err := e.node(b, controlNode, "tag", definition.Name+"Tag", "\t\t"); err != nil {
			return err
		}
		b.WriteString("\t\tvalue.Marshal(io)\n")
	}
	b.WriteString("\tdefault:\n\t\tio.InvalidValue(*x, \"unknown union value\")\n\t}\n}\n\n")
	return nil
}

func mustGoType(g *generator, node manifest.Node, hint string) string {
	typ, err := g.goType(node, hint)
	if err != nil {
		panic(err)
	}
	return typ
}

func varuint32Prefix(node manifest.Node) bool {
	return node.Prefix != nil && node.Prefix.Kind == manifest.KindPrimitive && node.Prefix.Primitive != nil && node.Prefix.Primitive.Code == "var_u32"
}

func primitiveIOMethod(code string) (string, error) {
	switch code {
	case "bool":
		return "Bool", nil
	case "i8":
		return "Int8", nil
	case "u8":
		return "Uint8", nil
	case "i16le":
		return "Int16", nil
	case "u16le":
		return "Uint16", nil
	case "i16be":
		return "BEInt16", nil
	case "u16be":
		return "BEUint16", nil
	case "i32le":
		return "Int32", nil
	case "u32le":
		return "Uint32", nil
	case "i32be":
		return "BEInt32", nil
	case "u32be":
		return "BEUint32", nil
	case "i64le":
		return "Int64", nil
	case "u64le":
		return "Uint64", nil
	case "i64be":
		return "BEInt64", nil
	case "u64be":
		return "BEUint64", nil
	case "f32le":
		return "Float32", nil
	case "f64le":
		return "Float64", nil
	case "f32be":
		return "BEFloat32", nil
	case "f64be":
		return "BEFloat64", nil
	case "zigzag_i32":
		return "Varint32", nil
	case "zigzag_i64":
		return "Varint64", nil
	case "var_u32":
		return "Varuint32", nil
	case "var_u64":
		return "Varuint64", nil
	case "var_i32":
		return "SignedVarint32", nil
	case "var_i64":
		return "SignedVarint64", nil
	default:
		return "", fmt.Errorf("primitive %q has no IO method", code)
	}
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
	value = collapseRedundantGroup(value)
	if strings.HasSuffix(value, "PayloadUnion") {
		value = strings.TrimSuffix(value, "PayloadUnion") + "Value"
	}
	value = strings.TrimSuffix(value, "Union")
	value = strings.TrimSuffix(value, "Payload")
	return strings.ReplaceAll(exportName(value), "Molang", "MoLang")
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
