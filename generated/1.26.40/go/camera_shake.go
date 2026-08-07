// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type CameraShake struct {
	Intensity   float32
	Seconds     float32
	ShakeType   CameraShakeType
	ShakeAction CameraShakeAction
}

// Marshal reads or writes CameraShake using its canonical wire layout.
func (x *CameraShake) Marshal(io IO) {
	io.Float32(&x.Intensity)
	io.Float32(&x.Seconds)
	IntegerFunc(&x.ShakeType, io.Uint8)
	IntegerFunc(&x.ShakeAction, io.Uint8)
}
