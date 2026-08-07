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

// Marshal reads or writes MovePlayer using its canonical wire layout.
func (x *MovePlayer) Marshal(io IO) {
	x.PlayerRuntimeID.Marshal(io)
	io.Vec3(&x.Position)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	enumValue1 := uint8(x.PositionMode)
	io.Uint8(&enumValue1)
	x.PositionMode = PlayerPositionModeComponentPositionMode(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 3:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.OnGround)
	x.RidingRuntimeID.Marshal(io)
	io.Bool(&x.TeleportData.set)
	if x.TeleportData.set {
		x.TeleportData.val.Marshal(io)
	} else if io.Reading() {
		var zero MovePlayerTeleportData
		x.TeleportData.val = zero
	}
	x.Tick.Marshal(io)
}
