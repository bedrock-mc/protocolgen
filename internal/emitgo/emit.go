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
	case "ActorRuntimeID":
		switch code, ok := primitiveStructCode(node); {
		case ok && code == "var_u64":
			return "uint64", true, nil
		case ok && code == "zigzag_i64":
			return "int64", true, nil
		case ok && code == "var_u32":
			return "uint32", true, nil
		case ok && code == "var_i64":
			return "int64", true, nil
		default:
			return "", true, fmt.Errorf("native ActorRuntimeID mapping requires a supported variable-width ID field")
		}
	case "ActorUniqueID":
		switch code, ok := primitiveStructCode(node); {
		case ok && code == "zigzag_i64":
			return "int64", true, nil
		case ok && code == "i64le":
			return "int64", true, nil
		case ok && code == "u64le":
			return "uint64", true, nil
		case ok && code == "var_u64":
			return "uint64", true, nil
		case ok && code == "var_i64":
			return "int64", true, nil
		default:
			return "", true, fmt.Errorf("native ActorUniqueID mapping requires a supported signed, unsigned, or fixed-width ID field")
		}
	case "PlayerInputTick":
		if !isPrimitiveStruct(node, "var_u64") {
			return "", true, fmt.Errorf("native PlayerInputTick mapping requires exactly one var_u64 field")
		}
		return "uint64", true, nil
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

func primitiveStructCode(node manifest.Node) (string, bool) {
	if len(node.Fields) != 1 {
		return "", false
	}
	field := node.Fields[0].Encode
	if field.Kind != manifest.KindPrimitive || field.Primitive == nil {
		return "", false
	}
	return field.Primitive.Code, true
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
	typesSource, err := emitTypeDefinitions(packageName)
	if err != nil {
		return nil, err
	}
	files["types.go"] = typesSource
	files["codec.go"] = emitCodecRuntime(packageName)
	files["reader.go"] = emitReader(packageName)
	files["writer.go"] = emitWriter(packageName)
	files["ids.go"] = emitPacketIDs(packageName, packets, packetNames)
	usedFiles := map[string]bool{
		"types.go": true, "codec.go": true, "reader.go": true, "writer.go": true, "ids.go": true,
	}
	for _, definition := range definitions {
		source, err := emitDefinition(packageName, g, definition)
		if err != nil {
			return nil, err
		}
		name := uniqueFileName(snakeName(definition.Name)+".go", 0, usedFiles)
		files[name] = source
	}
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

func emitDefinition(packageName string, g *generator, definition typeDefinition) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	switch definition.Kind {
	case manifest.KindStruct:
		var types []string
		for _, field := range definition.Fields {
			types = append(types, field.Type)
		}
		writeGoImports(&b, goImportsForTypes(types))
		fmt.Fprintf(&b, "type %s struct {\n", definition.Name)
		for _, field := range definition.Fields {
			fmt.Fprintf(&b, "\t%s %s\n", field.Name, field.Type)
		}
		b.WriteString("}\n\n")
		for _, union := range definition.Implements {
			fmt.Fprintf(&b, "func (%s) is%s() {}\n\n", definition.Name, union)
		}
		fmt.Fprintf(&b, "// Marshal reads or writes %s using its canonical wire layout.\n", definition.Name)
		fmt.Fprintf(&b, "func (x *%s) Marshal(io IO) {\n", definition.Name)
		emitter := marshalEmitter{g: g}
		for _, field := range definition.Fields {
			if err := emitter.node(&b, field.Node, "x."+field.Name, definition.Name+field.Name, "\t"); err != nil {
				return "", fmt.Errorf("type %s field %s marshal: %w", definition.Name, field.Name, err)
			}
		}
		b.WriteString("}\n")
	case manifest.KindUnion:
		fmt.Fprintf(&b, "type %s interface {\n\tis%s()\n}\n\n", definition.Name, definition.Name)
		emitter := marshalEmitter{g: g}
		if err := emitter.union(&b, definition); err != nil {
			return "", err
		}
	case manifest.KindEnum:
		fmt.Fprintf(&b, "type %s %s\n\nconst (\n", definition.Name, definition.Underlying)
		variantNames := map[string]string{}
		for _, variant := range definition.Variants {
			name := enumVariantName(variant.Name)
			if previous, exists := variantNames[name]; exists {
				return "", fmt.Errorf("enum %s variants %q and %q both map to %s", definition.Name, previous, variant.Name, name)
			}
			variantNames[name] = variant.Name
			fmt.Fprintf(&b, "\t%s%s %s = %d\n", definition.Name, name, definition.Name, variant.Value)
		}
		b.WriteString(")\n")
	case manifest.KindBitset:
		fmt.Fprintf(&b, "// %s stores the %d-bit value used by the wire bitset encoding.\n", definition.Name, definition.BitLength)
		fmt.Fprintf(&b, "type %s [%d]uint64\n", definition.Name, (definition.BitLength+63)/64)
	default:
		return "", fmt.Errorf("unsupported definition kind %q", definition.Kind)
	}
	return formatGoSource(b.String())
}

func emitTypeDefinitions(packageName string) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage %s\n\n", packageName)
	b.WriteString("// Optional holds a value that may be absent from the wire.\n")
	b.WriteString("type Optional[T any] struct {\n\tset bool\n\tval T\n}\n\n")
	b.WriteString("// Option creates a present Optional containing value.\n")
	b.WriteString("func Option[T any](value T) Optional[T] {\n\treturn Optional[T]{set: true, val: value}\n}\n\n")
	b.WriteString("// Value returns the optional value and whether it is present.\n")
	b.WriteString("func (o Optional[T]) Value() (T, bool) {\n\treturn o.val, o.set\n}\n\n")
	b.WriteString("// OrderedEntry preserves the source order and duplicate keys of a wire map.\n")
	b.WriteString("type OrderedEntry[K, V any] struct {\n\tKey K\n\tValue V\n}\n\n")
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
	return mustFormatGoSource(fmt.Sprintf(`// Code generated from canonical protocol manifest v2. DO NOT EDIT.

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
	// Semantic operations preserve the distinction between identifiers and
	// ordinary integers for downstream translation and validation.
	ActorRuntimeID(*uint64)
	ActorRuntimeIDVarint64(*int64)
	ActorRuntimeIDVaruint32(*uint32)
	ActorUniqueID(*int64)
	ActorUniqueIDInt64(*int64)
	ActorUniqueIDUint64(*uint64)
	ActorUniqueIDVaruint64(*uint64)
	PlayerInputTick(*uint64)

	String(*string)
	Bytes(*[]byte)
	NBT(*[]byte)
	UUID(*uuid.UUID)
	Vec2(*mgl32.Vec2)
	Vec3(*mgl32.Vec3)
	RGBA(*color.RGBA)
	Bitset(words []uint64, bits uint64)
}

`, packageName) + emitCodecHelpers())
}

func emitReader(packageName string) string {
	const source = `// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package {{PACKAGE}}

import (
	"encoding/binary"
	"fmt"
	"image/color"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// Reader decodes one generated protocol value from a byte slice. Decode
// failures are retained by Err instead of panicking.
type Reader struct {
	data           []byte
	pos            int
	err            error
	maxSliceLength uint64
}

// NewReader creates a Reader over data.
func NewReader(data []byte) *Reader {
	return NewReaderWithLimit(data, maxSliceLength)
}

// NewReaderWithLimit creates a Reader with an explicit maximum decoded
// collection length. A zero limit rejects every non-empty collection.
func NewReaderWithLimit(data []byte, max uint64) *Reader {
	return &Reader{data: data, maxSliceLength: max}
}

// Reading reports that this IO implementation decodes values.
func (*Reader) Reading() bool { return true }

// Err returns the first decode error, if any.
func (r *Reader) Err() error { return r.err }

// Remaining reports the number of unread bytes.
func (r *Reader) Remaining() int {
	if r.pos >= len(r.data) {
		return 0
	}
	return len(r.data) - r.pos
}

func (r *Reader) fail(err error) {
	if r.err == nil {
		r.err = err
	}
}

// InvalidValue records a malformed wire value.
func (r *Reader) InvalidValue(value any, context string) {
	r.fail(fmt.Errorf("invalid value %v for %s", value, context))
}

func (r *Reader) readByte() byte {
	if r.err != nil {
		return 0
	}
	if r.pos >= len(r.data) {
		r.fail(fmt.Errorf("unexpected end of input"))
		return 0
	}
	v := r.data[r.pos]
	r.pos++
	return v
}

func (r *Reader) readN(n int) []byte {
	if r.err != nil {
		return nil
	}
	if n < 0 || n > len(r.data)-r.pos {
		r.fail(fmt.Errorf("unexpected end of input reading %d bytes", n))
		return nil
	}
	v := r.data[r.pos : r.pos+n]
	r.pos += n
	return v
}

func (r *Reader) readUvarint() uint64 {
	var value uint64
	for shift := uint(0); shift < 64; shift += 7 {
		b := r.readByte()
		if r.err != nil {
			return 0
		}
		if shift == 63 && b > 1 {
			r.fail(fmt.Errorf("varint overflows uint64"))
			return 0
		}
		value |= uint64(b&0x7f) << shift
		if b&0x80 == 0 {
			return value
		}
	}
	r.fail(fmt.Errorf("varint exceeds ten bytes"))
	return 0
}

func (r *Reader) Bool(x *bool) {
	value := r.readByte()
	if r.err != nil {
		return
	}
	if value > 1 {
		r.InvalidValue(value, "boolean must be 0 or 1")
		return
	}
	*x = value == 1
}
func (r *Reader) Int8(x *int8) { *x = int8(r.readByte()) }
func (r *Reader) Uint8(x *uint8) { *x = r.readByte() }

func (r *Reader) Int16(x *int16) {
	data := r.readN(2)
	if data != nil {
		*x = int16(binary.LittleEndian.Uint16(data))
	}
}
func (r *Reader) Uint16(x *uint16) {
	data := r.readN(2)
	if data != nil {
		*x = binary.LittleEndian.Uint16(data)
	}
}
func (r *Reader) BEInt16(x *int16) {
	data := r.readN(2)
	if data != nil {
		*x = int16(binary.BigEndian.Uint16(data))
	}
}
func (r *Reader) BEUint16(x *uint16) {
	data := r.readN(2)
	if data != nil {
		*x = binary.BigEndian.Uint16(data)
	}
}

func (r *Reader) Int32(x *int32) {
	data := r.readN(4)
	if data != nil {
		*x = int32(binary.LittleEndian.Uint32(data))
	}
}
func (r *Reader) Uint32(x *uint32) {
	data := r.readN(4)
	if data != nil {
		*x = binary.LittleEndian.Uint32(data)
	}
}
func (r *Reader) BEInt32(x *int32) {
	data := r.readN(4)
	if data != nil {
		*x = int32(binary.BigEndian.Uint32(data))
	}
}
func (r *Reader) BEUint32(x *uint32) {
	data := r.readN(4)
	if data != nil {
		*x = binary.BigEndian.Uint32(data)
	}
}

func (r *Reader) Int64(x *int64) {
	data := r.readN(8)
	if data != nil {
		*x = int64(binary.LittleEndian.Uint64(data))
	}
}
func (r *Reader) Uint64(x *uint64) {
	data := r.readN(8)
	if data != nil {
		*x = binary.LittleEndian.Uint64(data)
	}
}
func (r *Reader) BEInt64(x *int64) {
	data := r.readN(8)
	if data != nil {
		*x = int64(binary.BigEndian.Uint64(data))
	}
}
func (r *Reader) BEUint64(x *uint64) {
	data := r.readN(8)
	if data != nil {
		*x = binary.BigEndian.Uint64(data)
	}
}

func (r *Reader) Float32(x *float32) {
	data := r.readN(4)
	if data != nil {
		*x = math.Float32frombits(binary.LittleEndian.Uint32(data))
	}
}
func (r *Reader) Float64(x *float64) {
	data := r.readN(8)
	if data != nil {
		*x = math.Float64frombits(binary.LittleEndian.Uint64(data))
	}
}
func (r *Reader) BEFloat32(x *float32) {
	data := r.readN(4)
	if data != nil {
		*x = math.Float32frombits(binary.BigEndian.Uint32(data))
	}
}
func (r *Reader) BEFloat64(x *float64) {
	data := r.readN(8)
	if data != nil {
		*x = math.Float64frombits(binary.BigEndian.Uint64(data))
	}
}

func (r *Reader) Varint32(x *int32) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	if value > uint64(^uint32(0)) {
		r.InvalidValue(value, "zigzag int32 overflows uint32")
		return
	}
	value32 := uint32(value)
	*x = int32(value32 >> 1)
	if value32&1 != 0 {
		*x = ^*x
	}
}

func (r *Reader) Varuint32(x *uint32) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	if value > uint64(^uint32(0)) {
		r.InvalidValue(value, "varuint32 overflows uint32")
		return
	}
	*x = uint32(value)
}

func (r *Reader) Varint64(x *int64) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	*x = int64(value >> 1)
	if value&1 != 0 {
		*x = ^*x
	}
}

func (r *Reader) Varuint64(x *uint64) { *x = r.readUvarint() }
func (r *Reader) SignedVarint32(x *int32) {
	value := r.readUvarint()
	if r.err != nil {
		return
	}
	if value > uint64(^uint32(0)) {
		r.InvalidValue(value, "signed varint32 overflows uint32")
		return
	}
	*x = int32(value)
}
func (r *Reader) SignedVarint64(x *int64) { *x = int64(r.readUvarint()) }

func (r *Reader) ActorRuntimeID(x *uint64) { r.Varuint64(x) }
func (r *Reader) ActorRuntimeIDVarint64(x *int64) { r.Varint64(x) }
func (r *Reader) ActorRuntimeIDVaruint32(x *uint32) { r.Varuint32(x) }
func (r *Reader) ActorUniqueID(x *int64) { r.Varint64(x) }
func (r *Reader) ActorUniqueIDInt64(x *int64) { r.Int64(x) }
func (r *Reader) ActorUniqueIDUint64(x *uint64) { r.Uint64(x) }
func (r *Reader) ActorUniqueIDVaruint64(x *uint64) { r.Varuint64(x) }
func (r *Reader) PlayerInputTick(x *uint64) { r.Varuint64(x) }

func (r *Reader) String(x *string) {
	length, ok := r.readLength("string")
	if !ok {
		return
	}
	*x = string(r.readN(length))
}

func (r *Reader) Bytes(x *[]byte) {
	length, ok := r.readLength("byte slice")
	if !ok {
		return
	}
	data := r.readN(length)
	if data == nil {
		return
	}
	*x = append((*x)[:0], data...)
}

func (r *Reader) readLength(context string) (int, bool) {
	var length uint32
	r.Varuint32(&length)
	if r.err != nil {
		return 0, false
	}
	if uint64(length) > uint64(^uint(0)>>1) {
		r.InvalidValue(length, context+" length overflows int")
		return 0, false
	}
	return int(length), true
}

func (r *Reader) NBT(x *[]byte) {
	if r.err != nil {
		return
	}
	start := r.pos
	if err := scanNBT(r.data, &r.pos, true); err != nil {
		r.fail(err)
		return
	}
	*x = append((*x)[:0], r.data[start:r.pos]...)
}

func (r *Reader) UUID(x *uuid.UUID) {
	data := r.readN(16)
	if data == nil {
		return
	}
	copy((*x)[:8], data[:8])
	copy((*x)[8:], data[8:])
	reverseBytes((*x)[:8])
	reverseBytes((*x)[8:])
}

func (r *Reader) Vec2(x *mgl32.Vec2) {
	r.Float32(&(*x)[0])
	r.Float32(&(*x)[1])
}

func (r *Reader) Vec3(x *mgl32.Vec3) {
	r.Float32(&(*x)[0])
	r.Float32(&(*x)[1])
	r.Float32(&(*x)[2])
}

func (r *Reader) RGBA(x *color.RGBA) {
	var value uint32
	r.Uint32(&value)
	*x = color.RGBA{R: byte(value), G: byte(value >> 8), B: byte(value >> 16), A: byte(value >> 24)}
}

func (r *Reader) Bitset(words []uint64, bits uint64) {
	wordCount := bits / 64
	if bits%64 != 0 {
		wordCount++
	}
	if uint64(len(words)) != wordCount {
		r.InvalidValue(len(words), "bitset word count does not match declared width")
		return
	}
	for i := range words {
		words[i] = 0
	}
	for offset := uint64(0); offset < bits; offset += 7 {
		value := r.readByte()
		if r.err != nil {
			return
		}
		width := bits - offset
		if width > 7 {
			width = 7
		}
		if width < 7 && uint64(value&0x7f) >= (uint64(1)<<width) {
			r.InvalidValue(value, "bitset contains bits outside its declared width")
			return
		}
		for bit := uint64(0); bit < width; bit++ {
			if value&(1<<bit) != 0 {
				index := offset + bit
				words[index/64] |= uint64(1) << (index % 64)
			}
		}
		if value&0x80 == 0 {
			return
		}
		if offset+7 >= bits {
			r.InvalidValue(value, "bitset exceeds its declared width")
			return
		}
	}
}

func (r *Reader) SliceLength(value uint64, max uint64) bool {
	if r.maxSliceLength < max {
		max = r.maxSliceLength
	}
	if value > max {
		r.InvalidValue(value, "collection length exceeds decoder limit")
		return false
	}
	return true
}

func reverseBytes(value []byte) {
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
}

func scanNBT(data []byte, pos *int, named bool) error {
	if *pos >= len(data) {
		return fmt.Errorf("NBT ended before its tag")
	}
	tag := data[*pos]
	*pos += 1
	if tag == 0 {
		return nil
	}
	if named {
		if _, err := nbtReadBytes(data, pos, 2); err != nil {
			return err
		}
		nameLength := int(binary.LittleEndian.Uint16(data[*pos-2 : *pos]))
		if _, err := nbtReadBytes(data, pos, nameLength); err != nil {
			return err
		}
	}
	return scanNBTPayload(data, pos, tag)
}

func scanNBTPayload(data []byte, pos *int, tag byte) error {
	switch tag {
	case 1:
		_, err := nbtReadBytes(data, pos, 1)
		return err
	case 2:
		_, err := nbtReadBytes(data, pos, 2)
		return err
	case 3, 5:
		_, err := nbtReadBytes(data, pos, 4)
		return err
	case 4, 6:
		_, err := nbtReadBytes(data, pos, 8)
		return err
	case 7:
		length, err := nbtReadInt32(data, pos)
		if err != nil || length < 0 {
			return fmt.Errorf("invalid NBT byte-array length %d", length)
		}
		_, err = nbtReadBytes(data, pos, int(length))
		return err
	case 8:
		length, err := nbtReadUint16(data, pos)
		if err != nil {
			return err
		}
		_, err = nbtReadBytes(data, pos, int(length))
		return err
	case 9:
		element, err := nbtReadByte(data, pos)
		if err != nil {
			return err
		}
		length, err := nbtReadInt32(data, pos)
		if err != nil || length < 0 {
			return fmt.Errorf("invalid NBT list length %d", length)
		}
		for i := int32(0); i < length; i++ {
			if err := scanNBTPayload(data, pos, element); err != nil {
				return err
			}
		}
		return nil
	case 10:
		for {
			if *pos >= len(data) {
				return fmt.Errorf("NBT compound has no end tag")
			}
			if data[*pos] == 0 {
				*pos += 1
				return nil
			}
			if err := scanNBT(data, pos, true); err != nil {
				return err
			}
		}
	case 11, 12:
		length, err := nbtReadInt32(data, pos)
		if err != nil || length < 0 {
			return fmt.Errorf("invalid NBT array length %d", length)
		}
		width := 4
		if tag == 12 {
			width = 8
		}
		_, err = nbtReadBytes(data, pos, int(length)*width)
		return err
	default:
		return fmt.Errorf("unknown NBT tag %d", tag)
	}
}

func nbtReadByte(data []byte, pos *int) (byte, error) {
	if *pos >= len(data) {
		return 0, fmt.Errorf("NBT ended unexpectedly")
	}
	value := data[*pos]
	*pos += 1
	return value, nil
}

func nbtReadBytes(data []byte, pos *int, length int) ([]byte, error) {
	if length < 0 || length > len(data)-*pos {
		return nil, fmt.Errorf("NBT ended unexpectedly")
	}
	value := data[*pos : *pos+length]
	*pos += length
	return value, nil
}

func nbtReadUint16(data []byte, pos *int) (uint16, error) {
	value, err := nbtReadBytes(data, pos, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(value), nil
}

func nbtReadInt32(data []byte, pos *int) (int32, error) {
	value, err := nbtReadBytes(data, pos, 4)
	if err != nil {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(value)), nil
}
`
	return mustFormatGoSource(strings.ReplaceAll(source, "{{PACKAGE}}", packageName))
}

func emitWriter(packageName string) string {
	const source = `// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package {{PACKAGE}}

import (
	"encoding/binary"
	"fmt"
	"image/color"
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// Writer encodes one generated protocol value into an in-memory byte slice.
// InvalidValue records a terminal error; callers should inspect Err after the
// generated Marshal method returns.
type Writer struct {
	data []byte
	err  error
}

// NewWriter creates an empty Writer.
func NewWriter() *Writer { return &Writer{} }

// Reading reports that this IO implementation encodes values.
func (*Writer) Reading() bool { return false }

// Err returns the first encoding error, if any.
func (w *Writer) Err() error { return w.err }

// Data returns the encoded bytes. The returned slice is owned by the Writer.
func (w *Writer) Data() []byte { return w.data }

func (w *Writer) fail(err error) {
	if w.err == nil {
		w.err = err
	}
}

// InvalidValue records a malformed model or an unsupported wire value.
func (w *Writer) InvalidValue(value any, context string) {
	w.fail(fmt.Errorf("invalid value %v for %s", value, context))
}

func (w *Writer) writeByte(value byte) {
	if w.err == nil {
		w.data = append(w.data, value)
	}
}

func (w *Writer) write(value []byte) {
	if w.err == nil {
		w.data = append(w.data, value...)
	}
}

func (w *Writer) writeUvarint(value uint64) {
	if w.err != nil {
		return
	}
	for value >= 0x80 {
		w.data = append(w.data, byte(value)|0x80)
		value >>= 7
	}
	w.data = append(w.data, byte(value))
}

func (w *Writer) Bool(x *bool) {
	if *x {
		w.writeByte(1)
	} else {
		w.writeByte(0)
	}
}

func (w *Writer) Int8(x *int8) { w.writeByte(byte(*x)) }
func (w *Writer) Uint8(x *uint8) { w.writeByte(*x) }

func (w *Writer) Int16(x *int16) {
	var data [2]byte
	binary.LittleEndian.PutUint16(data[:], uint16(*x))
	w.write(data[:])
}

func (w *Writer) Uint16(x *uint16) {
	var data [2]byte
	binary.LittleEndian.PutUint16(data[:], *x)
	w.write(data[:])
}

func (w *Writer) BEInt16(x *int16) {
	var data [2]byte
	binary.BigEndian.PutUint16(data[:], uint16(*x))
	w.write(data[:])
}

func (w *Writer) BEUint16(x *uint16) {
	var data [2]byte
	binary.BigEndian.PutUint16(data[:], *x)
	w.write(data[:])
}

func (w *Writer) Int32(x *int32) {
	var data [4]byte
	binary.LittleEndian.PutUint32(data[:], uint32(*x))
	w.write(data[:])
}

func (w *Writer) Uint32(x *uint32) {
	var data [4]byte
	binary.LittleEndian.PutUint32(data[:], *x)
	w.write(data[:])
}

func (w *Writer) BEInt32(x *int32) {
	var data [4]byte
	binary.BigEndian.PutUint32(data[:], uint32(*x))
	w.write(data[:])
}

func (w *Writer) BEUint32(x *uint32) {
	var data [4]byte
	binary.BigEndian.PutUint32(data[:], *x)
	w.write(data[:])
}

func (w *Writer) Int64(x *int64) {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], uint64(*x))
	w.write(data[:])
}

func (w *Writer) Uint64(x *uint64) {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], *x)
	w.write(data[:])
}

func (w *Writer) BEInt64(x *int64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], uint64(*x))
	w.write(data[:])
}

func (w *Writer) BEUint64(x *uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], *x)
	w.write(data[:])
}

func (w *Writer) Float32(x *float32) {
	var data [4]byte
	binary.LittleEndian.PutUint32(data[:], math.Float32bits(*x))
	w.write(data[:])
}

func (w *Writer) Float64(x *float64) {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], math.Float64bits(*x))
	w.write(data[:])
}

func (w *Writer) BEFloat32(x *float32) {
	var data [4]byte
	binary.BigEndian.PutUint32(data[:], math.Float32bits(*x))
	w.write(data[:])
}

func (w *Writer) BEFloat64(x *float64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], math.Float64bits(*x))
	w.write(data[:])
}

func (w *Writer) Varint32(x *int32) {
	value := uint32(*x)<<1 ^ uint32(*x>>31)
	w.writeUvarint(uint64(value))
}

func (w *Writer) Varuint32(x *uint32) { w.writeUvarint(uint64(*x)) }

func (w *Writer) Varint64(x *int64) {
	value := uint64(*x<<1) ^ uint64(*x>>63)
	w.writeUvarint(value)
}

func (w *Writer) Varuint64(x *uint64) { w.writeUvarint(*x) }
func (w *Writer) SignedVarint32(x *int32) { w.writeUvarint(uint64(uint32(*x))) }
func (w *Writer) SignedVarint64(x *int64) { w.writeUvarint(uint64(*x)) }

func (w *Writer) ActorRuntimeID(x *uint64) { w.Varuint64(x) }
func (w *Writer) ActorRuntimeIDVarint64(x *int64) { w.Varint64(x) }
func (w *Writer) ActorRuntimeIDVaruint32(x *uint32) { w.Varuint32(x) }
func (w *Writer) ActorUniqueID(x *int64) { w.Varint64(x) }
func (w *Writer) ActorUniqueIDInt64(x *int64) { w.Int64(x) }
func (w *Writer) ActorUniqueIDUint64(x *uint64) { w.Uint64(x) }
func (w *Writer) ActorUniqueIDVaruint64(x *uint64) { w.Varuint64(x) }
func (w *Writer) PlayerInputTick(x *uint64) { w.Varuint64(x) }

func (w *Writer) String(x *string) {
	data := []byte(*x)
	if uint64(len(data)) > uint64(^uint32(0)) {
		w.InvalidValue(len(data), "string exceeds uint32 length")
		return
	}
	length := uint32(len(data))
	w.Varuint32(&length)
	w.write(data)
}

func (w *Writer) Bytes(x *[]byte) {
	if uint64(len(*x)) > uint64(^uint32(0)) {
		w.InvalidValue(len(*x), "byte slice exceeds uint32 length")
		return
	}
	length := uint32(len(*x))
	w.Varuint32(&length)
	w.write(*x)
}

func (w *Writer) NBT(x *[]byte) {
	if w.err != nil {
		return
	}
	position := 0
	if err := scanNBT(*x, &position, true); err != nil {
		w.fail(err)
		return
	}
	if position != len(*x) {
		w.fail(fmt.Errorf("NBT has %d trailing bytes", len(*x)-position))
		return
	}
	w.write(*x)
}

func (w *Writer) UUID(x *uuid.UUID) {
	data := append([]byte(nil), x[:]...)
	reverseBytes(data[:8])
	reverseBytes(data[8:])
	w.write(data)
}

func (w *Writer) Vec2(x *mgl32.Vec2) {
	w.Float32(&(*x)[0])
	w.Float32(&(*x)[1])
}

func (w *Writer) Vec3(x *mgl32.Vec3) {
	w.Float32(&(*x)[0])
	w.Float32(&(*x)[1])
	w.Float32(&(*x)[2])
}

func (w *Writer) RGBA(x *color.RGBA) {
	value := uint32(x.R) | uint32(x.G)<<8 | uint32(x.B)<<16 | uint32(x.A)<<24
	w.Uint32(&value)
}

func (w *Writer) Bitset(words []uint64, bits uint64) {
	wordCount := bits / 64
	if bits%64 != 0 {
		wordCount++
	}
	if uint64(len(words)) != wordCount {
		w.InvalidValue(len(words), "bitset word count does not match declared width")
		return
	}
	if bits%64 != 0 && words[len(words)-1]>> (bits % 64) != 0 {
		w.InvalidValue(words[len(words)-1], "bitset contains bits outside its declared width")
		return
	}
	last := uint64(0)
	found := false
	for index, word := range words {
		if word == 0 {
			continue
		}
		found = true
		for bit := uint64(64); bit > 0; bit-- {
			if word&(uint64(1)<<(bit-1)) != 0 {
				last = uint64(index)*64 + bit - 1
				break
			}
		}
		if index == len(words)-1 {
			break
		}
	}
	if !found {
		w.writeByte(0)
		return
	}
	groups := last/7 + 1
	for group := uint64(0); group < groups; group++ {
		offset := group * 7
		width := bits - offset
		if width > 7 {
			width = 7
		}
		var value byte
		for bit := uint64(0); bit < width; bit++ {
			index := offset + bit
			if words[index/64]&(uint64(1)<<(index%64)) != 0 {
				value |= 1 << bit
			}
		}
		if group+1 < groups {
			value |= 0x80
		}
		w.writeByte(value)
	}
}
`
	return mustFormatGoSource(strings.ReplaceAll(source, "{{PACKAGE}}", packageName))
}

func mustFormatGoSource(source string) string {
	formatted, err := formatGoSource(source)
	if err != nil {
		panic(err)
	}
	return formatted
}

func emitCodecHelpers() string {
	return `// Marshaler is implemented by every generated struct and union payload.
type Marshaler interface {
	Marshal(IO)
}

// PtrMarshaler constrains a value whose pointer implements Marshaler.
type PtrMarshaler[T any] interface {
	Marshaler
	*T
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// IntegerFunc marshals a value through a wire integer type while preserving
// the public type used by the generated model.
func IntegerFunc[S, W integer](x *S, f func(*W)) {
	w := W(*x)
	f(&w)
	*x = S(w)
}

// OptionalFunc marshals a bool-prefixed optional value.
func OptionalFunc[T any](io IO, x *Optional[T], f func(*T)) {
	io.Bool(&x.set)
	if x.set {
		f(&x.val)
	}
}

// DoubleOptionalFunc marshals an optional nested inside an always-present
// outer optional, as used by several Cereal fields.
func DoubleOptionalFunc[T any](io IO, x *Optional[T], f func(*T)) {
	outer := true
	io.Bool(&outer)
	if outer {
		OptionalFunc(io, x, f)
	} else {
		*x = Optional[T]{}
	}
}

// OptionalMarshaler marshals an Optional whose value has a Marshal method.
func OptionalMarshaler[T any, A PtrMarshaler[T]](io IO, x *Optional[T]) {
	io.Bool(&x.set)
	if x.set {
		A(&x.val).Marshal(io)
	}
}

// UnionFunc dispatches a generated union codec without exposing the IO
// direction branch in every packet definition.
func UnionFunc(io IO, read, write func()) {
	if io.Reading() {
		read()
	} else {
		write()
	}
}

// sliceReader is implemented by decoders that allocate slices from a wire
// count. Writers intentionally do not implement it, so the same helpers work
// in both directions without exposing direction checks in packet methods.
type sliceReader interface {
	SliceLength(value uint64, max uint64) bool
}

const maxSliceLength = 4096

// FuncSlice marshals a length-prefixed slice using a count encoder and element
// callback. The count encoder determines the exact wire prefix type.
func FuncSlice[T any, C integer](io IO, x *[]T, count func(*C), f func(*T)) {
	if io.Reading() {
		reader, ok := io.(sliceReader)
		if !ok {
			io.InvalidValue(io, "reader does not implement SliceLength")
			return
		}
		var n C
		count(&n)
		length, valid := sliceLength(io, n)
		if !valid {
			return
		}
		if !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]T, length)
	} else {
		n, valid := sliceCount[C](io, len(*x))
		if !valid {
			return
		}
		count(&n)
	}
	for i := range *x {
		f(&(*x)[i])
	}
}

// OrderedMap marshals an ordered map representation while preserving duplicate
// keys and source order.
func OrderedMap[K, V any, C integer](io IO, x *[]OrderedEntry[K, V], count func(*C), key func(*K), value func(*V)) {
	if io.Reading() {
		reader, ok := io.(sliceReader)
		if !ok {
			io.InvalidValue(io, "reader does not implement SliceLength")
			return
		}
		var n C
		count(&n)
		length, valid := sliceLength(io, n)
		if !valid {
			return
		}
		if !reader.SliceLength(uint64(length), maxSliceLength) {
			return
		}
		*x = make([]OrderedEntry[K, V], length)
	} else {
		n, valid := sliceCount[C](io, len(*x))
		if !valid {
			return
		}
		count(&n)
	}
	for i := range *x {
		key(&(*x)[i].Key)
		value(&(*x)[i].Value)
	}
}

func sliceCount[C integer](io IO, length int) (C, bool) {
	var zero C
	max := ^uint64(0)
	switch zero := any(zero).(type) {
	case int:
		max = uint64(^uint(0) >> 1)
	case int8:
		max = uint64(^uint8(0) >> 1)
	case int16:
		max = uint64(^uint16(0) >> 1)
	case int32:
		max = uint64(^uint32(0) >> 1)
	case int64:
		max = uint64(^uint64(0) >> 1)
	case uint:
		max = uint64(^uint(0))
	case uint8:
		max = uint64(^uint8(0))
	case uint16:
		max = uint64(^uint16(0))
	case uint32:
		max = uint64(^uint32(0))
	case uint64:
		max = ^uint64(0)
	case uintptr:
		max = uint64(^uintptr(0))
	default:
		_ = zero
	}
	if uint64(length) > max {
		io.InvalidValue(length, "collection length exceeds its wire count")
		return zero, false
	}
	return C(length), true
}

func sliceLength[C integer](io IO, value C) (int, bool) {
	var n int64
	var unsigned uint64
	switch value := any(value).(type) {
	case int:
		n = int64(value)
	case int8:
		n = int64(value)
	case int16:
		n = int64(value)
	case int32:
		n = int64(value)
	case int64:
		n = value
	case uint:
		unsigned = uint64(value)
	case uint8:
		unsigned = uint64(value)
	case uint16:
		unsigned = uint64(value)
	case uint32:
		unsigned = uint64(value)
	case uint64:
		unsigned = value
	case uintptr:
		unsigned = uint64(value)
	}
	if n < 0 {
		io.InvalidValue(value, "negative collection length")
		return 0, false
	}
	if n > 0 {
		if uint64(n) > uint64(^uint(0)>>1) {
			io.InvalidValue(value, "collection length overflows int")
			return 0, false
		}
		return int(n), true
	}
	if unsigned > uint64(^uint(0)>>1) {
		io.InvalidValue(value, "collection length overflows int")
		return 0, false
	}
	return int(unsigned), true
}
`
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
	if method, ok := semanticIOCall(node); ok {
		fmt.Fprintf(b, "%sio.%s(&%s)\n", indent, method, expression)
		return nil
	}
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
			return e.optionalCall(b, "DoubleOptionalFunc", *value.Value, expression, hint+"Value", indent)
		}
		return e.optionalCall(b, "OptionalFunc", value, expression, hint+"Value", indent)
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

// nodePointer emits the codec for a value supplied as a pointer by one of the
// runtime collection/optional callbacks. Keeping the pointer intact is both
// simpler and consistent with the hand-written gophertunnel codecs: callback
// bodies should marshal value directly instead of copying it to a temporary.
func (e *marshalEmitter) nodePointer(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	if method, ok := semanticIOCall(node); ok {
		fmt.Fprintf(b, "%sio.%s(%s)\n", indent, method, expression)
		return nil
	}
	if native, matched, err := nativeGoType(node); matched || err != nil {
		if err != nil {
			return err
		}
		switch native {
		case "uuid.UUID":
			fmt.Fprintf(b, "%sio.UUID(%s)\n", indent, expression)
		case "mgl32.Vec2":
			fmt.Fprintf(b, "%sio.Vec2(%s)\n", indent, expression)
		case "mgl32.Vec3":
			fmt.Fprintf(b, "%sio.Vec3(%s)\n", indent, expression)
		case "color.RGBA":
			fmt.Fprintf(b, "%sio.RGBA(%s)\n", indent, expression)
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
			fmt.Fprintf(b, "%sio.NBT(%s)\n", indent, expression)
			return nil
		}
		method, err := primitiveIOMethod(node.Primitive.Code)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%sio.%s(%s)\n", indent, method, expression)
		return nil
	case manifest.KindString:
		if !varuint32Prefix(node) {
			return fmt.Errorf("string has unsupported length prefix")
		}
		fmt.Fprintf(b, "%sio.String(%s)\n", indent, expression)
		return nil
	case manifest.KindBytes:
		if !varuint32Prefix(node) {
			return fmt.Errorf("bytes have unsupported length prefix")
		}
		fmt.Fprintf(b, "%sio.Bytes(%s)\n", indent, expression)
		return nil
	case manifest.KindBitset:
		fmt.Fprintf(b, "%sio.Bitset((*%s)[:], %d)\n", indent, expression, node.Length)
		return nil
	case manifest.KindStruct:
		fmt.Fprintf(b, "%s%s.Marshal(io)\n", indent, pointerMethodReceiver(expression))
		return nil
	case manifest.KindRecursive:
		fmt.Fprintf(b, "%smarshal%s(io, %s)\n", indent, e.g.identity[node.Target], expression)
		return nil
	case manifest.KindEnum:
		if node.Primitive == nil {
			return fmt.Errorf("enum has no primitive")
		}
		if len(node.Variants) == 0 {
			return fmt.Errorf("enum has no variants")
		}
		if _, err := primitiveGoType(node.Primitive.Code); err != nil {
			return err
		}
		method, err := primitiveIOMethod(node.Primitive.Code)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%sIntegerFunc(%s, io.%s)\n", indent, expression, method)
		return nil
	case manifest.KindOptional:
		if node.Value == nil {
			return fmt.Errorf("optional has no value")
		}
		value := *node.Value
		if value.Kind == manifest.KindOptional {
			if value.Value == nil {
				return fmt.Errorf("nested optional has no value")
			}
			return e.optionalCallPointer(b, "DoubleOptionalFunc", *value.Value, expression, hint+"Value", indent)
		}
		return e.optionalCallPointer(b, "OptionalFunc", value, expression, hint+"Value", indent)
	case manifest.KindArray:
		if node.Element == nil || node.Prefix == nil {
			return fmt.Errorf("array has no element or prefix")
		}
		return e.collectionPointer(b, *node.Prefix, *node.Element, expression, hint+"Item", indent)
	case manifest.KindFixedArray:
		if node.Element == nil {
			return fmt.Errorf("fixed array has no element")
		}
		index := e.temporary("index")
		element := e.temporary("item")
		fmt.Fprintf(b, "%sfor %s := range *%s {\n", indent, index, expression)
		fmt.Fprintf(b, "%s\t%s := &(*%s)[%s]\n", indent, element, expression, index)
		if err := e.nodePointer(b, *node.Element, element, hint+"Item", indent+"\t"); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil
	case manifest.KindMap:
		if node.Key == nil || node.Value == nil || node.Prefix == nil {
			return fmt.Errorf("map has no key, value, or prefix")
		}
		return e.mapEntriesPointer(b, node, expression, hint, indent)
	case manifest.KindUnion:
		name, err := e.g.goType(node, hint)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%smarshal%s(io, %s)\n", indent, name, expression)
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

func pointerMethodReceiver(expression string) string {
	if expression == "" {
		return expression
	}
	for index, r := range expression {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || (index > 0 && r >= '0' && r <= '9') {
			continue
		}
		return "(" + expression + ")"
	}
	return expression
}

func (e *marshalEmitter) optionalCall(b *strings.Builder, helper string, value manifest.Node, expression, hint, indent string) error {
	fmt.Fprintf(b, "%s%s(io, &%s, ", indent, helper, expression)
	if method, ok := directIOCall(value); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := mustGoType(e.g, value, hint)
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.nodePointer(b, value, "value", hint, indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) optionalCallPointer(b *strings.Builder, helper string, value manifest.Node, expression, hint, indent string) error {
	fmt.Fprintf(b, "%s%s(io, %s, ", indent, helper, expression)
	if method, ok := directIOCall(value); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := mustGoType(e.g, value, hint)
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.nodePointer(b, value, "value", hint, indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) collection(b *strings.Builder, prefix, element manifest.Node, expression, hint, indent string) error {
	if prefix.Kind != manifest.KindPrimitive || prefix.Primitive == nil {
		return fmt.Errorf("collection prefix must be a primitive")
	}
	countMethod, err := primitiveIOMethod(prefix.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sFuncSlice(io, &%s, io.%s, ", indent, expression, countMethod)
	if method, ok := directIOCall(element); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := mustGoType(e.g, element, hint)
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.nodePointer(b, element, "value", hint, indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) collectionPointer(b *strings.Builder, prefix, element manifest.Node, expression, hint, indent string) error {
	if prefix.Kind != manifest.KindPrimitive || prefix.Primitive == nil {
		return fmt.Errorf("collection prefix must be a primitive")
	}
	countMethod, err := primitiveIOMethod(prefix.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sFuncSlice(io, %s, io.%s, ", indent, expression, countMethod)
	if method, ok := directIOCall(element); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := mustGoType(e.g, element, hint)
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.nodePointer(b, element, "value", hint, indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) mapEntries(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	if node.Prefix.Kind != manifest.KindPrimitive || node.Prefix.Primitive == nil {
		return fmt.Errorf("map prefix must be a primitive")
	}
	countMethod, err := primitiveIOMethod(node.Prefix.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sOrderedMap(io, &%s, io.%s, ", indent, expression, countMethod)
	if method, ok := directIOCall(*node.Key); ok {
		fmt.Fprintf(b, "io.%s, ", method)
	} else {
		typ := mustGoType(e.g, *node.Key, hint+"Key")
		fmt.Fprintf(b, "func(value *%s) {\n", typ)
		if err := e.nodePointer(b, *node.Key, "value", hint+"Key", indent+"\t"); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}, ", indent)
	}
	if method, ok := directIOCall(*node.Value); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := mustGoType(e.g, *node.Value, hint+"Value")
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.nodePointer(b, *node.Value, "value", hint+"Value", indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) mapEntriesPointer(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	if node.Prefix.Kind != manifest.KindPrimitive || node.Prefix.Primitive == nil {
		return fmt.Errorf("map prefix must be a primitive")
	}
	countMethod, err := primitiveIOMethod(node.Prefix.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sOrderedMap(io, %s, io.%s, ", indent, expression, countMethod)
	if method, ok := directIOCall(*node.Key); ok {
		fmt.Fprintf(b, "io.%s, ", method)
	} else {
		typ := mustGoType(e.g, *node.Key, hint+"Key")
		fmt.Fprintf(b, "func(value *%s) {\n", typ)
		if err := e.nodePointer(b, *node.Key, "value", hint+"Key", indent+"\t"); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}, ", indent)
	}
	if method, ok := directIOCall(*node.Value); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := mustGoType(e.g, *node.Value, hint+"Value")
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.nodePointer(b, *node.Value, "value", hint+"Value", indent+"\t"); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) enum(b *strings.Builder, node manifest.Node, expression, hint, indent string) error {
	if node.Primitive == nil {
		return fmt.Errorf("enum has no primitive")
	}
	if len(node.Variants) == 0 {
		return fmt.Errorf("enum has no variants")
	}
	if _, err := primitiveGoType(node.Primitive.Code); err != nil {
		return err
	}
	method, err := primitiveIOMethod(node.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%sIntegerFunc(&%s, io.%s)\n", indent, expression, method)
	return nil
}

func directIOCall(node manifest.Node) (string, bool) {
	if method, ok := semanticIOCall(node); ok {
		return method, true
	}
	if native, matched, err := nativeGoType(node); matched && err == nil {
		switch native {
		case "uuid.UUID":
			return "UUID", true
		case "mgl32.Vec2":
			return "Vec2", true
		case "mgl32.Vec3":
			return "Vec3", true
		case "color.RGBA":
			return "RGBA", true
		}
	}
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return "", false
		}
		if node.Primitive.Code == "nbt_le" {
			return "NBT", true
		}
		method, err := primitiveIOMethod(node.Primitive.Code)
		return method, err == nil
	case manifest.KindString:
		if varuint32Prefix(node) {
			return "String", true
		}
	case manifest.KindBytes:
		if varuint32Prefix(node) {
			return "Bytes", true
		}
	}
	return "", false
}

func semanticIOCall(node manifest.Node) (string, bool) {
	native, matched, err := nativeGoType(node)
	if err != nil || !matched {
		return "", false
	}
	code, hasCode := primitiveStructCode(node)
	switch node.TypeID {
	case "ActorRuntimeID":
		switch {
		case native == "uint64" && hasCode && code == "var_u64":
			return "ActorRuntimeID", true
		case native == "int64" && hasCode && code == "zigzag_i64":
			return "ActorRuntimeIDVarint64", true
		case native == "uint32" && hasCode && code == "var_u32":
			return "ActorRuntimeIDVaruint32", true
		case native == "int64" && hasCode && code == "var_i64":
			return "SignedVarint64", true
		}
	case "ActorUniqueID":
		switch {
		case native == "int64" && hasCode && code == "zigzag_i64":
			return "ActorUniqueID", true
		case native == "int64" && hasCode && code == "i64le":
			return "ActorUniqueIDInt64", true
		case native == "uint64" && hasCode && code == "u64le":
			return "ActorUniqueIDUint64", true
		case native == "uint64" && hasCode && code == "var_u64":
			return "ActorUniqueIDVaruint64", true
		case native == "int64" && hasCode && code == "var_i64":
			return "SignedVarint64", true
		}
	case "PlayerInputTick":
		if native == "uint64" {
			return "PlayerInputTick", true
		}
	default:
		return "", false
	}
	return "", false
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
	if definition.Underlying == "" {
		return fmt.Errorf("union %s has no discriminator primitive", definition.Name)
	}
	controlNode := manifest.Primitive(definition.Underlying)
	tagType, err := primitiveGoType(definition.Underlying)
	if err != nil {
		return err
	}
	b.WriteString("\tUnionFunc(io,\n\t\tfunc() {\n")
	fmt.Fprintf(b, "\t\tvar tag %s\n", tagType)
	if err := e.node(b, controlNode, "tag", definition.Name+"Tag", "\t\t"); err != nil {
		return err
	}
	b.WriteString("\t\tswitch int64(tag) {\n")
	for _, member := range definition.Union {
		fmt.Fprintf(b, "\t\tcase %d:\n\t\t\tvar value %s\n\t\t\tvalue.Marshal(io)\n\t\t\t*x = value\n", member.Value, member.Name)
	}
	b.WriteString("\t\tdefault:\n\t\t\tio.InvalidValue(tag, \"unknown union tag\")\n\t\t}\n\t\t},\n\t\tfunc() {\n\t\tswitch value := (*x).(type) {\n")
	for _, member := range definition.Union {
		fmt.Fprintf(b, "\t\tcase %s:\n\t\t\ttag := %s(%d)\n", member.Name, tagType, member.Value)
		if err := e.node(b, controlNode, "tag", definition.Name+"Tag", "\t\t\t"); err != nil {
			return err
		}
		b.WriteString("\t\t\tvalue.Marshal(io)\n")
	}
	b.WriteString("\t\tdefault:\n\t\t\tio.InvalidValue(*x, \"unknown union value\")\n\t\t}\n\t\t},\n\t)\n}\n\n")
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

func enumVariantName(value string) string {
	if value == "" {
		return "Unknown"
	}
	allUpper := true
	for _, r := range value {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			allUpper = false
			break
		}
	}
	if !allUpper {
		return normalizeEnumInitialisms(exportName(value))
	}
	var b strings.Builder
	for _, token := range strings.FieldsFunc(value, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }) {
		if token == "" {
			continue
		}
		if initialism, ok := enumInitialisms[token]; ok {
			b.WriteString(initialism)
			continue
		}
		lower := strings.ToLower(token)
		b.WriteString(exportName(lower))
	}
	if b.Len() == 0 {
		return "Unknown"
	}
	return b.String()
}

func normalizeEnumInitialisms(value string) string {
	for _, replacement := range []struct{ from, to string }{
		{from: "Tntcart", to: "TNTCart"},
		{from: "Fishpos", to: "FishPosition"},
		{from: "Hooktime", to: "HookTime"},
		{from: "Tnt", to: "TNT"},
		{from: "Nbt", to: "NBT"},
		{from: "Uuid", to: "UUID"},
		{from: "Argb", to: "ARGB"},
		{from: "Rgba", to: "RGBA"},
		{from: "Rgb", to: "RGB"},
		{from: "Uwp", to: "UWP"},
		{from: "Osx", to: "OSX"},
	} {
		value = strings.ReplaceAll(value, replacement.from, replacement.to)
	}
	return value
}

var enumInitialisms = map[string]string{
	"ANIM": "Animation", "FISHPOS": "FishPosition", "HOOKTIME": "HookTime", "ID": "ID", "NBT": "NBT", "OSX": "OSX", "RGBA": "RGBA", "RGB": "RGB", "TNT": "TNT", "TNTCART": "TNTCart", "UWP": "UWP", "URL": "URL", "URI": "URI", "UUID": "UUID", "X": "X", "Y": "Y", "Z": "Z",
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
