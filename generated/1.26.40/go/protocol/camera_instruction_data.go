// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionData struct {
	Set              Optional[CameraInstructionSet]
	Clear            Optional[bool]
	Fade             Optional[CameraInstructionFade]
	Target           Optional[CameraInstructionTargetData]
	RemoveTarget     Optional[bool]
	FieldOfView      Optional[CameraInstructionFieldOfView]
	Spline           Optional[CameraSplineInstruction]
	AttachToEntity   Optional[CameraInstructionTarget]
	DetachFromEntity Optional[bool]
}

// Marshal reads or writes CameraInstructionData using its canonical wire layout.
func (x *CameraInstructionData) Marshal(io IO) {
	OptionalFunc(io, &x.Set, func(value *CameraInstructionSet) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Clear, io.Bool)
	OptionalFunc(io, &x.Fade, func(value *CameraInstructionFade) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Target, func(value *CameraInstructionTargetData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.RemoveTarget, io.Bool)
	OptionalFunc(io, &x.FieldOfView, func(value *CameraInstructionFieldOfView) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Spline, func(value *CameraSplineInstruction) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.AttachToEntity, func(value *CameraInstructionTarget) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.DetachFromEntity, io.Bool)
}
