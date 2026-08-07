// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetSpawnPosition struct {
	SpawnPositionType SpawnPositionType
	BlockPosition     BlockPos
	DimensionType     DimensionType
	SpawnBlockPos     BlockPos
}

func (p *SetSpawnPosition) Encode(w Encoder) error {
	if err := w.Write("SetSpawnPositionPacket.Spawn Position Type", Shape{Kind: "enum", Semantic: "SpawnPositionType", TypeID: "enums/SpawnPositionType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "PlayerRespawn", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "WorldSpawn", Shape: Shape{Kind: "void"}}}}, p.SpawnPositionType); err != nil {
		return err
	}
	if err := w.Write("SetSpawnPositionPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("SetSpawnPositionPacket.Dimension type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.DimensionType); err != nil {
		return err
	}
	if err := w.Write("SetSpawnPositionPacket.Spawn Block Pos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.SpawnBlockPos); err != nil {
		return err
	}
	return nil
}

func DecodeSetSpawnPosition(r Decoder) (SetSpawnPosition, error) {
	var p SetSpawnPosition
	{
		raw, err := r.Read("SetSpawnPositionPacket.Spawn Position Type", Shape{Kind: "enum", Semantic: "SpawnPositionType", TypeID: "enums/SpawnPositionType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "PlayerRespawn", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "WorldSpawn", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SpawnPositionType)
		if !ok {
			return p, fmt.Errorf("field SetSpawnPositionPacket.Spawn Position Type has unexpected decoded type %T", raw)
		}
		p.SpawnPositionType = value
	}
	{
		raw, err := r.Read("SetSpawnPositionPacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field SetSpawnPositionPacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("SetSpawnPositionPacket.Dimension type", Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(DimensionType)
		if !ok {
			return p, fmt.Errorf("field SetSpawnPositionPacket.Dimension type has unexpected decoded type %T", raw)
		}
		p.DimensionType = value
	}
	{
		raw, err := r.Read("SetSpawnPositionPacket.Spawn Block Pos", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field SetSpawnPositionPacket.Spawn Block Pos has unexpected decoded type %T", raw)
		}
		p.SpawnBlockPos = value
	}
	return p, nil
}
