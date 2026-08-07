// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RequestChunkRadius struct {
	ChunkRadius    int32
	MaxChunkRadius uint8
}

func (p *RequestChunkRadius) Encode(w Encoder) error {
	if err := w.Write("RequestChunkRadiusPacket.Chunk Radius", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.ChunkRadius); err != nil {
		return err
	}
	if err := w.Write("RequestChunkRadiusPacket.Max ChunkRadius", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.MaxChunkRadius); err != nil {
		return err
	}
	return nil
}

func DecodeRequestChunkRadius(r Decoder) (RequestChunkRadius, error) {
	var p RequestChunkRadius
	{
		raw, err := r.Read("RequestChunkRadiusPacket.Chunk Radius", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field RequestChunkRadiusPacket.Chunk Radius has unexpected decoded type %T", raw)
		}
		p.ChunkRadius = value
	}
	{
		raw, err := r.Read("RequestChunkRadiusPacket.Max ChunkRadius", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field RequestChunkRadiusPacket.Max ChunkRadius has unexpected decoded type %T", raw)
		}
		p.MaxChunkRadius = value
	}
	return p, nil
}
