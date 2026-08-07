// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ShowProfile struct {
	PlayerXUID string
}

// Marshal reads or writes ShowProfile using its canonical wire layout.
func (x *ShowProfile) Marshal(io protocol.IO) {
	io.String(&x.PlayerXUID)
}

// ID returns the protocol ID for ShowProfile.
func (*ShowProfile) ID() uint32 { return IDShowProfile }
