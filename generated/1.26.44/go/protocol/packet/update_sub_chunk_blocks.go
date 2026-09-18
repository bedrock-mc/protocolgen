// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// UpdateSubChunkBlocks is essentially just UpdateBlock packet, however for a set of blocks in a sub-chunk.
type UpdateSubChunkBlocks struct {
	SubChunkBlockPosition protocol.BlockPos
	BlocksChanged         protocol.UpdateSubChunkBlocksChangedInfo
}

// ID returns the protocol ID for UpdateSubChunkBlocks.
func (*UpdateSubChunkBlocks) ID() uint32 { return IDUpdateSubChunkBlocks }

// Marshal reads or writes UpdateSubChunkBlocks using its canonical wire layout.
func (pk *UpdateSubChunkBlocks) Marshal(io protocol.IO) {
	pk.SubChunkBlockPosition.Marshal(io)
	pk.BlocksChanged.Marshal(io)
}
