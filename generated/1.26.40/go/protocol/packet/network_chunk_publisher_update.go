// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type NetworkChunkPublisherUpdate struct {
	NewPositionForView    protocol.BlockPos
	NewRadiusForView      uint32
	ServerBuiltChunksList []protocol.ChunkPos
}

// Marshal reads or writes NetworkChunkPublisherUpdate using its canonical wire layout.
func (x *NetworkChunkPublisherUpdate) Marshal(io protocol.IO) {
	x.NewPositionForView.Marshal(io)
	io.Varuint32(&x.NewRadiusForView)
	protocol.FuncSlice(io, &x.ServerBuiltChunksList, io.Uint32, func(value *protocol.ChunkPos) {
		value.Marshal(io)
	})
}
