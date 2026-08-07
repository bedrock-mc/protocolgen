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
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Pos, func(value *CameraInstructionOptionsSetInstructionPosOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Rot, func(value *CameraInstructionOptionsSetInstructionRotOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Facing, func(value *CameraInstructionOptionsSetInstructionFacingOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.ViewOffset, func(value *CameraInstructionOptionsSetInstructionViewOffsetOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.EntityOffset, func(value *CameraInstructionOptionsSetInstructionEntityOffsetOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Default, io.Bool)
	io.Bool(&x.RemoveIgnoreStartingValuesComponent)
}
