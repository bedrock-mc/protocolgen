// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type StructureTemplateDataRequest struct {
	StructureName      string
	StructurePosition  BlockPos
	StructureSettings  StructureSettings
	RequestedOperation StructureTemplateRequestOperation
}

func (p *StructureTemplateDataRequest) Encode(w Encoder) error {
	if err := w.Write("StructureTemplateDataRequestPacket.Structure Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.StructureName); err != nil {
		return err
	}
	if err := w.Write("StructureTemplateDataRequestPacket.Structure Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}, p.StructurePosition); err != nil {
		return err
	}
	if err := w.Write("StructureTemplateDataRequestPacket.Structure Settings", Shape{Kind: "struct", Semantic: "StructureSettings", TypeID: "StructureSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Structure Palette Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Should ignore entities?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Should ignore blocks?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Should Allow Non Ticking Player and Ticking Area Chunks", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Structure Size", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 5, Name: "Structure Offset", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 6, Name: "Last Edit Player", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 7, Name: "Rotation", Shape: Shape{Kind: "enum", Semantic: "Rotation", TypeID: "enums/Rotation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotate90", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Rotate180", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Rotate270", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 8, Name: "Mirror", Shape: Shape{Kind: "enum", Semantic: "Mirror", TypeID: "enums/Mirror", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "X", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Z", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "XZ", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 9, Name: "Animation Mode", Shape: Shape{Kind: "enum", Semantic: "AnimationMode", TypeID: "enums/AnimationMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Layers", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Blocks", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 10, Name: "Animation Seconds", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 11, Name: "Integrity Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 12, Name: "Integrity Seed", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 13, Name: "Rotation Pivot", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}, p.StructureSettings); err != nil {
		return err
	}
	if err := w.Write("StructureTemplateDataRequestPacket.Requested Operation", Shape{Kind: "enum", Semantic: "StructureTemplateRequestOperation", TypeID: "enums/StructureTemplateRequestOperation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "ExportFromSaveMode", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ExportFromLoadMode", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "QuerySavedStructure", Shape: Shape{Kind: "void"}}}}, p.RequestedOperation); err != nil {
		return err
	}
	return nil
}

func DecodeStructureTemplateDataRequest(r Decoder) (StructureTemplateDataRequest, error) {
	var p StructureTemplateDataRequest
	{
		raw, err := r.Read("StructureTemplateDataRequestPacket.Structure Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataRequestPacket.Structure Name has unexpected decoded type %T", raw)
		}
		p.StructureName = value
	}
	{
		raw, err := r.Read("StructureTemplateDataRequestPacket.Structure Position", Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(BlockPos)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataRequestPacket.Structure Position has unexpected decoded type %T", raw)
		}
		p.StructurePosition = value
	}
	{
		raw, err := r.Read("StructureTemplateDataRequestPacket.Structure Settings", Shape{Kind: "struct", Semantic: "StructureSettings", TypeID: "StructureSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Structure Palette Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Should ignore entities?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Should ignore blocks?", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Should Allow Non Ticking Player and Ticking Area Chunks", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 4, Name: "Structure Size", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 5, Name: "Structure Offset", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 6, Name: "Last Edit Player", Shape: Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}}, {Ordinal: 7, Name: "Rotation", Shape: Shape{Kind: "enum", Semantic: "Rotation", TypeID: "enums/Rotation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Rotate90", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Rotate180", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Rotate270", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 8, Name: "Mirror", Shape: Shape{Kind: "enum", Semantic: "Mirror", TypeID: "enums/Mirror", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "X", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Z", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "XZ", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 9, Name: "Animation Mode", Shape: Shape{Kind: "enum", Semantic: "AnimationMode", TypeID: "enums/AnimationMode", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Layers", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Blocks", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 10, Name: "Animation Seconds", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 11, Name: "Integrity Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 12, Name: "Integrity Seed", Shape: Shape{Kind: "primitive", PrimitiveCode: "u32le"}}, {Ordinal: 13, Name: "Rotation Pivot", Shape: Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(StructureSettings)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataRequestPacket.Structure Settings has unexpected decoded type %T", raw)
		}
		p.StructureSettings = value
	}
	{
		raw, err := r.Read("StructureTemplateDataRequestPacket.Requested Operation", Shape{Kind: "enum", Semantic: "StructureTemplateRequestOperation", TypeID: "enums/StructureTemplateRequestOperation", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "ExportFromSaveMode", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "ExportFromLoadMode", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "QuerySavedStructure", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(StructureTemplateRequestOperation)
		if !ok {
			return p, fmt.Errorf("field StructureTemplateDataRequestPacket.Requested Operation has unexpected decoded type %T", raw)
		}
		p.RequestedOperation = value
	}
	return p, nil
}
