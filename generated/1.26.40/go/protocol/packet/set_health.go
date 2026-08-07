// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetHealth struct {
	Health int32
}

// Marshal reads or writes SetHealth using its canonical wire layout.
func (x *SetHealth) Marshal(io protocol.IO) {
	io.Varint32(&x.Health)
}

// ID returns the protocol ID for SetHealth.
func (*SetHealth) ID() uint32 { return IDSetHealth }
