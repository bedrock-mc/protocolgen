package domains

import (
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func domainFixture() manifest.Manifest {
	first := manifest.Node{Kind: manifest.KindStruct, TypeID: "First", Fields: []manifest.Field{{Name: "Value", Encode: manifest.Primitive("u8")}}}
	second := manifest.Node{Kind: manifest.KindStruct, TypeID: "Second", Fields: []manifest.Field{{Name: "Value", Encode: manifest.Primitive("u8")}}}
	return manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Packets: []manifest.Packet{{Name: "FixturePacket", Fields: []manifest.Field{
			{TypeID: "First", Encode: first},
			{TypeID: "Second", Encode: second},
		}}},
	}
}

func TestValidateOverlayRejectsStaleTypeID(t *testing.T) {
	m := domainFixture()
	document := Document{
		SchemaVersion: 1,
		Target:        m.Target,
		Entries:       []Entry{{TypeID: "Missing", Domain: "misc", Rationale: "fixture"}},
	}
	if err := ValidateOverlay(m, document); err == nil || !strings.Contains(err.Error(), "Missing") {
		t.Fatalf("ValidateOverlay error = %v, want stale TypeID", err)
	}
}

func TestValidateOverlayListsUnassignedSharedTypes(t *testing.T) {
	m := domainFixture()
	document := Document{
		SchemaVersion: 1,
		Target:        m.Target,
		Entries:       []Entry{{TypeID: "First", Domain: "misc", Rationale: "fixture"}},
	}
	if err := ValidateOverlay(m, document); err == nil || !strings.Contains(err.Error(), "Second") {
		t.Fatalf("ValidateOverlay error = %v, want unassigned Second", err)
	}
}
