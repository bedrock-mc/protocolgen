// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type SetActorMotion struct {
	TargetRuntimeID ActorRuntimeID
	Motion          mgl32.Vec3
	Tick            PlayerInputTick
}

// Marshal reads or writes SetActorMotion using its canonical wire layout.
func (x *SetActorMotion) Marshal(io IO) {
	x.TargetRuntimeID.Marshal(io)
	io.Vec3(&x.Motion)
	x.Tick.Marshal(io)
}
