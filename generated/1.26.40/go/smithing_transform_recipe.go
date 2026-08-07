// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SmithingTransformRecipe struct {
	RecipeId           string
	TemplateIngredient CerealizerRecipeIngredientSerializedData
	BaseIngredient     CerealizerRecipeIngredientSerializedData
	AdditionIngredient CerealizerRecipeIngredientSerializedData
	Result             CerealizerNetworkItemInstanceDescriptorSerializedData
	Tag                string
	NetId              TypedServerNetIdStructRecipeNetIdTag
}

// Marshal reads or writes SmithingTransformRecipe using its canonical wire layout.
func (x *SmithingTransformRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	x.TemplateIngredient.Marshal(io)
	x.BaseIngredient.Marshal(io)
	x.AdditionIngredient.Marshal(io)
	x.Result.Marshal(io)
	io.String(&x.Tag)
	x.NetId.Marshal(io)
}
