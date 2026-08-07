// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ChunkRadiusUpdated struct {
	ChunkRadius int32
}

func (p *ChunkRadiusUpdated) Encode(w Encoder) error {
	if err := w.Write("ChunkRadiusUpdatedPacket.Chunk Radius", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.ChunkRadius); err != nil {
		return err
	}
	return nil
}

func DecodeChunkRadiusUpdated(r Decoder) (ChunkRadiusUpdated, error) {
	var p ChunkRadiusUpdated
	{
		raw, err := r.Read("ChunkRadiusUpdatedPacket.Chunk Radius", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field ChunkRadiusUpdatedPacket.Chunk Radius has unexpected decoded type %T", raw)
		}
		p.ChunkRadius = value
	}
	return p, nil
}
