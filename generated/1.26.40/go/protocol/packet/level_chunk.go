// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type LevelChunk struct {
	ChunkPosition              protocol.ChunkPos
	DimensionID                protocol.DimensionType
	SubChunksCount             uint32
	ClientRequestSubChunkLimit protocol.Optional[int32]
	CacheEnabled               bool
	CacheMetadata              []protocol.SubChunkMetadata
	SerializedChunkData        []byte
}

// Marshal reads or writes LevelChunk using its canonical wire layout.
func (x *LevelChunk) Marshal(io protocol.IO) {
	x.ChunkPosition.Marshal(io)
	x.DimensionID.Marshal(io)
	io.Varuint32(&x.SubChunksCount)
	protocol.OptionalFunc(io, &x.ClientRequestSubChunkLimit, io.Varint32)
	io.Bool(&x.CacheEnabled)
	protocol.Slice(io, &x.CacheMetadata)
	io.Bytes(&x.SerializedChunkData)
}

// ID returns the protocol ID for LevelChunk.
func (*LevelChunk) ID() uint32 { return IDLevelChunk }
