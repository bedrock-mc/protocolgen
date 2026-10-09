package generator

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// TestPrepareBlockShapesPreservesFacts checks that only reviewed liquid geometry
// and zero-volume outlines change; collision and scalar facts remain intact.
func TestPrepareBlockShapesPreservesFacts(t *testing.T) {
	states := []rawBlockState{{
		Name: "minecraft:stone", Hash: 100, Hardness: 3, Friction: 0.4,
		CollisionShape: json.RawMessage(`[[0.25,0,0.25,0.75,1,0.75]]`),
		OutlineShape:   json.RawMessage(`[0,0,0,1,0,1]`),
		VisualShape:    json.RawMessage(`[0,0,0,1,1,1]`), TintMethod: "Grass",
		LiquidClipShape: json.RawMessage(`[0.000065,0,0,0,0,0]`),
	}}
	want := states[0]
	want.OutlineShape = json.RawMessage(`[]`)
	want.LiquidClipShape = nil
	omissions := []LiquidClipOmission{{Name: "minecraft:stone", Hash: 100, SourceBox: []float64{0.000065, 0, 0, 0, 0, 0}}}
	if err := prepareBlockShapes(states, omissions); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(states[0], want) {
		t.Fatalf("prepared state = %+v, want %+v", states[0], want)
	}
	if _, available, err := decodeShape(states[0].LiquidClipShape); err != nil || available {
		t.Fatalf("omitted liquid shape must be unavailable, got available=%t, error=%v", available, err)
	}
	if boxes, available, err := decodeShape(states[0].OutlineShape); err != nil || !available || len(boxes) != 0 {
		t.Fatalf("outline must stay known-empty, got %v, available=%t, error=%v", boxes, available, err)
	}
}

// TestPrepareBlockShapesRejectsSourceDrift ensures exceptions cannot hide new,
// changed, or already-fixed source problems.
func TestPrepareBlockShapesRejectsSourceDrift(t *testing.T) {
	valid := rawBlockState{Name: "minecraft:stone", Hash: 100, CollisionShape: json.RawMessage(`[]`), OutlineShape: json.RawMessage(`[]`)}
	for _, test := range []struct {
		name      string
		shape     string
		omissions []LiquidClipOmission
		want      string
	}{
		{"unreviewed", `[1,0,0,0,1,1]`, nil, "liquid clip shape"},
		{"changed_box", `[2,0,0,0,1,1]`, []LiquidClipOmission{{"minecraft:stone", 100, []float64{1, 0, 0, 0, 1, 1}}}, "no longer matches"},
		{"wrong_name", `[1,0,0,0,1,1]`, []LiquidClipOmission{{"minecraft:air", 100, []float64{1, 0, 0, 0, 1, 1}}}, "no longer matches"},
		{"missing_state", `[]`, []LiquidClipOmission{{"minecraft:stone", 99, []float64{1, 0, 0, 0, 1, 1}}}, "source state is missing"},
		{"valid_box", `[0,0,0,1,1,1]`, []LiquidClipOmission{{"minecraft:stone", 100, []float64{0, 0, 0, 1, 1, 1}}}, "source shape is valid"},
		{"invalid_entry", `[]`, []LiquidClipOmission{{Name: "minecraft:stone", Hash: 100}}, "incomplete"},
		{"duplicate", `[1,0,0,0,1,1]`, []LiquidClipOmission{{"minecraft:stone", 100, []float64{1, 0, 0, 0, 1, 1}}, {"minecraft:stone", 100, []float64{1, 0, 0, 0, 1, 1}}}, "duplicate"},
	} {
		t.Run(test.name, func(t *testing.T) {
			state := valid
			state.LiquidClipShape = json.RawMessage(test.shape)
			if err := prepareBlockShapes([]rawBlockState{state}, test.omissions); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
	for _, field := range []string{"collision", "outline"} {
		t.Run(field, func(t *testing.T) {
			state := valid
			if field == "collision" {
				state.CollisionShape = nil
			} else {
				state.OutlineShape = json.RawMessage(`[1,0,0,0,1,1]`)
			}
			if err := prepareBlockShapes([]rawBlockState{state}, nil); err == nil || !strings.Contains(err.Error(), field+" shape") {
				t.Fatalf("error = %v, want required geometry failure", err)
			}
		})
	}
}
