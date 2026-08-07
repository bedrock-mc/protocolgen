// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SubChunk struct {
	CacheEnabled  bool
	DimensionType DimensionType
	CenterPos     SubChunkPos
	SubChunkData  []SubChunkSubChunkPacketData
}

func (p *SubChunk) Encode(w Encoder) error {
	if err := w.Write("SubChunkPacket.Cache Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.CacheEnabled); err != nil {
		return err
	}
	if err := w.Write("SubChunkPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionType); err != nil {
		return err
	}
	if err := w.Write("SubChunkPacket.Center Pos", Shape{Kind: "struct", Semantic: "SubChunkPos", TypeID: "SubChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Position X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Subchunk Position Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 2, Name: "Subchunk Position Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}, p.CenterPos); err != nil {
		return err
	}
	if err := w.Write("SubChunkPacket.SubChunk Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::SubChunkPacketData", TypeID: "SubChunkPacketPayload::SubChunkPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "SubChunk Pos Offset", Shape: Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::SubChunkPosOffset", TypeID: "SubChunkPacketPayload::SubChunkPosOffset", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Offset X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 1, Name: "Subchunk Offset Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 2, Name: "Subchunk Offset Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Ordinal: 1, Name: "SubChunk Request Result", Shape: Shape{Kind: "enum", Semantic: "SubChunkPacketPayload::SubChunkRequestResult", TypeID: "enums/SubChunkPacketPayload::SubChunkRequestResult", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 1, Name: "Success", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "LevelChunkDoesntExist", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "WrongDimension", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "PlayerDoesntExist", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "IndexOutOfBounds", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "SuccessAllAir", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Serialized Sub Chunk", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 3, Name: "Height Map Data", Shape: Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::HeightmapData", TypeID: "SubChunkPacketPayload::HeightmapData", Fields: []ShapeField{{Ordinal: 0, Name: "Height Map Type", Shape: Shape{Kind: "enum", Semantic: "SubChunkPacketPayload::HeightMapDataType", TypeID: "enums/SubChunkPacketPayload::HeightMapDataType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NoData", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "HasData", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "AllTooHigh", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "AllTooLow", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Subchunk Height Map", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Ordinal: 2, Name: "Render Height Map Type", Shape: Shape{Kind: "enum", Semantic: "SubChunkPacketPayload::HeightMapDataType", TypeID: "enums/SubChunkPacketPayload::HeightMapDataType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NoData", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "HasData", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "AllTooHigh", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "AllTooLow", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "AllCopied", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Subchunk Render Height Map", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}}}}, {Ordinal: 4, Name: "Blob Id", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}}, p.SubChunkData); err != nil {
		return err
	}
	return nil
}

func DecodeSubChunk(r Decoder) (SubChunk, error) {
	var p SubChunk
	{
		raw, err := r.Read("SubChunkPacket.Cache Enabled", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field SubChunkPacket.Cache Enabled has unexpected decoded type %T", raw)
		}
		p.CacheEnabled = value
	}
	{
		raw, err := r.Read("SubChunkPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field SubChunkPacket.Dimension Type has unexpected decoded type %T", raw)
		}
		p.DimensionType = value
	}
	{
		raw, err := r.Read("SubChunkPacket.Center Pos", Shape{Kind: "struct", Semantic: "SubChunkPos", TypeID: "SubChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Position X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Subchunk Position Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 2, Name: "Subchunk Position Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SubChunkPos)
		if !ok {
			return p, fmt.Errorf("field SubChunkPacket.Center Pos has unexpected decoded type %T", raw)
		}
		p.CenterPos = value
	}
	{
		raw, err := r.Read("SubChunkPacket.SubChunk Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::SubChunkPacketData", TypeID: "SubChunkPacketPayload::SubChunkPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "SubChunk Pos Offset", Shape: Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::SubChunkPosOffset", TypeID: "SubChunkPacketPayload::SubChunkPosOffset", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Offset X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 1, Name: "Subchunk Offset Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 2, Name: "Subchunk Offset Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Ordinal: 1, Name: "SubChunk Request Result", Shape: Shape{Kind: "enum", Semantic: "SubChunkPacketPayload::SubChunkRequestResult", TypeID: "enums/SubChunkPacketPayload::SubChunkRequestResult", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 1, Name: "Success", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "LevelChunkDoesntExist", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "WrongDimension", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "PlayerDoesntExist", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "IndexOutOfBounds", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "SuccessAllAir", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Serialized Sub Chunk", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 3, Name: "Height Map Data", Shape: Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::HeightmapData", TypeID: "SubChunkPacketPayload::HeightmapData", Fields: []ShapeField{{Ordinal: 0, Name: "Height Map Type", Shape: Shape{Kind: "enum", Semantic: "SubChunkPacketPayload::HeightMapDataType", TypeID: "enums/SubChunkPacketPayload::HeightMapDataType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NoData", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "HasData", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "AllTooHigh", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "AllTooLow", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Subchunk Height Map", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, {Ordinal: 2, Name: "Render Height Map Type", Shape: Shape{Kind: "enum", Semantic: "SubChunkPacketPayload::HeightMapDataType", TypeID: "enums/SubChunkPacketPayload::HeightMapDataType", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "NoData", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "HasData", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "AllTooHigh", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "AllTooLow", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "AllCopied", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Subchunk Render Height Map", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "fixed_array", Length: 16, Element: &Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}}}}, {Ordinal: 4, Name: "Blob Id", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u64le"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]SubChunkSubChunkPacketData)
		if !ok {
			return p, fmt.Errorf("field SubChunkPacket.SubChunk Data has unexpected decoded type %T", raw)
		}
		p.SubChunkData = value
	}
	return p, nil
}
