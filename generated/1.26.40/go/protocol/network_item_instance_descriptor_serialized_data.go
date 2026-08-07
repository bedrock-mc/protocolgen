// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type NetworkItemInstanceDescriptorSerializedData struct {
	ID             int32
	StackSize      uint16
	AuxValue       uint32
	BlockRuntimeID int32
	UserDataBuffer []byte
}

// Marshal reads or writes NetworkItemInstanceDescriptorSerializedData using its canonical wire layout.
func (x *NetworkItemInstanceDescriptorSerializedData) Marshal(io IO) {
	io.Varint32(&x.ID)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.AuxValue)
	io.Varint32(&x.BlockRuntimeID)
	io.Bytes(&x.UserDataBuffer)
}
