// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SmithingTrimRecipe struct {
	RecipeID           string
	TemplateIngredient RecipeIngredientSerializedData
	BaseIngredient     RecipeIngredientSerializedData
	AdditionIngredient RecipeIngredientSerializedData
	Tag                string
	NetID              RecipeNetID
}

// Marshal reads or writes SmithingTrimRecipe using its canonical wire layout.
func (x *SmithingTrimRecipe) Marshal(io IO) {
	io.String(&x.RecipeID)
	x.TemplateIngredient.Marshal(io)
	x.BaseIngredient.Marshal(io)
	x.AdditionIngredient.Marshal(io)
	io.String(&x.Tag)
	x.NetID.Marshal(io)
}
