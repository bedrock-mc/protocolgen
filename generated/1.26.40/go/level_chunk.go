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
	io.Bool(&x.ClientRequestSubChunkLimit.set)
	if x.ClientRequestSubChunkLimit.set {
		io.Varint32(&x.ClientRequestSubChunkLimit.val)
	} else if io.Reading() {
		var zero int32
		x.ClientRequestSubChunkLimit.val = zero
	}
	io.Bool(&x.CacheEnabled)
	if !io.Reading() && uint64(len(x.CacheMetadata)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.CacheMetadata), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.CacheMetadata))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.CacheMetadata = make([]LevelChunkSubChunkMetadata, int(count1))
	}
	for index2 := range x.CacheMetadata {
		x.CacheMetadata[index2].Marshal(io)
	}
	io.Bytes(&x.SerializedChunkData)
}
