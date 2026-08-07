// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraFacing struct {
	Pos mgl32.Vec3
}

// Marshal reads or writes CameraFacing using its canonical wire layout.
func (x *CameraFacing) Marshal(io IO) {
	io.Vec3(&x.Pos)
}
