// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type ClientboundDebugRenderer struct {
	Type            string
	DebugMarkerData protocol.Optional[protocol.DebugMarkerData]
}

// ID ...
func (*ClientboundDebugRenderer) ID() uint32 {
	return IDClientboundDebugRenderer
}

func (pk *ClientboundDebugRenderer) Marshal(io protocol.IO) {
	io.String(&pk.Type)
	protocol.OptionalMarshaler(io, &pk.DebugMarkerData)
}
