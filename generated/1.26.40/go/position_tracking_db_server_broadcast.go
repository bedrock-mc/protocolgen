// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PositionTrackingDBServerBroadcast struct {
	Action               PositionTrackingDBServerBroadcastAction
	Id                   PositionTrackingId
	PositionTrackingData []byte
}

func (p *PositionTrackingDBServerBroadcast) Encode(w Encoder) error {
	if err := w.Write("PositionTrackingDBServerBroadcastPacket.Action", Shape{Kind: "enum", Semantic: "PositionTrackingDBServerBroadcastPacketPayload::Action", TypeID: "enums/PositionTrackingDBServerBroadcastPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Update", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Destroy", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "NotFound", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("PositionTrackingDBServerBroadcastPacket.Id", Shape{Kind: "struct", Semantic: "PositionTrackingId", TypeID: "PositionTrackingId", Fields: []ShapeField{{Ordinal: 0, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Id); err != nil {
		return err
	}
	if err := w.Write("PositionTrackingDBServerBroadcastPacket.Position tracking data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.PositionTrackingData); err != nil {
		return err
	}
	return nil
}

func DecodePositionTrackingDBServerBroadcast(r Decoder) (PositionTrackingDBServerBroadcast, error) {
	var p PositionTrackingDBServerBroadcast
	{
		raw, err := r.Read("PositionTrackingDBServerBroadcastPacket.Action", Shape{Kind: "enum", Semantic: "PositionTrackingDBServerBroadcastPacketPayload::Action", TypeID: "enums/PositionTrackingDBServerBroadcastPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Update", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Destroy", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "NotFound", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PositionTrackingDBServerBroadcastAction)
		if !ok {
			return p, fmt.Errorf("field PositionTrackingDBServerBroadcastPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("PositionTrackingDBServerBroadcastPacket.Id", Shape{Kind: "struct", Semantic: "PositionTrackingId", TypeID: "PositionTrackingId", Fields: []ShapeField{{Ordinal: 0, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PositionTrackingId)
		if !ok {
			return p, fmt.Errorf("field PositionTrackingDBServerBroadcastPacket.Id has unexpected decoded type %T", raw)
		}
		p.Id = value
	}
	{
		raw, err := r.Read("PositionTrackingDBServerBroadcastPacket.Position tracking data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field PositionTrackingDBServerBroadcastPacket.Position tracking data has unexpected decoded type %T", raw)
		}
		p.PositionTrackingData = value
	}
	return p, nil
}
