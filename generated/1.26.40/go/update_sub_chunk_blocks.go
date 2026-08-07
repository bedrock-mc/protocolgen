// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateSubChunkBlocks struct {
	SubChunkBlockPosition BlockPos
	BlocksChanged         UpdateSubChunkBlocksChangedInfo
}

// Marshal reads or writes UpdateSubChunkBlocks using its canonical wire layout.
func (x *UpdateSubChunkBlocks) Marshal(io IO) {
	x.SubChunkBlockPosition.Marshal(io)
	x.BlocksChanged.Marshal(io)
}
