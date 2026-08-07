// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type BlockPickRequest struct {
	Position protocol.BlockPos
	WithData bool
	MaxSlots uint8
}

// Marshal reads or writes BlockPickRequest using its canonical wire layout.
func (x *BlockPickRequest) Marshal(io protocol.IO) {
	x.Position.Marshal(io)
	io.Bool(&x.WithData)
	io.Uint8(&x.MaxSlots)
}

// ID returns the protocol ID for BlockPickRequest.
func (*BlockPickRequest) ID() uint32 { return IDBlockPickRequest }
