// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionOptionsSplineInstructionSplineProgressOption struct {
	KeyFrameValue      float32
	KeyFrameTime       float32
	KeyFrameEasingFunc string
}

// Marshal reads or writes CameraInstructionOptionsSplineInstructionSplineProgressOption using its canonical wire layout.
func (x *CameraInstructionOptionsSplineInstructionSplineProgressOption) Marshal(io IO) {
	io.Float32(&x.KeyFrameValue)
	io.Float32(&x.KeyFrameTime)
	io.String(&x.KeyFrameEasingFunc)
}
