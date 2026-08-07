// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MapInfoRequest struct {
	MapUniqueID      int64
	ClientPixelsList []protocol.PixelRequest
}

// Marshal reads or writes MapInfoRequest using its canonical wire layout.
func (x *MapInfoRequest) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.MapUniqueID)
	protocol.FuncSlice(io, &x.ClientPixelsList, io.Uint32, func(value *protocol.PixelRequest) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for MapInfoRequest.
func (*MapInfoRequest) ID() uint32 { return IDMapInfoRequest }
