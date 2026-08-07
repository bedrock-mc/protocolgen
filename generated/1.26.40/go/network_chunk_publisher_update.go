// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type NetworkChunkPublisherUpdate struct {
	NewPositionForView    BlockPos
	NewRadiusForView      uint32
	ServerBuiltChunksList []ChunkPos
}

func (p *NetworkChunkPublisherUpdate) Encode(w Encoder) error {
	if err := w.Write("NetworkChunkPublisherUpdatePacket.New position for view", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.NewPositionForView); err != nil {
		return err
	}
	if err := w.Write("NetworkChunkPublisherUpdatePacket.New radius for view", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.NewRadiusForView); err != nil {
		return err
	}
	if err := w.Write("NetworkChunkPublisherUpdatePacket.Server Built Chunks List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "ChunkPos", TypeID: "ChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, p.ServerBuiltChunksList); err != nil {
		return err
	}
	return nil
}

func DecodeNetworkChunkPublisherUpdate(r Decoder) (NetworkChunkPublisherUpdate, error) {
	var p NetworkChunkPublisherUpdate
	{
		raw, err := r.Read("NetworkChunkPublisherUpdatePacket.New position for view", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field NetworkChunkPublisherUpdatePacket.New position for view has unexpected decoded type %T", raw)
		}
		p.NewPositionForView = value
	}
	{
		raw, err := r.Read("NetworkChunkPublisherUpdatePacket.New radius for view", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field NetworkChunkPublisherUpdatePacket.New radius for view has unexpected decoded type %T", raw)
		}
		p.NewRadiusForView = value
	}
	{
		raw, err := r.Read("NetworkChunkPublisherUpdatePacket.Server Built Chunks List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "ChunkPos", TypeID: "ChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ChunkPos)
		if !ok {
			return p, fmt.Errorf("field NetworkChunkPublisherUpdatePacket.Server Built Chunks List has unexpected decoded type %T", raw)
		}
		p.ServerBuiltChunksList = value
	}
	return p, nil
}
