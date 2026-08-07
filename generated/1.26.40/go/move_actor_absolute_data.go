// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type MoveActorAbsoluteData struct {
	ActorRuntimeID uint64
	Header         uint8
	Position       mgl32.Vec3
	RotationX      uint8
	RotationY      uint8
	RotationYHead  uint8
}

// Marshal reads or writes MoveActorAbsoluteData using its canonical wire layout.
func (x *MoveActorAbsoluteData) Marshal(io IO) {
	io.ActorRuntimeID(&x.ActorRuntimeID)
	io.Uint8(&x.Header)
	io.Vec3(&x.Position)
	io.Uint8(&x.RotationX)
	io.Uint8(&x.RotationY)
	io.Uint8(&x.RotationYHead)
}
