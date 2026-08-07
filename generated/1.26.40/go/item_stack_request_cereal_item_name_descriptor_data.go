// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealItemNameDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
	FullName       string
	AuxValue       int32
}

func (ItemStackRequestCerealItemNameDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

// Marshal reads or writes ItemStackRequestCerealItemNameDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealItemNameDescriptorData) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.String(&x.FullName)
	io.Varint32(&x.AuxValue)
}
