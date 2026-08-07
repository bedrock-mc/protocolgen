// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type SetPlayerInventoryOptions struct {
	InventoryOptions InventoryOptions
}

func (p *SetPlayerInventoryOptions) Encode(w Encoder) error {
	if err := w.Write("SetPlayerInventoryOptionsPacket.Inventory Options", Shape{Kind: "struct", Semantic: "InventoryOptions", TypeID: "InventoryOptions", Fields: []ShapeField{{Ordinal: 0, Name: "Left Inventory Tab", Shape: Shape{Kind: "enum", Semantic: "InventoryLeftTabIndex", TypeID: "enums/InventoryLeftTabIndex", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "RecipeConstruction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "RecipeEquipment", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RecipeItems", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "RecipeNature", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "RecipeSearch", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Survival", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Right Inventory Tab", Shape: Shape{Kind: "enum", Semantic: "InventoryRightTabIndex", TypeID: "enums/InventoryRightTabIndex", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "FullScreen", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Crafting", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Armor", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Filtering", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Layout Inv", Shape: Shape{Kind: "enum", Semantic: "InventoryLayout", TypeID: "enums/InventoryLayout", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InventoryOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RecipeBookOnly", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 4, Name: "Layout Craft", Shape: Shape{Kind: "enum", Semantic: "InventoryLayout", TypeID: "enums/InventoryLayout", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InventoryOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RecipeBookOnly", Shape: Shape{Kind: "void"}}}}}}}, p.InventoryOptions); err != nil {
		return err
	}
	return nil
}

func DecodeSetPlayerInventoryOptions(r Decoder) (SetPlayerInventoryOptions, error) {
	var p SetPlayerInventoryOptions
	{
		raw, err := r.Read("SetPlayerInventoryOptionsPacket.Inventory Options", Shape{Kind: "struct", Semantic: "InventoryOptions", TypeID: "InventoryOptions", Fields: []ShapeField{{Ordinal: 0, Name: "Left Inventory Tab", Shape: Shape{Kind: "enum", Semantic: "InventoryLeftTabIndex", TypeID: "enums/InventoryLeftTabIndex", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "RecipeConstruction", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "RecipeEquipment", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RecipeItems", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "RecipeNature", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "RecipeSearch", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Survival", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Right Inventory Tab", Shape: Shape{Kind: "enum", Semantic: "InventoryRightTabIndex", TypeID: "enums/InventoryRightTabIndex", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "FullScreen", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Crafting", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Armor", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 2, Name: "Filtering", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 3, Name: "Layout Inv", Shape: Shape{Kind: "enum", Semantic: "InventoryLayout", TypeID: "enums/InventoryLayout", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InventoryOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RecipeBookOnly", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 4, Name: "Layout Craft", Shape: Shape{Kind: "enum", Semantic: "InventoryLayout", TypeID: "enums/InventoryLayout", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InventoryOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "RecipeBookOnly", Shape: Shape{Kind: "void"}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(InventoryOptions)
		if !ok {
			return p, fmt.Errorf("field SetPlayerInventoryOptionsPacket.Inventory Options has unexpected decoded type %T", raw)
		}
		p.InventoryOptions = value
	}
	return p, nil
}
