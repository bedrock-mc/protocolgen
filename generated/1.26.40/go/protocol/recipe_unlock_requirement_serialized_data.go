// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RecipeUnlockRequirementSerializedData struct {
	UnlockingContext     RecipeUnlockingRequirementUnlockingContext
	UnlockingIngredients Optional[[]RecipeIngredientSerializedData]
}

// Marshal reads or writes RecipeUnlockRequirementSerializedData using its canonical wire layout.
func (x *RecipeUnlockRequirementSerializedData) Marshal(io IO) {
	IntegerFunc(&x.UnlockingContext, io.Varint32)
	OptionalFunc(io, &x.UnlockingIngredients, func(value *[]RecipeIngredientSerializedData) {
		Slice(io, value)
	})
}
