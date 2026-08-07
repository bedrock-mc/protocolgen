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
	protocol.FuncSlice(io, &x.ShapedRecipes, io.Varuint32, func(value *protocol.ShapedRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.ShapelessRecipes, io.Varuint32, func(value *protocol.ShapelessRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.MultiRecipes, io.Varuint32, func(value *protocol.MultiRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.UserDataShapelessRecipes, io.Varuint32, func(value *protocol.ShapelessRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.ShapelessChemistryRecipes, io.Varuint32, func(value *protocol.ShapelessRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.ShapedChemistryRecipes, io.Varuint32, func(value *protocol.ShapedRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.SmithingTransformRecipes, io.Varuint32, func(value *protocol.SmithingTransformRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.SmithingTrimRecipes, io.Varuint32, func(value *protocol.SmithingTrimRecipe) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.PotionMixes, io.Varuint32, func(value *protocol.PotionMixDataEntry) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.ContainerMixes, io.Varuint32, func(value *protocol.ContainerMixDataEntry) {
		value.Marshal(io)
	})
	protocol.FuncSlice(io, &x.MaterialReducers, io.Varuint32, func(value *protocol.MaterialReducerDataEntry) {
		value.Marshal(io)
	})
	io.Bool(&x.ClearRecipes)
}
