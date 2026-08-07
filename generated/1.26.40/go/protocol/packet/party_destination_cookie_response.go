// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PartyDestinationCookieResponse struct {
	Cookie   string
	Accepted bool
}

// Marshal reads or writes PartyDestinationCookieResponse using its canonical wire layout.
func (x *PartyDestinationCookieResponse) Marshal(io protocol.IO) {
	io.String(&x.Cookie)
	io.Bool(&x.Accepted)
}
