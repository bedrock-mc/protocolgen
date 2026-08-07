// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CerealizerNetworkItemInstanceDescriptorSerializedData struct {
	Id             int32
	StackSize      uint16
	AuxValue       uint32
	BlockRuntimeId int32
	UserDataBuffer []byte
}

// Marshal reads or writes CerealizerNetworkItemInstanceDescriptorSerializedData using its canonical wire layout.
func (x *CerealizerNetworkItemInstanceDescriptorSerializedData) Marshal(io IO) {
	io.Varint32(&x.Id)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.AuxValue)
	io.Varint32(&x.BlockRuntimeId)
	io.Bytes(&x.UserDataBuffer)
}
