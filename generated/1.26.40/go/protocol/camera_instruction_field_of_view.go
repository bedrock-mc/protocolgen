// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionFieldOfView struct {
	FieldOfView      float32
	FOVEaseTime      float32
	FOVEaseType      string
	FieldOfViewClear bool
}

// Marshal reads or writes CameraInstructionFieldOfView using its canonical wire layout.
func (x *CameraInstructionFieldOfView) Marshal(io IO) {
	io.Float32(&x.FieldOfView)
	io.Float32(&x.FOVEaseTime)
	io.String(&x.FOVEaseType)
	io.Bool(&x.FieldOfViewClear)
}
