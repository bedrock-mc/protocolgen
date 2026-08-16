// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// InventorySlot is sent by the server to update a single slot in one of the inventory windows that
// the client currently has opened. Usually this is the main inventory, but it may also be the off
// hand or, for example, a chest inventory.
type InventorySlot struct {
	// ContainerID is the ID of the window that the packet modifies. It must point to one of the windows
	// that the client currently has opened.
	ContainerID uint8
	// Slot is the index of the slot that the packet modifies. The new item will be set to the slot at
	// this index.
	Slot uint32
	// FullContainerName is the protocol.FullContainerName that describes the container that the content
	// is for.
	FullContainerName protocol.Optional[protocol.FullContainerName]
	// StorageItem is the item that is acting as the storage container for the inventory. If the
	// inventory is not a dynamic container then this field should be left empty. When set, only the
	// item type is used by the client and none of the other stack info.
	StorageItem protocol.Optional[protocol.NetworkItemStackDescriptorSerializedData]
	// Item is the item to be put in the slot at Slot. It will overwrite any item that may currently be
	// present in that slot.
	Item protocol.NetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventorySlot using its canonical wire layout.
func (x *InventorySlot) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	protocol.Minimum(io, &x.ContainerID, 0)
	protocol.Maximum(io, &x.ContainerID, 255)
	io.Varuint32(&x.Slot)
	protocol.Minimum(io, &x.Slot, 0)
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
