// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type SetActorMotion struct {
	TargetRuntimeID uint64
	Motion          mgl32.Vec3
	Tick            uint64
}

// Marshal reads or writes SetActorMotion using its canonical wire layout.
func (x *SetActorMotion) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.Vec3(&x.Motion)
	io.PlayerInputTick(&x.Tick)
}
