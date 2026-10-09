package generator

import (
	"encoding/json"
	"fmt"
)

// rawBlockShapes holds the fields supplied by the pinned block geometry source.
type rawBlockShapes struct {
	Name           string          `json:"name"`
	Hash           uint32          `json:"blockStateHash"`
	CollisionShape json.RawMessage `json:"collisionShape"`
	OutlineShape   json.RawMessage `json:"shape"`
	TintMethod     string          `json:"tintMethod"`
}

// applyBlockShapes joins the geometry source to every state by both hash and
// name. Other block properties stay with their original source.
func applyBlockShapes(states []rawBlockState, path string) error {
	var shapes []rawBlockShapes
	if err := readJSON(path, &shapes, false); err != nil {
		return fmt.Errorf("read block shapes: %w", err)
	}
	if len(shapes) != len(states) || len(shapes) == 0 {
		return fmt.Errorf("block shape coverage: got %d states, want %d", len(shapes), len(states))
	}
	byHash := make(map[uint32]rawBlockShapes, len(shapes))
	for _, shape := range shapes {
		if _, ok := byHash[shape.Hash]; ok {
			return fmt.Errorf("duplicate block shape hash %d", shape.Hash)
		}
		byHash[shape.Hash] = shape
	}
	for i := range states {
		state := &states[i]
		shape, ok := byHash[state.Hash]
		if !ok {
			return fmt.Errorf("block shapes missing %s hash %d", state.Name, state.Hash)
		}
		if shape.Name != state.Name {
			return fmt.Errorf("block shape hash %d names %s, want %s", state.Hash, shape.Name, state.Name)
		}
		if _, present, err := decodeShape(shape.CollisionShape); err != nil || !present {
			return fmt.Errorf("%s block collision shape is missing or invalid: %v", state.Name, err)
		}
		outline, present, err := decodeShape(shape.OutlineShape)
		if err != nil || !present {
			return fmt.Errorf("%s block outline shape is missing or invalid: %v", state.Name, err)
		}
		// A zero-volume outline has no visible extent. Retain it as a known
		// empty shape, which is different from an unavailable shape.
		visible := make([]generatedBox, 0, len(outline))
		for _, box := range outline {
			if box[3] > box[0] && box[4] > box[1] && box[5] > box[2] {
				visible = append(visible, box)
			}
		}
		encoded, err := json.Marshal(visible)
		if err != nil {
			return fmt.Errorf("%s block outline shape: %w", state.Name, err)
		}
		state.CollisionShape = shape.CollisionShape
		state.OutlineShape = encoded
		state.TintMethod = shape.TintMethod
		delete(byHash, state.Hash)
	}
	return nil
}
