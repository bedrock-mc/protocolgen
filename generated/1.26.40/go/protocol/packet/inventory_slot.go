// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type InventorySlot struct {
	ContainerID       uint8
	Slot              uint32
	FullContainerName protocol.Optional[protocol.FullContainerName]
	StorageItem       protocol.Optional[protocol.NetworkItemStackDescriptorSerializedData]
	Item              protocol.NetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventorySlot using its canonical wire layout.
func (x *InventorySlot) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Varuint32(&x.Slot)
	protocol.OptionalFunc(io, &x.FullContainerName, func(value *protocol.FullContainerName) {
		value.Marshal(io)
	})
	protocol.OptionalFunc(io, &x.StorageItem, func(value *protocol.NetworkItemStackDescriptorSerializedData) {
		value.Marshal(io)
	})
	x.Item.Marshal(io)
}

// ID returns the protocol ID for InventorySlot.
func (*InventorySlot) ID() uint32 { return IDInventorySlot }
