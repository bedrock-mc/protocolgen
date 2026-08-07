// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type CameraSplineRotationKeyFrame struct {
	Rotation mgl32.Vec3
	Time     float32
	Easing   Optional[string]
}

// Marshal reads or writes CameraSplineRotationKeyFrame using its canonical wire layout.
func (x *CameraSplineRotationKeyFrame) Marshal(io IO) {
	io.Vec3(&x.Rotation)
	io.Float32(&x.Time)
	OptionalFunc(io, &x.Easing, io.String)
}
