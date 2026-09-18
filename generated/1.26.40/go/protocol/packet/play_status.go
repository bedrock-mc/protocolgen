// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// PlayStatus is sent by the server to update a player on the play status. This includes failed statuses due
// to a mismatched version, but also success statuses.
type PlayStatus struct {
	// Status is the status of the packet. It is one of the constants found above.
	Status protocol.PlayStatusType
}

// ID returns the protocol ID for PlayStatus.
func (*PlayStatus) ID() uint32 { return IDPlayStatus }

// Marshal reads or writes PlayStatus using its canonical wire layout.
func (pk *PlayStatus) Marshal(io protocol.IO) {
	pk.Status.Marshal(io)
}
