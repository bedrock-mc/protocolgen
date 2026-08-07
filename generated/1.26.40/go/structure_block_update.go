// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type StructureBlockUpdate struct {
	BlockPosition BlockPos
	StructureData StructureEditorData
	Trigger       bool
	IsWaterlogged bool
}

func (p *StructureBlockUpdate) Encode(w Encoder) error {
	if err := w.Write("StructureBlockUpdatePacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.BlockPosition); err != nil {
		return err
	}
	if err := w.Write("StructureBlockUpdatePacket.Structure Data", Shape{Kind: "struct", Semantic: "StructureEditorData", TypeID: "StructureEditorData", Fields: []ShapeField{{Ordinal: 0, Name: "Structure Name", Shape: Shape{Kind: "struct", Semantic: "Bedrock::Safety::RedactableString", TypeID: "Bedrock::Safety::RedactableString", Fields: []ShapeField{{Ordinal: 0, Name: "Unredacted", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Redacted", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Ordinal: 1, Name: "Data Field", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Should include players?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Should show bounding box?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Structure Block Type", Shape: Shape{Kind: "enum", Semantic: "StructureBlockType", TypeID: "enums/StructureBlockType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Data", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Save", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Load", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Corner", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Export", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 5, Name: "Structure Settings", Shape: Shape{Kind: "struct", Semantic: "StructureSettings", TypeID: "StructureSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Structure Palette Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Should ignore entities?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Should ignore blocks?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Should Allow Non Ticking Player and Ticking Area Chunks", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Structure Size", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 5, Name: "Structure Offset", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 6, Name: "Last Edit Player", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 7, Name: "Rotation", Shape: Shape{Kind: "enum", Semantic: "Rotation", TypeID: "enums/Rotation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotate90", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Rotate180", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Rotate270", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 8, Name: "Mirror", Shape: Shape{Kind: "enum", Semantic: "Mirror", TypeID: "enums/Mirror", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "X", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Z", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "XZ", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 9, Name: "Animation Mode", Shape: Shape{Kind: "enum", Semantic: "AnimationMode", TypeID: "enums/AnimationMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Layers", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Blocks", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 10, Name: "Animation Seconds", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 11, Name: "Integrity Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 12, Name: "Integrity Seed", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 13, Name: "Rotation Pivot", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, {Ordinal: 6, Name: "Redstone Save Mode", Shape: Shape{Kind: "enum", Semantic: "StructureRedstoneSaveMode", TypeID: "enums/StructureRedstoneSaveMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SavesToMemory", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "SavesToDisk", Shape: Shape{Kind: "void"}}}}}}}, p.StructureData); err != nil {
		return err
	}
	if err := w.Write("StructureBlockUpdatePacket.Trigger?", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.Trigger); err != nil {
		return err
	}
	if err := w.Write("StructureBlockUpdatePacket.IsWaterlogged", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsWaterlogged); err != nil {
		return err
	}
	return nil
}

func DecodeStructureBlockUpdate(r Decoder) (StructureBlockUpdate, error) {
	var p StructureBlockUpdate
	{
		raw, err := r.Read("StructureBlockUpdatePacket.Block Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field StructureBlockUpdatePacket.Block Position has unexpected decoded type %T", raw)
		}
		p.BlockPosition = value
	}
	{
		raw, err := r.Read("StructureBlockUpdatePacket.Structure Data", Shape{Kind: "struct", Semantic: "StructureEditorData", TypeID: "StructureEditorData", Fields: []ShapeField{{Ordinal: 0, Name: "Structure Name", Shape: Shape{Kind: "struct", Semantic: "Bedrock::Safety::RedactableString", TypeID: "Bedrock::Safety::RedactableString", Fields: []ShapeField{{Ordinal: 0, Name: "Unredacted", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Redacted", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Ordinal: 1, Name: "Data Field", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Should include players?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Should show bounding box?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Structure Block Type", Shape: Shape{Kind: "enum", Semantic: "StructureBlockType", TypeID: "enums/StructureBlockType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Data", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Save", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Load", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Corner", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Invalid", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Export", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 5, Name: "Structure Settings", Shape: Shape{Kind: "struct", Semantic: "StructureSettings", TypeID: "StructureSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Structure Palette Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Should ignore entities?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Should ignore blocks?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Should Allow Non Ticking Player and Ticking Area Chunks", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Structure Size", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 5, Name: "Structure Offset", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 6, Name: "Last Edit Player", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 7, Name: "Rotation", Shape: Shape{Kind: "enum", Semantic: "Rotation", TypeID: "enums/Rotation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotate90", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Rotate180", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Rotate270", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 8, Name: "Mirror", Shape: Shape{Kind: "enum", Semantic: "Mirror", TypeID: "enums/Mirror", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "X", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Z", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "XZ", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 9, Name: "Animation Mode", Shape: Shape{Kind: "enum", Semantic: "AnimationMode", TypeID: "enums/AnimationMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Layers", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Blocks", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 10, Name: "Animation Seconds", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 11, Name: "Integrity Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 12, Name: "Integrity Seed", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 13, Name: "Rotation Pivot", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}, {Ordinal: 6, Name: "Redstone Save Mode", Shape: Shape{Kind: "enum", Semantic: "StructureRedstoneSaveMode", TypeID: "enums/StructureRedstoneSaveMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "SavesToMemory", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "SavesToDisk", Shape: Shape{Kind: "void"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(StructureEditorData)
		if !ok {
			return p, fmt.Errorf("field StructureBlockUpdatePacket.Structure Data has unexpected decoded type %T", raw)
		}
		p.StructureData = value
	}
	{
		raw, err := r.Read("StructureBlockUpdatePacket.Trigger?", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StructureBlockUpdatePacket.Trigger? has unexpected decoded type %T", raw)
		}
		p.Trigger = value
	}
	{
		raw, err := r.Read("StructureBlockUpdatePacket.IsWaterlogged", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StructureBlockUpdatePacket.IsWaterlogged has unexpected decoded type %T", raw)
		}
		p.IsWaterlogged = value
	}
	return p, nil
}
