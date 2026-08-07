// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetPlayerInventoryOptions struct {
	InventoryOptions InventoryOptions
}

// Marshal reads or writes SetPlayerInventoryOptions using its canonical wire layout.
func (x *SetPlayerInventoryOptions) Marshal(io IO) {
	x.InventoryOptions.Marshal(io)
}
