// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SubChunk struct {
	CacheEnabled  bool
	DimensionType protocol.DimensionType
	CenterPos     protocol.SubChunkPos
	SubChunkData  []protocol.SubChunkData
}

// Marshal reads or writes SubChunk using its canonical wire layout.
func (x *SubChunk) Marshal(io protocol.IO) {
	io.Bool(&x.CacheEnabled)
	x.DimensionType.Marshal(io)
	x.CenterPos.Marshal(io)
	protocol.Slice(io, &x.SubChunkData)
}

// ID returns the protocol ID for SubChunk.
func (*SubChunk) ID() uint32 { return IDSubChunk }
