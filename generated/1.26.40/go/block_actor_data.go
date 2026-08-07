// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type BlockActorData struct {
	BlockPosition BlockPos
	ActorDataTags []byte
}

func (p *BlockActorData) Encode(w Encoder) error {
	if err := w.Write("BlockActorDataPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("BlockActorDataPacket.Actor Data Tags", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.ActorDataTags); err != nil {
		return err
	}
	return nil
}

func DecodeBlockActorData(r Decoder) (BlockActorData, error) {
	var p BlockActorData
	{
		raw, err := r.Read("BlockActorDataPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field BlockActorDataPacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("BlockActorDataPacket.Actor Data Tags", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field BlockActorDataPacket.Actor Data Tags has unexpected decoded type %T", raw)
		}
		p.ActorDataTags = value
	}
	return p, nil
}
