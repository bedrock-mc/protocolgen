// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemInstance struct {
	ItemDescriptor ItemDescriptor
	StackSize      uint16
	BlockRuntimeID uint32
	UserDataBuffer []byte
}

// Marshal reads or writes ItemInstance using its canonical wire layout.
func (x *ItemInstance) Marshal(io IO) {
	MarshalItemDescriptor(io, &x.ItemDescriptor)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.BlockRuntimeID)
	io.Bytes(&x.UserDataBuffer)
}
