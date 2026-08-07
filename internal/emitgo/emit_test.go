package emitgo

import (
	"go/format"
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

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
		"value.Marshal(io)",
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
		"protocol/codec.go":  {"type IO interface", "func IntegerFunc", "func OptionalFunc", "func UnionFunc", "func FuncSlice"},
		"protocol/reader.go": {"type Reader struct", "func NewReader", "func (r *Reader) NBT", "func (r *Reader) SliceLength"},
		"protocol/writer.go": {"type Writer struct", "func NewWriter", "func (w *Writer) NBT", "func (w *Writer) Data"},
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
}

func TestGenerateWrapsRepeatedUnionPayloadTypes(t *testing.T) {
	message := manifest.Struct(manifest.Field{Ordinal: 0, Name: "Text", Encode: manifest.String(manifest.Primitive("var_u32")), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}})
	message.TypeID = "Message"
	union := manifest.Union(manifest.Primitive("u8"),
		manifest.Variant{Value: 0, Name: "First", Encode: message},
		manifest.Variant{Value: 1, Name: "Second", Encode: message},
	)
	union.TypeID = "Choice"
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "1", Digest: "fixture", MinecraftVersion: "fixture", ProtocolVersion: 1}}, Packets: []manifest.Packet{{ID: 1, Name: "ChoicePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Choice", Encode: union, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}
	files, err := Generate(m, "fixture")
	if err != nil {
		t.Fatal(err)
	}
	marshal := generatedSource(files)
	if strings.Count(marshal, "case Message:") != 1 || !strings.Contains(marshal, "case ChoiceChoiceSecond:") {
		t.Fatalf("repeated union payloads were not made tag-distinct:\n%s", marshal)
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
	if !strings.Contains(source, "type SoundDataEvent interface") || !strings.Contains(source, "func (SoundDataEventSetVolume) isSoundDataEvent()") || strings.Contains(source, "Tag int64") {
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
		"PlayerVideoCapturePacketPayload::Action":       "PlayerVideoCaptureAction",
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
		{Ordinal: 3, Name: "Data", Encode: manifest.Primitive("nbt_le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
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
