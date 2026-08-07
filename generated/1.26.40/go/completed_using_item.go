// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type CompletedUsingItem struct {
	ItemId        int16
	ItemUseMethod int32
}

func (p *CompletedUsingItem) Encode(w Encoder) error {
	if err := w.Write("CompletedUsingItemPacket.Item Id", Shape{Kind: "primitive", PrimitiveCode: "i16le"}, p.ItemId); err != nil {
		return err
	}
	if err := w.Write("CompletedUsingItemPacket.Item Use Method", Shape{Kind: "primitive", PrimitiveCode: "i32le"}, p.ItemUseMethod); err != nil {
		return err
	}
	return nil
}

func DecodeCompletedUsingItem(r Decoder) (CompletedUsingItem, error) {
	var p CompletedUsingItem
	{
		raw, err := r.Read("CompletedUsingItemPacket.Item Id", Shape{Kind: "primitive", PrimitiveCode: "i16le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int16)
		if !ok {
			return p, fmt.Errorf("field CompletedUsingItemPacket.Item Id has unexpected decoded type %T", raw)
		}
		p.ItemId = value
	}
	{
		raw, err := r.Read("CompletedUsingItemPacket.Item Use Method", Shape{Kind: "primitive", PrimitiveCode: "i32le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field CompletedUsingItemPacket.Item Use Method has unexpected decoded type %T", raw)
		}
		p.ItemUseMethod = value
	}
	return p, nil
}
