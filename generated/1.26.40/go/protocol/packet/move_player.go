// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type MovePlayer struct {
	PlayerRuntimeID uint64
	Position        mgl32.Vec3
	Rotation        mgl32.Vec2
	YHeadRotation   float32
	PositionMode    protocol.PlayerPositionModeComponentPositionMode
	OnGround        bool
	RidingRuntimeID uint64
	TeleportData    protocol.Optional[protocol.MovePlayerTeleportData]
	Tick            uint64
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
