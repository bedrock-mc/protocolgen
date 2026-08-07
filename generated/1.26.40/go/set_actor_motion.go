// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type SetActorMotion struct {
	TargetRuntimeID uint64
	Motion          mgl32.Vec3
	Tick            uint64
}

// Marshal reads or writes SetActorMotion using its canonical wire layout.
func (x *SetActorMotion) Marshal(io IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.Vec3(&x.Motion)
	io.PlayerInputTick(&x.Tick)
}
