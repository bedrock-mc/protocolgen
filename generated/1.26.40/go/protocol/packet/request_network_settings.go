// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// RequestNetworkSettings is sent by the client to request network settings, such as compression,
// from the server.
type RequestNetworkSettings struct {
	// ClientNetworkVersion is the protocol version of the player. The player is disconnected if the
	// protocol is incompatible with the protocol of the server.
	ClientNetworkVersion int32
}

// Marshal reads or writes RequestNetworkSettings using its canonical wire layout.
func (x *RequestNetworkSettings) Marshal(io protocol.IO) {
	io.BEInt32(&x.ClientNetworkVersion)
}

// ID returns the protocol ID for RequestNetworkSettings.
func (*RequestNetworkSettings) ID() uint32 { return IDRequestNetworkSettings }
