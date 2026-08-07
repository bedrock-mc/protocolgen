// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientboundDebugRenderer struct {
	Type            string
	DebugMarkerData protocol.Optional[protocol.ClientboundDebugRendererDebugMarkerData]
}

// Marshal reads or writes ClientboundDebugRenderer using its canonical wire layout.
func (x *ClientboundDebugRenderer) Marshal(io protocol.IO) {
	io.String(&x.Type)
	protocol.OptionalFunc(io, &x.DebugMarkerData, func(value *protocol.ClientboundDebugRendererDebugMarkerData) {
		value.Marshal(io)
	})
}
