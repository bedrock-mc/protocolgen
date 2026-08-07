// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PurchaseReceipt struct {
	PurchaseReceipts []string
}

// Marshal reads or writes PurchaseReceipt using its canonical wire layout.
func (x *PurchaseReceipt) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.PurchaseReceipts, io.Varuint32, io.String)
}
