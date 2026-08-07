// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionData struct {
	Set              Optional[CameraInstructionOptionsSetInstruction]
	Clear            Optional[bool]
	Fade             Optional[CameraInstructionOptionsFadeInstruction]
	Target           Optional[CameraInstructionOptionsTargetInstruction]
	RemoveTarget     Optional[bool]
	FieldOfView      Optional[CameraInstructionOptionsFovInstruction]
	Spline           Optional[CameraInstructionOptionsSplineInstruction]
	AttachToEntity   Optional[CameraInstructionOptionsAttachToEntityInstruction]
	DetachFromEntity Optional[bool]
}

// Marshal reads or writes CameraInstructionData using its canonical wire layout.
func (x *CameraInstructionData) Marshal(io IO) {
	OptionalFunc(io, &x.Set, func(value *CameraInstructionOptionsSetInstruction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Clear, io.Bool)
	OptionalFunc(io, &x.Fade, func(value *CameraInstructionOptionsFadeInstruction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Target, func(value *CameraInstructionOptionsTargetInstruction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.RemoveTarget, io.Bool)
	OptionalFunc(io, &x.FieldOfView, func(value *CameraInstructionOptionsFovInstruction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Spline, func(value *CameraInstructionOptionsSplineInstruction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.AttachToEntity, func(value *CameraInstructionOptionsAttachToEntityInstruction) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.DetachFromEntity, io.Bool)
}
