// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type ShapedRecipe struct {
	RecipeID             string
	Width                int32
	Height               int32
	Ingredients          []RecipeIngredientSerializedData
	Results              []NetworkItemInstanceDescriptorSerializedData
	UUID                 uuid.UUID
	Tag                  string
	Priority             int32
	AssumeSymmetry       bool
	UnlockingRequirement Optional[RecipeUnlockRequirementSerializedData]
	NetID                RecipeNetID
}

// Marshal reads or writes ShapedRecipe using its canonical wire layout.
func (x *ShapedRecipe) Marshal(io IO) {
	io.String(&x.RecipeID)
	io.Varint32(&x.Width)
	io.Varint32(&x.Height)
	Slice(io, &x.Ingredients)
	Slice(io, &x.Results)
	io.UUID(&x.UUID)
	io.String(&x.Tag)
	io.Varint32(&x.Priority)
	io.Bool(&x.AssumeSymmetry)
	OptionalFunc(io, &x.UnlockingRequirement, func(value *RecipeUnlockRequirementSerializedData) {
		value.Marshal(io)
	})
	x.NetID.Marshal(io)
}
