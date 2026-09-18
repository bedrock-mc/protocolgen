// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"
)

// LevelChunk is sent by the server to provide the client with a chunk of a world data (16xYx16 blocks).
// Typically, a certain amount of chunks is sent to the client before sending it the spawn PlayStatus packet,
// so that the client spawns in a loaded world.
type LevelChunk struct {
	// Position contains the X and Z coordinates of the chunk sent. You can convert a block coordinate to a chunk
	// coordinate by right-shifting it four bits.
	ChunkPosition protocol.ChunkPos
	DimensionID   protocol.DimensionType
	// Dimension is the ID of the dimension that the chunk belongs to. This must always be set otherwise the
	// client will always assume the chunk is part of the overworld dimension.
	SubChunksCount uint32
	// SubChunkLimit is the maximum amount of sub-chunks a client will request when in request mode. A value of -1
	// means there is no limit.
	ClientRequestSubChunkLimit protocol.Optional[int32]
	// CacheEnabled specifies if the client blob cache should be enabled. This system is based on hashes of blobs
	// which are consistent and saved by the client in combination with that blob, so that the server does not
	// have the same chunk multiple times. If the client does not yet have a blob with the hash sent, it will send
	// a ClientCacheBlobStatus packet containing the hashes is does not have the data of.
	CacheEnabled bool
	// BlobHashes is a list of all blob hashes used in the chunk. It is composed of SubChunkCount + 1 hashes, with
	// the first SubChunkCount hashes being those of the sub-chunks and the last one that of the biome of the
	// chunk. If CacheEnabled is set to false, BlobHashes can be left empty.
	CacheMetadata []protocol.SubChunkMetadata
	// RawPayload is a serialised string of chunk data. The data held depends on if CacheEnabled is set to true.
	// If set to false, the payload is composed of multiple sub-chunks, each of which carry a version which
	// indicates the way they are serialised, followed by biomes, border blocks and tile entities. If CacheEnabled
	// is true, the payload consists out of the border blocks and tile entities only.
	SerializedChunkData []byte
}

// ID ...
func (*LevelChunk) ID() uint32 {
	return IDLevelChunk
}

func (pk *LevelChunk) Marshal(io protocol.IO) {
	pk.ChunkPosition.Marshal(io)
	pk.DimensionID.Marshal(io)
	io.Varuint32(&pk.SubChunksCount)
	protocol.Maximum(io, &pk.SubChunksCount, 64)
	protocol.OptionalFunc(io, &pk.ClientRequestSubChunkLimit, func(value *int32) {
		io.Varint32(value)
		protocol.Minimum(io, value, -1)
		protocol.Maximum(io, value, 64)
	})
	io.Bool(&pk.CacheEnabled)
	protocol.SliceLimits(io, &pk.CacheMetadata, 0, 65)
	io.Bytes(&pk.SerializedChunkData)
}
