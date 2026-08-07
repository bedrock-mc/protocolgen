// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PositionTrackingDBClientRequest struct {
	Action PositionTrackingDBClientRequestAction
	Id     PositionTrackingId
}

func (p *PositionTrackingDBClientRequest) Encode(w Encoder) error {
	if err := w.Write("PositionTrackingDBClientRequestPacket.Action", Shape{Kind: "enum", Semantic: "PositionTrackingDBClientRequestPacketPayload::Action", TypeID: "enums/PositionTrackingDBClientRequestPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Query", Shape: Shape{Kind: "void"}}}}, p.Action); err != nil {
		return err
	}
	if err := w.Write("PositionTrackingDBClientRequestPacket.Id", Shape{Kind: "struct", Semantic: "PositionTrackingId", TypeID: "PositionTrackingId", Fields: []ShapeField{{Ordinal: 0, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Id); err != nil {
		return err
	}
	return nil
}

func DecodePositionTrackingDBClientRequest(r Decoder) (PositionTrackingDBClientRequest, error) {
	var p PositionTrackingDBClientRequest
	{
		raw, err := r.Read("PositionTrackingDBClientRequestPacket.Action", Shape{Kind: "enum", Semantic: "PositionTrackingDBClientRequestPacketPayload::Action", TypeID: "enums/PositionTrackingDBClientRequestPacketPayload::Action", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "Query", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PositionTrackingDBClientRequestAction)
		if !ok {
			return p, fmt.Errorf("field PositionTrackingDBClientRequestPacket.Action has unexpected decoded type %T", raw)
		}
		p.Action = value
	}
	{
		raw, err := r.Read("PositionTrackingDBClientRequestPacket.Id", Shape{Kind: "struct", Semantic: "PositionTrackingId", TypeID: "PositionTrackingId", Fields: []ShapeField{{Ordinal: 0, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PositionTrackingId)
		if !ok {
			return p, fmt.Errorf("field PositionTrackingDBClientRequestPacket.Id has unexpected decoded type %T", raw)
		}
		p.Id = value
	}
	return p, nil
}
