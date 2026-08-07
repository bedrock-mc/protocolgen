// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealRecipeIngredientData struct {
	ItemDescriptor ItemStackRequestCerealRecipeIngredientDataItemDescriptor
	StackSize      uint16
}

// Marshal reads or writes ItemStackRequestCerealRecipeIngredientData using its canonical wire layout.
func (x *ItemStackRequestCerealRecipeIngredientData) Marshal(io IO) {
	MarshalItemStackRequestCerealRecipeIngredientDataItemDescriptor(io, &x.ItemDescriptor)
	io.Uint16(&x.StackSize)
}
