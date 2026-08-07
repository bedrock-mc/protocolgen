// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type InventoryContent struct {
	ContainerId       uint32
	Slots             []CerealizerNetworkItemStackDescriptorSerializedData
	FullContainerName FullContainerName
	StorageItem       CerealizerNetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventoryContent using its canonical wire layout.
func (x *InventoryContent) Marshal(io IO) {
	io.Varuint32(&x.ContainerId)
	FuncSlice(io, &x.Slots, io.Varuint32, func(value *CerealizerNetworkItemStackDescriptorSerializedData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.FullContainerName.Marshal(io)
	x.StorageItem.Marshal(io)
}
