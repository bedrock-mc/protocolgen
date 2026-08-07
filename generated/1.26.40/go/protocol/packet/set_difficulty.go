// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetDifficulty struct {
	Difficulty uint32
}

// Marshal reads or writes SetDifficulty using its canonical wire layout.
func (x *SetDifficulty) Marshal(io protocol.IO) {
	io.Varuint32(&x.Difficulty)
}

// ID returns the protocol ID for SetDifficulty.
func (*SetDifficulty) ID() uint32 { return IDSetDifficulty }
