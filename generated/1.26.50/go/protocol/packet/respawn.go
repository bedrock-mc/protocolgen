// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// Respawn is sent by the server to make a player respawn client-side. It is sent in response to a
// PlayerAction packet with ActionType PlayerActionRespawn. As of 1.13, the server sends two of
// these packets with different states, and the client sends one of these back in order to complete
// the respawn.
type Respawn struct {
	// Position is the position on which the player should be respawned. The position might be in a
	// different dimension, in which case the client should first be sent a ChangeDimension packet.
	Position mgl32.Vec3
	// State is the 'state' of the respawn. It is one of the constants that may be found above, and the
	// value the packet contains depends on whether the server or client sends it.
	State           protocol.PlayerRespawnState
	PlayerRuntimeID uint64
}

// Marshal reads or writes Respawn using its canonical wire layout.
func (x *Respawn) Marshal(io protocol.IO) {
	io.Vec3(&x.Position)
	protocol.IntegerFunc(&x.State, io.Uint8)
	io.ActorRuntimeID(&x.PlayerRuntimeID)
}

// ID returns the protocol ID for Respawn.
func (*Respawn) ID() uint32 { return IDRespawn }
