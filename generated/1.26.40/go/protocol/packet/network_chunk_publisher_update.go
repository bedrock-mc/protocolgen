// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// NetworkChunkPublisherUpdate is sent by the server to change the point around which chunks are and
// remain loaded. This is useful for mini-game servers, where only one area is ever loaded, in which
// case the NetworkChunkPublisherUpdate packet can be sent in the middle of it, so that no chunks
// ever need to be additionally sent during the course of the game. In reality, the packet is not
// extraordinarily useful, and most servers just send it constantly at the position of the player.
// If the packet is not sent at all, no chunks will be shown to the player, regardless of where they
// are sent.
type NetworkChunkPublisherUpdate struct {
	// NewPositionForView is the block position around which chunks loaded will remain shown to the
	// client. Most servers set this position to the position of the player itself.
	NewPositionForView protocol.BlockPos
	// NewRadiusForView is the radius in blocks around Position that chunks sent show up in and will
	// remain loaded in. Unlike the RequestChunkRadius and ChunkRadiusUpdated packets, this radius is in
	// blocks rather than chunks, so the chunk radius needs to be multiplied by 16. (Or shifted to the
	// left by 4.)
	NewRadiusForView uint32
	// ServerBuiltChunksList ... TODO: Figure out what this field is used for.
	ServerBuiltChunksList []protocol.ChunkPos
}

// Marshal reads or writes NetworkChunkPublisherUpdate using its canonical wire layout.
func (x *NetworkChunkPublisherUpdate) Marshal(io protocol.IO) {
	x.NewPositionForView.Marshal(io)
	io.Varuint32(&x.NewRadiusForView)
	protocol.Minimum(io, &x.NewRadiusForView, 0)
	protocol.FuncSliceLimits(io, &x.ServerBuiltChunksList, io.Uint32, 0, 9216, func(value *protocol.ChunkPos) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for NetworkChunkPublisherUpdate.
func (*NetworkChunkPublisherUpdate) ID() uint32 { return IDNetworkChunkPublisherUpdate }
