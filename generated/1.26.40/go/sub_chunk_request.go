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
	FuncSlice(io, &x.SubChunkPositionOffsetList, io.Varuint32, func(value *SubChunkSubChunkPosOffset) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.CenterPos.Marshal(io)
}
