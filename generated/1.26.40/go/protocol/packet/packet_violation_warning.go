// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// PacketViolationWarning is sent by the client when it receives an invalid packet from the server. It holds
// some information on the error that occurred. noinspection GoNameStartsWithPackageName
type PacketViolationWarning struct {
	ViolationType     protocol.PacketViolationType
	ViolationSeverity protocol.PacketViolationSeverity
	ViolationPacketID int32
	// ViolationContext holds a description on the violation of the packet.
	ViolationContext string
}

// ID ...
func (*PacketViolationWarning) ID() uint32 {
	return IDPacketViolationWarning
}

func (pk *PacketViolationWarning) Marshal(io protocol.IO) {
	pk.ViolationType.Marshal(io)
	pk.ViolationSeverity.Marshal(io)
	io.Varint32(&pk.ViolationPacketID)
	io.String(&pk.ViolationContext)
}
