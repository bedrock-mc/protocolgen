// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsSetInstructionEaseOption struct {
	Type uint8
	Time float32
}

// Marshal reads or writes CameraInstructionOptionsSetInstructionEaseOption using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstructionEaseOption) Marshal(io IO) {
	io.Uint8(&x.Type)
	io.Float32(&x.Time)
}
