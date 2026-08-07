// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CompletedUsingItem struct {
	ItemId        int16
	ItemUseMethod int32
}

// Marshal reads or writes CompletedUsingItem using its canonical wire layout.
func (x *CompletedUsingItem) Marshal(io protocol.IO) {
	io.Int16(&x.ItemId)
	io.Int32(&x.ItemUseMethod)
}
