// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// ChunkRadiusUpdated is sent by the server in response to a RequestChunkRadius packet. It defines
// the chunk radius that the server allows the client to have. This may be lower than the chunk
// radius requested by the client in the RequestChunkRadius packet.
type ChunkRadiusUpdated struct {
	// ChunkRadius is the final chunk radius that the client will adapt when it receives the packet. It
	// does not have to be the same as the requested chunk radius.
	ChunkRadius int32
}

// Marshal reads or writes ChunkRadiusUpdated using its canonical wire layout.
func (x *ChunkRadiusUpdated) Marshal(io protocol.IO) {
	io.Varint32(&x.ChunkRadius)
}

// ID returns the protocol ID for ChunkRadiusUpdated.
func (*ChunkRadiusUpdated) ID() uint32 { return IDChunkRadiusUpdated }
