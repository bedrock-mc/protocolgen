package ingest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func TestMojangIngestionRetainsWireVocabulary(t *testing.T) {
	root := t.TempDir()
	writeJSON(t, filepath.Join(root, "Vocabulary.json"), map[string]any{
		"title": "VocabularyPacket", "$metaProperties": map[string]any{"[cereal:packet]": 1},
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
		"properties": map[string]any{
			"Optional": map[string]any{"type": "integer", "x-underlying-type": "int32", "x-ordinal-index": 0},
			"Double":   map[string]any{"type": "integer", "x-underlying-type": "int32", "x-serialization-options": []string{"+double-optional"}, "x-ordinal-index": 1},
			"Bytes":    map[string]any{"type": "string", "x-underlying-type": "bytearray", "x-ordinal-index": 2},
			"Fixed":    map[string]any{"type": "array", "minItems": 2, "maxItems": 2, "items": map[string]any{"type": "integer", "x-underlying-type": "uint16"}, "x-ordinal-index": 3},
			"Choice": map[string]any{"oneOf": []any{
				map[string]any{"title": "None", "type": "null", "x-ordinal-index": 0},
				map[string]any{"title": "Payload", "type": "string", "x-ordinal-index": 7},
			}, "x-control-value-type": "uint8", "x-ordinal-index": 4},
			"Mode": map[string]any{"type": "integer", "x-underlying-type": "uint8", "enum": []string{"Ready", "Later"}, "x-enum-values": []int{4, 9}, "x-serialization-options": []string{"Enum-as-Value"}, "x-ordinal-index": 5},
		},
		"required": []string{"Bytes", "Fixed", "Choice", "Mode"},
	})
	result, err := ParseMojang(root, fixturePin("mojang"), "")
	if err != nil {
		t.Fatalf("ParseMojang: %v", err)
	}
	if len(result.Claims) != 6 {
		t.Fatalf("claims = %d, want 6", len(result.Claims))
	}
	if got := result.Claims[0].Encode.Kind; got != manifest.KindOptional {
		t.Errorf("optional kind = %q, want optional", got)
	}
	if got := result.Claims[1].Encode.Kind; got != manifest.KindOptional || result.Claims[1].Encode.Value.Kind != manifest.KindOptional {
		t.Errorf("double optional = %+v", result.Claims[1].Encode)
	}
	if got := result.Claims[2].Encode.Kind; got != manifest.KindBytes {
		t.Errorf("bytes kind = %q, want bytes", got)
	}
	if got := result.Claims[3].Encode.Kind; got != manifest.KindFixedArray || result.Claims[3].Encode.Length != 2 {
		t.Errorf("fixed array = %+v", result.Claims[3].Encode)
	}
	if got := result.Claims[4].Encode.Kind; got != manifest.KindUnion || result.Claims[4].Encode.Variants[1].Value != 7 {
		t.Errorf("union = %+v", result.Claims[4].Encode)
	}
	if got := result.Claims[5].Encode.Kind; got != manifest.KindEnum || result.Claims[5].Encode.Variants[0].Value != 4 {
		t.Errorf("enum = %+v", result.Claims[5].Encode)
	}
}

func TestEndstoneIngestionRetainsMapsSwitchesAndFixedArrays(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "packets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "types"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "enums"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("**Minecraft Version:** `1.26.40.1`\n**Network Version:** `2168`\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(root, "enums", "Mode.json"), map[string]any{"name": "Mode", "values": []any{map[string]any{"name": "None", "value": 0}, map[string]any{"name": "Payload", "value": 7}}})
	writeJSON(t, filepath.Join(root, "types", "Entry.json"), map[string]any{"name": "Entry", "fields": []any{map[string]any{"name": "Value", "type": "uint8"}}})
	writeJSON(t, filepath.Join(root, "packets", "Vocabulary.json"), map[string]any{"name": "Vocabulary", "id": 1, "fields": []any{
		map[string]any{"name": "Optional", "type": "int32", "optional": true},
		map[string]any{"name": "Fixed", "type": map[string]any{"type": "uint16", "repeat": map[string]any{"count": 2}}},
		map[string]any{"name": "Choice", "type": map[string]any{"switch": map[string]any{"type": "uint8", "enum": "Mode"}, "cases": []any{nil, "Entry"}}},
		map[string]any{"name": "Text", "type": map[string]any{"type": "string", "encoding": "utf8"}},
		map[string]any{"name": "Bytes", "type": "restBuffer"},
		map[string]any{"name": "Entries", "type": map[string]any{"key": "string", "value": "uint8"}},
	}})
	result, err := ParseEndstone(root, fixturePin("endstone"), "")
	if err != nil {
		t.Fatalf("ParseEndstone: %v", err)
	}
	if len(result.Claims) != 6 {
		t.Fatalf("claims = %d, want 6", len(result.Claims))
	}
	if result.Claims[1].Encode.Kind != manifest.KindFixedArray || result.Claims[2].Encode.Kind != manifest.KindUnion || result.Claims[4].Encode.Kind != manifest.KindBytes || result.Claims[5].Encode.Kind != manifest.KindMap {
		t.Fatalf("unexpected Endstone shapes: %#v", result.Claims)
	}
}

func TestCorrectionRequiresPrePatchFingerprintAndEvidence(t *testing.T) {
	documents := map[string]any{"Packet.json": map[string]any{"value": "string"}}
	before, err := canonicalDigest(documents["Packet.json"].(map[string]any)["value"])
	if err != nil {
		t.Fatal(err)
	}
	corrections := t.TempDir()
	writeJSON(t, filepath.Join(corrections, "bytes.json"), map[string]any{"operations": []any{map[string]any{
		"id": "bytes-field", "file": "Packet.json", "pointer": "/value", "pre_patch_sha256": before, "replace": "bytearray", "why": "fixture byte evidence", "evidence": []any{map[string]any{"locator": "fixture/runtime"}},
	}}})
	proofs, err := applyCorrections(documents, corrections, fixturePin("mojang"))
	if err != nil {
		t.Fatalf("applyCorrections: %v", err)
	}
	if len(proofs) != 1 || proofs[0].PostPatchNodeSHA256 == before {
		t.Fatalf("proofs = %#v", proofs)
	}
	if proofs[0].PrePatchContextSHA256 == "" {
		t.Fatal("correction proof has no complete pre-patch context fingerprint")
	}
	documents["Packet.json"].(map[string]any)["value"] = "mutated"
	if _, err := applyCorrections(documents, corrections, fixturePin("mojang")); err == nil || !strings.Contains(err.Error(), "stale") {
		t.Fatalf("stale correction error = %v", err)
	}
}

func TestMojangRejectsMixedProtocolDocuments(t *testing.T) {
	root := t.TempDir()
	writeJSON(t, filepath.Join(root, "A.json"), map[string]any{
		"title": "A", "$metaProperties": map[string]any{"[cereal:packet]": 1},
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
		"properties": map[string]any{"Value": map[string]any{"type": "integer", "x-underlying-type": "uint8", "x-ordinal-index": 0}},
		"required":   []string{"Value"},
	})
	writeJSON(t, filepath.Join(root, "B.json"), map[string]any{
		"title": "B", "$metaProperties": map[string]any{"[cereal:packet]": 2},
		"x-minecraft-version": "1.26.50", "x-protocol-version": 2169,
		"properties": map[string]any{"Value": map[string]any{"type": "integer", "x-underlying-type": "uint8", "x-ordinal-index": 0}},
		"required":   []string{"Value"},
	})
	if _, err := ParseMojang(root, fixturePin("mojang"), ""); err == nil || !strings.Contains(err.Error(), "mixes") {
		t.Fatalf("ParseMojang error = %v, want mixed-version rejection", err)
	}
}

func TestEndstoneSpecialUnknownsStayReachableUnresolved(t *testing.T) {
	lowerer := &endstoneLowerer{types: map[string]any{}, enums: map[string]any{}, active: map[string]bool{}}
	for _, name := range []string{"cereal::DynamicValue", "brstd::bitset<128>"} {
		node := lowerer.lowerNamed(name, "fixture")
		if node.Kind != manifest.KindUnresolved || !node.Reachable {
			t.Fatalf("%s lowered to %#v, want reachable unresolved", name, node)
		}
		if node.Element != nil || node.Prefix != nil {
			t.Fatalf("%s invented a byte shape: %#v", name, node)
		}
	}
}

func TestUUIDUsesCanonicalPrimitive(t *testing.T) {
	node := primitive("uuid", nil, "")
	if node.Kind != manifest.KindPrimitive || node.Primitive == nil || node.Primitive.Code != "uuid" {
		t.Fatalf("uuid node = %#v, want canonical uuid primitive", node)
	}
}

func fixturePin(id string) manifest.SourcePin {
	return manifest.SourcePin{ID: id, Kind: id, Revision: "fixture-2168", Digest: "fixture:" + id, MinecraftVersion: "1.26.40", ProtocolVersion: 2168}
}

func writeJSON(t *testing.T, path string, value any) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
}
