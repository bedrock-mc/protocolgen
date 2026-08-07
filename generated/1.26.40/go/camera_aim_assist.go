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
	IntegerFunc(&x.TargetMode, io.Uint8)
	IntegerFunc(&x.Action, io.Uint8)
	io.Bool(&x.ShowDebugRender)
}
