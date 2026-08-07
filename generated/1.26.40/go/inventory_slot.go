// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type InventorySlot struct {
	ContainerId       uint8
	Slot              uint32
	FullContainerName *FullContainerName
	StorageItem       *CerealizerNetworkItemStackDescriptorSerializedData
	Item              CerealizerNetworkItemStackDescriptorSerializedData
}

func (p *InventorySlot) Encode(w Encoder) error {
	if err := w.Write("InventorySlotPacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerId); err != nil {
		return err
	}
	if err := w.Write("InventorySlotPacket.Slot", Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, p.Slot); err != nil {
		return err
	}
	if err := w.Write("InventorySlotPacket.Full Container Name", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "FullContainerName", TypeID: "FullContainerName", Fields: []ShapeField{{Ordinal: 0, Name: "Container Name", Shape: Shape{Kind: "enum", Semantic: "ContainerEnumName", TypeID: "enums/ContainerEnumName", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "AnvilInputContainer", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "AnvilMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "AnvilResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "SmithingTableInputContainer", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "SmithingTableMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "SmithingTableResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "ArmorContainer", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "LevelEntityContainer", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "BeaconPaymentContainer", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "BrewingStandInputContainer", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "BrewingStandResultContainer", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "BrewingStandFuelContainer", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "CombinedHotbarAndInventoryContainer", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "CraftingInputContainer", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "CraftingOutputPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "RecipeConstructionContainer", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "RecipeNatureContainer", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "RecipeItemsContainer", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "RecipeSearchContainer", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "RecipeSearchBarContainer", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "RecipeEquipmentContainer", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "RecipeBookContainer", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "EnchantingInputContainer", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "EnchantingMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "FurnaceFuelContainer", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "FurnaceIngredientContainer", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "FurnaceResultContainer", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "HorseEquipContainer", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "HotbarContainer", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "InventoryContainer", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "ShulkerBoxContainer", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "TradeIngredient1Container", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "TradeIngredient2Container", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "TradeResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "OffhandContainer", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "CompoundCreatorInput", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "CompoundCreatorOutputPreview", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "ElementConstructorOutputPreview", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "MaterialReducerInput", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "MaterialReducerOutput", Shape: Shape{Kind: "void"}}, {Value: 40, Name: "LabTableInput", Shape: Shape{Kind: "void"}}, {Value: 41, Name: "LoomInputContainer", Shape: Shape{Kind: "void"}}, {Value: 42, Name: "LoomDyeContainer", Shape: Shape{Kind: "void"}}, {Value: 43, Name: "LoomMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 44, Name: "LoomResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 45, Name: "BlastFurnaceIngredientContainer", Shape: Shape{Kind: "void"}}, {Value: 46, Name: "SmokerIngredientContainer", Shape: Shape{Kind: "void"}}, {Value: 47, Name: "Trade2Ingredient1Container", Shape: Shape{Kind: "void"}}, {Value: 48, Name: "Trade2Ingredient2Container", Shape: Shape{Kind: "void"}}, {Value: 49, Name: "Trade2ResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 50, Name: "GrindstoneInputContainer", Shape: Shape{Kind: "void"}}, {Value: 51, Name: "GrindstoneAdditionalContainer", Shape: Shape{Kind: "void"}}, {Value: 52, Name: "GrindstoneResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 53, Name: "StonecutterInputContainer", Shape: Shape{Kind: "void"}}, {Value: 54, Name: "StonecutterResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 55, Name: "CartographyInputContainer", Shape: Shape{Kind: "void"}}, {Value: 56, Name: "CartographyAdditionalContainer", Shape: Shape{Kind: "void"}}, {Value: 57, Name: "CartographyResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 58, Name: "BarrelContainer", Shape: Shape{Kind: "void"}}, {Value: 59, Name: "CursorContainer", Shape: Shape{Kind: "void"}}, {Value: 60, Name: "CreatedOutputContainer", Shape: Shape{Kind: "void"}}, {Value: 61, Name: "SmithingTableTemplateContainer", Shape: Shape{Kind: "void"}}, {Value: 62, Name: "CrafterLevelEntityContainer", Shape: Shape{Kind: "void"}}, {Value: 63, Name: "DynamicContainer", Shape: Shape{Kind: "void"}}, {Value: 64, Name: "RecipeFoodContainer", Shape: Shape{Kind: "void"}}, {Value: 65, Name: "RecipeBlocksContainer", Shape: Shape{Kind: "void"}}, {Value: 66, Name: "RecipeFurnaceItemsContainer", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Dynamic ID", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}}}}, p.FullContainerName); err != nil {
		return err
	}
	if err := w.Write("InventorySlotPacket.Storage Item", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, p.StorageItem); err != nil {
		return err
	}
	if err := w.Write("InventorySlotPacket.Item", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.Item); err != nil {
		return err
	}
	return nil
}

func DecodeInventorySlot(r Decoder) (InventorySlot, error) {
	var p InventorySlot
	{
		raw, err := r.Read("InventorySlotPacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field InventorySlotPacket.Container Id has unexpected decoded type %T", raw)
		}
		p.ContainerId = value
	}
	{
		raw, err := r.Read("InventorySlotPacket.Slot", Shape{Kind: "primitive", PrimitiveCode: "var_u32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint32)
		if !ok {
			return p, fmt.Errorf("field InventorySlotPacket.Slot has unexpected decoded type %T", raw)
		}
		p.Slot = value
	}
	{
		raw, err := r.Read("InventorySlotPacket.Full Container Name", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "FullContainerName", TypeID: "FullContainerName", Fields: []ShapeField{{Ordinal: 0, Name: "Container Name", Shape: Shape{Kind: "enum", Semantic: "ContainerEnumName", TypeID: "enums/ContainerEnumName", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "AnvilInputContainer", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "AnvilMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "AnvilResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "SmithingTableInputContainer", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "SmithingTableMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "SmithingTableResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "ArmorContainer", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "LevelEntityContainer", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "BeaconPaymentContainer", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "BrewingStandInputContainer", Shape: Shape{Kind: "void"}}, {Value: 10, Name: "BrewingStandResultContainer", Shape: Shape{Kind: "void"}}, {Value: 11, Name: "BrewingStandFuelContainer", Shape: Shape{Kind: "void"}}, {Value: 12, Name: "CombinedHotbarAndInventoryContainer", Shape: Shape{Kind: "void"}}, {Value: 13, Name: "CraftingInputContainer", Shape: Shape{Kind: "void"}}, {Value: 14, Name: "CraftingOutputPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 15, Name: "RecipeConstructionContainer", Shape: Shape{Kind: "void"}}, {Value: 16, Name: "RecipeNatureContainer", Shape: Shape{Kind: "void"}}, {Value: 17, Name: "RecipeItemsContainer", Shape: Shape{Kind: "void"}}, {Value: 18, Name: "RecipeSearchContainer", Shape: Shape{Kind: "void"}}, {Value: 19, Name: "RecipeSearchBarContainer", Shape: Shape{Kind: "void"}}, {Value: 20, Name: "RecipeEquipmentContainer", Shape: Shape{Kind: "void"}}, {Value: 21, Name: "RecipeBookContainer", Shape: Shape{Kind: "void"}}, {Value: 22, Name: "EnchantingInputContainer", Shape: Shape{Kind: "void"}}, {Value: 23, Name: "EnchantingMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 24, Name: "FurnaceFuelContainer", Shape: Shape{Kind: "void"}}, {Value: 25, Name: "FurnaceIngredientContainer", Shape: Shape{Kind: "void"}}, {Value: 26, Name: "FurnaceResultContainer", Shape: Shape{Kind: "void"}}, {Value: 27, Name: "HorseEquipContainer", Shape: Shape{Kind: "void"}}, {Value: 28, Name: "HotbarContainer", Shape: Shape{Kind: "void"}}, {Value: 29, Name: "InventoryContainer", Shape: Shape{Kind: "void"}}, {Value: 30, Name: "ShulkerBoxContainer", Shape: Shape{Kind: "void"}}, {Value: 31, Name: "TradeIngredient1Container", Shape: Shape{Kind: "void"}}, {Value: 32, Name: "TradeIngredient2Container", Shape: Shape{Kind: "void"}}, {Value: 33, Name: "TradeResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 34, Name: "OffhandContainer", Shape: Shape{Kind: "void"}}, {Value: 35, Name: "CompoundCreatorInput", Shape: Shape{Kind: "void"}}, {Value: 36, Name: "CompoundCreatorOutputPreview", Shape: Shape{Kind: "void"}}, {Value: 37, Name: "ElementConstructorOutputPreview", Shape: Shape{Kind: "void"}}, {Value: 38, Name: "MaterialReducerInput", Shape: Shape{Kind: "void"}}, {Value: 39, Name: "MaterialReducerOutput", Shape: Shape{Kind: "void"}}, {Value: 40, Name: "LabTableInput", Shape: Shape{Kind: "void"}}, {Value: 41, Name: "LoomInputContainer", Shape: Shape{Kind: "void"}}, {Value: 42, Name: "LoomDyeContainer", Shape: Shape{Kind: "void"}}, {Value: 43, Name: "LoomMaterialContainer", Shape: Shape{Kind: "void"}}, {Value: 44, Name: "LoomResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 45, Name: "BlastFurnaceIngredientContainer", Shape: Shape{Kind: "void"}}, {Value: 46, Name: "SmokerIngredientContainer", Shape: Shape{Kind: "void"}}, {Value: 47, Name: "Trade2Ingredient1Container", Shape: Shape{Kind: "void"}}, {Value: 48, Name: "Trade2Ingredient2Container", Shape: Shape{Kind: "void"}}, {Value: 49, Name: "Trade2ResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 50, Name: "GrindstoneInputContainer", Shape: Shape{Kind: "void"}}, {Value: 51, Name: "GrindstoneAdditionalContainer", Shape: Shape{Kind: "void"}}, {Value: 52, Name: "GrindstoneResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 53, Name: "StonecutterInputContainer", Shape: Shape{Kind: "void"}}, {Value: 54, Name: "StonecutterResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 55, Name: "CartographyInputContainer", Shape: Shape{Kind: "void"}}, {Value: 56, Name: "CartographyAdditionalContainer", Shape: Shape{Kind: "void"}}, {Value: 57, Name: "CartographyResultPreviewContainer", Shape: Shape{Kind: "void"}}, {Value: 58, Name: "BarrelContainer", Shape: Shape{Kind: "void"}}, {Value: 59, Name: "CursorContainer", Shape: Shape{Kind: "void"}}, {Value: 60, Name: "CreatedOutputContainer", Shape: Shape{Kind: "void"}}, {Value: 61, Name: "SmithingTableTemplateContainer", Shape: Shape{Kind: "void"}}, {Value: 62, Name: "CrafterLevelEntityContainer", Shape: Shape{Kind: "void"}}, {Value: 63, Name: "DynamicContainer", Shape: Shape{Kind: "void"}}, {Value: 64, Name: "RecipeFoodContainer", Shape: Shape{Kind: "void"}}, {Value: 65, Name: "RecipeBlocksContainer", Shape: Shape{Kind: "void"}}, {Value: 66, Name: "RecipeFurnaceItemsContainer", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "Dynamic ID", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*FullContainerName)
		if !ok {
			return p, fmt.Errorf("field InventorySlotPacket.Full Container Name has unexpected decoded type %T", raw)
		}
		p.FullContainerName = value
	}
	{
		raw, err := r.Read("InventorySlotPacket.Storage Item", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field InventorySlotPacket.Storage Item has unexpected decoded type %T", raw)
		}
		p.StorageItem = value
	}
	{
		raw, err := r.Read("InventorySlotPacket.Item", Shape{Kind: "struct", Semantic: "cerealizer<NetworkItemStackDescriptor>::SerializedData", TypeID: "cerealizer<NetworkItemStackDescriptor>::SerializedData", Fields: []ShapeField{{Ordinal: 0, Name: "Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "i16le"}}, {Ordinal: 1, Name: "Stack size", Shape: Shape{Kind: "primitive", PrimitiveCode: "u16le"}}, {Ordinal: 2, Name: "Aux value", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 3, Name: "Net Id Variant", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}, {Ordinal: 4, Name: "Block Runtime Id", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, {Ordinal: 5, Name: "User Data Buffer", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(CerealizerNetworkItemStackDescriptorSerializedData)
		if !ok {
			return p, fmt.Errorf("field InventorySlotPacket.Item has unexpected decoded type %T", raw)
		}
		p.Item = value
	}
	return p, nil
}
