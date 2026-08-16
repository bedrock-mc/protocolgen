// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// UpdateAdventureSettings is a packet sent from the server to the client to update the adventure
// settings of the player. It, along with the UpdateAbilities packet, are replacements of the
// AdventureSettings packet since v1.19.10.
type UpdateAdventureSettings struct {
	AdventureSettings protocol.AdventureSettings
}

// Marshal reads or writes UpdateAdventureSettings using its canonical wire layout.
func (x *UpdateAdventureSettings) Marshal(io protocol.IO) {
	x.AdventureSettings.Marshal(io)
}

// ID returns the protocol ID for UpdateAdventureSettings.
func (*UpdateAdventureSettings) ID() uint32 { return IDUpdateAdventureSettings }
