// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// ClientboundDebugRenderer is sent by the server to spawn an outlined cube on client-side.
type ClientboundDebugRenderer struct {
	// Type is the type of action. It is one of the constants above.
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
