// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type ClientboundDebugRenderer struct {
	Type            string
	DebugMarkerData protocol.Optional[protocol.DebugMarkerData]
}

// ID returns the protocol ID for ClientboundDebugRenderer.
func (*ClientboundDebugRenderer) ID() uint32 { return IDClientboundDebugRenderer }

// Marshal reads or writes ClientboundDebugRenderer using its canonical wire layout.
func (pk *ClientboundDebugRenderer) Marshal(io protocol.IO) {
	io.String(&pk.Type)
	protocol.OptionalMarshaler(io, &pk.DebugMarkerData)
}
