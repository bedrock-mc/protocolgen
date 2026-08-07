// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CraftingData struct {
	ShapedRecipes             []protocol.ShapedRecipe
	ShapelessRecipes          []protocol.ShapelessRecipe
	MultiRecipes              []protocol.MultiRecipe
	UserDataShapelessRecipes  []protocol.ShapelessRecipe
	ShapelessChemistryRecipes []protocol.ShapelessRecipe
	ShapedChemistryRecipes    []protocol.ShapedRecipe
	SmithingTransformRecipes  []protocol.SmithingTransformRecipe
	SmithingTrimRecipes       []protocol.SmithingTrimRecipe
	PotionMixes               []protocol.PotionMixDataEntry
	ContainerMixes            []protocol.ContainerMixDataEntry
	MaterialReducers          []protocol.MaterialReducerDataEntry
	ClearRecipes              bool
}

// Marshal reads or writes CraftingData using its canonical wire layout.
func (x *CraftingData) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.ShapedRecipes)
	protocol.Slice(io, &x.ShapelessRecipes)
	protocol.Slice(io, &x.MultiRecipes)
	protocol.Slice(io, &x.UserDataShapelessRecipes)
	protocol.Slice(io, &x.ShapelessChemistryRecipes)
	protocol.Slice(io, &x.ShapedChemistryRecipes)
	protocol.Slice(io, &x.SmithingTransformRecipes)
	protocol.Slice(io, &x.SmithingTrimRecipes)
	protocol.Slice(io, &x.PotionMixes)
	protocol.Slice(io, &x.ContainerMixes)
	protocol.Slice(io, &x.MaterialReducers)
	io.Bool(&x.ClearRecipes)
}

// ID returns the protocol ID for CraftingData.
func (*CraftingData) ID() uint32 { return IDCraftingData }
