// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkData struct {
	SubChunkPosOffset     SubChunkPosOffset
	SubChunkRequestResult SubChunkRequestResult
	SerializedSubChunk    Optional[string]
	HeightMapData         HeightmapData
	BlobID                Optional[uint64]
}

// Marshal reads or writes SubChunkData using its canonical wire layout.
func (x *SubChunkData) Marshal(io IO) {
	x.SubChunkPosOffset.Marshal(io)
	IntegerFunc(&x.SubChunkRequestResult, io.Uint8)
	OptionalFunc(io, &x.SerializedSubChunk, io.String)
	x.HeightMapData.Marshal(io)
	OptionalFunc(io, &x.BlobID, io.Uint64)
}
