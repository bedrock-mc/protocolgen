// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type Disconnect struct {
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
