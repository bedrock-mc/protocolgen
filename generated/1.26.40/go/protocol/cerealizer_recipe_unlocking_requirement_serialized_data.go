// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealizerRecipeUnlockingRequirementSerializedData struct {
	UnlockingContext     RecipeUnlockingRequirementUnlockingContext
	UnlockingIngredients Optional[[]CerealizerRecipeIngredientSerializedData]
}

// Marshal reads or writes CerealizerRecipeUnlockingRequirementSerializedData using its canonical wire layout.
func (x *CerealizerRecipeUnlockingRequirementSerializedData) Marshal(io IO) {
	IntegerFunc(&x.UnlockingContext, io.Varint32)
	OptionalFunc(io, &x.UnlockingIngredients, func(value *[]CerealizerRecipeIngredientSerializedData) {
		Slice(io, value)
	})
}
