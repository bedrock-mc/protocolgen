// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ItemRegistry struct {
	ItemData []ItemData
}

func (p *ItemRegistry) Encode(w Encoder) error {
	if err := w.Write("ItemRegistryPacket.Item Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ItemData", TypeID: "ItemData", Fields: []ShapeField{{Ordinal: 0, Name: "Item Name", Shape: Shape{Kind: "string", Semantic: "hashed_string", TypeID: "hashed_string.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Item Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 2, Name: "Is Component Based", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Item Version", Shape: Shape{Kind: "enum", Semantic: "ItemVersion", TypeID: "enums/ItemVersion", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Legacy", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DataDriven", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "None", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 4, Name: "Item Component Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}}, p.ItemData); err != nil {
		return err
	}
	return nil
}

func DecodeItemRegistry(r Decoder) (ItemRegistry, error) {
	var p ItemRegistry
	{
		raw, err := r.Read("ItemRegistryPacket.Item Data", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ItemData", TypeID: "ItemData", Fields: []ShapeField{{Ordinal: 0, Name: "Item Name", Shape: Shape{Kind: "string", Semantic: "hashed_string", TypeID: "hashed_string.json#", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Item Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 2, Name: "Is Component Based", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Item Version", Shape: Shape{Kind: "enum", Semantic: "ItemVersion", TypeID: "enums/ItemVersion", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Legacy", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "DataDriven", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "None", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 4, Name: "Item Component Data", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ItemData)
		if !ok {
			return p, fmt.Errorf("field ItemRegistryPacket.Item Data has unexpected decoded type %T", raw)
		}
		p.ItemData = value
	}
	return p, nil
}
