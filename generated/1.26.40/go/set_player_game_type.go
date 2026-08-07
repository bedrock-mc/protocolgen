// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetPlayerGameType struct {
	PlayerGameType GameType
}

func (p *SetPlayerGameType) Encode(w Encoder) error {
	if err := w.Write("SetPlayerGameTypePacket.Player Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}, p.PlayerGameType); err != nil {
		return err
	}
	return nil
}

func DecodeSetPlayerGameType(r Decoder) (SetPlayerGameType, error) {
	var p SetPlayerGameType
	{
		raw, err := r.Read("SetPlayerGameTypePacket.Player Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GameType)
		if !ok {
			return p, fmt.Errorf("field SetPlayerGameTypePacket.Player Game Type has unexpected decoded type %T", raw)
		}
		p.PlayerGameType = value
	}
	return p, nil
}
