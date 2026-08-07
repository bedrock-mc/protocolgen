// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdatePlayerGameType struct {
	PlayerGameType GameType
	TargetPlayer   ActorUniqueID
	Tick           PlayerInputTick
}

func (p *UpdatePlayerGameType) Encode(w Encoder) error {
	if err := w.Write("UpdatePlayerGameTypePacket.Player Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}, p.PlayerGameType); err != nil {
		return err
	}
	if err := w.Write("UpdatePlayerGameTypePacket.Target player", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetPlayer); err != nil {
		return err
	}
	if err := w.Write("UpdatePlayerGameTypePacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.Tick); err != nil {
		return err
	}
	return nil
}

func DecodeUpdatePlayerGameType(r Decoder) (UpdatePlayerGameType, error) {
	var p UpdatePlayerGameType
	{
		raw, err := r.Read("UpdatePlayerGameTypePacket.Player Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GameType)
		if !ok {
			return p, fmt.Errorf("field UpdatePlayerGameTypePacket.Player Game Type has unexpected decoded type %T", raw)
		}
		p.PlayerGameType = value
	}
	{
		raw, err := r.Read("UpdatePlayerGameTypePacket.Target player", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field UpdatePlayerGameTypePacket.Target player has unexpected decoded type %T", raw)
		}
		p.TargetPlayer = value
	}
	{
		raw, err := r.Read("UpdatePlayerGameTypePacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerInputTick)
		if !ok {
			return p, fmt.Errorf("field UpdatePlayerGameTypePacket.Tick has unexpected decoded type %T", raw)
		}
		p.Tick = value
	}
	return p, nil
}
