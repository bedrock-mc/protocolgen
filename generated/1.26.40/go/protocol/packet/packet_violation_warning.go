// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PacketViolationWarning struct {
	ViolationType     protocol.PacketViolationType
	ViolationSeverity protocol.PacketViolationSeverity
	ViolationPacketId int32
	ViolationContext  string
}

// Marshal reads or writes PacketViolationWarning using its canonical wire layout.
func (x *PacketViolationWarning) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.ViolationType, io.Varint32)
	protocol.IntegerFunc(&x.ViolationSeverity, io.Varint32)
	io.Varint32(&x.ViolationPacketId)
	io.String(&x.ViolationContext)
}
