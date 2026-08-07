// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type InventoryContent struct {
	ContainerId       uint32
	Slots             []protocol.CerealizerNetworkItemStackDescriptorSerializedData
	FullContainerName protocol.FullContainerName
	StorageItem       protocol.CerealizerNetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventoryContent using its canonical wire layout.
func (x *InventoryContent) Marshal(io protocol.IO) {
	io.Varuint32(&x.ContainerId)
	protocol.FuncSlice(io, &x.Slots, io.Varuint32, func(value *protocol.CerealizerNetworkItemStackDescriptorSerializedData) {
		value.Marshal(io)
	})
	x.FullContainerName.Marshal(io)
	x.StorageItem.Marshal(io)
}
