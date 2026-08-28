// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

type ClientboundDataDrivenUICloseScreen struct {
	FormID protocol.Optional[uint32]
}

// Marshal reads or writes ClientboundDataDrivenUICloseScreen using its canonical wire layout.
func (x *ClientboundDataDrivenUICloseScreen) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &x.FormID, io.Uint32)
}

// ID returns the protocol ID for ClientboundDataDrivenUICloseScreen.
func (*ClientboundDataDrivenUICloseScreen) ID() uint32 { return IDClientboundDataDrivenUICloseScreen }
