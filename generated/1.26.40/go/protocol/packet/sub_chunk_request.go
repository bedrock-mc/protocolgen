// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SubChunkRequest struct {
	DimensionType              protocol.DimensionType
	SubChunkPositionOffsetList []protocol.SubChunkPosOffset
	CenterPos                  protocol.SubChunkPos
}

// Marshal reads or writes SubChunkRequest using its canonical wire layout.
func (x *SubChunkRequest) Marshal(io protocol.IO) {
	x.DimensionType.Marshal(io)
	protocol.Slice(io, &x.SubChunkPositionOffsetList)
	x.CenterPos.Marshal(io)
}

// ID returns the protocol ID for SubChunkRequest.
func (*SubChunkRequest) ID() uint32 { return IDSubChunkRequest }
