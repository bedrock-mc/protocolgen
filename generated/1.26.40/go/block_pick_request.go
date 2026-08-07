// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type BlockPickRequest struct {
	Position BlockPos
	WithData bool
	MaxSlots uint8
}

func (p *BlockPickRequest) Encode(w Encoder) error {
	if err := w.Write("BlockPickRequestPacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("BlockPickRequestPacket.With Data?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.WithData); err != nil {
		return err
	}
	if err := w.Write("BlockPickRequestPacket.Max Slots", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.MaxSlots); err != nil {
		return err
	}
	return nil
}

func DecodeBlockPickRequest(r Decoder) (BlockPickRequest, error) {
	var p BlockPickRequest
	{
		raw, err := r.Read("BlockPickRequestPacket.Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field BlockPickRequestPacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("BlockPickRequestPacket.With Data?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field BlockPickRequestPacket.With Data? has unexpected decoded type %T", raw)
		}
		p.WithData = value
	}
	{
		raw, err := r.Read("BlockPickRequestPacket.Max Slots", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field BlockPickRequestPacket.Max Slots has unexpected decoded type %T", raw)
		}
		p.MaxSlots = value
	}
	return p, nil
}
