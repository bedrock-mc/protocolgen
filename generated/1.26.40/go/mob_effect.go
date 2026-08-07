// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MobEffect struct {
	TargetRuntimeID     ActorRuntimeID
	EventID             MobEffectEvent
	EffectID            int32
	EffectAmplifier     int32
	ShowParticles       bool
	EffectDurationTicks int32
	Tick                PlayerInputTick
	Ambient             bool
}

func (p *MobEffect) Encode(w Encoder) error {
	if err := w.Write("MobEffectPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetRuntimeID); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Event ID", Shape{Kind: "enum", Semantic: "MobEffectPacketPayload::Event", TypeID: "enums/MobEffectPacketPayload::Event", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Update", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Remove", Shape: Shape{Kind: "void"}}}}, p.EventID); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Effect ID", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EffectID); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Effect Amplifier", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EffectAmplifier); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Show Particles", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ShowParticles); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Effect Duration Ticks", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EffectDurationTicks); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.Tick); err != nil {
		return err
	}
	if err := w.Write("MobEffectPacket.Ambient", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Ambient); err != nil {
		return err
	}
	return nil
}

func DecodeMobEffect(r Decoder) (MobEffect, error) {
	var p MobEffect
	{
		raw, err := r.Read("MobEffectPacket.Target Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Target Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetRuntimeID = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Event ID", Shape{Kind: "enum", Semantic: "MobEffectPacketPayload::Event", TypeID: "enums/MobEffectPacketPayload::Event", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Add", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Update", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Remove", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(MobEffectEvent)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Event ID has unexpected decoded type %T", raw)
		}
		p.EventID = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Effect ID", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Effect ID has unexpected decoded type %T", raw)
		}
		p.EffectID = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Effect Amplifier", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Effect Amplifier has unexpected decoded type %T", raw)
		}
		p.EffectAmplifier = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Show Particles", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Show Particles has unexpected decoded type %T", raw)
		}
		p.ShowParticles = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Effect Duration Ticks", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Effect Duration Ticks has unexpected decoded type %T", raw)
		}
		p.EffectDurationTicks = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Tick", Shape{Kind: "struct", Semantic: "PlayerInputTick", TypeID: "PlayerInputTick", Fields: []ShapeField{{Ordinal: 0, Name: "Input tick", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayerInputTick)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Tick has unexpected decoded type %T", raw)
		}
		p.Tick = value
	}
	{
		raw, err := r.Read("MobEffectPacket.Ambient", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field MobEffectPacket.Ambient has unexpected decoded type %T", raw)
		}
		p.Ambient = value
	}
	return p, nil
}
