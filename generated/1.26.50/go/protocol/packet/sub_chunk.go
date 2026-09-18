// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// SubChunk sends data about multiple sub-chunks around a center point.
type SubChunk struct {
	// CacheEnabled is whether the sub-chunk caching is enabled or not.
	CacheEnabled  bool
	DimensionType protocol.DimensionType
	CenterPos     protocol.SubChunkPos
	SubChunkData  []protocol.SubChunkData
}

// ID returns the protocol ID for SubChunk.
func (*SubChunk) ID() uint32 { return IDSubChunk }

// Marshal reads or writes SubChunk using its canonical wire layout.
func (pk *SubChunk) Marshal(io protocol.IO) {
	io.Bool(&pk.CacheEnabled)
	pk.DimensionType.Marshal(io)
	pk.CenterPos.Marshal(io)
	protocol.SliceLimits(io, &pk.SubChunkData, 0, 8192)
}
