// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type Camera struct {
	CameraID       ActorUniqueID
	TargetPlayerID ActorUniqueID
}

func (p *Camera) Encode(w Encoder) error {
	if err := w.Write("CameraPacket.Camera ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.CameraID); err != nil {
		return err
	}
	if err := w.Write("CameraPacket.Target Player ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.TargetPlayerID); err != nil {
		return err
	}
	return nil
}

func DecodeCamera(r Decoder) (Camera, error) {
	var p Camera
	{
		raw, err := r.Read("CameraPacket.Camera ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field CameraPacket.Camera ID has unexpected decoded type %T", raw)
		}
		p.CameraID = value
	}
	{
		raw, err := r.Read("CameraPacket.Target Player ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field CameraPacket.Target Player ID has unexpected decoded type %T", raw)
		}
		p.TargetPlayerID = value
	}
	return p, nil
}
