// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ServerToClientHandshake is sent by the server to the client to complete the key exchange in order
// to initialise encryption on client and server side. It is followed up by a
// ClientToServerHandshake packet from the client.
type ServerToClientHandshake struct {
	HandshakeWebToken string
}

// Marshal reads or writes ServerToClientHandshake using its canonical wire layout.
func (x *ServerToClientHandshake) Marshal(io protocol.IO) {
	io.String(&x.HandshakeWebToken)
}

// ID returns the protocol ID for ServerToClientHandshake.
func (*ServerToClientHandshake) ID() uint32 { return IDServerToClientHandshake }
