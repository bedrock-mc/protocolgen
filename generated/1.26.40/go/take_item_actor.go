// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type TakeItemActor struct {
	ItemRuntimeID  ActorRuntimeID
	ActorRuntimeID ActorRuntimeID
}

func (p *TakeItemActor) Encode(w Encoder) error {
	if err := w.Write("TakeItemActorPacket.Item Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.ItemRuntimeID); err != nil {
		return err
	}
	if err := w.Write("TakeItemActorPacket.Actor Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.ActorRuntimeID); err != nil {
		return err
	}
	return nil
}

func DecodeTakeItemActor(r Decoder) (TakeItemActor, error) {
	var p TakeItemActor
	{
		raw, err := r.Read("TakeItemActorPacket.Item Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field TakeItemActorPacket.Item Runtime ID has unexpected decoded type %T", raw)
		}
		p.ItemRuntimeID = value
	}
	{
		raw, err := r.Read("TakeItemActorPacket.Actor Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field TakeItemActorPacket.Actor Runtime ID has unexpected decoded type %T", raw)
		}
		p.ActorRuntimeID = value
	}
	return p, nil
}
