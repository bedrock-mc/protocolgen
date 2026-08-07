// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type ShapedRecipe struct {
	RecipeId             string
	Width                int32
	Height               int32
	Ingredients          []CerealizerRecipeIngredientSerializedData
	Results              []CerealizerNetworkItemInstanceDescriptorSerializedData
	UUID                 uuid.UUID
	Tag                  string
	Priority             int32
	AssumeSymmetry       bool
	UnlockingRequirement Optional[CerealizerRecipeUnlockingRequirementSerializedData]
	NetId                TypedServerNetIdStructRecipeNetIdTag
}

// Marshal reads or writes ShapedRecipe using its canonical wire layout.
func (x *ShapedRecipe) Marshal(io IO) {
	io.String(&x.RecipeId)
	io.Varint32(&x.Width)
	io.Varint32(&x.Height)
	FuncSlice(io, &x.Ingredients, io.Varuint32, func(value *CerealizerRecipeIngredientSerializedData) {
		value.Marshal(io)
	})
	FuncSlice(io, &x.Results, io.Varuint32, func(value *CerealizerNetworkItemInstanceDescriptorSerializedData) {
		value.Marshal(io)
	})
	io.UUID(&x.UUID)
	io.String(&x.Tag)
	io.Varint32(&x.Priority)
	io.Bool(&x.AssumeSymmetry)
	OptionalFunc(io, &x.UnlockingRequirement, func(value *CerealizerRecipeUnlockingRequirementSerializedData) {
		value.Marshal(io)
	})
	x.NetId.Marshal(io)
}
