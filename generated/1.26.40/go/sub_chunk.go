// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SubChunk struct {
	CacheEnabled  bool
	DimensionType DimensionType
	CenterPos     SubChunkPos
	SubChunkData  []SubChunkSubChunkPacketData
}

// Marshal reads or writes SubChunk using its canonical wire layout.
func (x *SubChunk) Marshal(io IO) {
	io.Bool(&x.CacheEnabled)
	x.DimensionType.Marshal(io)
	x.CenterPos.Marshal(io)
	FuncSlice(io, &x.SubChunkData, io.Varuint32, func(value *SubChunkSubChunkPacketData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
