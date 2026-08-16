// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// LecternUpdate is sent by the client to update the server on which page was opened in a book on a
// lectern, or if the book should be removed from it.
type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate protocol.BlockPos
}

// Marshal reads or writes LecternUpdate using its canonical wire layout.
func (x *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&x.NewPageToShow)
	protocol.Minimum(io, &x.NewPageToShow, 0)
	protocol.Maximum(io, &x.NewPageToShow, 255)
	io.Uint8(&x.TotalPages)
	protocol.Minimum(io, &x.TotalPages, 0)
	protocol.Maximum(io, &x.TotalPages, 255)
	x.PositionOfLecternToUpdate.Marshal(io)
}

// ID returns the protocol ID for LecternUpdate.
func (*LecternUpdate) ID() uint32 { return IDLecternUpdate }
