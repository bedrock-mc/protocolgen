// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CerealizerRecipeUnlockingRequirementSerializedData struct {
	UnlockingContext     RecipeUnlockingRequirementUnlockingContext
	UnlockingIngredients Optional[[]CerealizerRecipeIngredientSerializedData]
}

// Marshal reads or writes CerealizerRecipeUnlockingRequirementSerializedData using its canonical wire layout.
func (x *CerealizerRecipeUnlockingRequirementSerializedData) Marshal(io IO) {
	IntegerFunc(&x.UnlockingContext, io.Varint32)
	OptionalFunc(io, &x.UnlockingIngredients, func(value *[]CerealizerRecipeIngredientSerializedData) {
		item := *value
		FuncSlice(io, &item, io.Varuint32, func(value *CerealizerRecipeIngredientSerializedData) {
			item := *value
			item.Marshal(io)
			*value = item
		})
		*value = item
	})
}
