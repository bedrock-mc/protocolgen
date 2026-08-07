// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealNetworkItemInstanceDescriptorData struct {
	ItemDescriptor ItemStackRequestCerealRecipeIngredientDataItemDescriptor
	StackSize      uint16
	BlockRuntimeId uint32
	UserDataBuffer []byte
}

// Marshal reads or writes ItemStackRequestCerealNetworkItemInstanceDescriptorData using its canonical wire layout.
func (x *ItemStackRequestCerealNetworkItemInstanceDescriptorData) Marshal(io IO) {
	marshalItemStackRequestCerealRecipeIngredientDataItemDescriptor(io, &x.ItemDescriptor)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.BlockRuntimeId)
	io.Bytes(&x.UserDataBuffer)
}
