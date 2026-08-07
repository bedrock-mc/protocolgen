// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type LevelChunk struct {
	ChunkPosition              ChunkPos
	DimensionId                DimensionType
	SubChunksCount             uint32
	ClientRequestSubChunkLimit *int32
	CacheEnabled               bool
	CacheMetadata              []LevelChunkSubChunkMetadata
	SerializedChunkData        string
}

func (p *LevelChunk) Encode(w Encoder) error {
	if err := w.Write("LevelChunkPacket.Chunk Position", Shape{Kind: "struct", Semantic: "ChunkPos", TypeID: "ChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.ChunkPosition); err != nil {
		return err
	}
	if err := w.Write("LevelChunkPacket.Dimension Id", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionId); err != nil {
		return err
	}
	if err := w.Write("LevelChunkPacket.Sub-chunks Count", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.SubChunksCount); err != nil {
		return err
	}
	if err := w.Write("LevelChunkPacket.Client Request SubChunk Limit", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, p.ClientRequestSubChunkLimit); err != nil {
		return err
	}
	if err := w.Write("LevelChunkPacket.Cache Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.CacheEnabled); err != nil {
		return err
	}
	if err := w.Write("LevelChunkPacket.Cache Metadata", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "LevelChunkPacketPayload::SubChunkMetadata", TypeID: "LevelChunkPacketPayload::SubChunkMetadata", Fields: []ShapeField{{Ordinal: 0, Name: "Blob Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}, p.CacheMetadata); err != nil {
		return err
	}
	if err := w.Write("LevelChunkPacket.Serialized Chunk Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.SerializedChunkData); err != nil {
		return err
	}
	return nil
}

func DecodeLevelChunk(r Decoder) (LevelChunk, error) {
	var p LevelChunk
	{
		raw, err := r.Read("LevelChunkPacket.Chunk Position", Shape{Kind: "struct", Semantic: "ChunkPos", TypeID: "ChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ChunkPos)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Chunk Position has unexpected decoded type %T", raw)
		}
		p.ChunkPosition = value
	}
	{
		raw, err := r.Read("LevelChunkPacket.Dimension Id", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Dimension Id has unexpected decoded type %T", raw)
		}
		p.DimensionId = value
	}
	{
		raw, err := r.Read("LevelChunkPacket.Sub-chunks Count", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Sub-chunks Count has unexpected decoded type %T", raw)
		}
		p.SubChunksCount = value
	}
	{
		raw, err := r.Read("LevelChunkPacket.Client Request SubChunk Limit", Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*int32)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Client Request SubChunk Limit has unexpected decoded type %T", raw)
		}
		p.ClientRequestSubChunkLimit = value
	}
	{
		raw, err := r.Read("LevelChunkPacket.Cache Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Cache Enabled has unexpected decoded type %T", raw)
		}
		p.CacheEnabled = value
	}
	{
		raw, err := r.Read("LevelChunkPacket.Cache Metadata", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "LevelChunkPacketPayload::SubChunkMetadata", TypeID: "LevelChunkPacketPayload::SubChunkMetadata", Fields: []ShapeField{{Ordinal: 0, Name: "Blob Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]LevelChunkSubChunkMetadata)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Cache Metadata has unexpected decoded type %T", raw)
		}
		p.CacheMetadata = value
	}
	{
		raw, err := r.Read("LevelChunkPacket.Serialized Chunk Data", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field LevelChunkPacket.Serialized Chunk Data has unexpected decoded type %T", raw)
		}
		p.SerializedChunkData = value
	}
	return p, nil
}
