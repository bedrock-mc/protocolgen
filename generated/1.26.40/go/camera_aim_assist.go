// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type CameraAimAssist struct {
	PresetId        string
	ViewAngle       mgl32.Vec2
	Distance        float32
	TargetMode      CameraAimAssistTargetModeType
	Action          CameraAimAssistAction
	ShowDebugRender bool
}

// Marshal reads or writes CameraAimAssist using its canonical wire layout.
func (x *CameraAimAssist) Marshal(io IO) {
	io.String(&x.PresetId)
	io.Vec2(&x.ViewAngle)
	io.Float32(&x.Distance)
	enumValue1 := uint8(x.TargetMode)
	io.Uint8(&enumValue1)
	x.TargetMode = CameraAimAssistTargetModeType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	enumValue2 := uint8(x.Action)
	io.Uint8(&enumValue2)
	x.Action = CameraAimAssistAction(enumValue2)
	switch int64(enumValue2) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
	io.Bool(&x.ShowDebugRender)
}
