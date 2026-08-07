// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/google/uuid"

type MultiRecipe struct {
	MultiRecipeUUID uuid.UUID
	NetId           TypedServerNetIdStructRecipeNetIdTag
}

// Marshal reads or writes MultiRecipe using its canonical wire layout.
func (x *MultiRecipe) Marshal(io IO) {
	io.UUID(&x.MultiRecipeUUID)
	x.NetId.Marshal(io)
}
