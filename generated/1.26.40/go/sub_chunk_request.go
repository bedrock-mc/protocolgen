// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SubChunkRequest struct {
	DimensionType              DimensionType
	SubChunkPositionOffsetList []SubChunkSubChunkPosOffset
	CenterPos                  SubChunkPos
}

func (p *SubChunkRequest) Encode(w Encoder) error {
	if err := w.Write("SubChunkRequestPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionType); err != nil {
		return err
	}
	if err := w.Write("SubChunkRequestPacket.SubChunk Position Offset List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::SubChunkPosOffset", TypeID: "SubChunkPacketPayload::SubChunkPosOffset", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Offset X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 1, Name: "Subchunk Offset Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 2, Name: "Subchunk Offset Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}}, p.SubChunkPositionOffsetList); err != nil {
		return err
	}
	if err := w.Write("SubChunkRequestPacket.Center Pos", Shape{Kind: "struct", Semantic: "SubChunkPos", TypeID: "SubChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Position X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Subchunk Position Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 2, Name: "Subchunk Position Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}}, p.CenterPos); err != nil {
		return err
	}
	return nil
}

func DecodeSubChunkRequest(r Decoder) (SubChunkRequest, error) {
	var p SubChunkRequest
	{
		raw, err := r.Read("SubChunkRequestPacket.Dimension Type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field SubChunkRequestPacket.Dimension Type has unexpected decoded type %T", raw)
		}
		p.DimensionType = value
	}
	{
		raw, err := r.Read("SubChunkRequestPacket.SubChunk Position Offset List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "SubChunkPacketPayload::SubChunkPosOffset", TypeID: "SubChunkPacketPayload::SubChunkPosOffset", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Offset X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 1, Name: "Subchunk Offset Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}, {Ordinal: 2, Name: "Subchunk Offset Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i8"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]SubChunkSubChunkPosOffset)
		if !ok {
			return p, fmt.Errorf("field SubChunkRequestPacket.SubChunk Position Offset List has unexpected decoded type %T", raw)
		}
		p.SubChunkPositionOffsetList = value
	}
	{
		raw, err := r.Read("SubChunkRequestPacket.Center Pos", Shape{Kind: "struct", Semantic: "SubChunkPos", TypeID: "SubChunkPos", Fields: []ShapeField{{Ordinal: 0, Name: "Subchunk Position X", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 1, Name: "Subchunk Position Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 2, Name: "Subchunk Position Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SubChunkPos)
		if !ok {
			return p, fmt.Errorf("field SubChunkRequestPacket.Center Pos has unexpected decoded type %T", raw)
		}
		p.CenterPos = value
	}
	return p, nil
}
