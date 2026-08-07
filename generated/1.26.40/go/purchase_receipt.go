// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PurchaseReceipt struct {
	PurchaseReceipts []string
}

// Marshal reads or writes PurchaseReceipt using its canonical wire layout.
func (x *PurchaseReceipt) Marshal(io IO) {
	FuncSlice(io, &x.PurchaseReceipts, io.Varuint32, io.String)
}
