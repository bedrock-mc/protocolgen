// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// Disconnect may be sent by the server to disconnect the client using an optional message to send
// as the disconnect screen.
type Disconnect struct {
	// Reason is the reason for the disconnection. This affects the error code displayed on the Ore UI
	// disconnection screen and is one of the constants above.
	Reason   protocol.ConnectionDisconnectFailReason
	Messages protocol.DisconnectMessages
}

// Marshal reads or writes Disconnect using its canonical wire layout.
func (x *Disconnect) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Reason, io.Varint32)
	protocol.MarshalDisconnectMessages(io, &x.Messages)
}

// ID returns the protocol ID for Disconnect.
func (*Disconnect) ID() uint32 { return IDDisconnect }
