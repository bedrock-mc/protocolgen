// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraRotationOption struct {
	KeyFrameValue      mgl32.Vec3
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

// Marshal reads or writes CameraRotationOption using its canonical wire layout.
func (x *CameraRotationOption) Marshal(io IO) {
	io.Vec3(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}
