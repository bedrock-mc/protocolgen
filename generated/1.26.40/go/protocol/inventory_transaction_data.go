// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type InventoryTransactionData struct {
	Actions Optional[[]InventoryAction]
}

// Marshal reads or writes InventoryTransactionData using its canonical wire layout.
func (x *InventoryTransactionData) Marshal(io IO) {
	OptionalFunc(io, &x.Actions, func(value *[]InventoryAction) {
		Slice(io, value)
	})
}
