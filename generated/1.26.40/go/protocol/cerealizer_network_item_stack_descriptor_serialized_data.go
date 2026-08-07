// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CerealizerNetworkItemStackDescriptorSerializedData struct {
	Id             int16
	StackSize      uint16
	AuxValue       uint32
	NetIdVariant   Optional[int32]
	BlockRuntimeId uint32
	UserDataBuffer []byte
}

// Marshal reads or writes CerealizerNetworkItemStackDescriptorSerializedData using its canonical wire layout.
func (x *CerealizerNetworkItemStackDescriptorSerializedData) Marshal(io IO) {
	io.Int16(&x.Id)
	io.Uint16(&x.StackSize)
	io.Varuint32(&x.AuxValue)
	OptionalFunc(io, &x.NetIdVariant, io.Varint32)
	io.Varuint32(&x.BlockRuntimeId)
	io.Bytes(&x.UserDataBuffer)
}
