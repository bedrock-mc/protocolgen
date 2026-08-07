// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type BlockEvent struct {
	BlockPosition BlockPos
	EventType     int32
	EventValue    int32
}

func (p *BlockEvent) Encode(w Encoder) error {
	if err := w.Write("BlockEventPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("BlockEventPacket.Event Type", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EventType); err != nil {
		return err
	}
	if err := w.Write("BlockEventPacket.Event Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EventValue); err != nil {
		return err
	}
	return nil
}

func DecodeBlockEvent(r Decoder) (BlockEvent, error) {
	var p BlockEvent
	{
		raw, err := r.Read("BlockEventPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field BlockEventPacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("BlockEventPacket.Event Type", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field BlockEventPacket.Event Type has unexpected decoded type %T", raw)
		}
		p.EventType = value
	}
	{
		raw, err := r.Read("BlockEventPacket.Event Value", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field BlockEventPacket.Event Value has unexpected decoded type %T", raw)
		}
		p.EventValue = value
	}
	return p, nil
}
