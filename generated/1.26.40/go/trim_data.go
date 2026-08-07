// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type TrimData struct {
	TrimPatternList  []TrimPattern
	TrimMaterialList []TrimMaterial
}

func (p *TrimData) Encode(w Encoder) error {
	if err := w.Write("TrimDataPacket.TrimPattern List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "TrimPattern", TypeID: "TrimPattern", Fields: []ShapeField{{Ordinal: 0, Name: "Item Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Pattern Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.TrimPatternList); err != nil {
		return err
	}
	if err := w.Write("TrimDataPacket.TrimMaterial List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "TrimMaterial", TypeID: "TrimMaterial", Fields: []ShapeField{{Ordinal: 0, Name: "Material Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Color", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Item Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.TrimMaterialList); err != nil {
		return err
	}
	return nil
}

func DecodeTrimData(r Decoder) (TrimData, error) {
	var p TrimData
	{
		raw, err := r.Read("TrimDataPacket.TrimPattern List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "TrimPattern", TypeID: "TrimPattern", Fields: []ShapeField{{Ordinal: 0, Name: "Item Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Pattern Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]TrimPattern)
		if !ok {
			return p, fmt.Errorf("field TrimDataPacket.TrimPattern List has unexpected decoded type %T", raw)
		}
		p.TrimPatternList = value
	}
	{
		raw, err := r.Read("TrimDataPacket.TrimMaterial List", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "TrimMaterial", TypeID: "TrimMaterial", Fields: []ShapeField{{Ordinal: 0, Name: "Material Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Color", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Item Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]TrimMaterial)
		if !ok {
			return p, fmt.Errorf("field TrimDataPacket.TrimMaterial List has unexpected decoded type %T", raw)
		}
		p.TrimMaterialList = value
	}
	return p, nil
}
