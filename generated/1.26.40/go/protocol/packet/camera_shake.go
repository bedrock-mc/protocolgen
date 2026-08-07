// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type CameraShake struct {
	Intensity   float32
	Seconds     float32
	ShakeType   protocol.CameraShakeType
	ShakeAction protocol.CameraShakeAction
}

// Marshal reads or writes CameraShake using its canonical wire layout.
func (x *CameraShake) Marshal(io protocol.IO) {
	io.Float32(&x.Intensity)
	io.Float32(&x.Seconds)
	protocol.IntegerFunc(&x.ShakeType, io.Uint8)
	protocol.IntegerFunc(&x.ShakeAction, io.Uint8)
}

// ID returns the protocol ID for CameraShake.
func (*CameraShake) ID() uint32 { return IDCameraShake }
