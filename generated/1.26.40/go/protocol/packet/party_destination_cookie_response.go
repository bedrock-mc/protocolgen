// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// PartyDestinationCookieResponse is sent by the client to the server in response to a
// SendPartyDestinationCookie packet.
type PartyDestinationCookieResponse struct {
	// Cookie is the opaque party destination cookie echoed back from the SendPartyDestinationCookie
	// packet.
	Cookie string
	// Accepted is true if the client accepted the party destination.
	Accepted bool
}

// Marshal reads or writes PartyDestinationCookieResponse using its canonical wire layout.
func (x *PartyDestinationCookieResponse) Marshal(io protocol.IO) {
	io.String(&x.Cookie)
	io.Bool(&x.Accepted)
}

// ID returns the protocol ID for PartyDestinationCookieResponse.
func (*PartyDestinationCookieResponse) ID() uint32 { return IDPartyDestinationCookieResponse }
