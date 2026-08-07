// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type GuiDataPickItem struct {
	ItemName       string
	ItemEffectName string
	Slot           int32
}

// Marshal reads or writes GuiDataPickItem using its canonical wire layout.
func (x *GuiDataPickItem) Marshal(io protocol.IO) {
	io.String(&x.ItemName)
	io.String(&x.ItemEffectName)
	io.Int32(&x.Slot)
}
