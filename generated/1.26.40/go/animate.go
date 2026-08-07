// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Animate struct {
	Action               AnimateAction
	TargetActorRuntimeID ActorRuntimeID
	Data                 float32
	SwingSource          *string
}

func (p *Animate) Encode(w Encoder) error {
	if err := w.Write("AnimatePacket.Action", Shape{Kind: "enum", Semantic: "AnimatePacketPayload::Action", TypeID: "enums/AnimatePacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NoAction", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Swing", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "WakeUp", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "CriticalHit", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "MagicCriticalHit", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("AnimatePacket.Target Actor Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.TargetActorRuntimeID); err != nil {
		return err
	}
	if err := w.Write("AnimatePacket.Data", Shape{Kind: "primitive", PrimitiveCode: "f32le"}, p.Data); err != nil {
		return err
	}
	if err := w.Write("AnimatePacket.Swing Source", Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.SwingSource); err != nil {
		return err
	}
	return nil
}

func DecodeAnimate(r Decoder) (Animate, error) {
	var p Animate
	{
		raw, err := r.Read("AnimatePacket.Action", Shape{Kind: "enum", Semantic: "AnimatePacketPayload::Action", TypeID: "enums/AnimatePacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NoAction", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Swing", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "WakeUp", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "CriticalHit", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "MagicCriticalHit", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(AnimateAction)
		if !ok {
			return p, fmt.Errorf("field AnimatePacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("AnimatePacket.Target Actor Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field AnimatePacket.Target Actor Runtime ID has unexpected decoded type %T", raw)
		}
		p.TargetActorRuntimeID = value
	}
	{
		raw, err := r.Read("AnimatePacket.Data", Shape{Kind: "primitive", PrimitiveCode: "f32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(float32)
		if !ok {
			return p, fmt.Errorf("field AnimatePacket.Data has unexpected decoded type %T", raw)
		}
		p.Data = value
	}
	{
		raw, err := r.Read("AnimatePacket.Swing Source", Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*string)
		if !ok {
			return p, fmt.Errorf("field AnimatePacket.Swing Source has unexpected decoded type %T", raw)
		}
		p.SwingSource = value
	}
	return p, nil
}
