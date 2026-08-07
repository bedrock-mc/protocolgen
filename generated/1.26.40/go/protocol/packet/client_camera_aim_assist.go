// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ClientCameraAimAssist struct {
	CameraPresetID string
	Action         protocol.ClientCameraAimAssistAction
	AllowAimAssist bool
}

// Marshal reads or writes ClientCameraAimAssist using its canonical wire layout.
func (x *ClientCameraAimAssist) Marshal(io protocol.IO) {
	io.String(&x.CameraPresetID)
	protocol.IntegerFunc(&x.Action, io.Uint8)
	io.Bool(&x.AllowAimAssist)
}

// ID returns the protocol ID for ClientCameraAimAssist.
func (*ClientCameraAimAssist) ID() uint32 { return IDClientCameraAimAssist }
