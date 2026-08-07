// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionSet struct {
	Preset                              uint32
	Ease                                Optional[CameraEase]
	Pos                                 Optional[CameraPosition]
	Rot                                 Optional[CameraRotation]
	Facing                              Optional[CameraFacing]
	ViewOffset                          Optional[CameraViewOffset]
	EntityOffset                        Optional[CameraEntityOffset]
	Default                             Optional[bool]
	RemoveIgnoreStartingValuesComponent bool
}

// Marshal reads or writes CameraInstructionSet using its canonical wire layout.
func (x *CameraInstructionSet) Marshal(io IO) {
	io.Uint32(&x.Preset)
	OptionalFunc(io, &x.Ease, func(value *CameraEase) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Pos, func(value *CameraPosition) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Rot, func(value *CameraRotation) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Facing, func(value *CameraFacing) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.ViewOffset, func(value *CameraViewOffset) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.EntityOffset, func(value *CameraEntityOffset) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Default, io.Bool)
	io.Bool(&x.RemoveIgnoreStartingValuesComponent)
}
