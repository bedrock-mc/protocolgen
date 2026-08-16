// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type ClientboundDataDrivenUIShowScreen struct {
	ScreenID       string
	FormID         uint32
	DataInstanceID protocol.Optional[uint32]
}

// Marshal reads or writes ClientboundDataDrivenUIShowScreen using its canonical wire layout.
func (x *ClientboundDataDrivenUIShowScreen) Marshal(io protocol.IO) {
	io.StringLimits(&x.ScreenID, 0, 500)
	io.Uint32(&x.FormID)
	protocol.Minimum(io, &x.FormID, 0)
	protocol.OptionalFunc(io, &x.DataInstanceID, io.Uint32)
}

// ID returns the protocol ID for ClientboundDataDrivenUIShowScreen.
func (*ClientboundDataDrivenUIShowScreen) ID() uint32 { return IDClientboundDataDrivenUIShowScreen }
