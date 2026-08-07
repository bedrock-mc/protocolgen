// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraInstructionTargetData struct {
	TargetCenterOffset Optional[mgl32.Vec3]
	TargetActorID      int64
}

// Marshal reads or writes CameraInstructionTargetData using its canonical wire layout.
func (x *CameraInstructionTargetData) Marshal(io IO) {
	OptionalFunc(io, &x.TargetCenterOffset, io.Vec3)
	io.Int64(&x.TargetActorID)
}
