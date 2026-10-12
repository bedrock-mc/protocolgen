package retail

import (
	"fmt"

	"github.com/google/uuid"
	gtprotocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	gtpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// CraftingInputs encodes, with the pinned gophertunnel, CraftingData variants
// covering every recipe list and ingredient descriptor kind it writes; the
// corpus only has shaped recipes with default descriptors.
func CraftingInputs() ([]Input, error) {
	var inputs []Input
	for _, variant := range craftingVariants() {
		payload, err := encodeGophertunnel(gtpacket.Header{PacketID: variant.packet.ID()}, variant.packet)
		if err != nil {
			return nil, fmt.Errorf("encode %s: %w", variant.name, err)
		}
		inputs = append(inputs, Input{
			Fixture: variant.name + ".bin",
			Case:    CaseCrafting,
			Sender:  SenderServer,
			Batch:   JoinBatch([]Entry{{Payload: payload}}),
		})
	}
	return inputs, nil
}

type craftingVariant struct {
	name   string
	packet *gtpacket.CraftingData
}

func craftingVariants() []craftingVariant {
	output := []gtprotocol.ItemStack{{ItemType: gtprotocol.ItemType{NetworkID: 7, MetadataValue: 1}, Count: 4}}
	recipeUUID := uuid.UUID{0: 0x11, 7: 0x40, 15: 0x22}
	named := func(name string) gtprotocol.ItemDescriptorCount {
		return gtprotocol.ItemDescriptorCount{Descriptor: &gtprotocol.DefaultItemDescriptor{Name: name, MetadataValue: 2}, Count: 1}
	}
	tagged := gtprotocol.ItemDescriptorCount{Descriptor: &gtprotocol.ItemTagItemDescriptor{Tag: "minecraft:logs"}, Count: 2}
	molang := gtprotocol.ItemDescriptorCount{Descriptor: &gtprotocol.MoLangItemDescriptor{Expression: "q.any_tag('minecraft:planks')", Version: 12}, Count: 1}
	invalid := gtprotocol.ItemDescriptorCount{Descriptor: &gtprotocol.InvalidItemDescriptor{}}
	shapeless := func(id uint32, inputs ...gtprotocol.ItemDescriptorCount) gtprotocol.ShapelessRecipe {
		return gtprotocol.ShapelessRecipe{
			RecipeID: fmt.Sprintf("harness:shapeless_%d", id), Input: inputs, Output: output,
			UUID: recipeUUID, Block: "crafting_table", Priority: 3, RecipeNetworkID: id,
		}
	}
	shaped := func(id uint32, inputs ...gtprotocol.ItemDescriptorCount) gtprotocol.ShapedRecipe {
		return gtprotocol.ShapedRecipe{
			RecipeID: fmt.Sprintf("harness:shaped_%d", id), Width: int32(len(inputs)), Height: 1, Input: inputs,
			Output: output, UUID: recipeUUID, Block: "crafting_table", Priority: -2, AssumeSymmetry: true,
			RecipeNetworkID: id,
		}
	}
	unlockedByIngredients := shapeless(5, named("minecraft:oak_planks"))
	unlockedByIngredients.UnlockRequirement = gtprotocol.Option(gtprotocol.RecipeUnlockRequirement{
		Context: gtprotocol.RecipeUnlockContextNone, Ingredients: []gtprotocol.ItemDescriptorCount{named("minecraft:stick")},
	})
	unlockedInWater := shaped(6, named("minecraft:kelp"), tagged)
	unlockedInWater.UnlockRequirement = gtprotocol.Option(gtprotocol.RecipeUnlockRequirement{
		Context: gtprotocol.RecipeUnlockContextPlayerInWater,
	})
	return []craftingVariant{
		{"crafting_shapeless_default", &gtpacket.CraftingData{ShapelessRecipes: []gtprotocol.ShapelessRecipe{shapeless(1, named("minecraft:oak_log"))}}},
		{"crafting_shapeless_invalid", &gtpacket.CraftingData{ShapelessRecipes: []gtprotocol.ShapelessRecipe{shapeless(2, invalid)}}},
		{"crafting_shapeless_molang", &gtpacket.CraftingData{ShapelessRecipes: []gtprotocol.ShapelessRecipe{shapeless(3, molang)}}},
		{"crafting_shapeless_item_tag", &gtpacket.CraftingData{ShapelessRecipes: []gtprotocol.ShapelessRecipe{shapeless(4, tagged)}}},
		{"crafting_shapeless_unlock_ingredients", &gtpacket.CraftingData{ShapelessRecipes: []gtprotocol.ShapelessRecipe{unlockedByIngredients}}},
		{"crafting_shaped_unlock_context", &gtpacket.CraftingData{ShapedRecipes: []gtprotocol.ShapedRecipe{unlockedInWater}}},
		{"crafting_user_data_and_chemistry", &gtpacket.CraftingData{
			UserDataShapelessRecipes:  []gtprotocol.UserDataShapelessRecipe{{ShapelessRecipe: shapeless(7, named("minecraft:white_dye"))}},
			ShapelessChemistryRecipes: []gtprotocol.ShapelessChemistryRecipe{{ShapelessRecipe: shapeless(8, named("minecraft:compound"))}},
			ShapedChemistryRecipes:    []gtprotocol.ShapedChemistryRecipe{{ShapedRecipe: shaped(9, named("minecraft:element_1"), named("minecraft:element_8"))}},
		}},
		{"crafting_multi", &gtpacket.CraftingData{MultiRecipes: []gtprotocol.MultiRecipe{{UUID: recipeUUID, RecipeNetworkID: 10}}}},
		{"crafting_smithing", &gtpacket.CraftingData{
			SmithingTransformRecipes: []gtprotocol.SmithingTransformRecipe{{
				RecipeNetworkID: 11, RecipeID: "harness:smithing_transform", Template: named("minecraft:netherite_upgrade_smithing_template"),
				Base: named("minecraft:diamond_sword"), Addition: tagged, Result: output[0], Block: "smithing_table",
			}},
			SmithingTrimRecipes: []gtprotocol.SmithingTrimRecipe{{
				RecipeNetworkID: 12, RecipeID: "harness:smithing_trim", Template: named("minecraft:coast_armor_trim_smithing_template"),
				Base: tagged, Addition: named("minecraft:iron_ingot"), Block: "smithing_table",
			}},
		}},
		{"crafting_potions", &gtpacket.CraftingData{
			PotionRecipes: []gtprotocol.PotionRecipe{{
				InputPotionID: 1, InputPotionMetadata: 0, ReagentItemID: 2, ReagentItemMetadata: 0, OutputPotionID: 3, OutputPotionMetadata: 1,
			}},
			PotionContainerChangeRecipes: []gtprotocol.PotionContainerChangeRecipe{{InputItemID: 4, ReagentItemID: 5, OutputItemID: 6}},
		}},
		{"crafting_material_reducers", &gtpacket.CraftingData{
			MaterialReducers: []gtprotocol.MaterialReducer{{
				InputItem: gtprotocol.ItemType{NetworkID: 8, MetadataValue: 2},
				Outputs:   []gtprotocol.MaterialReducerOutput{{NetworkID: 9, Count: 3}},
			}},
			ClearRecipes: true,
		}},
	}
}
