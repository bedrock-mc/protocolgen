// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// SubClientLogin is sent when a sub-client joins the server while another client is already
// connected to it. The packet is sent as a result of split-screen game play, and allows up to four
// players to play using the same network connection. After an initial Login packet from the 'main'
// client, each sub-client that connects sends a SubClientLogin to request their own login.
type SubClientLogin struct {
	SubClientConnectionRequest []byte
}

// Marshal reads or writes SubClientLogin using its canonical wire layout.
func (x *SubClientLogin) Marshal(io protocol.IO) {
	io.Bytes(&x.SubClientConnectionRequest)
}

// ID returns the protocol ID for SubClientLogin.
func (*SubClientLogin) ID() uint32 { return IDSubClientLogin }
