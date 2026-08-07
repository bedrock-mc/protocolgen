// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundDataDrivenUIShowScreen struct {
	ScreenId       string
	FormId         uint32
	DataInstanceId protocol.Optional[uint32]
}

// Marshal reads or writes ClientboundDataDrivenUIShowScreen using its canonical wire layout.
func (x *ClientboundDataDrivenUIShowScreen) Marshal(io protocol.IO) {
	io.String(&x.ScreenId)
	io.Uint32(&x.FormId)
	protocol.OptionalFunc(io, &x.DataInstanceId, io.Uint32)
}

// ID returns the protocol ID for ClientboundDataDrivenUIShowScreen.
func (*ClientboundDataDrivenUIShowScreen) ID() uint32 { return IDClientboundDataDrivenUIShowScreen }
