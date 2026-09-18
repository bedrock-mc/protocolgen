// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type GuiDataPickItem struct {
	ItemName       string
	ItemEffectName string
	Slot           int32
}

// ID returns the protocol ID for GuiDataPickItem.
func (*GuiDataPickItem) ID() uint32 { return IDGuiDataPickItem }

// Marshal reads or writes GuiDataPickItem using its canonical wire layout.
func (pk *GuiDataPickItem) Marshal(io protocol.IO) {
	io.String(&pk.ItemName)
	io.String(&pk.ItemEffectName)
	io.Int32(&pk.Slot)
}
