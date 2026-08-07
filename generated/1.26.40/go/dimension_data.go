// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type DimensionData struct {
	Definitions []OrderedEntry[string, DimensionDefinitionGroupDimensionDefinition]
}

func (p *DimensionData) Encode(w Encoder) error {
	if err := w.Write("DimensionDataPacket.Definitions", Shape{Kind: "map", Representation: "ordered_entries", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Value: &Shape{Kind: "struct", Semantic: "DimensionDefinitionGroup::DimensionDefinition", TypeID: "DimensionDefinitionGroup::DimensionDefinition", Fields: []ShapeField{{Ordinal: 0, Name: "Height Maximum", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Height Minimum", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Generator Type", Shape: Shape{Kind: "enum", Semantic: "GeneratorType", TypeID: "enums/GeneratorType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Legacy", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Overworld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Flat", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Nether", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "TheEnd", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Void", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Undefined", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Dimension Type", Shape: Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 4, Name: "Pack Id", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}}}, Key: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, p.Definitions); err != nil {
		return err
	}
	return nil
}

func DecodeDimensionData(r Decoder) (DimensionData, error) {
	var p DimensionData
	{
		raw, err := r.Read("DimensionDataPacket.Definitions", Shape{Kind: "map", Representation: "ordered_entries", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Value: &Shape{Kind: "struct", Semantic: "DimensionDefinitionGroup::DimensionDefinition", TypeID: "DimensionDefinitionGroup::DimensionDefinition", Fields: []ShapeField{{Ordinal: 0, Name: "Height Maximum", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Height Minimum", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Generator Type", Shape: Shape{Kind: "enum", Semantic: "GeneratorType", TypeID: "enums/GeneratorType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Legacy", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Overworld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Flat", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Nether", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "TheEnd", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Void", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Undefined", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Dimension Type", Shape: Shape{Kind: "struct", Semantic: "DimensionType", TypeID: "DimensionType", Fields: []ShapeField{{Ordinal: 0, Name: "value", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 4, Name: "Pack Id", Shape: Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}}}}, Key: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]OrderedEntry[string, DimensionDefinitionGroupDimensionDefinition])
		if !ok {
			return p, fmt.Errorf("field DimensionDataPacket.Definitions has unexpected decoded type %T", raw)
		}
		p.Definitions = value
	}
	return p, nil
}
