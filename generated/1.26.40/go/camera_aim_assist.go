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
