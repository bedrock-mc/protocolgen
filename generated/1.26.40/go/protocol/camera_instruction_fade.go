// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraInstructionFade struct {
	Time  Optional[CameraFadeTimeData]
	Color Optional[CameraFadeColor]
}

// Marshal reads or writes CameraInstructionFade using its canonical wire layout.
func (x *CameraInstructionFade) Marshal(io IO) {
	OptionalFunc(io, &x.Time, func(value *CameraFadeTimeData) {
		value.Marshal(io)
	})
	OptionalFunc(io, &x.Color, func(value *CameraFadeColor) {
		value.Marshal(io)
	})
}
