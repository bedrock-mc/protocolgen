// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayerStartItemCooldown struct {
	ItemCategory  string
	DurationTicks int32
}

func (p *PlayerStartItemCooldown) Encode(w Encoder) error {
	if err := w.Write("PlayerStartItemCooldownPacket.Item Category", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ItemCategory); err != nil {
		return err
	}
	if err := w.Write("PlayerStartItemCooldownPacket.Duration Ticks", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.DurationTicks); err != nil {
		return err
	}
	return nil
}

func DecodePlayerStartItemCooldown(r Decoder) (PlayerStartItemCooldown, error) {
	var p PlayerStartItemCooldown
	{
		raw, err := r.Read("PlayerStartItemCooldownPacket.Item Category", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field PlayerStartItemCooldownPacket.Item Category has unexpected decoded type %T", raw)
		}
		p.ItemCategory = value
	}
	{
		raw, err := r.Read("PlayerStartItemCooldownPacket.Duration Ticks", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field PlayerStartItemCooldownPacket.Duration Ticks has unexpected decoded type %T", raw)
		}
		p.DurationTicks = value
	}
	return p, nil
}
