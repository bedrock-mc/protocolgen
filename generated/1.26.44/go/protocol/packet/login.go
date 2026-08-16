// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// Login is sent when the client initially tries to join the server. It is the first packet sent and
// contains information specific to the player.
type Login struct {
	ClientNetworkVersion int32
	// ConnectionRequest is a string containing information about the player and JWTs that may be used
	// to verify if the player is connected to XBOX Live. The connection request also contains the
	// necessary client public key to initiate encryption.
	ConnectionRequest []byte
}

// Marshal reads or writes Login using its canonical wire layout.
func (x *Login) Marshal(io protocol.IO) {
	io.BEInt32(&x.ClientNetworkVersion)
	io.Bytes(&x.ConnectionRequest)
}

// ID returns the protocol ID for Login.
func (*Login) ID() uint32 { return IDLogin }
