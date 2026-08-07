package manifest

import (
	"strings"
	"testing"
)

func TestValidateAcceptsCompleteV2WireVocabulary(t *testing.T) {
	manifest := Manifest{
		SchemaVersion: 2,
		Target:        Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture-2168", Digest: "sha256:fixture"}},
		Packets: []Packet{{
			ID:        1,
			Name:      "Vocabulary",
			Direction: DirectionClientbound,
			Fields: []Field{
				{Ordinal: 0, Name: "Nested", Semantic: "NestedValue", Encode: Optional(Optional(Struct(
					Field{Ordinal: 0, Name: "Values", Encode: Array(Primitive("var_u32"), FixedArray(2, Primitive("u16le"))), Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}},
				))), Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}},
				{Ordinal: 1, Name: "Choice", Encode: Union(
					Primitive("var_u32"),
					Variant{Value: 0, Name: "None", Encode: Void()},
					Variant{Value: 7, Name: "Payload", Encode: Primitive("bytes")},
				), Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}},
				{Ordinal: 2, Name: "Mode", Encode: Enum("u8", EnumValue{Name: "Ready", Value: 4}), Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}},
				{Ordinal: 3, Name: "Compatibility", Encode: Reserved(Primitive("u8")), Reserved: true, Ignored: true, Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}},
				{Ordinal: 4, Name: "Inline", Encode: Sequence(Primitive("u8"), String(Primitive("var_u32"))), Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}},
			},
		}},
	}

	if err := Validate(manifest); err != nil {
		t.Fatalf("Validate: %v", err)
	}

	data, err := MarshalStable(manifest)
	if err != nil {
		t.Fatalf("MarshalStable: %v", err)
	}
	if !strings.Contains(string(data), `"schema_version": 2`) || !strings.Contains(string(data), `"kind": "optional"`) {
		t.Fatalf("stable manifest omitted canonical fields:\n%s", data)
	}
}

func TestValidateRejectsReachableUnresolvedNode(t *testing.T) {
	manifest := Manifest{
		SchemaVersion: 2,
		Target:        Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture-2168", Digest: "sha256:fixture"}},
		Packets: []Packet{{
			ID: 1, Name: "Blocked", Direction: DirectionClientbound,
			Fields: []Field{{Ordinal: 0, Name: "Unknown", Encode: Unresolved("dynamic-value", true), Symmetry: Symmetric, Provenance: Provenance{Pins: []string{"fixture"}}}},
		}},
	}

	err := Validate(manifest)
	if err == nil || !strings.Contains(err.Error(), "unresolved") {
		t.Fatalf("Validate error = %v, want reachable unresolved failure", err)
	}
	manifest.Packets[0].Fields[0].Encode.Reachable = false
	if err := Validate(manifest); err == nil || !strings.Contains(err.Error(), "unresolved") {
		t.Fatalf("Validate non-reachable flag error = %v, want fail-closed unresolved failure", err)
	}
}

func TestValidateRejectsUnknownPacketDirection(t *testing.T) {
	value := Manifest{
		SchemaVersion: SchemaVersion,
		Target:        Target{MinecraftVersion: "fixture", ProtocolVersion: 2168},
		Sources:       []SourcePin{{ID: "fixture", Kind: "synthetic", Revision: "fixture-2168", Digest: "sha256:fixture"}},
		Packets:       []Packet{{ID: 1, Name: "UnknownDirection", Direction: DirectionUnknown}},
	}
	if err := Validate(value); err == nil || !strings.Contains(err.Error(), "unknown direction") {
		t.Fatalf("Validate error = %v, want unknown direction failure", err)
	}
}

func TestPrimitiveVarintAndZigZagRemainDistinct(t *testing.T) {
	varint, err := PrimitiveForCode("var_i32")
	if err != nil {
		t.Fatal(err)
	}
	zigzag, err := PrimitiveForCode("zigzag_i32")
	if err != nil {
		t.Fatal(err)
	}
	if varint.Signed != zigzag.Signed || varint.Width != zigzag.Width || varint.ZigZag || !zigzag.ZigZag {
		t.Fatalf("varint=%+v zigzag=%+v", varint, zigzag)
	}
}

func TestValidateRejectsMixedProtocolTargetSources(t *testing.T) {
	manifest := Manifest{
		SchemaVersion: 2,
		Target:        Target{MinecraftVersion: "1.26.40", ProtocolVersion: 2168},
		Sources: []SourcePin{
			{ID: "old", Kind: "mojang", Revision: "old", Digest: "sha256:old", MinecraftVersion: "1.26.40", ProtocolVersion: 2168},
			{ID: "new", Kind: "mojang", Revision: "new", Digest: "sha256:new", MinecraftVersion: "1.26.50", ProtocolVersion: 2169},
		},
		Packets: []Packet{{ID: 1, Name: "P", Direction: DirectionUnknown}},
	}
	if err := Validate(manifest); err == nil || !strings.Contains(err.Error(), "mix") {
		t.Fatalf("Validate error = %v, want mixed-source failure", err)
	}
}
