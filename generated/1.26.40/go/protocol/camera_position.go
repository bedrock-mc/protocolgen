// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraPosition struct {
	Pos mgl32.Vec3
}

// Marshal reads or writes CameraPosition using its canonical wire layout.
func (x *CameraPosition) Marshal(io IO) {
	io.Vec3(&x.Pos)
}
