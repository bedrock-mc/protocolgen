// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type LevelChunk struct {
	ChunkPosition              ChunkPos
	DimensionId                DimensionType
	SubChunksCount             uint32
	ClientRequestSubChunkLimit *int32
	CacheEnabled               bool
	CacheMetadata              []LevelChunkSubChunkMetadata
	SerializedChunkData        []byte
}
