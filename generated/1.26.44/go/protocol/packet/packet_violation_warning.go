// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// PacketViolationWarning is sent by the client when it receives an invalid packet from the server.
// It holds some information on the error that occurred. noinspection GoNameStartsWithPackageName
type PacketViolationWarning struct {
	ViolationType     protocol.PacketViolationType
	ViolationSeverity protocol.PacketViolationSeverity
	ViolationPacketID int32
	// ViolationContext holds a description on the violation of the packet.
	ViolationContext string
}

// Marshal reads or writes PacketViolationWarning using its canonical wire layout.
func (x *PacketViolationWarning) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.ViolationType, io.Varint32)
	protocol.IntegerFunc(&x.ViolationSeverity, io.Varint32)
	io.Varint32(&x.ViolationPacketID)
	io.String(&x.ViolationContext)
}

// ID returns the protocol ID for PacketViolationWarning.
func (*PacketViolationWarning) ID() uint32 { return IDPacketViolationWarning }
