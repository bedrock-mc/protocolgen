// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerHotbar struct {
	SelectedSlot     uint32
	ContainerID      uint8
	ShouldSelectSlot bool
}

func (p *PlayerHotbar) Encode(w Encoder) error {
	if err := w.Write("PlayerHotbarPacket.Selected Slot", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.SelectedSlot); err != nil {
		return err
	}
	if err := w.Write("PlayerHotbarPacket.Container ID", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerID); err != nil {
		return err
	}
	if err := w.Write("PlayerHotbarPacket.Should select slot?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ShouldSelectSlot); err != nil {
		return err
	}
	return nil
}

func DecodePlayerHotbar(r Decoder) (PlayerHotbar, error) {
	var p PlayerHotbar
	{
		raw, err := r.Read("PlayerHotbarPacket.Selected Slot", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field PlayerHotbarPacket.Selected Slot has unexpected decoded type %T", raw)
		}
		p.SelectedSlot = value
	}
	{
		raw, err := r.Read("PlayerHotbarPacket.Container ID", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field PlayerHotbarPacket.Container ID has unexpected decoded type %T", raw)
		}
		p.ContainerID = value
	}
	{
		raw, err := r.Read("PlayerHotbarPacket.Should select slot?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field PlayerHotbarPacket.Should select slot? has unexpected decoded type %T", raw)
		}
		p.ShouldSelectSlot = value
	}
	return p, nil
}
