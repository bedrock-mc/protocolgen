// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// LecternUpdate is sent by the client to update the server on which page was opened in a book on a lectern,
// or if the book should be removed from it.
type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate protocol.BlockPos
}

// ID returns the protocol ID for LecternUpdate.
func (*LecternUpdate) ID() uint32 { return IDLecternUpdate }

// Marshal reads or writes LecternUpdate using its canonical wire layout.
func (pk *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&pk.NewPageToShow)
	io.Uint8(&pk.TotalPages)
	pk.PositionOfLecternToUpdate.Marshal(io)
}
