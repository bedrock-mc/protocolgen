// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkSubChunkPacketData struct {
	SubChunkPosOffset     SubChunkSubChunkPosOffset
	SubChunkRequestResult SubChunkSubChunkRequestResult
	SerializedSubChunk    Optional[string]
	HeightMapData         SubChunkHeightmapData
	BlobId                Optional[uint64]
}

// Marshal reads or writes SubChunkSubChunkPacketData using its canonical wire layout.
func (x *SubChunkSubChunkPacketData) Marshal(io IO) {
	x.SubChunkPosOffset.Marshal(io)
	IntegerFunc(&x.SubChunkRequestResult, io.Uint8)
	OptionalFunc(io, &x.SerializedSubChunk, io.String)
	x.HeightMapData.Marshal(io)
	OptionalFunc(io, &x.BlobId, io.Uint64)
}
