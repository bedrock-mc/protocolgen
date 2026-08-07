// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type InventorySlot struct {
	ContainerId       uint8
	Slot              uint32
	FullContainerName Optional[FullContainerName]
	StorageItem       Optional[CerealizerNetworkItemStackDescriptorSerializedData]
	Item              CerealizerNetworkItemStackDescriptorSerializedData
}
