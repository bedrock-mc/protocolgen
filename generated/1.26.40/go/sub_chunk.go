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
	if !io.Reading() && uint64(len(x.SubChunkData)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SubChunkData), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.SubChunkData))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.SubChunkData = make([]SubChunkSubChunkPacketData, int(count1))
	}
	for index2 := range x.SubChunkData {
		x.SubChunkData[index2].Marshal(io)
	}
}
