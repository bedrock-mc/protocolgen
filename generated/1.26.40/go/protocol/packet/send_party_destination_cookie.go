// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// SendPartyDestinationCookie is sent by the server to a client with a party destination cookie.
type SendPartyDestinationCookie struct {
	// Cookie is the opaque party destination cookie.
	Cookie string
	// Intent is the intent of the cookie. It is one of the PartyDestinationCookieIntent constants.
	Intent string
	// DestinationName is the name of the destination the cookie refers to.
	DestinationName string
}

// Marshal reads or writes SendPartyDestinationCookie using its canonical wire layout.
func (x *SendPartyDestinationCookie) Marshal(io protocol.IO) {
	io.StringLimits(&x.Cookie, 0, 2048)
	io.String(&x.Intent)
	io.StringLimits(&x.DestinationName, 0, 64)
}

// ID returns the protocol ID for SendPartyDestinationCookie.
func (*SendPartyDestinationCookie) ID() uint32 { return IDSendPartyDestinationCookie }
