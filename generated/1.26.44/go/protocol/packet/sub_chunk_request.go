// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// SubChunkRequest requests specific sub-chunks from the server using a center point.
type SubChunkRequest struct {
	DimensionType protocol.DimensionType
	// Offsets contains all requested offsets around the center point.
	SubChunkPositionOffsetList []protocol.SubChunkPosOffset
	// Position is an absolute sub-chunk center point used as a base point for all sub-chunks requested. The X and
	// Z coordinates represent the chunk coordinates, while the Y coordinate is the absolute sub-chunk index.
	CenterPos protocol.SubChunkPos
}

// ID ...
func (*SubChunkRequest) ID() uint32 {
	return IDSubChunkRequest
}

func (pk *SubChunkRequest) Marshal(io protocol.IO) {
	pk.DimensionType.Marshal(io)
	protocol.SliceLimits(io, &pk.SubChunkPositionOffsetList, 0, 8192)
	pk.CenterPos.Marshal(io)
}
