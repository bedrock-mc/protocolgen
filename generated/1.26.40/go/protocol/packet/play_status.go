// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PlayStatus struct {
	Status protocol.PlayStatusType
}

// Marshal reads or writes PlayStatus using its canonical wire layout.
func (x *PlayStatus) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Status, io.BEInt32)
}

// ID returns the protocol ID for PlayStatus.
func (*PlayStatus) ID() uint32 { return IDPlayStatus }
