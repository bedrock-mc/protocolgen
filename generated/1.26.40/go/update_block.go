// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateBlock struct {
	BlockPosition  BlockPos
	BlockRuntimeID uint32
	Flags          uint32
	Layer          uint32
}

func (p *UpdateBlock) Encode(w Encoder) error {
	if err := w.Write("UpdateBlockPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockPacket.Block Runtime ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.BlockRuntimeID); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockPacket.Flags", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.Flags); err != nil {
		return err
	}
	if err := w.Write("UpdateBlockPacket.Layer", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.Layer); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateBlock(r Decoder) (UpdateBlock, error) {
	var p UpdateBlock
	{
		raw, err := r.Read("UpdateBlockPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockPacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("UpdateBlockPacket.Block Runtime ID", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockPacket.Block Runtime ID has unexpected decoded type %T", raw)
		}
		p.BlockRuntimeID = value
	}
	{
		raw, err := r.Read("UpdateBlockPacket.Flags", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockPacket.Flags has unexpected decoded type %T", raw)
		}
		p.Flags = value
	}
	{
		raw, err := r.Read("UpdateBlockPacket.Layer", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field UpdateBlockPacket.Layer has unexpected decoded type %T", raw)
		}
		p.Layer = value
	}
	return p, nil
}
