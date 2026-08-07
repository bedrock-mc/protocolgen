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
	if !io.Reading() && uint64(len(x.Slots)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.Slots), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.Slots))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.Slots = make([]CerealizerNetworkItemStackDescriptorSerializedData, int(count1))
	}
	for index2 := range x.Slots {
		x.Slots[index2].Marshal(io)
	}
	x.FullContainerName.Marshal(io)
	x.StorageItem.Marshal(io)
}
