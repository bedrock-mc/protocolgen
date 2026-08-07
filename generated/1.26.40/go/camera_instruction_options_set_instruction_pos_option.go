// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type CameraInstructionOptionsSetInstructionPosOption struct {
	Pos mgl32.Vec3
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionPosOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionPosOption) Marshal(io IO) {
	io.Vec3(&x.Pos)
}
