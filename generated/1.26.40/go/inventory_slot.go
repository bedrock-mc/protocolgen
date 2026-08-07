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
	io.Bool(&x.FullContainerName.set)
	if x.FullContainerName.set {
		x.FullContainerName.val.Marshal(io)
	} else if io.Reading() {
		var zero FullContainerName
		x.FullContainerName.val = zero
	}
	io.Bool(&x.StorageItem.set)
	if x.StorageItem.set {
		x.StorageItem.val.Marshal(io)
	} else if io.Reading() {
		var zero CerealizerNetworkItemStackDescriptorSerializedData
		x.StorageItem.val = zero
	}
	x.Item.Marshal(io)
}
