// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraSplineControlPoint struct {
	Position mgl32.Vec3
}

// Marshal reads or writes CameraSplineControlPoint using its canonical wire layout.
func (x *CameraSplineControlPoint) Marshal(io IO) {
	io.Vec3(&x.Position)
}
