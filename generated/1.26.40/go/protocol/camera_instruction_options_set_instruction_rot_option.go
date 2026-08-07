// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionOptionsSetInstructionRotOption struct {
	X float32
	Y float32
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionRotOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionRotOption) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}
