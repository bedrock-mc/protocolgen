// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SubChunkRequest struct {
	DimensionType              protocol.DimensionType
	SubChunkPositionOffsetList []protocol.SubChunkSubChunkPosOffset
	CenterPos                  protocol.SubChunkPos
}

// Marshal reads or writes SubChunkRequest using its canonical wire layout.
func (x *SubChunkRequest) Marshal(io protocol.IO) {
	x.DimensionType.Marshal(io)
	protocol.FuncSlice(io, &x.SubChunkPositionOffsetList, io.Varuint32, func(value *protocol.SubChunkSubChunkPosOffset) {
		value.Marshal(io)
	})
	x.CenterPos.Marshal(io)
}
