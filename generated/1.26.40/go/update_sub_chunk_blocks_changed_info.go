// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateSubChunkBlocksChangedInfo struct {
	BlocksChangedStandards []UpdateSubChunkNetworkBlockInfo
	BlocksChangedExtras    []UpdateSubChunkNetworkBlockInfo
}

// Marshal reads or writes UpdateSubChunkBlocksChangedInfo using its canonical wire layout.
func (x *UpdateSubChunkBlocksChangedInfo) Marshal(io IO) {
	FuncSlice(io, &x.BlocksChangedStandards, io.Varuint32, func(value *UpdateSubChunkNetworkBlockInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	FuncSlice(io, &x.BlocksChangedExtras, io.Varuint32, func(value *UpdateSubChunkNetworkBlockInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
