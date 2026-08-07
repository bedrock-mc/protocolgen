// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventoryOptions struct {
	LeftInventoryTab  InventoryLeftTabIndex
	RightInventoryTab InventoryRightTabIndex
	Filtering         bool
	LayoutInv         InventoryLayout
	LayoutCraft       InventoryLayout
}

// Marshal reads or writes InventoryOptions using its canonical wire layout.
func (x *InventoryOptions) Marshal(io IO) {
	IntegerFunc(&x.LeftInventoryTab, io.Varint32)
	IntegerFunc(&x.RightInventoryTab, io.Varint32)
	io.Bool(&x.Filtering)
	IntegerFunc(&x.LayoutInv, io.Varint32)
	IntegerFunc(&x.LayoutCraft, io.Varint32)
}
