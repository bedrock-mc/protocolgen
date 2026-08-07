// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type AnvilDamage struct {
	BlockPosition protocol.BlockPos
}

// Marshal reads or writes AnvilDamage using its canonical wire layout.
func (x *AnvilDamage) Marshal(io protocol.IO) {
	x.BlockPosition.Marshal(io)
}

// ID returns the protocol ID for AnvilDamage.
func (*AnvilDamage) ID() uint32 { return IDAnvilDamage }
