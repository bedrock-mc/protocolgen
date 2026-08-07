// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type LevelSoundEvent struct {
	SoundEvent      string
	Position        Vec3
	Data            int32
	ActorIdentifier string
	IsBaby          bool
	IsGlobal        bool
	ActorUniqueId   int64
	FireAtPosition  *Vec3
}

func (p *LevelSoundEvent) Encode(w Encoder) error {
	if err := w.Write("LevelSoundEventPacket.Sound Event", Shape{Kind: "string", Semantic: "SoundEventIdentifier", TypeID: "SoundEventIdentifier.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SoundEvent); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.Data); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Actor Identifier", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ActorIdentifier); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Is Baby", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsBaby); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Is Global", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsGlobal); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Actor Unique Id", Shape{Kind: "primitive", PrimitiveCode: "i64le"}, p.ActorUniqueId); err != nil {
		return err
	}
	if err := w.Write("LevelSoundEventPacket.Fire At Position", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}, p.FireAtPosition); err != nil {
		return err
	}
	return nil
}

func DecodeLevelSoundEvent(r Decoder) (LevelSoundEvent, error) {
	var p LevelSoundEvent
	{
		raw, err := r.Read("LevelSoundEventPacket.Sound Event", Shape{Kind: "string", Semantic: "SoundEventIdentifier", TypeID: "SoundEventIdentifier.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Sound Event has unexpected decoded type %T", raw)
		}
		p.SoundEvent = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Data", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Actor Identifier", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Actor Identifier has unexpected decoded type %T", raw)
		}
		p.ActorIdentifier = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Is Baby", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Is Baby has unexpected decoded type %T", raw)
		}
		p.IsBaby = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Is Global", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Is Global has unexpected decoded type %T", raw)
		}
		p.IsGlobal = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Actor Unique Id", Shape{Kind: "primitive", PrimitiveCode: "i64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int64)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Actor Unique Id has unexpected decoded type %T", raw)
		}
		p.ActorUniqueId = value
	}
	{
		raw, err := r.Read("LevelSoundEventPacket.Fire At Position", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*Vec3)
		if !ok {
			return p, fmt.Errorf("field LevelSoundEventPacket.Fire At Position has unexpected decoded type %T", raw)
		}
		p.FireAtPosition = value
	}
	return p, nil
}
