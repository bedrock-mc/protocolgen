// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraInstructionOptionsFadeInstruction struct {
	Time  Optional[CameraInstructionOptionsFadeInstructionTimeOption]
	Color Optional[CameraInstructionOptionsFadeInstructionColorOption]
}

// Marshal reads or writes CameraInstructionOptionsFadeInstruction using its canonical wire layout.
func (x *CameraInstructionOptionsFadeInstruction) Marshal(io IO) {
	OptionalFunc(io, &x.Time, func(value *CameraInstructionOptionsFadeInstructionTimeOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	OptionalFunc(io, &x.Color, func(value *CameraInstructionOptionsFadeInstructionColorOption) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
