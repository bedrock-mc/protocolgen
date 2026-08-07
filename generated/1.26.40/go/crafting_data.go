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
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.ShapelessRecipes, io.Varuint32, func(value *ShapelessRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.MultiRecipes, io.Varuint32, func(value *MultiRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.UserDataShapelessRecipes, io.Varuint32, func(value *ShapelessRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.ShapelessChemistryRecipes, io.Varuint32, func(value *ShapelessRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.ShapedChemistryRecipes, io.Varuint32, func(value *ShapedRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.SmithingTransformRecipes, io.Varuint32, func(value *SmithingTransformRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.SmithingTrimRecipes, io.Varuint32, func(value *SmithingTrimRecipe) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.PotionMixes, io.Varuint32, func(value *PotionMixDataEntry) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.ContainerMixes, io.Varuint32, func(value *ContainerMixDataEntry) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.MaterialReducers, io.Varuint32, func(value *MaterialReducerDataEntry) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.Bool(&x.ClearRecipes)
}
