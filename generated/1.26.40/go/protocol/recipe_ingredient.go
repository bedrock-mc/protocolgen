// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RecipeIngredient struct {
	ItemDescriptor ItemDescriptor
	StackSize      uint16
}

// Marshal reads or writes RecipeIngredient using its canonical wire layout.
func (x *RecipeIngredient) Marshal(io IO) {
	MarshalItemDescriptor(io, &x.ItemDescriptor)
	io.Uint16(&x.StackSize)
}
