// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// PartyChanged is sent by the client to the server to indicate that the player's party ID has
// changed.
type PartyChanged struct {
	PartyInfo protocol.Optional[protocol.PlayerPartyInfo]
}

// Marshal reads or writes PartyChanged using its canonical wire layout.
func (x *PartyChanged) Marshal(io protocol.IO) {
	protocol.OptionalFunc(io, &x.PartyInfo, func(value *protocol.PlayerPartyInfo) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for PartyChanged.
func (*PartyChanged) ID() uint32 { return IDPartyChanged }
