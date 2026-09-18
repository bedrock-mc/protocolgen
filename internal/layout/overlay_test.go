package layout

import (
	"strings"
	"testing"

	"protocolgen/internal/manifest"
)

func fixture() manifest.Manifest {
	kind := manifest.Enum("u8", manifest.EnumValue{Value: 0, Name: "Survival"}, manifest.EnumValue{Value: 1, Name: "Creative"})
	kind.TypeID = "enums/GameType"
	return manifest.Manifest{
		SchemaVersion: 2,
		Target:        manifest.Target{MinecraftVersion: "fixture", ProtocolVersion: 1},
		Packets: []manifest.Packet{{ID: 1, Name: "SetGameTypePacket", Fields: []manifest.Field{
			{Ordinal: 0, Name: "Game Type", Encode: kind, Symmetry: manifest.Symmetric},
		}}},
	}
}

func TestValidateRejectsUnknownTargets(t *testing.T) {
	m := fixture()
	base := Document{SchemaVersion: 1, Target: m.Target}
	for name, document := range map[string]Document{
		"unknown enum":    {SchemaVersion: 1, Target: m.Target, Constants: []ConstantEntry{{TypeID: "enums/Nope", Package: "packet", File: "x", Names: map[string]string{}}}},
		"unknown variant": {SchemaVersion: 1, Target: m.Target, Constants: []ConstantEntry{{TypeID: "enums/GameType", Package: "packet", File: "x", Names: map[string]string{"Spectator": "GameTypeSpectator"}}}},
		"bad package":     {SchemaVersion: 1, Target: m.Target, Constants: []ConstantEntry{{TypeID: "enums/GameType", Package: "minecraft", File: "x"}}},
		"unexported name": {SchemaVersion: 1, Target: m.Target, Constants: []ConstantEntry{{TypeID: "enums/GameType", Package: "packet", File: "x", Names: map[string]string{"Creative": "creative"}}}},
		"unknown field":   {SchemaVersion: 1, Target: m.Target, Fields: []FieldEntry{{TypeID: "SetGameTypePacket", Field: "Mode", Name: "Mode"}}},
		"duplicate field": {SchemaVersion: 1, Target: m.Target, Fields: []FieldEntry{{TypeID: "SetGameTypePacket", Field: "Game Type", Name: "A"}, {TypeID: "SetGameTypePacket", Field: "Game Type", Name: "B"}}},
		"wrong target":    {SchemaVersion: 1, Target: manifest.Target{MinecraftVersion: "other", ProtocolVersion: 1}},
		"unknown schema":  {SchemaVersion: 2, Target: m.Target},
	} {
		if err := ValidateOverlay(m, document); err == nil {
			t.Errorf("%s was accepted", name)
		}
	}
	if err := ValidateOverlay(m, base); err != nil {
		t.Fatal(err)
	}
	valid := Document{SchemaVersion: 1, Target: m.Target,
		Constants: []ConstantEntry{{TypeID: "enums/GameType", Package: "packet", File: "set_game_type", Names: map[string]string{"Creative": "GameTypeCreative"}}},
		Fields:    []FieldEntry{{TypeID: "SetGameTypePacket", Field: "Game Type", Name: "PlayerGameMode"}}}
	if err := ValidateOverlay(m, valid); err != nil {
		t.Fatal(err)
	}
}

func TestFieldNameLooksUpByOwnerAndWireName(t *testing.T) {
	overlay := Overlay{Fields: map[string]string{FieldKey("SetGameTypePacket", "Game Type"): "PlayerGameMode"}}
	if got := overlay.FieldName("SetGameTypePacket", "Game Type"); got != "PlayerGameMode" {
		t.Fatalf("FieldName = %q", got)
	}
	if got := overlay.FieldName("SetGameTypePacket", "Other"); got != "" {
		t.Fatalf("unmapped field returned %q", got)
	}
	if !strings.Contains(FieldKey("a", "b"), "\x00") {
		t.Fatal("field key does not separate owner and field")
	}
}
