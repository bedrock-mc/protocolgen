// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetDefaultGameType struct {
	DefaultGameType GameType
}

func (p *SetDefaultGameType) Encode(w Encoder) error {
	if err := w.Write("SetDefaultGameTypePacket.Default Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}, p.DefaultGameType); err != nil {
		return err
	}
	return nil
}

func DecodeSetDefaultGameType(r Decoder) (SetDefaultGameType, error) {
	var p SetDefaultGameType
	{
		raw, err := r.Read("SetDefaultGameTypePacket.Default Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GameType)
		if !ok {
			return p, fmt.Errorf("field SetDefaultGameTypePacket.Default Game Type has unexpected decoded type %T", raw)
		}
		p.DefaultGameType = value
	}
	return p, nil
}
