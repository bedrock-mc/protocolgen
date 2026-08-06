package emitrust

import (
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func TestGenerateRustConsumesCanonicalManifest(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "FixturePacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Optional(manifest.Primitive("u8")), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	source, err := Generate(m)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(source, "pub struct FixturePacket") || !strings.Contains(source, "pub const FIXTUREPACKET_VALUE_SHAPE") {
		t.Fatalf("generated Rust omitted packet/shape:\n%s", source)
	}
	if strings.Contains(source, "gophertunnel") || strings.Contains(source, "bedrock-protocol-docs") {
		t.Fatalf("generated Rust contains source/profile lookup text")
	}
}

func TestGenerateRustFailsClosedForUnresolved(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "Blocked", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Unresolved("dynamic", true), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	if _, err := Generate(m); err == nil || !strings.Contains(err.Error(), "unresolved") {
		t.Fatalf("Generate error = %v, want unresolved failure", err)
	}
}

func TestGenerateRustEscapesKeywordFieldNames(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}},
		Packets:       []manifest.Packet{{ID: 1, Name: "KeywordPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Type", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	}
	source, err := Generate(m)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(source, "pub r#type: u8") {
		t.Fatalf("generated Rust did not escape keyword field:\n%s", source)
	}
}

func TestRustFieldNamesUseSnakeCase(t *testing.T) {
	tests := map[string]string{
		"Actor Unique ID": "actor_unique_id",
		"FrameRate":       "frame_rate",
		"Pack UUID":       "pack_uuid",
	}
	for input, want := range tests {
		if got := fieldName(input); got != want {
			t.Fatalf("fieldName(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestGenerateRustUsesCanonicalNamedTypesAndOrderedMapTuples(t *testing.T) {
	biome := manifest.Node{Kind: manifest.KindStruct, Semantic: "BiomeDefinitionData", TypeID: "BiomeDefinitionData", Fields: []manifest.Field{{Ordinal: 0, Name: "id", Encode: manifest.Primitive("u16le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}
	mapping := manifest.Map(manifest.Primitive("var_u32"), manifest.Primitive("u16le"), biome)
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 122, Name: "BiomeDefinitionListPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Map of Biome names to data", Encode: mapping, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}

	source, err := Generate(m)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(source, "pub struct BiomeDefinitionData") {
		t.Fatalf("generated Rust did not use canonical type identity:\n%s", source)
	}
	if !strings.Contains(source, "pub map_of_biome_names_to_data: Vec<(u16, BiomeDefinitionData)>") {
		t.Fatalf("generated Rust did not use ordered map tuples:\n%s", source)
	}
	if strings.Contains(source, "BiomeDefinitionListPacketMapOfBiomeNamesToDataValueStruct") {
		t.Fatalf("generated Rust retained a path-derived type name:\n%s", source)
	}
}

func TestGenerateRustUsesTypedUnionEnum(t *testing.T) {
	union := manifest.Union(manifest.Primitive("var_u32"),
		manifest.Variant{Value: 0, Name: "SoundDataEvent::Stop", Encode: manifest.Void()},
		manifest.Variant{Value: 1, Name: "SoundDataEvent::SetVolume", Encode: manifest.Node{Kind: manifest.KindStruct, Semantic: "SoundDataEvent::SetVolume", TypeID: "SoundDataEvent::SetVolume", Fields: []manifest.Field{{Ordinal: 0, Name: "Volume", Encode: manifest.Primitive("f32le"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}},
	)
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 348, Name: "SoundPacket", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Event", Encode: union, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}

	source, err := Generate(m)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(source, "pub enum SoundDataEvent") || !strings.Contains(source, "SetVolume(SoundDataEventSetVolume)") || strings.Contains(source, "pub tag: i64") {
		t.Fatalf("generated Rust did not emit a typed union enum:\n%s", source)
	}
}

func TestGenerateRustEmptyPacketDoesNotEmitUnusedParameters(t *testing.T) {
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 4, Name: "EmptyPacket", Direction: manifest.DirectionServerbound, Fields: []manifest.Field{}}}}
	source, err := Generate(m)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(source, "_encoder: &mut E") || !strings.Contains(source, "_decoder: &mut D") {
		t.Fatalf("empty packet parameters are not intentionally unused:\n%s", source)
	}
}

func TestGenerateRustKeepsUnrelatedAnonymousUnionsDistinct(t *testing.T) {
	first := manifest.Union(manifest.Primitive("u8"), manifest.Variant{Value: 0, Name: "First", Encode: manifest.Void()})
	second := manifest.Union(manifest.Primitive("u8"), manifest.Variant{Value: 0, Name: "Second", Encode: manifest.Void()})
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:rust", MinecraftVersion: "fixture", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 1, Name: "Packet", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Alpha", Encode: first, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}, {Ordinal: 1, Name: "Beta", Encode: second, Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}
	source, err := Generate(m)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if !strings.Contains(source, "pub alpha: PacketAlphaUnion") || !strings.Contains(source, "pub beta: PacketBetaUnion") {
		t.Fatalf("anonymous unions were incorrectly deduplicated:\n%s", source)
	}
}
