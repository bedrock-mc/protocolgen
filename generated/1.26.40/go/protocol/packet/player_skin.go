// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/google/uuid"
)

// PlayerSkin is sent by the client to the server when it updates its own skin using the in-game
// skin picker. It is relayed by the server, or sent if the server changes the skin of a player on
// its own accord. Note that the packet can only be sent for players that are in the player list at
// the time of sending.
type PlayerSkin struct {
	// UUID is the UUID of the player as sent in the Login packet when the client joined the server. It
	// must match this UUID exactly for the skin to show up on the player.
	UUID                 uuid.UUID
	SerializedSkin       protocol.SerializedSkinRef
	LocalizedNewSkinName string
	LocalizedOldSkinName string
}

// Marshal reads or writes PlayerSkin using its canonical wire layout.
func (x *PlayerSkin) Marshal(io protocol.IO) {
	io.UUID(&x.UUID)
	x.SerializedSkin.Marshal(io)
	io.String(&x.LocalizedNewSkinName)
	io.String(&x.LocalizedOldSkinName)
}

// ID returns the protocol ID for PlayerSkin.
func (*PlayerSkin) ID() uint32 { return IDPlayerSkin }
