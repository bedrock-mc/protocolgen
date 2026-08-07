// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SmithingTransformRecipe struct {
	RecipeID           string
	TemplateIngredient RecipeIngredientSerializedData
	BaseIngredient     RecipeIngredientSerializedData
	AdditionIngredient RecipeIngredientSerializedData
	Result             NetworkItemInstanceDescriptorSerializedData
	Tag                string
	NetID              RecipeNetID
}

// Marshal reads or writes SmithingTransformRecipe using its canonical wire layout.
func (x *SmithingTransformRecipe) Marshal(io IO) {
	io.String(&x.RecipeID)
	x.TemplateIngredient.Marshal(io)
	x.BaseIngredient.Marshal(io)
	x.AdditionIngredient.Marshal(io)
	x.Result.Marshal(io)
	io.String(&x.Tag)
	x.NetID.Marshal(io)
}
