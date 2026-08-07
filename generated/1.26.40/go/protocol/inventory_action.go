// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventoryAction struct {
	Source   InventorySource
	Slot     uint32
	FromItem NetworkItemStackDescriptorSerializedData
	ToItem   NetworkItemStackDescriptorSerializedData
}

// Marshal reads or writes InventoryAction using its canonical wire layout.
func (x *InventoryAction) Marshal(io IO) {
	x.Source.Marshal(io)
	io.Varuint32(&x.Slot)
	x.FromItem.Marshal(io)
	x.ToItem.Marshal(io)
}
