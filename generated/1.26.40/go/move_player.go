// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type MovePlayer struct {
	PlayerRuntimeID ActorRuntimeID
	Position        mgl32.Vec3
	Rotation        mgl32.Vec2
	YHeadRotation   float32
	PositionMode    PlayerPositionModeComponentPositionMode
	OnGround        bool
	RidingRuntimeID ActorRuntimeID
	TeleportData    Optional[MovePlayerTeleportData]
	Tick            PlayerInputTick
}
