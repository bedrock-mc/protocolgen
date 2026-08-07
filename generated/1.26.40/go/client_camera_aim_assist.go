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
	enumValue1 := uint8(x.Action)
	io.Uint8(&enumValue1)
	x.Action = ClientCameraAimAssistPacketAction(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Bool(&x.AllowAimAssist)
}
