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
	protocol.Slice(io, &x.Slots)
	x.FullContainerName.Marshal(io)
	x.StorageItem.Marshal(io)
}

// ID returns the protocol ID for InventoryContent.
func (*InventoryContent) ID() uint32 { return IDInventoryContent }
