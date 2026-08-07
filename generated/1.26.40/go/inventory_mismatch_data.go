// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type InventoryMismatchData struct {
	Actions InventoryTransactionData
}

func (InventoryMismatchData) isInventoryTransactionTransactionValue() {}

// Marshal reads or writes InventoryMismatchData using its canonical wire layout.
func (x *InventoryMismatchData) Marshal(io IO) {
	x.Actions.Marshal(io)
}
