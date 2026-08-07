// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LevelChunk struct {
	ChunkPosition              ChunkPos
	DimensionId                DimensionType
	SubChunksCount             uint32
	ClientRequestSubChunkLimit Optional[int32]
	CacheEnabled               bool
	CacheMetadata              []LevelChunkSubChunkMetadata
	SerializedChunkData        []byte
}

// Marshal reads or writes LevelChunk using its canonical wire layout.
func (x *LevelChunk) Marshal(io IO) {
	x.ChunkPosition.Marshal(io)
	x.DimensionId.Marshal(io)
	io.Varuint32(&x.SubChunksCount)
	OptionalFunc(io, &x.ClientRequestSubChunkLimit, io.Varint32)
	io.Bool(&x.CacheEnabled)
	FuncSlice(io, &x.CacheMetadata, io.Varuint32, func(value *LevelChunkSubChunkMetadata) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.Bytes(&x.SerializedChunkData)
}
