// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ClientboundDataDrivenUIShowScreen struct {
	ScreenID       string
	FormID         uint32
	DataInstanceID protocol.Optional[uint32]
}

// ID ...
func (*ClientboundDataDrivenUIShowScreen) ID() uint32 {
	return IDClientboundDataDrivenUIShowScreen
}

func (pk *ClientboundDataDrivenUIShowScreen) Marshal(io protocol.IO) {
	io.StringLimits(&pk.ScreenID, 0, 500)
	io.Uint32(&pk.FormID)
	protocol.OptionalFunc(io, &pk.DataInstanceID, io.Uint32)
}
