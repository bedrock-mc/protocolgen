// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MovePlayer struct {
	PlayerRuntimeID ActorRuntimeID
	Position        Vec3
	Rotation        Vec2
	YHeadRotation   float32
	PositionMode    PlayerPositionModeComponentPositionMode
	OnGround        bool
	RidingRuntimeID ActorRuntimeID
	TeleportData    *MovePlayerTeleportData
	Tick            PlayerInputTick
}

func (p *MovePlayer) Encode(w Encoder) error {
	if err := w.Write("MovePlayerPacket.Player Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.PlayerRuntimeID); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Rotation); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Y-Head Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.YHeadRotation); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Position Mode", Shape{Kind: "enum", Semantic: "PlayerPositionModeComponent::PositionMode", TypeID: "enums/PlayerPositionModeComponent::PositionMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Normal", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Respawn", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Teleport", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "OnlyHeadRot", Shape: Shape{Kind: "void"}}}}, p.PositionMode); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.On Ground", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.OnGround); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Riding Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.RidingRuntimeID); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Teleport Data", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "MovePlayerTeleportData", TypeID: "MovePlayerTeleportData", Fields: []ShapeField{{Ordinal: 0, Name: "Teleportation Cause", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Source Actor Type", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}}, p.TeleportData); err != nil {
		return err
	}
	if err := w.Write("MovePlayerPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.Tick); err != nil {
		return err
	}
	return nil
}

func DecodeMovePlayer(r Decoder) (MovePlayer, error) {
	var p MovePlayer
	{
		raw, err := r.Read("MovePlayerPacket.Player Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Player Runtime ID has unexpected decoded type %T", raw)
		}
		p.PlayerRuntimeID = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec2)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Rotation has unexpected decoded type %T", raw)
		}
		p.Rotation = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Y-Head Rotation", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Y-Head Rotation has unexpected decoded type %T", raw)
		}
		p.YHeadRotation = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Position Mode", Shape{Kind: "enum", Semantic: "PlayerPositionModeComponent::PositionMode", TypeID: "enums/PlayerPositionModeComponent::PositionMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Normal", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Respawn", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Teleport", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "OnlyHeadRot", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerPositionModeComponentPositionMode)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Position Mode has unexpected decoded type %T", raw)
		}
		p.PositionMode = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.On Ground", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.On Ground has unexpected decoded type %T", raw)
		}
		p.OnGround = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Riding Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Riding Runtime ID has unexpected decoded type %T", raw)
		}
		p.RidingRuntimeID = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Teleport Data", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "MovePlayerTeleportData", TypeID: "MovePlayerTeleportData", Fields: []ShapeField{{Ordinal: 0, Name: "Teleportation Cause", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Source Actor Type", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*MovePlayerTeleportData)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Teleport Data has unexpected decoded type %T", raw)
		}
		p.TeleportData = value
	}
	{
		raw, err := r.Read("MovePlayerPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerInputTick)
		if !ok {
			return p, fmt.Errorf("field MovePlayerPacket.Tick has unexpected decoded type %T", raw)
		}
		p.Tick = value
	}
	return p, nil
}
