// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type CameraAimAssist struct {
	PresetId        string
	ViewAngle       mgl32.Vec2
	Distance        float32
	TargetMode      protocol.CameraAimAssistTargetModeType
	Action          protocol.CameraAimAssistAction
	ShowDebugRender bool
}

// Marshal reads or writes CameraAimAssist using its canonical wire layout.
func (x *CameraAimAssist) Marshal(io protocol.IO) {
	io.String(&x.PresetId)
	io.Vec2(&x.ViewAngle)
	io.Float32(&x.Distance)
	protocol.IntegerFunc(&x.TargetMode, io.Uint8)
	protocol.IntegerFunc(&x.Action, io.Uint8)
	io.Bool(&x.ShowDebugRender)
}

// ID returns the protocol ID for CameraAimAssist.
func (*CameraAimAssist) ID() uint32 { return IDCameraAimAssist }
