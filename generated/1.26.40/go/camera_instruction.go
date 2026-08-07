// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstruction struct {
	CameraInstruction CameraInstructionData
}

// Marshal reads or writes CameraInstruction using its canonical wire layout.
func (x *CameraInstruction) Marshal(io IO) {
	x.CameraInstruction.Marshal(io)
}
