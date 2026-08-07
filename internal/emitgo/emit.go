// Package emitgo emits a standalone Go view from the canonical v2 manifest.
// It receives no source documents; target-specific conveniences are selected
// explicitly through Options.
package emitgo

import (
	"embed"
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
	definitions        map[string]typeDefinition
	identity           map[string]string
	usedNames          map[string]bool
	protocolImportPath string
	nativeTypes        bool
	emitPacketRuntime  bool
	emitPacketPools    bool
}

//go:embed runtime/codec.go runtime/helpers.go runtime/reader.go runtime/types.go runtime/writer.go
var runtimeSource embed.FS

func emitRuntimeFiles() (map[string]string, error) {
	files := make(map[string]string, 5)
	for _, name := range []string{"codec.go", "helpers.go", "reader.go", "types.go", "writer.go"} {
		data, err := runtimeSource.ReadFile("runtime/" + name)
		if err != nil {
			return nil, fmt.Errorf("read embedded runtime %s: %w", name, err)
		}
		files["protocol/"+name] = string(data)
	}
	return files, nil
}

// Options controls the target-specific conveniences emitted alongside the
// canonical wire representation. NativeTypes enables established Go types
// such as uuid.UUID and mgl32 vectors; disabling it keeps every named shape in
// generated protocol structs. Packet runtime and pool emission are enabled by
// Generate and can be disabled by callers that only need definitions.
type Options struct {
	ProtocolImportPath string
	NativeTypes        bool
	EmitPacketRuntime  bool
	EmitPacketPools    bool
}

// Generate emits a protocol package and its packet subpackage. The protocol
// import path is embedded in packet files so the generated tree can be used as
// an ordinary Go package from any parent module.
func Generate(m manifest.Manifest, protocolImportPath string) (map[string]string, error) {
	return GenerateWithOptions(m, Options{
		ProtocolImportPath: protocolImportPath,
		NativeTypes:        true,
		EmitPacketRuntime:  true,
		EmitPacketPools:    true,
	})
}

// GenerateWithOptions emits a protocol package using explicit target options.
// The canonical manifest remains the source of truth; options only affect the
// generated API surface and Go type mappings.
func GenerateWithOptions(m manifest.Manifest, options Options) (map[string]string, error) {
	if err := manifest.Validate(m); err != nil {
		return nil, err
	}
	if options.ProtocolImportPath == "" || strings.ContainsAny(options.ProtocolImportPath, " \t\r\n") {
		return nil, fmt.Errorf("invalid protocol import path %q", options.ProtocolImportPath)
	}
	g := &generator{
		definitions:        map[string]typeDefinition{},
		identity:           map[string]string{},
		usedNames:          map[string]bool{},
		protocolImportPath: options.ProtocolImportPath,
		nativeTypes:        options.NativeTypes,
		emitPacketRuntime:  options.EmitPacketRuntime,
		emitPacketPools:    options.EmitPacketPools,
	}
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
	files, err := g.emitFiles(packets, packetNames)
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
	if typ, matched, err := g.nativeGoType(node); matched || err != nil {
		return typ, err
	}
	switch node.Kind {
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return "", fmt.Errorf("primitive has no shape")
		}
		if node.Primitive.Code == "uuid" {
			return "[16]byte", nil
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
	return nativeGoTypeWithOptions(node, true)
}

func (g *generator) nativeGoType(node manifest.Node) (string, bool, error) {
	return nativeGoTypeWithOptions(node, g.nativeTypes)
}

func nativeGoTypeWithOptions(node manifest.Node, enabled bool) (string, bool, error) {
	if !enabled {
		return "", false, nil
	}
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

func (g *generator) emitFiles(packets []manifest.Packet, packetNames map[uint32]string) (map[string]string, error) {
	definitions := make([]typeDefinition, 0, len(g.definitions))
	for _, definition := range g.definitions {
		definitions = append(definitions, definition)
	}
	sort.Slice(definitions, func(i, j int) bool { return definitions[i].Name < definitions[j].Name })

	files, err := emitRuntimeFiles()
	if err != nil {
		return nil, err
	}
	usedFiles := map[string]bool{
		"types.go": true, "codec.go": true, "helpers.go": true, "reader.go": true, "writer.go": true,
	}
	for _, definition := range definitions {
		source, err := emitDefinition(g, definition)
		if err != nil {
			return nil, err
		}
		name := uniqueFileName(snakeName(definition.Name)+".go", 0, usedFiles)
		files["protocol/"+name] = source
	}
	packetFiles := map[string]string{}
	packetUsed := map[string]bool{"ids.go": true}
	if g.emitPacketRuntime {
		packetUsed["packet.go"] = true
	}
	if g.emitPacketPools {
		packetUsed["pool.go"] = true
	}
	packetFiles["ids.go"] = emitPacketIDs(packets, packetNames)
	if g.emitPacketRuntime {
		packetFiles["packet.go"] = emitPacketRuntime(g.protocolImportPath)
	}
	if g.emitPacketPools {
		packetFiles["pool.go"] = emitPacketPools(packets, packetNames)
	}
	for _, packet := range packets {
		packetName := packetNames[packet.ID]
		base := snakeName(packetName) + ".go"
		name := uniqueFileName(base, packet.ID, packetUsed)
		source, err := g.emitPacket(packet, packetName)
		if err != nil {
			return nil, err
		}
		packetFiles[name] = source
	}
	for name, source := range packetFiles {
		files["protocol/packet/"+name] = source
	}
	return files, nil
}

func emitDefinition(g *generator, definition typeDefinition) (string, error) {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage protocol\n\n")
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
			if err := emitter.node(&b, field.Node, "x."+field.Name, definition.Name+field.Name, "\t", addressStrategy{}); err != nil {
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
		fmt.Fprintf(&b, "\nconst %sLength = %d\n\n", definition.Name, definition.BitLength)
		fmt.Fprintf(&b, "// Set marks bit index i. It panics when i is outside [0, %d).\n", definition.BitLength)
		fmt.Fprintf(&b, "func (b *%s) Set(i int) {\n", definition.Name)
		fmt.Fprintf(&b, "\tif i < 0 || i >= %sLength {\n\t\tpanic(\"index out of bounds\")\n\t}\n", definition.Name)
		fmt.Fprintf(&b, "\tb[i/64] |= uint64(1) << uint(i%%64)\n}\n\n")
		fmt.Fprintf(&b, "// Unset clears bit index i. It panics when i is outside [0, %d).\n", definition.BitLength)
		fmt.Fprintf(&b, "func (b *%s) Unset(i int) {\n", definition.Name)
		fmt.Fprintf(&b, "\tif i < 0 || i >= %sLength {\n\t\tpanic(\"index out of bounds\")\n\t}\n", definition.Name)
		fmt.Fprintf(&b, "\tb[i/64] &^= uint64(1) << uint(i%%64)\n}\n\n")
		fmt.Fprintf(&b, "// Load reports whether bit index i is set. It panics when i is outside [0, %d).\n", definition.BitLength)
		fmt.Fprintf(&b, "func (b %s) Load(i int) bool {\n", definition.Name)
		fmt.Fprintf(&b, "\tif i < 0 || i >= %sLength {\n\t\tpanic(\"index out of bounds\")\n\t}\n", definition.Name)
		fmt.Fprintf(&b, "\treturn b[i/64]&(uint64(1)<<uint(i%%64)) != 0\n}\n\n")
		fmt.Fprintf(&b, "// Len returns the number of bits in the bitset.\nfunc (b %s) Len() int { return %sLength }\n", definition.Name, definition.Name)
	default:
		return "", fmt.Errorf("unsupported definition kind %q", definition.Kind)
	}
	return formatGoSource(b.String())
}

type packetField struct {
	name string
	typ  string
	node manifest.Node
}

func (g *generator) emitPacket(packet manifest.Packet, packetName string) (string, error) {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage packet\n\n")
	used := map[string]bool{}
	fields := make([]packetField, 0, len(packet.Fields))
	for _, field := range packet.Fields {
		baseName := exportName(field.Name)
		if g.emitPacketRuntime && baseName == "ID" {
			// ID is reserved by the generated packet runtime method. Keep the
			// wire field explicit without making the struct fail to compile.
			baseName = "IDValue"
		}
		name := uniqueFieldName(baseName, used)
		typ := mustGoType(g, field.Encode, packetName+name)
		fields = append(fields, packetField{name: name, typ: qualifyGoType(typ, g.definitions), node: field.Encode})
	}
	imports := append([]string{g.protocolImportPath}, goImportsForFields(fields)...)
	writeGoImports(&b, imports)
	fmt.Fprintf(&b, "type %s struct {\n", packetName)
	for _, field := range fields {
		fmt.Fprintf(&b, "\t%s %s\n", field.name, field.typ)
	}
	b.WriteString("}\n\n")
	fmt.Fprintf(&b, "// Marshal reads or writes %s using its canonical wire layout.\n", packetName)
	fmt.Fprintf(&b, "func (x *%s) Marshal(io protocol.IO) {\n", packetName)
	emitter := marshalEmitter{g: g, qualifier: "protocol."}
	for _, field := range fields {
		if err := emitter.node(&b, field.node, "x."+field.name, packetName+field.name, "\t", addressStrategy{}); err != nil {
			return "", fmt.Errorf("packet %s field %s marshal: %w", packet.Name, field.name, err)
		}
	}
	b.WriteString("}\n")
	if g.emitPacketRuntime {
		fmt.Fprintf(&b, "\n// ID returns the protocol ID for %s.\nfunc (*%s) ID() uint32 { return ID%s }\n", packetName, packetName, packetName)
	}
	return formatGoSource(b.String())
}

func emitPacketIDs(packets []manifest.Packet, packetNames map[uint32]string) string {
	var b strings.Builder
	b.WriteString("// Code generated from canonical protocol manifest v2. DO NOT EDIT.\n\npackage packet\n\nconst (\n")
	for _, packet := range packets {
		fmt.Fprintf(&b, "\tID%s uint32 = %d\n", packetNames[packet.ID], packet.ID)
	}
	b.WriteString(")\n")
	return b.String()
}

func emitPacketRuntime(protocolImportPath string) string {
	return mustFormatGoSource(fmt.Sprintf(`// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import %q

// Packet is the common runtime contract for every generated Bedrock packet.
// Marshal reads from or writes to the supplied protocol IO implementation.
type Packet interface {
	ID() uint32
	Marshal(protocol.IO)
}
`, protocolImportPath))
}

func emitPacketPools(packets []manifest.Packet, packetNames map[uint32]string) string {
	var b strings.Builder
	b.WriteString(`// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

// Factory constructs a fresh packet value for decoding.
type Factory func() Packet

// Pool maps protocol IDs to packet constructors.
type Pool map[uint32]Factory

`)
	emit := func(name string, include func(manifest.Direction) bool) {
		fmt.Fprintf(&b, "var %s = Pool{\n", name)
		for _, packet := range packets {
			if !include(packet.Direction) {
				continue
			}
			packetName := packetNames[packet.ID]
			fmt.Fprintf(&b, "\tID%s: func() Packet { return &%s{} },\n", packetName, packetName)
		}
		b.WriteString("}\n\n")
	}
	emit("allPacketFactories", func(manifest.Direction) bool { return true })
	emit("clientPacketFactories", func(direction manifest.Direction) bool {
		return direction == manifest.DirectionServerbound || direction == manifest.DirectionBidirectional
	})
	emit("serverPacketFactories", func(direction manifest.Direction) bool {
		return direction == manifest.DirectionClientbound || direction == manifest.DirectionBidirectional
	})
	b.WriteString(`// NewPool returns a copy of the complete packet factory table.
func NewPool() Pool { return clonePool(allPacketFactories) }

// NewClientPool returns factories for packets sent by a client, including
// packets explicitly marked bidirectional in the manifest.
func NewClientPool() Pool { return clonePool(clientPacketFactories) }

// NewServerPool returns factories for packets sent by a server, including
// packets explicitly marked bidirectional in the manifest.
func NewServerPool() Pool { return clonePool(serverPacketFactories) }

// NewPacket constructs the packet with id, if it is known.
func NewPacket(id uint32) (Packet, bool) { return newFromPool(allPacketFactories, id) }

// NewClientPacket constructs a known packet sent by a client.
func NewClientPacket(id uint32) (Packet, bool) { return newFromPool(clientPacketFactories, id) }

// NewServerPacket constructs a known packet sent by a server.
func NewServerPacket(id uint32) (Packet, bool) { return newFromPool(serverPacketFactories, id) }

func clonePool(source Pool) Pool {
	copy := make(Pool, len(source))
	for id, factory := range source { copy[id] = factory }
	return copy
}

func newFromPool(pool Pool, id uint32) (Packet, bool) {
	factory, ok := pool[id]
	if !ok { return nil, false }
	return factory(), true
}
`)
	return mustFormatGoSource(b.String())
}

func mustFormatGoSource(source string) string {
	formatted, err := formatGoSource(source)
	if err != nil {
		panic(err)
	}
	return formatted
}

type marshalEmitter struct {
	g         *generator
	qualifier string
	counter   int
}

func (e *marshalEmitter) mustGoType(node manifest.Node, hint string) string {
	typ := mustGoType(e.g, node, hint)
	if e.qualifier == "" {
		return typ
	}
	return qualifyGoType(typ, e.g.definitions)
}

func (e *marshalEmitter) runtime(name string) string {
	return e.qualifier + name
}

func (e *marshalEmitter) ioType() string {
	return e.qualifier + "IO"
}

func (e *marshalEmitter) temporary(prefix string) string {
	e.counter++
	return fmt.Sprintf("%s%d", prefix, e.counter)
}

type addressStrategy struct {
	pointer bool
}

func (a addressStrategy) address(expression string) string {
	if a.pointer {
		return expression
	}
	return "&" + expression
}

func (a addressStrategy) container(expression string) string {
	if a.pointer {
		return "*" + expression
	}
	return expression
}

func (a addressStrategy) element(expression, index string) string {
	if a.pointer {
		return "&(*" + expression + ")[" + index + "]"
	}
	return expression + "[" + index + "]"
}

func (a addressStrategy) bits(expression string) string {
	if a.pointer {
		return "(*" + expression + ")[:]"
	}
	return expression + "[:]"
}

func (e *marshalEmitter) node(b *strings.Builder, node manifest.Node, expression, hint, indent string, address addressStrategy) error {
	if method, ok := e.semanticIOCall(node); ok {
		fmt.Fprintf(b, "%sio.%s(%s)\n", indent, method, address.address(expression))
		return nil
	}
	if native, matched, err := e.g.nativeGoType(node); matched || err != nil {
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%sio.%s(%s)\n", indent, nativeMethod(native), address.address(expression))
		return nil
	}
	switch node.Kind {
	case manifest.KindVoid:
		return nil
	case manifest.KindPrimitive:
		if node.Primitive == nil {
			return fmt.Errorf("primitive has no shape")
		}
		if node.Primitive.Code == "uuid" {
			fmt.Fprintf(b, "%sio.UUIDBytes(%s)\n", indent, address.address(expression))
			return nil
		}
		if node.Primitive.Code == "nbt_le" {
			fmt.Fprintf(b, "%sio.NBT(%s)\n", indent, address.address(expression))
			return nil
		}
		method, err := primitiveIOMethod(node.Primitive.Code)
		if err != nil {
			return err
		}
		fmt.Fprintf(b, "%sio.%s(%s)\n", indent, method, address.address(expression))
		return nil
	case manifest.KindString:
		if !varuint32Prefix(node) {
			return fmt.Errorf("string has unsupported length prefix")
		}
		fmt.Fprintf(b, "%sio.String(%s)\n", indent, address.address(expression))
		return nil
	case manifest.KindBytes:
		if !varuint32Prefix(node) {
			return fmt.Errorf("bytes have unsupported length prefix")
		}
		fmt.Fprintf(b, "%sio.Bytes(%s)\n", indent, address.address(expression))
		return nil
	case manifest.KindBitset:
		fmt.Fprintf(b, "%sio.Bitset(%s, %d)\n", indent, address.bits(expression), node.Length)
		return nil
	case manifest.KindStruct:
		fmt.Fprintf(b, "%s%s.Marshal(io)\n", indent, expression)
		return nil
	case manifest.KindRecursive:
		fmt.Fprintf(b, "%s%s(io, %s)\n", indent, e.runtime("Marshal"+e.g.identity[node.Target]), address.address(expression))
		return nil
	case manifest.KindEnum:
		return e.enum(b, node, expression, hint, indent, address)
	case manifest.KindOptional:
		if node.Value == nil {
			return fmt.Errorf("optional has no value")
		}
		value := *node.Value
		if value.Kind == manifest.KindOptional {
			if value.Value == nil {
				return fmt.Errorf("nested optional has no value")
			}
			return e.optionalCall(b, "DoubleOptionalFunc", *value.Value, expression, hint+"Value", indent, address)
		}
		return e.optionalCall(b, "OptionalFunc", value, expression, hint+"Value", indent, address)
	case manifest.KindArray:
		if node.Element == nil || node.Prefix == nil {
			return fmt.Errorf("array has no element or prefix")
		}
		return e.collection(b, *node.Prefix, *node.Element, expression, hint+"Item", indent, address)
	case manifest.KindFixedArray:
		if node.Element == nil {
			return fmt.Errorf("fixed array has no element")
		}
		index := e.temporary("index")
		fmt.Fprintf(b, "%sfor %s := range %s {\n", indent, index, address.container(expression))
		if err := e.node(b, *node.Element, address.element(expression, index), hint+"Item", indent+"\t", address); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}\n", indent)
		return nil
	case manifest.KindMap:
		if node.Key == nil || node.Value == nil || node.Prefix == nil {
			return fmt.Errorf("map has no key, value, or prefix")
		}
		return e.mapEntries(b, node, expression, hint, indent, address)
	case manifest.KindUnion:
		name := e.mustGoType(node, hint)
		fmt.Fprintf(b, "%s%s(io, %s)\n", indent, e.runtime("Marshal"+name), address.address(expression))
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

func nativeMethod(native string) string {
	switch native {
	case "uuid.UUID":
		return "UUID"
	case "mgl32.Vec2":
		return "Vec2"
	case "mgl32.Vec3":
		return "Vec3"
	case "color.RGBA":
		return "RGBA"
	default:
		panic("native type " + native + " has no codec operation")
	}
}

func (e *marshalEmitter) optionalCall(b *strings.Builder, helper string, value manifest.Node, expression, hint, indent string, address addressStrategy) error {
	fmt.Fprintf(b, "%s%s(io, %s, ", indent, e.runtime(helper), address.address(expression))
	if method, ok := e.directIOCall(value); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := e.mustGoType(value, hint)
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.node(b, value, "value", hint, indent+"\t", addressStrategy{pointer: true}); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) collection(b *strings.Builder, prefix, element manifest.Node, expression, hint, indent string, address addressStrategy) error {
	if prefix.Kind != manifest.KindPrimitive || prefix.Primitive == nil {
		return fmt.Errorf("collection prefix must be a primitive")
	}
	countMethod, err := primitiveIOMethod(prefix.Primitive.Code)
	if err != nil {
		return err
	}
	if prefix.Primitive.Code == "var_u32" && e.marshalableElement(element, hint) {
		fmt.Fprintf(b, "%s%s(io, %s)\n", indent, e.runtime("Slice"), address.address(expression))
		return nil
	}
	fmt.Fprintf(b, "%s%s(io, %s, io.%s, ", indent, e.runtime("FuncSlice"), address.address(expression), countMethod)
	if method, ok := e.directIOCall(element); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := e.mustGoType(element, hint)
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.node(b, element, "value", hint, indent+"\t", addressStrategy{pointer: true}); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) marshalableElement(element manifest.Node, hint string) bool {
	if element.Kind != manifest.KindStruct && element.Kind != manifest.KindRecursive {
		return false
	}
	typ := mustGoType(e.g, element, hint)
	definition, ok := e.g.definitions[typ]
	return ok && definition.Kind == manifest.KindStruct
}

func (e *marshalEmitter) mapEntries(b *strings.Builder, node manifest.Node, expression, hint, indent string, address addressStrategy) error {
	if node.Prefix.Kind != manifest.KindPrimitive || node.Prefix.Primitive == nil {
		return fmt.Errorf("map prefix must be a primitive")
	}
	countMethod, err := primitiveIOMethod(node.Prefix.Primitive.Code)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "%s%s(io, %s, io.%s, ", indent, e.runtime("OrderedMap"), address.address(expression), countMethod)
	if method, ok := e.directIOCall(*node.Key); ok {
		fmt.Fprintf(b, "io.%s, ", method)
	} else {
		typ := e.mustGoType(*node.Key, hint+"Key")
		fmt.Fprintf(b, "func(value *%s) {\n", typ)
		if err := e.node(b, *node.Key, "value", hint+"Key", indent+"\t", addressStrategy{pointer: true}); err != nil {
			return err
		}
		fmt.Fprintf(b, "%s}, ", indent)
	}
	if method, ok := e.directIOCall(*node.Value); ok {
		fmt.Fprintf(b, "io.%s)\n", method)
		return nil
	}
	typ := e.mustGoType(*node.Value, hint+"Value")
	fmt.Fprintf(b, "func(value *%s) {\n", typ)
	if err := e.node(b, *node.Value, "value", hint+"Value", indent+"\t", addressStrategy{pointer: true}); err != nil {
		return err
	}
	fmt.Fprintf(b, "%s})\n", indent)
	return nil
}

func (e *marshalEmitter) enum(b *strings.Builder, node manifest.Node, expression, hint, indent string, address addressStrategy) error {
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
	fmt.Fprintf(b, "%s%s(%s, io.%s)\n", indent, e.runtime("IntegerFunc"), address.address(expression), method)
	return nil
}

func (e *marshalEmitter) directIOCall(node manifest.Node) (string, bool) {
	if method, ok := e.semanticIOCall(node); ok {
		return method, true
	}
	if native, matched, err := e.g.nativeGoType(node); matched && err == nil {
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
		if node.Primitive.Code == "uuid" {
			return "UUIDBytes", true
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

func (e *marshalEmitter) semanticIOCall(node manifest.Node) (string, bool) {
	native, matched, err := e.g.nativeGoType(node)
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

// semanticIOCall retains the package-local helper used by emitter tests and
// older integrations; generated code uses the option-aware method above.
func semanticIOCall(node manifest.Node) (string, bool) {
	return (&marshalEmitter{g: &generator{nativeTypes: true}}).semanticIOCall(node)
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
	fmt.Fprintf(b, "// Marshal%s reads or writes the %s union using its canonical wire layout.\n", definition.Name, definition.Name)
	fmt.Fprintf(b, "func Marshal%s(io %s, x *%s) {\n", definition.Name, e.ioType(), definition.Name)
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
	fmt.Fprintf(b, "\t%s(io,\n\t\tfunc() {\n", e.runtime("UnionFunc"))
	fmt.Fprintf(b, "\t\tvar tag %s\n", tagType)
	if err := e.node(b, controlNode, "tag", definition.Name+"Tag", "\t\t", addressStrategy{}); err != nil {
		return err
	}
	b.WriteString("\t\tswitch int64(tag) {\n")
	for _, member := range definition.Union {
		fmt.Fprintf(b, "\t\tcase %d:\n\t\t\tvar value %s\n\t\t\tvalue.Marshal(io)\n\t\t\t*x = value\n", member.Value, member.Name)
	}
	b.WriteString("\t\tdefault:\n\t\t\tio.InvalidValue(tag, \"unknown union tag\")\n\t\t}\n\t\t},\n\t\tfunc() {\n\t\tswitch value := (*x).(type) {\n")
	for _, member := range definition.Union {
		fmt.Fprintf(b, "\t\tcase %s:\n\t\t\ttag := %s(%d)\n", member.Name, tagType, member.Value)
		if err := e.node(b, controlNode, "tag", definition.Name+"Tag", "\t\t\t", addressStrategy{}); err != nil {
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

func qualifyGoType(typ string, definitions map[string]typeDefinition) string {
	var b strings.Builder
	for index := 0; index < len(typ); {
		if !isGoIdentifierByte(typ[index]) {
			b.WriteByte(typ[index])
			index++
			continue
		}
		end := index + 1
		for end < len(typ) && isGoIdentifierByte(typ[end]) {
			end++
		}
		name := typ[index:end]
		_, generated := definitions[name]
		if (generated || name == "Optional" || name == "OrderedEntry") && (index == 0 || typ[index-1] != '.') {
			b.WriteString("protocol.")
		}
		b.WriteString(name)
		index = end
	}
	return b.String()
}

func isGoIdentifierByte(value byte) bool {
	return value == '_' || value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9'
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
	case "uuid":
		return "[16]byte", nil
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
