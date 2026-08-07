// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerToggleCrafterSlotRequest struct {
	PosX       int32
	PosY       int32
	PosZ       int32
	SlotIndex  uint8
	IsDisabled bool
}

func (p *PlayerToggleCrafterSlotRequest) Encode(w Encoder) error {
	if err := w.Write("PlayerToggleCrafterSlotRequestPacket.Pos X", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.PosX); err != nil {
		return err
	}
	if err := w.Write("PlayerToggleCrafterSlotRequestPacket.Pos Y", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.PosY); err != nil {
		return err
	}
	if err := w.Write("PlayerToggleCrafterSlotRequestPacket.Pos Z", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.PosZ); err != nil {
		return err
	}
	if err := w.Write("PlayerToggleCrafterSlotRequestPacket.Slot Index", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.SlotIndex); err != nil {
		return err
	}
	if err := w.Write("PlayerToggleCrafterSlotRequestPacket.Is Disabled", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsDisabled); err != nil {
		return err
	}
	return nil
}

func DecodePlayerToggleCrafterSlotRequest(r Decoder) (PlayerToggleCrafterSlotRequest, error) {
	var p PlayerToggleCrafterSlotRequest
	{
		raw, err := r.Read("PlayerToggleCrafterSlotRequestPacket.Pos X", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PlayerToggleCrafterSlotRequestPacket.Pos X has unexpected decoded type %T", raw)
		}
		p.PosX = value
	}
	{
		raw, err := r.Read("PlayerToggleCrafterSlotRequestPacket.Pos Y", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PlayerToggleCrafterSlotRequestPacket.Pos Y has unexpected decoded type %T", raw)
		}
		p.PosY = value
	}
	{
		raw, err := r.Read("PlayerToggleCrafterSlotRequestPacket.Pos Z", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PlayerToggleCrafterSlotRequestPacket.Pos Z has unexpected decoded type %T", raw)
		}
		p.PosZ = value
	}
	{
		raw, err := r.Read("PlayerToggleCrafterSlotRequestPacket.Slot Index", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field PlayerToggleCrafterSlotRequestPacket.Slot Index has unexpected decoded type %T", raw)
		}
		p.SlotIndex = value
	}
	{
		raw, err := r.Read("PlayerToggleCrafterSlotRequestPacket.Is Disabled", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field PlayerToggleCrafterSlotRequestPacket.Is Disabled has unexpected decoded type %T", raw)
		}
		p.IsDisabled = value
	}
	return p, nil
}
