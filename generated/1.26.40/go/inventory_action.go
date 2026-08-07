// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type InventoryAction struct {
	Source   InventorySource
	Slot     uint32
	FromItem CerealizerNetworkItemStackDescriptorSerializedData
	ToItem   CerealizerNetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventoryAction using its canonical wire layout.
func (x *InventoryAction) Marshal(io IO) {
	x.Source.Marshal(io)
	io.Varuint32(&x.Slot)
	x.FromItem.Marshal(io)
	x.ToItem.Marshal(io)
}
