package emitgo

import (
	"go/format"
	"os"
	"strings"
	"testing"

	"protocolgen/internal/docs"
	"protocolgen/internal/domains"
	"protocolgen/internal/manifest"
)

func TestMarshalEmitterUsesOneAddressParameterizedWalk(t *testing.T) {
	source, err := os.ReadFile("emit.go")
	if err != nil {
		t.Fatal(err)
	}
	text := string(source)
	if strings.Count(text, "func (e *marshalEmitter) node(") != 1 {
		t.Fatalf("marshal emitter has more than one node walk")
	}
	for _, duplicate := range []string{"nodePointer", "optionalCallPointer", "collectionPointer", "mapEntriesPointer"} {
		if strings.Contains(text, duplicate) {
			t.Fatalf("marshal emitter retains duplicate helper %q", duplicate)
		}
	}
}

func TestGenerateConsumesOnlyCanonicalManifest(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture-2168", Digest: "fixture:golden", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets: []manifest.Packet{{ID: 1, Name: "VocabularyPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Value", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 1, Name: "Maybe", Encode: manifest.Optional(manifest.String(manifest.Primitive("var_u32"))), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	for name, source := range files {
		if _, err := format.Source([]byte(source)); err != nil {
			t.Fatalf("%s is not valid Go: %v\n%s", name, err, source)
		}
		if strings.Contains(source, "gophertunnel") || strings.Contains(source, "bedrock-protocol-docs") {
			t.Errorf("%s contains source/profile lookup text", name)
		}
	}
	packet := files["protocol/packet/vocabulary.go"]
	version := files["protocol/version.go"]
	if !strings.Contains(version, `GAME_VERSION     = "fixture"`) || !strings.Contains(version, "PROTOCOL_VERSION = 2168") {
		t.Fatalf("generated version constants are incomplete:\n%s", version)
	}
	if !strings.Contains(packet, "type Vocabulary struct") || !strings.Contains(packet, "Maybe protocol.Optional[string]") || !strings.Contains(files["protocol/packet/ids.go"], "IDVocabulary uint32 = 1") {
		t.Fatalf("generated output omitted packet definition or ID:\n%s\n%s", packet, files["protocol/packet/ids.go"])
	}
	for _, want := range []string{"func (x *Vocabulary) Marshal(io protocol.IO)", "io.Uint8(&x.Value)", "protocol.OptionalFunc(io, &x.Maybe, io.String)"} {
		if !strings.Contains(packet, want) {
			t.Fatalf("generated packet marshal omits %q:\n%s", want, packet)
		}
	}
	if strings.Contains(packet, "Shape{") || strings.Contains(packet, "func (p *Vocabulary) Encode") {
		t.Fatalf("generated packet exposed runtime schema or placeholder codecs:\n%s", packet)
	}
}

func TestGenerateGroupsSharedDefinitionsByReviewedDomain(t *testing.T) {
	alpha := manifest.Node{Kind: manifest.KindStruct, TypeID: "Alpha", Fields: []manifest.Field{{Name: "Value", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	beta := manifest.Node{Kind: manifest.KindStruct, TypeID: "Beta", Fields: []manifest.Field{{Name: "Value", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets: []manifest.Packet{{ID: 1, Name: "FixturePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Beta", Encode: beta, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 1, Name: "Alpha", Encode: alpha, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
	files, err := GenerateWithOptions(m, Options{
		ProtocolImportPath: "fixture",
		NativeTypes:        false,
		EmitPacketRuntime:  true,
		EmitPacketPools:    true,
		Domains: domains.Overlay{Domains: map[string]string{
			"Alpha": "shared",
			"Beta":  "shared",
		}},
	})
	if err != nil {
		t.Fatalf("GenerateWithOptions: %v", err)
	}
	source := files["protocol/shared.go"]
	if !strings.Contains(source, "type Alpha struct") || !strings.Contains(source, "type Beta struct") {
		t.Fatalf("shared domain file omitted definitions:\n%s", source)
	}
	if strings.Index(source, "type Alpha struct") > strings.Index(source, "type Beta struct") {
		t.Fatalf("definitions are not sorted by name within domain:\n%s", source)
	}
	if _, ok := files["protocol/alpha.go"]; ok {
		t.Fatalf("domain grouping retained per-type file")
	}
}

func TestGenerateEmitsReviewedGoDocs(t *testing.T) {
	shared := manifest.Node{Kind: manifest.KindStruct, TypeID: "Shared", Fields: []manifest.Field{{Name: "Value", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets:       []manifest.Packet{{ID: 1, Name: "FixturePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Name: "Shared Value", Encode: shared, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	files, err := GenerateWithOptions(m, Options{ProtocolImportPath: "fixture", NativeTypes: false, Docs: docs.Overlay{
		Types:  map[string]string{"Shared": "Shared docs.", "FixturePacket": "Packet docs."},
		Fields: map[string]string{docs.FieldKey("Shared", "Value"): "Value docs.", docs.FieldKey("FixturePacket", "Shared Value"): "Packet value docs."},
	}})
	if err != nil {
		t.Fatalf("GenerateWithOptions: %v", err)
	}
	if !strings.Contains(files["protocol/shared.go"], "// Shared docs.") || !strings.Contains(files["protocol/shared.go"], "// Value docs.") {
		t.Fatalf("shared docs were not emitted:\n%s", files["protocol/shared.go"])
	}
	if !strings.Contains(files["protocol/packet/fixture.go"], "// Packet docs.") || !strings.Contains(files["protocol/packet/fixture.go"], "// Packet value docs.") {
		t.Fatalf("packet docs were not emitted:\n%s", files["protocol/packet/fixture.go"])
	}
}

func TestGenerateUsesRuntimeHelpersForEnumsAndOptionals(t *testing.T) {
	value := manifest.Enum("u8", manifest.EnumValue{Name: "Zero", Value: 0}, manifest.EnumValue{Name: "One", Value: 1})
	entry := manifest.Struct(manifest.Field{
		Ordinal:    0,
		Name:       "Value",
		Encode:     manifest.Primitive("u8"),
		Symmetry:   manifest.Symmetric,
		Provenance: manifest.Provenance{Pins: []string{"fixture"}},
	})
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:helpers", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets: []manifest.Packet{{ID: 1, Name: "HelperPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Kind", Encode: value, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 1, Name: "Maybe", Encode: manifest.Optional(manifest.Primitive("i32le")), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 2, Name: "Kinds", Encode: manifest.Array(manifest.Primitive("var_u32"), value), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 3, Name: "Entries", Encode: manifest.Array(manifest.Primitive("var_u32"), entry), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 4, Name: "MaybeKinds", Encode: manifest.Optional(manifest.Array(manifest.Primitive("var_u32"), value)), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	source := files["protocol/packet/helper.go"]
	for _, want := range []string{
		"protocol.IntegerFunc(&x.Kind, io.Uint8)",
		"protocol.OptionalFunc(io, &x.Maybe, io.Int32)",
		"protocol.IntegerFunc(value, io.Uint8)",
		"protocol.Slice(io, &x.Entries)",
	} {
		if !strings.Contains(source, want) {
			t.Fatalf("generated packet omitted runtime helper %q:\n%s", want, source)
		}
	}
	if strings.Contains(source, "io.Reading()") || strings.Contains(source, "unknown enum value") {
		t.Fatalf("ordinary enum/optional mechanics leaked into packet code:\n%s", source)
	}
	if strings.Contains(source, "item := *value") || strings.Contains(source, "*value = item") {
		t.Fatalf("callback codec copied values instead of using the supplied pointer:\n%s", source)
	}
	if !strings.Contains(source, "FuncSlice(io, value, io.Varuint32") {
		t.Fatalf("nested collection callback did not retain the supplied slice pointer:\n%s", source)
	}
	if strings.Contains(source, "FuncSlice(io, &x.Entries") {
		t.Fatalf("struct slice still emits an escaping callback:\n%s", source)
	}
}

func TestGenerateKeepsUnionValidation(t *testing.T) {
	union := manifest.Union(manifest.Primitive("u8"),
		manifest.Variant{Value: 0, Name: "First", Encode: manifest.Void()},
		manifest.Variant{Value: 1, Name: "Second", Encode: manifest.Void()},
	)
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:union", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "UnionPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: union, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	marshal := generatedSource(files)
	if !strings.Contains(marshal, "unknown union tag") {
		t.Fatalf("union discriminator validation was removed:\n%s", marshal)
	}
	for name, source := range files {
		if strings.Contains(source, "unknown union tag") && strings.Contains(source, "io.Reading()") {
			t.Fatalf("union direction mechanics leaked into %s:\n%s", name, source)
		}
	}
}

func TestGenerateIncludesConcreteCodecRuntime(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:runtime", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "RuntimePacket", Direction: manifest.DirectionClientbound}},
	}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatal(err)
	}
	for name, wants := range map[string][]string{
		"protocol/codec.go":   {"type IO interface"},
		"protocol/helpers.go": {"func IntegerFunc", "func OptionalFunc", "func UnionFunc", "func FuncSlice"},
		"protocol/reader.go":  {"type Reader struct", "func NewReader", "func (r *Reader) NBT", "func (r *Reader) SliceLength"},
		"protocol/writer.go":  {"type Writer struct", "func NewWriter", "func (w *Writer) NBT", "func (w *Writer) Data"},
		"protocol/types.go":   {"type Optional", "type OrderedEntry"},
	} {
		source, ok := files[name]
		if !ok {
			t.Fatalf("generated output omits %s", name)
		}
		for _, want := range wants {
			if !strings.Contains(source, want) {
				t.Fatalf("%s omits %q:\n%s", name, want, source)
			}
		}
	}
	packetRuntime := files["protocol/packet/packet.go"]
	for _, want := range []string{"func Decode(data []byte, pk Packet) error", "func Encode(pk Packet) ([]byte, error)"} {
		if !strings.Contains(packetRuntime, want) {
			t.Fatalf("packet runtime omits %q:\n%s", want, packetRuntime)
		}
	}
}

func TestGenerateWrapsRepeatedUnionPayloadTypes(t *testing.T) {
	message := manifest.Struct(manifest.Field{Ordinal: 0, Name: "Text", Encode: manifest.String(manifest.Primitive("var_u32")), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}})
	message.TypeID = "Message"
	union := manifest.Union(manifest.Primitive("u8"),
		manifest.Variant{Value: 0, Name: "First", Encode: message},
		manifest.Variant{Value: 1, Name: "Second", Encode: message},
	)
	union.TypeID = "Choice"
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}}, Packets: []manifest.Packet{{ID: 1, Name: "EnvelopePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Choice", Encode: union, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}
	files, err := Generate(m, "fixture")
	if err != nil {
		t.Fatal(err)
	}
	marshal := generatedSource(files)
	if strings.Count(marshal, "case *Message:") != 1 || !strings.Contains(marshal, "case *ChoiceSecond:") {
		t.Fatalf("repeated union payloads were not made tag-distinct:\n%s", marshal)
	}
	if !strings.Contains(marshal, "value := new(Message)") || !strings.Contains(marshal, "*x = value") {
		t.Fatalf("union decode does not allocate pointer payloads:\n%s", marshal)
	}
}

func TestGenerateCodecFailsClosedForSequence(t *testing.T) {
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}}, Packets: []manifest.Packet{{ID: 1, Name: "SequencePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Sequence(manifest.Primitive("u8")), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}
	if _, err := Generate(m, "fixture"); err == nil || !strings.Contains(err.Error(), "sequence") {
		t.Fatalf("Generate error = %v, want sequence failure", err)
	}
}

func TestGenerateCodecFailsClosedForAsymmetricField(t *testing.T) {
	decode := manifest.Primitive("u16le")
	field := manifest.Field{Ordinal: 0, Name: "Value", Encode: manifest.Primitive("u8"), Decode: &decode, Symmetry: manifest.Asymmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}}, Packets: []manifest.Packet{{ID: 1, Name: "AsymmetricPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{field}}}}
	if _, err := Generate(m, "fixture"); err == nil || !strings.Contains(err.Error(), "asymmetric") {
		t.Fatalf("Generate error = %v, want asymmetric failure", err)
	}
}

func TestGenerateCollapsesCerealDoubleOptionalToOnePublicOptional(t *testing.T) {
	doubleOptional := manifest.Optional(manifest.Optional(manifest.Primitive("i32le")))
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:optional", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "OptionalPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: doubleOptional, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatal(err)
	}
	packet := files["protocol/packet/optional.go"]
	if !strings.Contains(packet, "Value protocol.Optional[int32]") || strings.Contains(packet, "**int32") || strings.Contains(packet, "Optional[Optional[") {
		t.Fatalf("double optional did not use one public Optional:\n%s", packet)
	}
	types := files["protocol/types.go"]
	for _, want := range []string{"type Optional[T any] struct", "func Option[T any]", "func (o Optional[T]) Value() (T, bool)"} {
		if !strings.Contains(types, want) {
			t.Fatalf("generated optional support omits %q:\n%s", want, types)
		}
	}
}

func TestGenerateFailsClosedForReachableUnresolvedNode(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture-2168", Digest: "fixture:golden", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "Blocked", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Unresolved("dynamic", true), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	if _, err := Generate(m, "wiregen"); err == nil || !strings.Contains(err.Error(), "unresolved") {
		t.Fatalf("Generate error = %v, want unresolved failure", err)
	}
}

func TestGenerateUsesCanonicalNamedTypesAndOrderedMapEntries(t *testing.T) {
	biome := manifest.Node{Kind: manifest.KindStruct, Semantic: "BiomeDefinitionData", TypeID: "BiomeDefinitionData", Fields: []manifest.Field{{Ordinal: 0, Name: "id", Encode: manifest.Primitive("u16le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	mapping := manifest.Map(manifest.Primitive("var_u32"), manifest.Primitive("u16le"), biome)
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:go", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 122, Name: "BiomeDefinitionListPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Map of Biome names to data", Encode: mapping, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}

	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	source := generatedSource(files)
	if !strings.Contains(source, "type BiomeDefinitionData struct") || !strings.Contains(source, "[]protocol.OrderedEntry[uint16, protocol.BiomeDefinitionData]") {
		t.Fatalf("generated Go did not use canonical ordered map definitions:\n%s", source)
	}
}

func TestGenerateUsesTypedUnionInterface(t *testing.T) {
	union := manifest.Union(manifest.Primitive("var_u32"),
		manifest.Variant{Value: 0, Name: "SoundDataEvent::Stop", Encode: manifest.Node{Kind: manifest.KindStruct, Semantic: "SoundDataEvent::Stop", TypeID: "SoundDataEvent::Stop"}},
		manifest.Variant{Value: 1, Name: "SoundDataEvent::SetVolume", Encode: manifest.Node{Kind: manifest.KindStruct, Semantic: "SoundDataEvent::SetVolume", TypeID: "SoundDataEvent::SetVolume", Fields: []manifest.Field{{Ordinal: 0, Name: "Volume", Encode: manifest.Primitive("f32le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	)
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:go", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 348, Name: "SoundPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Event", Encode: union, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	source := generatedSource(files)
	if !strings.Contains(source, "type SoundDataEvent interface") || !strings.Contains(source, "func (*SoundDataEventSetVolume) isSoundDataEvent()") || strings.Contains(source, "Tag int64") {
		t.Fatalf("generated Go did not emit a typed union interface:\n%s", source)
	}
}

func TestGenerateSplitsPacketsAndSharedDefinitions(t *testing.T) {
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:go", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 1, Name: "LoginPacket", Direction: manifest.DirectionServerbound}}}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"protocol/packet/ids.go", "protocol/types.go", "protocol/packet/login.go"} {
		if _, ok := files[name]; !ok {
			t.Fatalf("generated files omit %s: %v", name, files)
		}
	}
	if !strings.Contains(files["protocol/types.go"], "package protocol") || !strings.Contains(files["protocol/packet/login.go"], "package packet") || !strings.Contains(files["protocol/packet/login.go"], `import "wiregen"`) {
		t.Fatalf("generated packages were not separated or wired together:\n%s", files["protocol/packet/login.go"])
	}
	if _, ok := files["protocol/marshal.go"]; ok {
		t.Fatalf("generated output unexpectedly contains monolithic marshal.go")
	}
	if _, ok := files["protocol/enums.go"]; ok {
		t.Fatalf("generated output unexpectedly contains monolithic enums.go")
	}
	if strings.Contains(files["protocol/packet/login.go"], "LoginPacket") || !strings.Contains(files["protocol/packet/login.go"], "type Login struct") {
		t.Fatalf("packet name was not cleaned:\n%s", files["protocol/packet/login.go"])
	}
	if _, ok := files["wire.go"]; ok {
		t.Fatalf("definition-only output unexpectedly contains wire.go")
	}
}

func generatedSource(files map[string]string) string {
	var b strings.Builder
	for _, source := range files {
		b.WriteString(source)
	}
	return b.String()
}

func TestPublicNamesDropSchemaScaffolding(t *testing.T) {
	tests := map[string]string{
		"enums/MoLangVersion":                           "MoLangVersion",
		"enums/MolangVersion":                           "MoLangVersion",
		"PlayerVideoCapturePacketPayload::Action":       "Action",
		"DataItemEntryPayloadUnion":                     "DataItemEntryValue",
		"DimensionDefinitionGroup::DimensionDefinition": "DimensionDefinition",
		"ServerWaypointGroup::Action":                   "ServerWaypointGroupAction",
		"SharedTypes::v1_26_0::CameraSplineDefinition":  "CameraSplineDefinition",
	}
	for input, want := range tests {
		if got := publicTypeName(input); got != want {
			t.Fatalf("publicTypeName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestGoNamesUseCommonInitialisms(t *testing.T) {
	for input, want := range map[string]string{
		"Container Id":   "ContainerID",
		"Json Uri Xz Ui": "JSONURIXZUI",
		"Identifier":     "Identifier",
	} {
		if got := exportName(input); got != want {
			t.Errorf("exportName(%q) = %q, want %q", input, got, want)
		}
	}
	if got := enumVariantName("VALUE_ID"); got != "ValueID" {
		t.Fatalf("enumVariantName initialism = %q, want ValueID", got)
	}
}

func TestEnumVariantNamesAreIdiomaticGo(t *testing.T) {
	for input, want := range map[string]string{
		"NONE":                    "None",
		"START_ATTACKING":         "StartAttacking",
		"FISHHOOK_FISHPOS":        "FishhookFishPosition",
		"FISHHOOK_HOOKTIME":       "FishhookHookTime",
		"DRAGON_START_DEATH_ANIM": "DragonStartDeathAnimation",
		"SILVERFISH_MERGE_ANIM":   "SilverfishMergeAnimation",
		"PRIMED_TNT":              "PrimedTNT",
		"PRIME_TNTCART":           "PrimeTNTCart",
		"Primed_Tnt":              "PrimedTNT",
		"OSX":                     "OSX",
		"UWP":                     "UWP",
	} {
		if got := enumVariantName(input); got != want {
			t.Fatalf("enumVariantName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestGenerateMapsCanonicalSemanticsToNativeGoTypes(t *testing.T) {
	vec3 := manifest.Node{Kind: manifest.KindStruct, Semantic: "Vec3", TypeID: "Vec3", Fields: []manifest.Field{
		{Ordinal: 0, Name: "X", Encode: manifest.Primitive("f32le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 1, Name: "Y", Encode: manifest.Primitive("f32le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 2, Name: "Z", Encode: manifest.Primitive("f32le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
	}}
	colour := manifest.Node{Kind: manifest.KindStruct, Semantic: "mce::Color", TypeID: "mce::Color", Fields: []manifest.Field{
		{Ordinal: 0, Name: "Color", Encode: manifest.Primitive("i32le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
	}}
	runtimeID := manifest.Node{Kind: manifest.KindStruct, Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Primitive("var_u64"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	uniqueID := manifest.Node{Kind: manifest.KindStruct, Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Primitive("zigzag_i64"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:native-go", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 1, Name: "NativePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
		{Ordinal: 0, Name: "ID", Encode: manifest.Primitive("uuid"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 1, Name: "Position", Encode: vec3, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 2, Name: "Colour", Encode: colour, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 3, Name: "Data", Encode: manifest.NBT(manifest.NBTNetwork), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 4, Name: "Runtime", Encode: runtimeID, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 5, Name: "Unique", Encode: uniqueID, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
	}}}}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatal(err)
	}
	source := files["protocol/packet/native.go"]
	for _, want := range []string{"uuid.UUID", "mgl32.Vec3", "color.RGBA", "[]byte", "Runtime  uint64", "Unique   int64", `"github.com/google/uuid"`, `"github.com/go-gl/mathgl/mgl32"`, `"image/color"`} {
		if !strings.Contains(source, want) {
			t.Fatalf("native Go output omits %q:\n%s", want, source)
		}
	}
	if strings.Contains(files["protocol/types.go"], "type Vec3 struct") || strings.Contains(files["protocol/types.go"], "type MceColor struct") {
		t.Fatalf("native Go types were redundantly regenerated:\n%s", files["protocol/types.go"])
	}
	if !strings.Contains(source, "io.ActorRuntimeID(&x.Runtime)") || !strings.Contains(source, "io.ActorUniqueID(&x.Unique)") {
		t.Fatalf("semantic ID IO methods were not used:\n%s", source)
	}
}

func TestNativeGoTypeRejectsMislabelledVector(t *testing.T) {
	node := manifest.Node{Kind: manifest.KindStruct, TypeID: "Vec3", Fields: []manifest.Field{{Encode: manifest.Primitive("f64le")}}}
	if _, matched, err := nativeGoType(node); !matched || err == nil {
		t.Fatalf("nativeGoType matched=%v error=%v, want a closed failure", matched, err)
	}
}

func TestSemanticIOCallUsesExactIdentifierCodecs(t *testing.T) {
	tests := []struct {
		typeID string
		code   string
		want   string
	}{
		{typeID: "ActorRuntimeID", code: "var_u64", want: "ActorRuntimeID"},
		{typeID: "ActorRuntimeID", code: "zigzag_i64", want: "ActorRuntimeIDVarint64"},
		{typeID: "ActorRuntimeID", code: "var_u32", want: "ActorRuntimeIDVaruint32"},
		{typeID: "ActorRuntimeID", code: "var_i64", want: "SignedVarint64"},
		{typeID: "ActorUniqueID", code: "zigzag_i64", want: "ActorUniqueID"},
		{typeID: "ActorUniqueID", code: "i64le", want: "ActorUniqueIDInt64"},
		{typeID: "ActorUniqueID", code: "u64le", want: "ActorUniqueIDUint64"},
		{typeID: "ActorUniqueID", code: "var_u64", want: "ActorUniqueIDVaruint64"},
		{typeID: "ActorUniqueID", code: "var_i64", want: "SignedVarint64"},
	}
	for _, test := range tests {
		t.Run(test.typeID+"/"+test.code, func(t *testing.T) {
			node := manifest.Node{Kind: manifest.KindStruct, TypeID: test.typeID, Fields: []manifest.Field{{Encode: manifest.Primitive(test.code)}}}
			if method, ok := semanticIOCall(node); !ok || method != test.want {
				t.Fatalf("semanticIOCall = (%q, %v), want (%q, true)", method, ok, test.want)
			}
		})
	}

	bad := manifest.Node{Kind: manifest.KindStruct, TypeID: "ActorUniqueID", Fields: []manifest.Field{{Encode: manifest.Primitive("f32le")}}}
	if method, ok := semanticIOCall(bad); ok {
		t.Fatalf("semanticIOCall returned %q for unsupported identifier layout", method)
	}
}

func TestGenerateUsesNamedRecursiveTypeAndBoundedBitset(t *testing.T) {
	dynamic := manifest.Union(manifest.Primitive("u8"),
		manifest.Variant{Value: 0, Name: "Empty", Encode: manifest.Void()},
		manifest.Variant{Value: 1, Name: "List", Encode: manifest.Array(manifest.Primitive("var_u32"), manifest.Recursive("cereal::DynamicValue"))},
	)
	dynamic.TypeID = "cereal::DynamicValue"
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "fixture", Revision: "1", Digest: "sha256:fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets: []manifest.Packet{{ID: 1, Name: "RecursivePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Value", Encode: dynamic, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 1, Name: "Flags", Encode: manifest.Bitset(131), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
	files, err := Generate(m, "fixture")
	if err != nil {
		t.Fatal(err)
	}
	var all strings.Builder
	for _, source := range files {
		all.WriteString(source)
	}
	for _, want := range []string{"[]CerealDynamicValue", "type Bitset131 [3]uint64", "Flags protocol.Bitset131"} {
		if !strings.Contains(all.String(), want) {
			t.Fatalf("generated output missing %q:\n%s", want, all.String())
		}
	}
}

func TestGenerateEmitsPacketRuntimePoolsAndOptionalNativeProfile(t *testing.T) {
	packets := []manifest.Packet{
		{ID: 1, Name: "ServerEnvelope", Direction: manifest.DirectionClientbound},
		{ID: 2, Name: "ClientEnvelope", Direction: manifest.DirectionServerbound},
		{ID: 3, Name: "SharedEnvelope", Direction: manifest.DirectionBidirectional},
	}
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "fixture", Revision: "1", Digest: "sha256:fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets:       packets,
	}
	files, err := Generate(m, "fixture")
	if err != nil {
		t.Fatal(err)
	}
	serverSource := ""
	for _, source := range files {
		if strings.Contains(source, "type ServerEnvelope struct") {
			serverSource = source
			break
		}
	}
	if !strings.Contains(files["protocol/packet/packet.go"], "type Packet interface") || !strings.Contains(serverSource, "func (*ServerEnvelope) ID() uint32 { return IDServerEnvelope }") {
		t.Fatalf("packet runtime contract was not emitted:\n%s\n%s", files["protocol/packet/packet.go"], serverSource)
	}
	pool := files["protocol/packet/pool.go"]
	for _, want := range []string{"IDServerEnvelope: func() Packet", "IDClientEnvelope: func() Packet", "IDSharedEnvelope: func() Packet", "func NewClientPacket", "func NewServerPacket"} {
		if !strings.Contains(pool, want) {
			t.Fatalf("packet pools omit %q:\n%s", want, pool)
		}
	}
	withoutNative, err := GenerateWithOptions(m, Options{ProtocolImportPath: "fixture", EmitPacketRuntime: true, EmitPacketPools: true})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(withoutNative["protocol/packet/server_packet.go"], "uuid.UUID") {
		t.Fatal("native profile unexpectedly leaked into disabled output")
	}
}

func TestGenerateWithoutNativeTypesUsesFixedUUIDBytes(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "fixture", Revision: "1", Digest: "sha256:fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}},
		Packets: []manifest.Packet{{ID: 1, Name: "UUIDPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{
			Ordinal: 0, Name: "UUID", Encode: manifest.Primitive("uuid"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}},
		}}}},
	}
	files, err := GenerateWithOptions(m, Options{ProtocolImportPath: "fixture", EmitPacketRuntime: true, EmitPacketPools: true})
	if err != nil {
		t.Fatal(err)
	}
	packet := ""
	for _, source := range files {
		if strings.Contains(source, "type UUID struct") {
			packet = source
			break
		}
	}
	if !strings.Contains(packet, "UUID [16]byte") || !strings.Contains(packet, "io.UUIDBytes(&x.UUID)") {
		t.Fatalf("disabled native profile did not emit fixed UUID bytes:\n%s", packet)
	}
}

func TestGenerateGoPreservesNBTEncodingOnIOCalls(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: manifest.SchemaVersion,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "sha256:fixture"}},
		Packets: []manifest.Packet{{ID: 1, Name: "NBTPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Network", Encode: manifest.NBT(manifest.NBTNetwork), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 1, Name: "Persistent", Encode: manifest.NBT(manifest.NBTPersistent), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
	files, err := Generate(m, "example.com/protocol")
	if err != nil {
		t.Fatal(err)
	}
	packet := files["protocol/packet/nbt.go"]
	for _, want := range []string{"io.NBT(&x.Network, protocol.NBTNetwork)", "io.NBT(&x.Persistent, protocol.NBTPersistent)"} {
		if !strings.Contains(packet, want) {
			t.Fatalf("packet omitted %q:\n%s", want, packet)
		}
	}
	if !strings.Contains(files["protocol/codec.go"], "NBT(*[]byte, NBTEncoding)") {
		t.Fatalf("runtime IO did not expose format-aware NBT:\n%s", files["protocol/codec.go"])
	}
}
