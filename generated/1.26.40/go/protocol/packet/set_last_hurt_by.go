// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetLastHurtBy struct {
	LastHurtBy protocol.ActorType
}

// Marshal reads or writes SetLastHurtBy using its canonical wire layout.
func (x *SetLastHurtBy) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.LastHurtBy, io.Varint32)
}

// ID returns the protocol ID for SetLastHurtBy.
func (*SetLastHurtBy) ID() uint32 { return IDSetLastHurtBy }
