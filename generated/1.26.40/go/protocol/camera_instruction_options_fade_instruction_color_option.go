// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionOptionsFadeInstructionColorOption struct {
	Red   float32
	Green float32
	Blue  float32
}

// Marshal reads or writes CameraInstructionOptionsFadeInstructionColorOption using its canonical wire layout.
func (x *CameraInstructionOptionsFadeInstructionColorOption) Marshal(io IO) {
	io.Float32(&x.Red)
	io.Float32(&x.Green)
	io.Float32(&x.Blue)
}
