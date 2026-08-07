// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type VoxelShapes struct {
	Shapes           []VoxelShapesSerializableVoxelShape
	NameMap          []OrderedEntry[string, VoxelShapesRegistryHandle]
	CustomShapeCount uint16
}

func (p *VoxelShapes) Encode(w Encoder) error {
	if err := w.Write("VoxelShapesPacket.Shapes", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "VoxelShapes::SerializableVoxelShape", TypeID: "VoxelShapes::SerializableVoxelShape", Fields: []ShapeField{{Ordinal: 0, Name: "Cells", Shape: Shape{Kind: "struct", Semantic: "VoxelShapes::SerializableCells", TypeID: "VoxelShapes::SerializableCells", Fields: []ShapeField{{Ordinal: 0, Name: "X Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 1, Name: "Y Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 2, Name: "Z Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 3, Name: "Storage", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}}, {Ordinal: 1, Name: "X Coordinates", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 2, Name: "Y Coordinates", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 3, Name: "Z Coordinates", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}, p.Shapes); err != nil {
		return err
	}
	if err := w.Write("VoxelShapesPacket.Name Map", Shape{Kind: "map", Representation: "ordered_entries", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Value: &Shape{Kind: "struct", Semantic: "VoxelShapes::RegistryHandle", TypeID: "VoxelShapes::RegistryHandle", Fields: []ShapeField{{Ordinal: 0, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}}}, Key: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.NameMap); err != nil {
		return err
	}
	if err := w.Write("VoxelShapesPacket.Custom Shape Count", Shape{Kind: "primitive", PrimitiveCode: "u16le"}, p.CustomShapeCount); err != nil {
		return err
	}
	return nil
}

func DecodeVoxelShapes(r Decoder) (VoxelShapes, error) {
	var p VoxelShapes
	{
		raw, err := r.Read("VoxelShapesPacket.Shapes", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "VoxelShapes::SerializableVoxelShape", TypeID: "VoxelShapes::SerializableVoxelShape", Fields: []ShapeField{{Ordinal: 0, Name: "Cells", Shape: Shape{Kind: "struct", Semantic: "VoxelShapes::SerializableCells", TypeID: "VoxelShapes::SerializableCells", Fields: []ShapeField{{Ordinal: 0, Name: "X Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 1, Name: "Y Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 2, Name: "Z Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u8"}}, {Ordinal: 3, Name: "Storage", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "u8"}}}}}}, {Ordinal: 1, Name: "X Coordinates", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 2, Name: "Y Coordinates", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}, {Ordinal: 3, Name: "Z Coordinates", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]VoxelShapesSerializableVoxelShape)
		if !ok {
			return p, fmt.Errorf("field VoxelShapesPacket.Shapes has unexpected decoded type %T", raw)
		}
		p.Shapes = value
	}
	{
		raw, err := r.Read("VoxelShapesPacket.Name Map", Shape{Kind: "map", Representation: "ordered_entries", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Value: &Shape{Kind: "struct", Semantic: "VoxelShapes::RegistryHandle", TypeID: "VoxelShapes::RegistryHandle", Fields: []ShapeField{{Ordinal: 0, Name: "Value", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}}}, Key: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]OrderedEntry[string, VoxelShapesRegistryHandle])
		if !ok {
			return p, fmt.Errorf("field VoxelShapesPacket.Name Map has unexpected decoded type %T", raw)
		}
		p.NameMap = value
	}
	{
		raw, err := r.Read("VoxelShapesPacket.Custom Shape Count", Shape{Kind: "primitive", PrimitiveCode: "u16le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint16)
		if !ok {
			return p, fmt.Errorf("field VoxelShapesPacket.Custom Shape Count has unexpected decoded type %T", raw)
		}
		p.CustomShapeCount = value
	}
	return p, nil
}
