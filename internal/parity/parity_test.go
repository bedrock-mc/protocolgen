package parity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func TestAxolotlV1NormalizationMatchesCanonicalByteShape(t *testing.T) {
	m := manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "1.26.40", ProtocolVersion: 2168},
		Sources:       []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:parity", MinecraftVersion: "1.26.40", ProtocolVersion: 2168}},
		Packets: []manifest.Packet{{ID: 7, Name: "Fixture", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{
			{Ordinal: 0, Name: "Value", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
			{Ordinal: 1, Name: "Maybe", Encode: manifest.Optional(manifest.String(manifest.Primitive("var_u32"))), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}},
		}}},
	}
	axolotl := `{"schema_version":1,"source":"fixture","minecraft_version":"1.26.40","protocol_version":2168,"packets":[{"id":7,"name":"Fixture","operations":[{"kind":"primitive","field":"Value","op":"U8"},{"kind":"option","field":"Maybe","presence":"Bool","value":[{"kind":"string","field":"Maybe","prefix":"VarInt","encoding":"utf8"}]}]}]}`
	path := filepath.Join(t.TempDir(), "axolotl.json")
	if err := os.WriteFile(path, []byte(axolotl), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := CompareFile(m, path); err != nil {
		t.Fatalf("CompareFile: %v", err)
	}
}

func TestAxolotlParityRejectsProtocolMix(t *testing.T) {
	m := manifest.Manifest{SchemaVersion: 2, Target: manifest.Target{MinecraftVersion: "1.26.40", ProtocolVersion: 2168}, Sources: []manifest.SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture", Digest: "fixture:parity", MinecraftVersion: "1.26.40", ProtocolVersion: 2168}}, Packets: []manifest.Packet{{ID: 1, Name: "P", Direction: manifest.DirectionClientbound, Fields: []manifest.Field{{Ordinal: 0, Name: "Value", Encode: manifest.Primitive("u8"), Symmetry: manifest.Symmetric, Provenance: manifest.Provenance{Pins: []string{"fixture"}}}}}}}
	path := filepath.Join(t.TempDir(), "axolotl.json")
	if err := os.WriteFile(path, []byte(`{"schema_version":1,"source":"fixture","minecraft_version":"1.26.50","protocol_version":2169,"packets":[]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := CompareFile(m, path); err == nil || !strings.Contains(err.Error(), "mix") {
		t.Fatalf("CompareFile error = %v, want protocol mix failure", err)
	}
}
