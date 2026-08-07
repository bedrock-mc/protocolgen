// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealEmptyItemDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
}

func (ItemStackRequestCerealEmptyItemDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

// Marshal reads or writes ItemStackRequestCerealEmptyItemDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealEmptyItemDescriptorData) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
}
