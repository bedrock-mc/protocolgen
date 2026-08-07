// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type UpdateSubChunkBlocks struct {
	SubChunkBlockPosition protocol.BlockPos
	BlocksChanged         protocol.UpdateSubChunkBlocksChangedInfo
}

// Marshal reads or writes UpdateSubChunkBlocks using its canonical wire layout.
func (x *UpdateSubChunkBlocks) Marshal(io protocol.IO) {
	x.SubChunkBlockPosition.Marshal(io)
	x.BlocksChanged.Marshal(io)
}
