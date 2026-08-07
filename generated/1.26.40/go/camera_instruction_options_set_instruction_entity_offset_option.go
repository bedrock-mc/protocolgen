// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsSetInstructionEntityOffsetOption struct {
	EntityOffsetX float32
	EntityOffsetY float32
	EntityOffsetZ float32
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionEntityOffsetOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionEntityOffsetOption) Marshal(io IO) {
	io.Float32(&x.EntityOffsetX)
	io.Float32(&x.EntityOffsetY)
	io.Float32(&x.EntityOffsetZ)
}
