// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type CameraInstructionOptionsSplineInstructionSplineRotationOption struct {
	KeyFrameValue      mgl32.Vec3
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

// Marshal reads or writes CameraInstructionOptionsSplineInstructionSplineRotationOption using its canonical wire layout.
func (x *CameraInstructionOptionsSplineInstructionSplineRotationOption) Marshal(io IO) {
	io.Vec3(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}
