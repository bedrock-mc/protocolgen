// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ActorPickRequest struct {
	ActorID  int64
	MaxSlots uint8
	WithData bool
}

func (p *ActorPickRequest) Encode(w Encoder) error {
	if err := w.Write("ActorPickRequestPacket.Actor ID", Shape{Kind: "primitive", PrimitiveCode: "i64le"}, p.ActorID); err != nil {
		return err
	}
	if err := w.Write("ActorPickRequestPacket.Max Slots", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.MaxSlots); err != nil {
		return err
	}
	if err := w.Write("ActorPickRequestPacket.With Data", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.WithData); err != nil {
		return err
	}
	return nil
}

func DecodeActorPickRequest(r Decoder) (ActorPickRequest, error) {
	var p ActorPickRequest
	{
		raw, err := r.Read("ActorPickRequestPacket.Actor ID", Shape{Kind: "primitive", PrimitiveCode: "i64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int64)
		if !ok {
			return p, fmt.Errorf("field ActorPickRequestPacket.Actor ID has unexpected decoded type %T", raw)
		}
		p.ActorID = value
	}
	{
		raw, err := r.Read("ActorPickRequestPacket.Max Slots", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ActorPickRequestPacket.Max Slots has unexpected decoded type %T", raw)
		}
		p.MaxSlots = value
	}
	{
		raw, err := r.Read("ActorPickRequestPacket.With Data", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ActorPickRequestPacket.With Data has unexpected decoded type %T", raw)
		}
		p.WithData = value
	}
	return p, nil
}
