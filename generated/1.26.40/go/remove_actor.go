// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RemoveActor struct {
	TargetActorID ActorUniqueID
}

func (p *RemoveActor) Encode(w Encoder) error {
	if err := w.Write("RemoveActorPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetActorID); err != nil {
		return err
	}
	return nil
}

func DecodeRemoveActor(r Decoder) (RemoveActor, error) {
	var p RemoveActor
	{
		raw, err := r.Read("RemoveActorPacket.Target Actor ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field RemoveActorPacket.Target Actor ID has unexpected decoded type %T", raw)
		}
		p.TargetActorID = value
	}
	return p, nil
}
