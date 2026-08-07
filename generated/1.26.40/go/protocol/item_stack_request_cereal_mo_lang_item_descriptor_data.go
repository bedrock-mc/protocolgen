// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealMoLangItemDescriptorData struct {
	DescriptorType ItemStackRequestCerealItemDescriptorType
	TagExpression  string
	MolangVersion  MoLangVersion
}

func (ItemStackRequestCerealMoLangItemDescriptorData) isItemStackRequestCerealRecipeIngredientDataItemDescriptor() {
}

// Marshal reads or writes ItemStackRequestCerealMoLangItemDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealMoLangItemDescriptorData) Marshal(io IO) {
	IntegerFunc(&x.DescriptorType, io.Uint8)
	io.String(&x.TagExpression)
	IntegerFunc(&x.MolangVersion, io.Int16)
}
