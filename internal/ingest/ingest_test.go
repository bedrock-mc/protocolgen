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

func TestMojangPacketRootReferenceUsesPayloadFields(t *testing.T) {
	root := t.TempDir()
	writeJSON(t, filepath.Join(root, "LoginPacket.json"), map[string]any{
		"title": "LoginPacket", "$metaProperties": map[string]any{"[cereal:packet]": 1},
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
		"$ref": "./LoginPacketPayload.json",
	})
	writeJSON(t, filepath.Join(root, "LoginPacketPayload.json"), map[string]any{
		"title": "LoginPacketPayload", "type": "object",
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
		"properties": map[string]any{
			"Client Network Version": map[string]any{"type": "integer", "x-underlying-type": "int32", "x-serialization-options": []string{"Big Endian"}, "x-ordinal-index": 0},
			"Connection Request":     map[string]any{"type": "string", "x-ordinal-index": 1},
		},
		"required": []string{"Client Network Version", "Connection Request"},
	})

	result, err := ParseMojang(root, fixturePin("mojang"), "")
	if err != nil {
		t.Fatalf("ParseMojang: %v", err)
	}
	if len(result.Claims) != 2 {
		t.Fatalf("claims = %d, want 2", len(result.Claims))
	}
	if got := result.Claims[0].Encode.Primitive.Code; got != "i32be" {
		t.Fatalf("client network version codec = %q, want i32be", got)
	}
}

func TestMojangPacketRootReferenceAllowsEmptyObject(t *testing.T) {
	root := t.TempDir()
	writeJSON(t, filepath.Join(root, "HandshakePacket.json"), map[string]any{
		"title": "HandshakePacket", "$metaProperties": map[string]any{"[cereal:packet]": 4},
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
		"$ref": "./HandshakePacketPayload.json",
	})
	writeJSON(t, filepath.Join(root, "HandshakePacketPayload.json"), map[string]any{
		"title": "HandshakePacketPayload", "type": "object",
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
	})

	result, err := ParseMojang(root, fixturePin("mojang"), "")
	if err != nil {
		t.Fatalf("ParseMojang: %v", err)
	}
	if len(result.Claims) != 0 {
		t.Fatalf("claims = %d, want 0", len(result.Claims))
	}
}

func TestMojangBareSelfReferenceIsUnresolved(t *testing.T) {
	root := t.TempDir()
	writeJSON(t, filepath.Join(root, "HandshakePacket.json"), map[string]any{
		"title": "HandshakePacket", "$metaProperties": map[string]any{"[cereal:packet]": 3},
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
		"properties": map[string]any{
			"Token": map[string]any{"$ref": "./WebToken.json", "x-ordinal-index": 0},
		},
		"required": []string{"Token"},
	})
	writeJSON(t, filepath.Join(root, "WebToken.json"), map[string]any{
		"title": "WebToken", "$ref": "./WebToken.json",
		"x-minecraft-version": "1.26.40", "x-protocol-version": 2168,
	})

	result, err := ParseMojang(root, fixturePin("mojang"), "")
	if err != nil {
		t.Fatalf("ParseMojang: %v", err)
	}
	if got := result.Claims[0].Encode.Kind; got != manifest.KindUnresolved {
		t.Fatalf("self reference kind = %q, want unresolved", got)
	}
}

func TestMojangUnionInfersMissingPositionalSelectorWhenPublishedSelectorsConfirmOrder(t *testing.T) {
	lowerer := &mojangLowerer{documents: map[string]any{}, active: map[string]bool{}}
	node := lowerer.lowerSchema(map[string]any{
		"oneOf": []any{
			map[string]any{"title": "Payload", "type": "string"},
			map[string]any{"title": "None", "type": "null", "x-ordinal-index": 1},
		},
		"x-control-value-type": "uint32",
	}, "Packet.json", "PacketChoice")
	if node.Kind != manifest.KindUnion || len(node.Variants) != 2 || node.Variants[0].Value != 0 || node.Variants[1].Value != 1 {
		t.Fatalf("union = %#v, want positional selectors 0 and 1", node)
	}
}

func TestMojangUUIDReferenceUsesCanonicalPrimitive(t *testing.T) {
	lowerer := &mojangLowerer{documents: map[string]any{
		"mce__UUID.json": map[string]any{
			"title": "mce::UUID", "type": "object",
			"properties": map[string]any{
				"Most Significant Bits":  map[string]any{"type": "integer", "x-underlying-type": "uint64", "x-ordinal-index": 0},
				"Least Significant Bits": map[string]any{"type": "integer", "x-underlying-type": "uint64", "x-ordinal-index": 1},
			},
		},
	}, active: map[string]bool{}}
	node := lowerer.lowerSchema(map[string]any{"$ref": "./mce__UUID.json"}, "Packet.json", "PacketUUID")
	if node.Kind != manifest.KindPrimitive || node.Primitive == nil || node.Primitive.Code != "uuid" {
		t.Fatalf("UUID reference = %#v, want canonical uuid primitive", node)
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

func TestEndstonePacketFieldAppliesRepeatPrefix(t *testing.T) {
	root := t.TempDir()
	for _, directory := range []string{"packets", "types", "enums"} {
		if err := os.Mkdir(filepath.Join(root, directory), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("**Minecraft Version:** `1.26.40.1`\n**Network Version:** `2168`\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(root, "packets", "ListPacket.json"), map[string]any{
		"name": "ListPacket", "id": 1,
		"fields": []any{map[string]any{"name": "Items", "type": "uint16", "repeat": map[string]any{"prefix": "uvarint32"}}},
	})

	result, err := ParseEndstone(root, fixturePin("endstone"), "")
	if err != nil {
		t.Fatalf("ParseEndstone: %v", err)
	}
	node := result.Claims[0].Encode
	if node.Kind != manifest.KindArray || node.Prefix == nil || node.Prefix.Primitive == nil || node.Prefix.Primitive.Code != "var_u32" {
		t.Fatalf("field = %#v, want var_u32-prefixed array", node)
	}
}

func TestEndstoneFieldEnumUsesItsCerealValueConstraints(t *testing.T) {
	lowerer := &endstoneLowerer{
		types: map[string]any{},
		enums: map[string]any{
			"GameType.json": map[string]any{"name": "GameType", "values": []any{
				map[string]any{"name": "Survival", "value": 0},
				map[string]any{"name": "Creative", "value": 1},
				map[string]any{"name": "WorldDefault", "value": 0},
			}},
		},
		active: map[string]bool{},
	}
	field := map[string]any{
		"name": "Game Type", "type": "varint32", "enum": "GameType",
		"constraints": map[string]any{"enum_values": []any{0, 1}},
	}
	node := lowerer.applyFieldWrappers(endstoneScalar("varint32"), field, "StartGamePacket.Game Type")
	if node.Kind != manifest.KindEnum || len(node.Variants) != 2 {
		t.Fatalf("enum = %#v, want two constrained values", node)
	}
	if node.Variants[0].Name != "Survival" || node.Variants[1].Name != "Creative" {
		t.Fatalf("variants = %#v, want first canonical names for values 0 and 1", node.Variants)
	}
}

func TestEndstoneSwitchUsesVariantCerealSelectorConstraints(t *testing.T) {
	lowerer := &endstoneLowerer{
		types: map[string]any{
			"MessageOnly.json": map[string]any{"name": "MessageOnly", "fields": []any{
				map[string]any{"name": "Message Type", "type": "uint8", "enum": "MessageType", "constraints": map[string]any{"enum_values": []any{0, 2}}},
				map[string]any{"name": "Message", "type": "string"},
			}},
		},
		enums: map[string]any{
			"MessageType.json": map[string]any{"name": "MessageType", "values": []any{
				map[string]any{"name": "Raw", "value": 0}, map[string]any{"name": "Chat", "value": 1}, map[string]any{"name": "System", "value": 2},
			}},
		},
		active: map[string]bool{},
	}
	node := lowerer.lowerTypeValue(map[string]any{
		"switch": map[string]any{"name": "Message Type", "type": "uint8", "enum": "MessageType"},
		"cases":  []any{"MessageOnly"},
	}, "PacketBody", "Packet.Body")
	if node.Kind != manifest.KindUnion || len(node.Variants) != 2 || node.Variants[0].Value != 0 || node.Variants[1].Value != 2 {
		t.Fatalf("union = %#v, want selectors 0 and 2", node)
	}
	for _, variant := range node.Variants {
		if variant.Encode.Kind != manifest.KindStruct || len(variant.Encode.Fields) != 1 || variant.Encode.Fields[0].Name != "Message" {
			t.Fatalf("variant payload = %#v, want discriminator removed", variant.Encode)
		}
	}
}

func TestEndstoneSwitchRetainsDifferentlyEncodedCompatibilityDiscriminator(t *testing.T) {
	lowerer := &endstoneLowerer{
		types: map[string]any{"Coordinates.json": map[string]any{"name": "Coordinates", "fields": []any{
			map[string]any{"name": "Packet Type", "type": "varint32", "constraints": map[string]any{"enum_values": []any{0}}},
			map[string]any{"name": "X", "type": "float32"},
		}}},
		enums:  map[string]any{},
		active: map[string]bool{},
	}
	node := lowerer.lowerTypeValue(map[string]any{
		"switch": map[string]any{"name": "Packet Type", "type": "uvarint32"},
		"cases":  []any{"Coordinates"},
	}, "Location", "PlayerLocationPacket.Location")
	if node.Kind != manifest.KindUnion || len(node.Variants) != 1 || len(node.Variants[0].Encode.Fields) != 2 {
		t.Fatalf("union = %#v, want compatibility discriminator retained before payload", node)
	}
}

func TestEndstoneSwitchUsesCaseOrderWithoutNamedEnum(t *testing.T) {
	lowerer := &endstoneLowerer{
		types: map[string]any{
			"Stop.json":      map[string]any{"name": "Stop", "fields": []any{}},
			"SetVolume.json": map[string]any{"name": "SetVolume", "fields": []any{map[string]any{"name": "Volume", "type": "float32"}}},
		},
		enums:  map[string]any{},
		active: map[string]bool{},
	}
	node := lowerer.lowerTypeValue(map[string]any{
		"switch": map[string]any{"type": "uvarint32"},
		"cases":  []any{"Stop", "SetVolume"},
	}, "SoundEvent", "Packet.Event")
	if node.Kind != manifest.KindUnion || len(node.Variants) != 2 {
		t.Fatalf("union = %#v, want two variants", node)
	}
	if node.Variants[0].Value != 0 || node.Variants[1].Value != 1 {
		t.Fatalf("selectors = [%d, %d], want [0, 1]", node.Variants[0].Value, node.Variants[1].Value)
	}
}

func TestEndstoneRetainsRepeatedCerealUnionFields(t *testing.T) {
	root := t.TempDir()
	for _, directory := range []string{"packets", "types", "enums"} {
		if err := os.Mkdir(filepath.Join(root, directory), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("**Minecraft Version:** `1.26.40.1`\n**Network Version:** `2168`\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	unionType := map[string]any{
		"switch": map[string]any{"type": "uvarint32"},
		"cases":  []any{"SoundDataEvent::Stop", "SoundDataEvent::SetVolume"},
	}
	writeJSON(t, filepath.Join(root, "packets", "SoundPacket.json"), map[string]any{
		"name": "SoundPacket", "id": 7,
		"fields": []any{
			map[string]any{"name": "Handle", "type": "uint64"},
			map[string]any{"name": "Stop", "type": unionType},
			map[string]any{"name": "SetVolume", "type": unionType},
			map[string]any{"name": "Tail", "type": "bool"},
		},
	})
	writeJSON(t, filepath.Join(root, "types", "Stop.json"), map[string]any{"name": "SoundDataEvent::Stop", "fields": []any{}})
	writeJSON(t, filepath.Join(root, "types", "SetVolume.json"), map[string]any{"name": "SoundDataEvent::SetVolume", "fields": []any{map[string]any{"name": "Volume", "type": "float32"}}})

	result, err := ParseEndstone(root, fixturePin("endstone"), "")
	if err != nil {
		t.Fatalf("ParseEndstone: %v", err)
	}
	if len(result.Claims) != 4 {
		t.Fatalf("claims = %d, want handle, both union slots, and tail", len(result.Claims))
	}
	if result.Claims[1].Encode.Kind != manifest.KindUnion {
		t.Fatalf("first union field = %#v, want union", result.Claims[1].Encode)
	}
	if result.Claims[1].Name != "Stop" || result.Claims[2].Name != "SetVolume" {
		t.Fatalf("union field names = %q, %q, want Stop and SetVolume", result.Claims[1].Name, result.Claims[2].Name)
	}
	if result.Claims[3].Name != "Tail" || result.Claims[3].Ordinal != 3 {
		t.Fatalf("tail = %#v, want source ordinal 3", result.Claims[3])
	}
}

func TestEndstoneRecordsEmptyPackets(t *testing.T) {
	root := t.TempDir()
	for _, directory := range []string{"packets", "types", "enums"} {
		if err := os.Mkdir(filepath.Join(root, directory), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("**Minecraft Version:** `1.26.40.1`\n**Network Version:** `2168`\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	writeJSON(t, filepath.Join(root, "packets", "EmptyPacket.json"), map[string]any{"name": "EmptyPacket", "id": 4, "fields": []any{}})

	result, err := ParseEndstone(root, fixturePin("endstone"), "")
	if err != nil {
		t.Fatalf("ParseEndstone: %v", err)
	}
	if len(result.Packets) != 1 || result.Packets[0].ID != 4 || result.Packets[0].Name != "EmptyPacket" {
		t.Fatalf("packet claims = %#v, want empty packet metadata", result.Packets)
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

func TestCorrectionCanAddMissingObjectMetadataWithoutOverwriting(t *testing.T) {
	documents := map[string]any{"Packet.json": map[string]any{"field": map[string]any{"name": "Value", "type": "uint8"}}}
	before, err := canonicalDigest(documents["Packet.json"].(map[string]any)["field"])
	if err != nil {
		t.Fatal(err)
	}
	corrections := t.TempDir()
	writeJSON(t, filepath.Join(corrections, "optional.json"), map[string]any{"operations": []any{map[string]any{
		"id": "optional-field", "file": "Packet.json", "pointer": "/field", "pre_patch_sha256": before,
		"merge": map[string]any{"optional": true}, "why": "fixture optional evidence", "evidence": []any{map[string]any{"locator": "fixture/runtime"}},
	}}})
	if _, err := applyCorrections(documents, corrections, fixturePin("endstone")); err != nil {
		t.Fatalf("applyCorrections: %v", err)
	}
	field := documents["Packet.json"].(map[string]any)["field"].(map[string]any)
	if field["optional"] != true {
		t.Fatalf("field = %#v, want optional metadata merged", field)
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

func TestMojangMissingOrdinalDiagnosticIsDeterministic(t *testing.T) {
	properties := map[string]any{
		"Value": map[string]any{"type": "integer"},
		"Key":   map[string]any{"type": "string"},
	}
	_, err := lowerMojangFields(nil, 0, "MapEntry", manifest.DirectionUnknown, "MapEntry.json", properties, nil, "")
	if err == nil || !strings.Contains(err.Error(), "property Key has no explicit ordinal") {
		t.Fatalf("lowerMojangFields error = %v, want the first lexical property", err)
	}
}

func TestMojangMapEntrySchemasLowerKeyAndValue(t *testing.T) {
	lowerer := &mojangLowerer{documents: map[string]any{}, active: map[string]bool{}}
	node := lowerer.lowerSchema(map[string]any{
		"type": "object",
		"additionalProperties": map[string]any{
			"type": "object",
			"properties": map[string]any{
				"key":   map[string]any{"type": "integer", "x-underlying-type": "uint16"},
				"value": map[string]any{"type": "number", "x-underlying-type": "float"},
			},
		},
	}, "Map.json", "Fixture")
	if node.Kind != manifest.KindMap || node.Key == nil || node.Value == nil {
		t.Fatalf("node = %#v, want map with key and value", node)
	}
	if node.Key.Primitive == nil || node.Key.Primitive.Code != "u16le" || node.Value.Primitive == nil || node.Value.Primitive.Code != "f32le" {
		t.Fatalf("map key/value = %#v/%#v, want u16le/f32le", node.Key, node.Value)
	}
}

func TestEndstoneSpecialUnknownsStayReachableUnresolved(t *testing.T) {
	lowerer := &endstoneLowerer{types: map[string]any{}, enums: map[string]any{}, active: map[string]bool{}}
	for _, name := range []string{"cereal::UnknownBuiltin"} {
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

func TestEndstoneScalarSpellingsUseCanonicalPrimitives(t *testing.T) {
	tests := map[string]string{
		"int32_be":    "i32be",
		"varint32":    "zigzag_i32",
		"varint64":    "zigzag_i64",
		"uvarint32":   "var_u32",
		"uvarint64":   "var_u64",
		"mce::UUID":   "uuid",
		"CompoundTag": "nbt_le",
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			node := endstoneScalar(input)
			if node.Kind != manifest.KindPrimitive || node.Primitive == nil || node.Primitive.Code != want {
				t.Fatalf("endstoneScalar(%q) = %#v, want %s", input, node, want)
			}
		})
	}
}

func TestEndstoneFixedWidthBitsetRetainsItsCerealBitCount(t *testing.T) {
	node := endstoneScalar("brstd::bitset<131>")
	if node.Kind != manifest.NodeKind("bitset") || node.Length != 131 {
		t.Fatalf("bitset = %#v, want 131-bit wire shape", node)
	}
}

func TestEndstoneAlwaysPresentMarkersFoldIntoOptionals(t *testing.T) {
	fields, err := canonicalEndstoneFields([]any{
		map[string]any{"type": "bool", "value": true},
		map[string]any{"name": "Value", "type": "uint32", "optional": true},
	}, "Fixture")
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || !fields[0].outerOptional || fields[0].sourceIndex != 1 {
		t.Fatalf("canonical fields = %#v, want one outer optional sourced from field 1", fields)
	}
	lowerer := &endstoneLowerer{}
	node := lowerer.lowerTypeValue(fields[0].object["type"], "Value", "Fixture.Value")
	node = lowerer.applyFieldWrappers(node, fields[0].object, "Fixture.Value")
	if fields[0].outerOptional {
		node = manifest.Optional(node)
	}
	if node.Kind != manifest.KindOptional || node.Value == nil || node.Value.Kind != manifest.KindOptional {
		t.Fatalf("node = %#v, want double optional", node)
	}
}

func TestEndstoneCerealDynamicValueLowersToRecursiveTaggedValue(t *testing.T) {
	lowerer := &endstoneLowerer{types: map[string]any{}, enums: map[string]any{}, active: map[string]bool{}}
	node := lowerer.lowerNamed("cereal::DynamicValue", "fixture")
	if node.Kind != manifest.KindUnion || node.TypeID != "cereal::DynamicValue" || len(node.Variants) != 7 {
		t.Fatalf("dynamic value = %#v, want seven-way recursive union", node)
	}
	if node.Control == nil || node.Control.Primitive == nil || node.Control.Primitive.Code != "i32le" {
		t.Fatalf("dynamic value control = %#v, want fixed i32", node.Control)
	}
	list := node.Variants[5].Encode
	if list.Kind != manifest.KindArray || list.Element == nil || list.Element.Kind != manifest.KindRecursive || list.Element.Target != "cereal::DynamicValue" {
		t.Fatalf("dynamic list = %#v, want recursive value array", list)
	}
	mapping := node.Variants[6].Encode
	if mapping.Kind != manifest.KindMap || mapping.Value == nil || mapping.Value.Kind != manifest.KindRecursive || mapping.Value.Target != "cereal::DynamicValue" {
		t.Fatalf("dynamic map = %#v, want recursive string/value map", mapping)
	}
}

func TestEndstoneNamedDocumentUsesDeclaredName(t *testing.T) {
	document := map[string]any{"name": "Connection::DisconnectFailReason", "values": []any{}}
	got, ok := namedDocument(map[string]any{"Connection__DisconnectFailReason.json": document}, "Connection::DisconnectFailReason")
	if !ok || got == nil {
		t.Fatal("namedDocument did not match the document's declared C++ name")
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
