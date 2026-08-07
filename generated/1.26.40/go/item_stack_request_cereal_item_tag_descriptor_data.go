// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealItemTagDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
	ItemTag        string
}

func (ItemStackRequestCerealItemTagDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

// Marshal reads or writes ItemStackRequestCerealItemTagDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealItemTagDescriptorData) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.String(&x.ItemTag)
}
