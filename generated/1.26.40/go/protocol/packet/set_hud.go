// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetHud struct {
	HudElement []protocol.HudElement
	HudVisible protocol.HudVisibility
}

// Marshal reads or writes SetHud using its canonical wire layout.
func (x *SetHud) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.HudElement, io.Varuint32, func(value *protocol.HudElement) {
		protocol.IntegerFunc(value, io.Varint32)
	})
	protocol.IntegerFunc(&x.HudVisible, io.Varint32)
}

// ID returns the protocol ID for SetHud.
func (*SetHud) ID() uint32 { return IDSetHud }
