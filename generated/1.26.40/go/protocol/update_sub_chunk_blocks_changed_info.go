// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type UpdateSubChunkBlocksChangedInfo struct {
	BlocksChangedStandards []UpdateSubChunkNetworkBlockInfo
	BlocksChangedExtras    []UpdateSubChunkNetworkBlockInfo
}

// Marshal reads or writes UpdateSubChunkBlocksChangedInfo using its canonical wire layout.
func (x *UpdateSubChunkBlocksChangedInfo) Marshal(io IO) {
	Slice(io, &x.BlocksChangedStandards)
	Slice(io, &x.BlocksChangedExtras)
}
