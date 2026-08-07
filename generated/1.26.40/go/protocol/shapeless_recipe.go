// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type ShapelessRecipe struct {
	RecipeID             string
	Ingredients          []RecipeIngredientSerializedData
	Results              []NetworkItemInstanceDescriptorSerializedData
	UUID                 uuid.UUID
	Tag                  string
	Priority             int32
	UnlockingRequirement Optional[RecipeUnlockRequirementSerializedData]
	NetID                RecipeNetID
}

// Marshal reads or writes ShapelessRecipe using its canonical wire layout.
func (x *ShapelessRecipe) Marshal(io IO) {
	io.String(&x.RecipeID)
	Slice(io, &x.Ingredients)
	Slice(io, &x.Results)
	io.UUID(&x.UUID)
	io.String(&x.Tag)
	io.Varint32(&x.Priority)
	OptionalFunc(io, &x.UnlockingRequirement, func(value *RecipeUnlockRequirementSerializedData) {
		value.Marshal(io)
	})
	x.NetID.Marshal(io)
}
