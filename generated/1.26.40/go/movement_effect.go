// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MovementEffect struct {
	TargetRuntimeID ActorRuntimeID
	EffectID        MovementEffectType
	EffectDuration  int32
	Tick            PlayerInputTick
}

func (p *MovementEffect) Encode(w Encoder) error {
	if err := w.Write("MovementEffectPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("MovementEffectPacket.Effect ID", Shape{Kind: "enum", Semantic: "MovementEffectType", TypeID: "enums/MovementEffectType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "GLIDE_BOOST", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DOLPHIN_BOOST", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "GEYSER_BOOST", Shape: Shape{Kind: "void"}}}}, p.EffectID); err != nil {
		return err
	}
	if err := w.Write("MovementEffectPacket.Effect Duration", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EffectDuration); err != nil {
		return err
	}
	if err := w.Write("MovementEffectPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.Tick); err != nil {
		return err
	}
	return nil
}

func DecodeMovementEffect(r Decoder) (MovementEffect, error) {
	var p MovementEffect
	{
		raw, err := r.Read("MovementEffectPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MovementEffectPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("MovementEffectPacket.Effect ID", Shape{Kind: "enum", Semantic: "MovementEffectType", TypeID: "enums/MovementEffectType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "GLIDE_BOOST", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DOLPHIN_BOOST", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "GEYSER_BOOST", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(MovementEffectType)
		if !ok {
			return p, fmt.Errorf("field MovementEffectPacket.Effect ID has unexpected decoded type %T", raw)
		}
		p.EffectID = value
	}
	{
		raw, err := r.Read("MovementEffectPacket.Effect Duration", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field MovementEffectPacket.Effect Duration has unexpected decoded type %T", raw)
		}
		p.EffectDuration = value
	}
	{
		raw, err := r.Read("MovementEffectPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerInputTick)
		if !ok {
			return p, fmt.Errorf("field MovementEffectPacket.Tick has unexpected decoded type %T", raw)
		}
		p.Tick = value
	}
	return p, nil
}
