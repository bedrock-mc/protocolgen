// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsFovInstruction struct {
	FieldOfView      float32
	FOVEaseTime      float32
	FOVEaseType      string
	FieldOfViewClear bool
}

// Marshal reads or writes CameraInstructionOptionsFovInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsFovInstruction) Marshal(io IO) {
	io.Float32(&x.FieldOfView)
	io.Float32(&x.FOVEaseTime)
	io.String(&x.FOVEaseType)
	io.Bool(&x.FieldOfViewClear)
}
