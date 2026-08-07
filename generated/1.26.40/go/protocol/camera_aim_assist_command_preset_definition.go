// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CameraAimAssistCommandPresetDefinition struct {
	PresetId   Optional[string]
	TargetMode Optional[CameraAimAssistTargetMode]
	ViewAngle  Optional[mgl32.Vec2]
	Distance   Optional[float32]
}

// Marshal reads or writes CameraAimAssistCommandPresetDefinition using its canonical wire layout.
func (x *CameraAimAssistCommandPresetDefinition) Marshal(io IO) {
	OptionalFunc(io, &x.PresetId, io.String)
	OptionalFunc(io, &x.TargetMode, func(value *CameraAimAssistTargetMode) {
		IntegerFunc(value, io.Int32)
	})
	OptionalFunc(io, &x.ViewAngle, io.Vec2)
	OptionalFunc(io, &x.Distance, io.Float32)
}
