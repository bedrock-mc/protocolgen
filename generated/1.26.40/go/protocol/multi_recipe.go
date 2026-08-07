// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/google/uuid"

type MultiRecipe struct {
	MultiRecipeUUID uuid.UUID
	NetID           RecipeNetID
}

// Marshal reads or writes MultiRecipe using its canonical wire layout.
func (x *MultiRecipe) Marshal(io IO) {
	io.UUID(&x.MultiRecipeUUID)
	x.NetID.Marshal(io)
}
