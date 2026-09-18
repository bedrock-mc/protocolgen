// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// UpdateAdventureSettings is a packet sent from the server to the client to update the adventure settings of
// the player. It, along with the UpdateAbilities packet, are replacements of the AdventureSettings packet
// since v1.19.10.
type UpdateAdventureSettings struct {
	NoPvM          bool
	NoMvP          bool
	ImmutableWorld bool
	ShowNameTags   bool
	AutoJump       bool
}

// ID returns the protocol ID for UpdateAdventureSettings.
func (*UpdateAdventureSettings) ID() uint32 { return IDUpdateAdventureSettings }

// Marshal reads or writes UpdateAdventureSettings using its canonical wire layout.
func (pk *UpdateAdventureSettings) Marshal(io protocol.IO) {
	io.Bool(&pk.NoPvM)
	io.Bool(&pk.NoMvP)
	io.Bool(&pk.ImmutableWorld)
	io.Bool(&pk.ShowNameTags)
	io.Bool(&pk.AutoJump)
}
