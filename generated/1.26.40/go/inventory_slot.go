// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type InventorySlot struct {
	ContainerId       uint8
	Slot              uint32
	FullContainerName Optional[FullContainerName]
	StorageItem       Optional[CerealizerNetworkItemStackDescriptorSerializedData]
	Item              CerealizerNetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventorySlot using its canonical wire layout.
func (x *InventorySlot) Marshal(io IO) {
	io.Uint8(&x.ContainerId)
	io.Varuint32(&x.Slot)
	OptionalFunc(io, &x.FullContainerName, func(value *FullContainerName) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.StorageItem, func(value *CerealizerNetworkItemStackDescriptorSerializedData) {
		value.Marshal(io)
	})
	x.Item.Marshal(io)
}
