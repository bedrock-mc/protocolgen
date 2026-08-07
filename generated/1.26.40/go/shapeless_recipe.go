// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type ShapelessRecipe struct {
	RecipeId             string
	Ingredients          []CerealizerRecipeIngredientSerializedData
	Results              []CerealizerNetworkItemInstanceDescriptorSerializedData
	UUID                 uuid.UUID
	Tag                  string
	Priority             int32
	UnlockingRequirement Optional[CerealizerRecipeUnlockingRequirementSerializedData]
	NetId                TypedServerNetIdStructRecipeNetIdTag
}

// Marshal reads or writes ShapelessRecipe using its canonical wire layout.
func (x *ShapelessRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	FuncSlice(io, &x.Ingredients, io.Varuint32, func(value *CerealizerRecipeIngredientSerializedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.Results, io.Varuint32, func(value *CerealizerNetworkItemInstanceDescriptorSerializedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.UUID(&x.UUID)
	io.String(&x.Tag)
	io.Varint32(&x.Priority)
	OptionalFunc(io, &x.UnlockingRequirement, func(value *CerealizerRecipeUnlockingRequirementSerializedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.NetId.Marshal(io)
}
