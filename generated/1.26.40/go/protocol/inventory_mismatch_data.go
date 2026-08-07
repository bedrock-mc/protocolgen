// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventoryMismatchData struct {
	Actions InventoryTransactionData
}

func (*InventoryMismatchData) isInventoryTransactionValue() {}

// Marshal reads or writes InventoryMismatchData using its canonical wire layout.
func (x *InventoryMismatchData) Marshal(io IO) {
	x.Actions.Marshal(io)
}
