// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SmithingTrimRecipe struct {
	RecipeId           string
	TemplateIngredient CerealizerRecipeIngredientSerializedData
	BaseIngredient     CerealizerRecipeIngredientSerializedData
	AdditionIngredient CerealizerRecipeIngredientSerializedData
	Tag                string
	NetId              TypedServerNetIdStructRecipeNetIdTag
}

// Marshal reads or writes SmithingTrimRecipe using its canonical wire layout.
func (x *SmithingTrimRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	x.TemplateIngredient.Marshal(io)
	x.BaseIngredient.Marshal(io)
	x.AdditionIngredient.Marshal(io)
	io.String(&x.Tag)
	x.NetId.Marshal(io)
}
