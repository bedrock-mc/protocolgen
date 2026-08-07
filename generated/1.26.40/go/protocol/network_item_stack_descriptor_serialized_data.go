// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type NetworkItemStackDescriptorSerializedData struct {
	ID             int16
	StackSize      uint16
	AuxValue       uint32
	NetIDVariant   Optional[int32]
	BlockRuntimeID uint32
	UserDataBuffer []byte
}

// Marshal reads or writes NetworkItemStackDescriptorSerializedData using its canonical wire layout.
func (x *NetworkItemStackDescriptorSerializedData) Marshal(io IO) {
	io.Int16(&x.ID)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.AuxValue)
	OptionalFunc(io, &x.NetIDVariant, io.Varint32)
	io.Varuint32(&x.BlockRuntimeID)
	io.Bytes(&x.UserDataBuffer)
}
