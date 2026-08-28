// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// SubChunk sends data about multiple sub-chunks around a center point.
type SubChunk struct {
	// CacheEnabled is whether the sub-chunk caching is enabled or not.
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
	protocol.SliceLimits(io, &x.SubChunkData, 0, 8192)
}

// ID returns the protocol ID for SubChunk.
func (*SubChunk) ID() uint32 { return IDSubChunk }
