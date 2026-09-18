// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// LecternUpdate is sent by the client to update the server on which page was opened in a book on a lectern,
// or if the book should be removed from it.
type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate protocol.BlockPos
}

// ID ...
func (*LecternUpdate) ID() uint32 {
	return IDLecternUpdate
}

func (pk *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&pk.NewPageToShow)
	io.Uint8(&pk.TotalPages)
	pk.PositionOfLecternToUpdate.Marshal(io)
}
