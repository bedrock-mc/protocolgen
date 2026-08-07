// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SubClientLogin struct {
	SubClientConnectionRequest string
}

// Marshal reads or writes SubClientLogin using its canonical wire layout.
func (x *SubClientLogin) Marshal(io protocol.IO) {
	io.String(&x.SubClientConnectionRequest)
}

// ID returns the protocol ID for SubClientLogin.
func (*SubClientLogin) ID() uint32 { return IDSubClientLogin }
