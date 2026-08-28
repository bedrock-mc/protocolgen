// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// SetLastHurtBy is sent by the server to let the client know what entity type it was last hurt by.
// At this moment, the packet is useless and should not be used. There is no behaviour that depends
// on if this packet is sent or not.
type SetLastHurtBy struct {
	LastHurtBy protocol.ActorType
}

// Marshal reads or writes SetLastHurtBy using its canonical wire layout.
func (x *SetLastHurtBy) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.LastHurtBy, io.Varint32)
}

// ID returns the protocol ID for SetLastHurtBy.
func (*SetLastHurtBy) ID() uint32 { return IDSetLastHurtBy }
