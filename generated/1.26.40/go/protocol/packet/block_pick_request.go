// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// BlockPickRequest is sent by the client when it requests to pick a block in the world and place
// its item in their inventory.
type BlockPickRequest struct {
	// Position is the position at which the client requested to pick the block. The block at that
	// position should have its item put in HotBarSlot if it is empty.
	Position protocol.BlockPos
	WithData bool
	MaxSlots uint8
}

// Marshal reads or writes BlockPickRequest using its canonical wire layout.
func (x *BlockPickRequest) Marshal(io protocol.IO) {
	x.Position.Marshal(io)
	io.Bool(&x.WithData)
	io.Uint8(&x.MaxSlots)
	protocol.Minimum(io, &x.MaxSlots, 0)
	protocol.Maximum(io, &x.MaxSlots, 255)
}

// ID returns the protocol ID for BlockPickRequest.
func (*BlockPickRequest) ID() uint32 { return IDBlockPickRequest }
