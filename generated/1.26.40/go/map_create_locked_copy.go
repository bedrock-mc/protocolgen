// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type MapCreateLockedCopy struct {
	OriginalMapId ActorUniqueID
	NewMapId      ActorUniqueID
}

func (p *MapCreateLockedCopy) Encode(w Encoder) error {
	if err := w.Write("MapCreateLockedCopyPacket.Original Map Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.OriginalMapId); err != nil {
		return err
	}
	if err := w.Write("MapCreateLockedCopyPacket.New Map Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.NewMapId); err != nil {
		return err
	}
	return nil
}

func DecodeMapCreateLockedCopy(r Decoder) (MapCreateLockedCopy, error) {
	var p MapCreateLockedCopy
	{
		raw, err := r.Read("MapCreateLockedCopyPacket.Original Map Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field MapCreateLockedCopyPacket.Original Map Id has unexpected decoded type %T", raw)
		}
		p.OriginalMapId = value
	}
	{
		raw, err := r.Read("MapCreateLockedCopyPacket.New Map Id", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field MapCreateLockedCopyPacket.New Map Id has unexpected decoded type %T", raw)
		}
		p.NewMapId = value
	}
	return p, nil
}
