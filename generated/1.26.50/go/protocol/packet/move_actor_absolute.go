// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// MoveActorAbsolute is sent by the server to move an entity to an absolute position. It is typically used for
// movements where high accuracy isn't needed, such as for long range teleporting.
type MoveActorAbsolute struct {
	ActorRuntimeID uint64
	Header         uint8
	Position       mgl32.Vec3
	RotationX      uint8
	RotationY      uint8
	RotationYHead  uint8
}

// ID returns the protocol ID for MoveActorAbsolute.
func (*MoveActorAbsolute) ID() uint32 { return IDMoveActorAbsolute }

// Marshal reads or writes MoveActorAbsolute using its canonical wire layout.
func (pk *MoveActorAbsolute) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.ActorRuntimeID)
	io.Uint8(&pk.Header)
	io.Vec3(&pk.Position)
	io.Uint8(&pk.RotationX)
	io.Uint8(&pk.RotationY)
	io.Uint8(&pk.RotationYHead)
}
