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
	if !strings.Contains(files["vocabulary.go"], "type Vocabulary struct") || !strings.Contains(files["vocabulary.go"], "Kind: \"optional\"") {
		t.Fatalf("generated packet source did not contain manifest shape:\n%s", files["vocabulary.go"])
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
	source := files["types.go"] + files["biome_definition_list.go"]
	if !strings.Contains(source, "type BiomeDefinitionData struct") || !strings.Contains(source, "[]OrderedEntry[uint16, BiomeDefinitionData]") {
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
	source := files["types.go"] + files["sound.go"]
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
	for _, name := range []string{"wire.go", "types.go", "enums.go", "login.go"} {
		if _, ok := files[name]; !ok {
			t.Fatalf("generated files omit %s: %v", name, files)
		}
	}
	if strings.Contains(files["login.go"], "LoginPacket") || !strings.Contains(files["login.go"], "type Login struct") {
		t.Fatalf("packet name was not cleaned:\n%s", files["login.go"])
	}
}

func TestPublicNamesDropSchemaScaffolding(t *testing.T) {
	tests := map[string]string{
		"enums/MoLangVersion":                          "MoLangVersion",
		"enums/MolangVersion":                          "MoLangVersion",
		"PlayerVideoCapturePacketPayload::Action":      "PlayerVideoCaptureAction",
		"DataItemEntryPayloadUnion":                    "DataItemEntryValue",
		"SharedTypes::v1_26_0::CameraSplineDefinition": "CameraSplineDefinition",
	}
	for input, want := range tests {
		if got := publicTypeName(input); got != want {
			t.Fatalf("publicTypeName(%q) = %q, want %q", input, got, want)
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
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:native-go", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 1, Name: "NativePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
		{Ordinal: 0, Name: "ID", Encode: manifest.Primitive("uuid"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 1, Name: "Position", Encode: vec3, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 2, Name: "Colour", Encode: colour, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		{Ordinal: 3, Name: "Data", Encode: manifest.Primitive("nbt_le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
	}}}}
	files, err := Generate(m, "wiregen")
	if err != nil {
		t.Fatal(err)
	}
	source := files["native.go"]
	for _, want := range []string{"uuid.UUID", "mgl32.Vec3", "color.RGBA", "[]byte", `"github.com/google/uuid"`, `"github.com/go-gl/mathgl/mgl32"`, `"image/color"`} {
		if !strings.Contains(source, want) {
			t.Fatalf("native Go output omits %q:\n%s", want, source)
		}
	}
	if strings.Contains(files["types.go"], "type Vec3 struct") || strings.Contains(files["types.go"], "type MceColor struct") {
		t.Fatalf("native Go types were redundantly regenerated:\n%s", files["types.go"])
	}
}

func TestNativeGoTypeRejectsMislabelledVector(t *testing.T) {
	node := manifest.Node{Kind: manifest.KindStruct, TypeID: "Vec3", Fields: []manifest.Field{{Encode: manifest.Primitive("f64le")}}}
	if _, matched, err := nativeGoType(node); !matched || err == nil {
		t.Fatalf("nativeGoType matched=%v error=%v, want a closed failure", matched, err)
	}
}
