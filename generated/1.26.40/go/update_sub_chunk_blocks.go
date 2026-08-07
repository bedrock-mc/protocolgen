// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type UpdateSubChunkBlocks struct {
	SubChunkBlockPosition BlockPos
	BlocksChanged         UpdateSubChunkBlocksChangedInfo
}

func (p *UpdateSubChunkBlocks) Encode(w Encoder) error {
	if err := w.Write("UpdateSubChunkBlocksPacket.Sub Chunk Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.SubChunkBlockPosition); err != nil {
		return err
	}
	if err := w.Write("UpdateSubChunkBlocksPacket.Blocks Changed", Shape{Kind: "struct", Semantic: "UpdateSubChunkBlocksChangedInfo", TypeID: "UpdateSubChunkBlocksChangedInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Blocks Changed - Standards", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "UpdateSubChunkNetworkBlockInfo", TypeID: "UpdateSubChunkNetworkBlockInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Pos", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 1, Name: "Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 2, Name: "Update Flags", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Sync Message - Entity Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}, {Ordinal: 4, Name: "Sync Message - Message", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 1, Name: "Blocks Changed - Extras", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "UpdateSubChunkNetworkBlockInfo", TypeID: "UpdateSubChunkNetworkBlockInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Pos", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 1, Name: "Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 2, Name: "Update Flags", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Sync Message - Entity Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}, {Ordinal: 4, Name: "Sync Message - Message", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, p.BlocksChanged); err != nil {
		return err
	}
	return nil
}

func DecodeUpdateSubChunkBlocks(r Decoder) (UpdateSubChunkBlocks, error) {
	var p UpdateSubChunkBlocks
	{
		raw, err := r.Read("UpdateSubChunkBlocksPacket.Sub Chunk Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field UpdateSubChunkBlocksPacket.Sub Chunk Block Position has unexpected decoded type %T", raw)
		}
		p.SubChunkBlockPosition = value
	}
	{
		raw, err := r.Read("UpdateSubChunkBlocksPacket.Blocks Changed", Shape{Kind: "struct", Semantic: "UpdateSubChunkBlocksChangedInfo", TypeID: "UpdateSubChunkBlocksChangedInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Blocks Changed - Standards", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "UpdateSubChunkNetworkBlockInfo", TypeID: "UpdateSubChunkNetworkBlockInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Pos", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 1, Name: "Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 2, Name: "Update Flags", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Sync Message - Entity Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}, {Ordinal: 4, Name: "Sync Message - Message", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 1, Name: "Blocks Changed - Extras", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "UpdateSubChunkNetworkBlockInfo", TypeID: "UpdateSubChunkNetworkBlockInfo", Fields: []ShapeField{{Ordinal: 0, Name: "Pos", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 1, Name: "Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 2, Name: "Update Flags", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Sync Message - Entity Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}, {Ordinal: 4, Name: "Sync Message - Message", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(UpdateSubChunkBlocksChangedInfo)
		if !ok {
			return p, fmt.Errorf("field UpdateSubChunkBlocksPacket.Blocks Changed has unexpected decoded type %T", raw)
		}
		p.BlocksChanged = value
	}
	return p, nil
}
