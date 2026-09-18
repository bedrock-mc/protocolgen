// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// CameraShake is sent by the server to make the camera shake client-side. This feature was added for map-
// making partners.
type CameraShake struct {
	// Intensity is the intensity of the shaking. The client limits this value to 4, so anything higher may not
	// work.
	Intensity   float32
	Seconds     float32
	ShakeType   protocol.CameraShakeType
	ShakeAction protocol.CameraShakeAction
}

// ID ...
func (*CameraShake) ID() uint32 {
	return IDCameraShake
}

func (pk *CameraShake) Marshal(io protocol.IO) {
	io.Float32(&pk.Intensity)
	io.Float32(&pk.Seconds)
	pk.ShakeType.Marshal(io)
	pk.ShakeAction.Marshal(io)
}
