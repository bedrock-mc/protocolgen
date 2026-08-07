// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsSetInstructionViewOffsetOption struct {
	X float32
	Y float32
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionViewOffsetOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionViewOffsetOption) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}
