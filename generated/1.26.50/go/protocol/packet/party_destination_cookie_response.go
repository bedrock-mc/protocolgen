// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// PartyDestinationCookieResponse is sent by the client to the server in response to a
// SendPartyDestinationCookie packet.
type PartyDestinationCookieResponse struct {
	// Cookie is the opaque party destination cookie echoed back from the SendPartyDestinationCookie packet.
	Cookie string
	// Accepted is true if the client accepted the party destination.
	Accepted bool
}

// ID returns the protocol ID for PartyDestinationCookieResponse.
func (*PartyDestinationCookieResponse) ID() uint32 { return IDPartyDestinationCookieResponse }

// Marshal reads or writes PartyDestinationCookieResponse using its canonical wire layout.
func (pk *PartyDestinationCookieResponse) Marshal(io protocol.IO) {
	io.StringLimits(&pk.Cookie, 0, 2048)
	io.Bool(&pk.Accepted)
}
