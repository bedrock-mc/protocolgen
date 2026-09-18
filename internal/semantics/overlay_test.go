package semantics

import (
	"testing"

	"protocolgen/internal/manifest"
)

func fixture() manifest.Manifest {
	return manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Packets: []manifest.Packet{{ID: 1, Name: "PickPacket", Fields: []manifest.Field{
			{Ordinal: 0, Name: "Actor ID", Encode: manifest.Primitive("i64le"), Symmetry: manifest.Symmetric},
			{Ordinal: 1, Name: "Name", Encode: manifest.String(manifest.Primitive("var_u32")), Symmetry: manifest.Symmetric},
		}}},
	}
}

func TestValidateRejectsFieldsThatCannotCarryAnIdentifier(t *testing.T) {
	m := fixture()
	for name, document := range map[string]Document{
		"unknown field": {SchemaVersion: 1, Target: m.Target, Entries: []Entry{{TypeID: "PickPacket", Field: "Nope", Semantic: ActorUniqueID}}},
		"string field":  {SchemaVersion: 1, Target: m.Target, Entries: []Entry{{TypeID: "PickPacket", Field: "Name", Semantic: ActorUniqueID}}},
		"runtime i64le": {SchemaVersion: 1, Target: m.Target, Entries: []Entry{{TypeID: "PickPacket", Field: "Actor ID", Semantic: ActorRuntimeID}}},
		"bad semantic":  {SchemaVersion: 1, Target: m.Target, Entries: []Entry{{TypeID: "PickPacket", Field: "Actor ID", Semantic: "PlayerID"}}},
	} {
		if err := ValidateOverlay(m, document); err == nil {
			t.Errorf("%s was accepted", name)
		}
	}
	if err := ValidateOverlay(m, Document{SchemaVersion: 1, Target: m.Target, Entries: []Entry{{TypeID: "PickPacket", Field: "Actor ID", Semantic: ActorUniqueID}}}); err != nil {
		t.Fatal(err)
	}
}

// Apply wraps a marked integer in the identifier struct and leaves other
// nodes alone.
func TestApplyWrapsOnlyMarkedIntegers(t *testing.T) {
	m := fixture()
	overlay := Overlay{Fields: map[string]string{FieldKey("PickPacket", "Actor ID"): ActorUniqueID}}
	wrapped := overlay.Apply("PickPacket", m.Packets[0].Fields[0])
	if wrapped.Kind != manifest.KindStruct || wrapped.TypeID != ActorUniqueID || len(wrapped.Fields) != 1 || wrapped.Fields[0].Encode.Primitive.Code != "i64le" {
		t.Fatalf("Apply = %#v", wrapped)
	}
	if plain := overlay.Apply("PickPacket", m.Packets[0].Fields[1]); plain.Kind != manifest.KindString {
		t.Fatalf("unmarked field was changed: %#v", plain)
	}
}
