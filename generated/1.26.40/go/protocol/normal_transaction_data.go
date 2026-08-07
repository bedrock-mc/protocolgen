// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type NormalTransactionData struct {
	Actions InventoryTransactionData
}

func (*NormalTransactionData) isInventoryTransactionValue() {}

// Marshal reads or writes NormalTransactionData using its canonical wire layout.
func (x *NormalTransactionData) Marshal(io IO) {
	x.Actions.Marshal(io)
}
