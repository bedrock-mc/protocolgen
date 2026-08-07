// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type SendPartyDestinationCookie struct {
	Cookie          string
	Intent          string
	DestinationName string
}

// Marshal reads or writes SendPartyDestinationCookie using its canonical wire layout.
func (x *SendPartyDestinationCookie) Marshal(io protocol.IO) {
	io.String(&x.Cookie)
	io.String(&x.Intent)
	io.String(&x.DestinationName)
}
