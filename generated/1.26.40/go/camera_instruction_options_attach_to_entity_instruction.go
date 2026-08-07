// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsAttachToEntityInstruction struct {
	EntityActorID int64
}

// Marshal reads or writes CameraInstructionOptionsAttachToEntityInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsAttachToEntityInstruction) Marshal(io IO) {
	io.Int64(&x.EntityActorID)
}
