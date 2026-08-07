// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type CameraInstructionOptionsTargetInstruction struct {
	TargetCenterOffset Optional[mgl32.Vec3]
	TargetActorID      int64
}

// Marshal reads or writes CameraInstructionOptionsTargetInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsTargetInstruction) Marshal(io IO) {
	OptionalFunc(io, &x.TargetCenterOffset, io.Vec3)
	io.Int64(&x.TargetActorID)
}
