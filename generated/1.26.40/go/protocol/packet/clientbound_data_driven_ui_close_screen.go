// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ClientboundDataDrivenUICloseScreen struct {
	FormID protocol.Optional[uint32]
}

// ID ...
func (*ClientboundDataDrivenUICloseScreen) ID() uint32 {
	return IDClientboundDataDrivenUICloseScreen
}

func (pk *ClientboundDataDrivenUICloseScreen) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &pk.FormID, io.Uint32)
}
