// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CraftingData struct {
	ShapedRecipes             []ShapedRecipe
	ShapelessRecipes          []ShapelessRecipe
	MultiRecipes              []MultiRecipe
	UserDataShapelessRecipes  []ShapelessRecipe
	ShapelessChemistryRecipes []ShapelessRecipe
	ShapedChemistryRecipes    []ShapedRecipe
	SmithingTransformRecipes  []SmithingTransformRecipe
	SmithingTrimRecipes       []SmithingTrimRecipe
	PotionMixes               []PotionMixDataEntry
	ContainerMixes            []ContainerMixDataEntry
	MaterialReducers          []MaterialReducerDataEntry
	ClearRecipes              bool
}

// Marshal reads or writes CraftingData using its canonical wire layout.
func (x *CraftingData) Marshal(io IO) {
	FuncSlice(io, &x.ShapedRecipes, io.Varuint32, func(value *ShapedRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.ShapelessRecipes, io.Varuint32, func(value *ShapelessRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.MultiRecipes, io.Varuint32, func(value *MultiRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.UserDataShapelessRecipes, io.Varuint32, func(value *ShapelessRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.ShapelessChemistryRecipes, io.Varuint32, func(value *ShapelessRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.ShapedChemistryRecipes, io.Varuint32, func(value *ShapedRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.SmithingTransformRecipes, io.Varuint32, func(value *SmithingTransformRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.SmithingTrimRecipes, io.Varuint32, func(value *SmithingTrimRecipe) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.PotionMixes, io.Varuint32, func(value *PotionMixDataEntry) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.ContainerMixes, io.Varuint32, func(value *ContainerMixDataEntry) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.MaterialReducers, io.Varuint32, func(value *MaterialReducerDataEntry) {
		value.Marshal(io)
	})
	io.Bool(&x.ClearRecipes)
}
