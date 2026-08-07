// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate protocol.BlockPos
}

// Marshal reads or writes LecternUpdate using its canonical wire layout.
func (x *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&x.NewPageToShow)
	io.Uint8(&x.TotalPages)
	x.PositionOfLecternToUpdate.Marshal(io)
}

// ID returns the protocol ID for LecternUpdate.
func (*LecternUpdate) ID() uint32 { return IDLecternUpdate }
