// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SubChunkRequest struct {
	DimensionType              DimensionType
	SubChunkPositionOffsetList []SubChunkSubChunkPosOffset
	CenterPos                  SubChunkPos
}

// Marshal reads or writes SubChunkRequest using its canonical wire layout.
func (x *SubChunkRequest) Marshal(io IO) {
	x.DimensionType.Marshal(io)
	if !io.Reading() && uint64(len(x.SubChunkPositionOffsetList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.SubChunkPositionOffsetList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.SubChunkPositionOffsetList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.SubChunkPositionOffsetList = make([]SubChunkSubChunkPosOffset, int(count1))
	}
	for index2 := range x.SubChunkPositionOffsetList {
		x.SubChunkPositionOffsetList[index2].Marshal(io)
	}
	x.CenterPos.Marshal(io)
}
