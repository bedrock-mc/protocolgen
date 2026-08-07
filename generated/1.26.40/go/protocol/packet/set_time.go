// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SetTime struct {
	Time int32
}

// Marshal reads or writes SetTime using its canonical wire layout.
func (x *SetTime) Marshal(io protocol.IO) {
	io.Varint32(&x.Time)
}

// ID returns the protocol ID for SetTime.
func (*SetTime) ID() uint32 { return IDSetTime }
