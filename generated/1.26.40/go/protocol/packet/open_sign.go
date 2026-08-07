// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type OpenSign struct {
	Pos         protocol.BlockPos
	IsFrontSide bool
}

// Marshal reads or writes OpenSign using its canonical wire layout.
func (x *OpenSign) Marshal(io protocol.IO) {
	x.Pos.Marshal(io)
	io.Bool(&x.IsFrontSide)
}
