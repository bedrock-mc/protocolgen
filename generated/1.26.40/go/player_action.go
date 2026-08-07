// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerAction struct {
	PlayerRuntimeID ActorRuntimeID
	Action          PlayerActionType
	BlockPosition   BlockPos
	ResultPos       BlockPos
	Face            int32
}

func (p *PlayerAction) Encode(w Encoder) error {
	if err := w.Write("PlayerActionPacket.Player Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.PlayerRuntimeID); err != nil {
		return err
	}
	if err := w.Write("PlayerActionPacket.Action", Shape{Kind: "enum", Semantic: "PlayerActionType", TypeID: "enums/PlayerActionType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "StartDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "AbortDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "StopDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "GetUpdatedBlock", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "DropItem", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "StartSleeping", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "StopSleeping", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Respawn", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "StartJump", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "StartSprinting", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "StopSprinting", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "StartSneaking", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "StopSneaking", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "CreativeDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "ChangeDimensionAck", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "StartGliding", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "StopGliding", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "DenyDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "CrackBlock", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "ChangeSkin", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "UpdatedEnchantingSeed", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "StartSwimming", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "StopSwimming", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "StartSpinAttack", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "StopSpinAttack", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "InteractWithBlock", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "PredictDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "ContinueDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "StartItemUseOn", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "StopItemUseOn", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "HandledTeleport", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "MissedSwing", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "StartCrawling", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "StopCrawling", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "StartFlying", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "StopFlying", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "ClientAckServerData", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "StartUsingItem", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "InternalUpdate", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "Count", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("PlayerActionPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("PlayerActionPacket.Result Pos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.ResultPos); err != nil {
		return err
	}
	if err := w.Write("PlayerActionPacket.Face", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Face); err != nil {
		return err
	}
	return nil
}

func DecodePlayerAction(r Decoder) (PlayerAction, error) {
	var p PlayerAction
	{
		raw, err := r.Read("PlayerActionPacket.Player Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field PlayerActionPacket.Player Runtime ID has unexpected decoded type %T", raw)
		}
		p.PlayerRuntimeID = value
	}
	{
		raw, err := r.Read("PlayerActionPacket.Action", Shape{Kind: "enum", Semantic: "PlayerActionType", TypeID: "enums/PlayerActionType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Unknown", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "StartDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "AbortDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "StopDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "GetUpdatedBlock", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "DropItem", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "StartSleeping", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "StopSleeping", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "Respawn", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "StartJump", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "StartSprinting", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "StopSprinting", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "StartSneaking", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "StopSneaking", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "CreativeDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "ChangeDimensionAck", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "StartGliding", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "StopGliding", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "DenyDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "CrackBlock", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "ChangeSkin", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "UpdatedEnchantingSeed", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "StartSwimming", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "StopSwimming", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "StartSpinAttack", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "StopSpinAttack", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "InteractWithBlock", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "PredictDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "ContinueDestroyBlock", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "StartItemUseOn", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "StopItemUseOn", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "HandledTeleport", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "MissedSwing", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "StartCrawling", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "StopCrawling", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "StartFlying", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "StopFlying", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "ClientAckServerData", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "StartUsingItem", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "InternalUpdate", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "Count", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerActionType)
		if !ok {
			return p, fmt.Errorf("field PlayerActionPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("PlayerActionPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field PlayerActionPacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("PlayerActionPacket.Result Pos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field PlayerActionPacket.Result Pos has unexpected decoded type %T", raw)
		}
		p.ResultPos = value
	}
	{
		raw, err := r.Read("PlayerActionPacket.Face", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PlayerActionPacket.Face has unexpected decoded type %T", raw)
		}
		p.Face = value
	}
	return p, nil
}
