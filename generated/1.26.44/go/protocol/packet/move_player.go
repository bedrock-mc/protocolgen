// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// MovePlayer is sent by players to send their movement to the server, and by the server to update
// the movement of player entities to other players.
type MovePlayer struct {
	PlayerRuntimeID uint64
	// Position is the position to spawn the player on. If the player is on a distance that the viewer
	// cannot see it, the player will still show up if the viewer moves closer.
	Position      mgl32.Vec3
	Rotation      mgl32.Vec2
	YHeadRotation float32
	PositionMode  protocol.PlayerPositionModeComponentPositionMode
	// OnGround specifies if the player is considered on the ground. Note that proxies or hacked clients
	// could fake this to always be true, so it should not be taken for granted.
	OnGround        bool
	RidingRuntimeID uint64
	TeleportData    protocol.Optional[protocol.MovePlayerTeleportData]
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
}

// Marshal reads or writes MovePlayer using its canonical wire layout.
func (x *MovePlayer) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.PlayerRuntimeID)
	io.Vec3(&x.Position)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	protocol.IntegerFunc(&x.PositionMode, io.Uint8)
	io.Bool(&x.OnGround)
	io.ActorRuntimeID(&x.RidingRuntimeID)
	protocol.OptionalFunc(io, &x.TeleportData, func(value *protocol.MovePlayerTeleportData) {
		value.Marshal(io)
	})
	io.PlayerInputTick(&x.Tick)
}

// ID returns the protocol ID for MovePlayer.
func (*MovePlayer) ID() uint32 { return IDMovePlayer }
