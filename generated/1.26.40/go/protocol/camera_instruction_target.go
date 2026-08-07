// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionTarget struct {
	EntityActorID int64
}

// Marshal reads or writes CameraInstructionTarget using its canonical wire layout.
func (x *CameraInstructionTarget) Marshal(io IO) {
	io.Int64(&x.EntityActorID)
}
