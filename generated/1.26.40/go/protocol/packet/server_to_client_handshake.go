// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ServerToClientHandshake struct {
	HandshakeWebToken string
}

// Marshal reads or writes ServerToClientHandshake using its canonical wire layout.
func (x *ServerToClientHandshake) Marshal(io protocol.IO) {
	io.String(&x.HandshakeWebToken)
}
