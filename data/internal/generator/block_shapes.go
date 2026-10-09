package generator

import (
	"encoding/json"
	"fmt"
	"slices"
)

// LiquidClipOmission identifies one reviewed invalid box in the locked extract.
// Its geometry is unavailable; it must not be replaced with an empty or guessed box.
type LiquidClipOmission struct {
	Name      string    `json:"name"`
	Hash      uint32    `json:"hash"`
	SourceBox []float64 `json:"source_box"`
}

// prepareBlockShapes validates required geometry and applies exact source omissions.
// Empty outlines keep their known-empty meaning without retaining zero-volume boxes.
func prepareBlockShapes(states []rawBlockState, omissions []LiquidClipOmission) error {
	remaining := make(map[uint32]LiquidClipOmission, len(omissions))
	for _, omission := range omissions {
		if omission.Name == "" || len(omission.SourceBox) != 6 {
			return fmt.Errorf("liquid clip omission %d has incomplete identity or geometry", omission.Hash)
		}
		if _, ok := remaining[omission.Hash]; ok {
			return fmt.Errorf("duplicate liquid clip omission %d", omission.Hash)
		}
		remaining[omission.Hash] = omission
	}
	for i := range states {
		state := &states[i]
		if _, present, err := decodeShape(state.CollisionShape); err != nil || !present {
			return fmt.Errorf("%s collision shape is missing or invalid: %v", state.Name, err)
		}
		outline, present, err := decodeShape(state.OutlineShape)
		if err != nil || !present {
			return fmt.Errorf("%s outline shape is missing or invalid: %v", state.Name, err)
		}
		visible := make([]generatedBox, 0, len(outline))
		for _, box := range outline {
			if box[3] > box[0] && box[4] > box[1] && box[5] > box[2] {
				visible = append(visible, box)
			}
		}
		state.OutlineShape, err = json.Marshal(visible)
		if err != nil {
			return fmt.Errorf("%s outline shape: %w", state.Name, err)
		}

		_, _, shapeErr := decodeShape(state.LiquidClipShape)
		omission, omitted := remaining[state.Hash]
		if omitted {
			var actual []float64
			if state.Name != omission.Name || json.Unmarshal(state.LiquidClipShape, &actual) != nil || !slices.Equal(actual, omission.SourceBox) {
				return fmt.Errorf("liquid clip omission %d no longer matches its source name and box", state.Hash)
			}
			if shapeErr == nil {
				return fmt.Errorf("liquid clip omission %d is obsolete: the source shape is valid", state.Hash)
			}
			state.LiquidClipShape = nil
			delete(remaining, state.Hash)
		} else if shapeErr != nil {
			return fmt.Errorf("%s hash %d liquid clip shape: %w", state.Name, state.Hash, shapeErr)
		}
	}
	for _, omission := range omissions {
		if _, ok := remaining[omission.Hash]; ok {
			return fmt.Errorf("liquid clip omission %d is obsolete: its source state is missing", omission.Hash)
		}
	}
	return nil
}
