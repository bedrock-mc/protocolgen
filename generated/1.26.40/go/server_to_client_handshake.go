// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerToClientHandshake struct {
	HandshakeWebToken string
}

// Marshal reads or writes ServerToClientHandshake using its canonical wire layout.
func (x *ServerToClientHandshake) Marshal(io IO) {
	io.String(&x.HandshakeWebToken)
}
