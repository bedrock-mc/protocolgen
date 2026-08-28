// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// OpenSign is sent by the server to open a sign for editing. As of 1.19.80, the player can interact
// with a sign to edit the text on both sides instead of just the front.
type OpenSign struct {
	// Pos is the position of the sign to edit. The client uses this position to get the data of the
	// sign, including the existing text and formatting etc.
	Pos protocol.BlockPos
	// IsFrontSide dictates whether the front side of the sign should be opened for editing. If false,
	// the back side is assumed to be edited.
	IsFrontSide bool
}

// Marshal reads or writes OpenSign using its canonical wire layout.
func (x *OpenSign) Marshal(io protocol.IO) {
	x.Pos.Marshal(io)
	io.Bool(&x.IsFrontSide)
}

// ID returns the protocol ID for OpenSign.
func (*OpenSign) ID() uint32 { return IDOpenSign }
