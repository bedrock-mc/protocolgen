// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// UpdateBlock is sent by the server to update a block client-side, without resending the entire
// chunk that the block is located in. It is particularly useful for small modifications like block
// breaking/placing.
type UpdateBlock struct {
	// BlockPosition is the block position at which a block is updated.
	BlockPosition protocol.BlockPos
	// BlockRuntimeID is the runtime ID of the block that is placed at Position after sending the packet
	// to the client.
	BlockRuntimeID uint32
	// Flags is a combination of flags that specify the way the block is updated client-side. It is a
	// combination of the flags above, but typically sending only the BlockUpdateNetwork flag is
	// sufficient.
	Flags uint32
	// Layer is the world layer on which the block is updated. For most blocks, this is the first layer,
	// as that layer is the default layer to place blocks on, but for blocks inside of each other, this
	// differs.
	Layer uint32
}

// Marshal reads or writes UpdateBlock using its canonical wire layout.
func (x *UpdateBlock) Marshal(io protocol.IO) {
	x.BlockPosition.Marshal(io)
	io.Varuint32(&x.BlockRuntimeID)
	protocol.Minimum(io, &x.BlockRuntimeID, 0)
	io.Varuint32(&x.Flags)
	protocol.Minimum(io, &x.Flags, 0)
	io.Varuint32(&x.Layer)
	protocol.Minimum(io, &x.Layer, 0)
}

// ID returns the protocol ID for UpdateBlock.
func (*UpdateBlock) ID() uint32 { return IDUpdateBlock }
