// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientCameraAimAssist struct {
	CameraPresetId string
	Action         ClientCameraAimAssistPacketAction
	AllowAimAssist bool
}

// Marshal reads or writes ClientCameraAimAssist using its canonical wire layout.
func (x *ClientCameraAimAssist) Marshal(io IO) {
	io.String(&x.CameraPresetId)
	IntegerFunc(&x.Action, io.Uint8)
	io.Bool(&x.AllowAimAssist)
}
