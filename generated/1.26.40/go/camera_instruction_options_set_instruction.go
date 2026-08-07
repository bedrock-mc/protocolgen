// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsSetInstruction struct {
	Preset                              uint32
	Ease                                Optional[CameraInstructionOptionsSetInstructionEaseOption]
	Pos                                 Optional[CameraInstructionOptionsSetInstructionPosOption]
	Rot                                 Optional[CameraInstructionOptionsSetInstructionRotOption]
	Facing                              Optional[CameraInstructionOptionsSetInstructionFacingOption]
	ViewOffset                          Optional[CameraInstructionOptionsSetInstructionViewOffsetOption]
	EntityOffset                        Optional[CameraInstructionOptionsSetInstructionEntityOffsetOption]
	Default                             Optional[bool]
	RemoveIgnoreStartingValuesComponent bool
}

// Marshal reads or writes CameraInstructionOptionsSetInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsSetInstruction) Marshal(io IO) {
	io.Uint32(&x.Preset)
	OptionalFunc(io, &x.Ease, func(value *CameraInstructionOptionsSetInstructionEaseOption) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Pos, func(value *CameraInstructionOptionsSetInstructionPosOption) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Rot, func(value *CameraInstructionOptionsSetInstructionRotOption) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Facing, func(value *CameraInstructionOptionsSetInstructionFacingOption) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ViewOffset, func(value *CameraInstructionOptionsSetInstructionViewOffsetOption) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.EntityOffset, func(value *CameraInstructionOptionsSetInstructionEntityOffsetOption) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Default, io.Bool)
	io.Bool(&x.RemoveIgnoreStartingValuesComponent)
}
