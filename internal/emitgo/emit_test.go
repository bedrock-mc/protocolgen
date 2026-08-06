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
	if !strings.Contains(files["packets.go"], "type VocabularyPacket struct") || !strings.Contains(files["packets.go"], "Kind: \"optional\"") {
		t.Fatalf("generated packet source did not contain manifest shape:\n%s", files["packets.go"])
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
	source := files["packets.go"]
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
	source := files["packets.go"]
	if !strings.Contains(source, "type SoundDataEvent interface") || !strings.Contains(source, "func (SoundDataEventSetVolume) isSoundDataEvent()") || strings.Contains(source, "Tag int64") {
		t.Fatalf("generated Go did not emit a typed union interface:\n%s", source)
	}
}
