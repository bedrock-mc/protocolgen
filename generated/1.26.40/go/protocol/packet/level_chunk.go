// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// LevelChunk is sent by the server to provide the client with a chunk of a world data (16xYx16
// blocks). Typically, a certain amount of chunks is sent to the client before sending it the spawn
// PlayStatus packet, so that the client spawns in a loaded world.
type LevelChunk struct {
	ChunkPosition              protocol.ChunkPos
	DimensionID                protocol.DimensionType
	SubChunksCount             uint32
	ClientRequestSubChunkLimit protocol.Optional[int32]
	// CacheEnabled specifies if the client blob cache should be enabled. This system is based on hashes
	// of blobs which are consistent and saved by the client in combination with that blob, so that the
	// server does not have the same chunk multiple times. If the client does not yet have a blob with
	// the hash sent, it will send a ClientCacheBlobStatus packet containing the hashes is does not have
	// the data of.
	CacheEnabled        bool
	CacheMetadata       []protocol.SubChunkMetadata
	SerializedChunkData []byte
}

// Marshal reads or writes LevelChunk using its canonical wire layout.
func (x *LevelChunk) Marshal(io protocol.IO) {
	x.ChunkPosition.Marshal(io)
	x.DimensionID.Marshal(io)
	io.Varuint32(&x.SubChunksCount)
	protocol.Minimum(io, &x.SubChunksCount, 0)
	protocol.Maximum(io, &x.SubChunksCount, 64)
	protocol.OptionalFunc(io, &x.ClientRequestSubChunkLimit, func(value *int32) {
		io.Varint32(value)
		protocol.Minimum(io, value, -1)
		protocol.Maximum(io, value, 64)
	})
	io.Bool(&x.CacheEnabled)
	protocol.SliceLimits(io, &x.CacheMetadata, 0, 65)
	io.Bytes(&x.SerializedChunkData)
}

// ID returns the protocol ID for LevelChunk.
func (*LevelChunk) ID() uint32 { return IDLevelChunk }
