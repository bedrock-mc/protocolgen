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
	enumValue1 := uint8(x.ShakeType)
	io.Uint8(&enumValue1)
	x.ShakeType = CameraShakeType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	enumValue2 := uint8(x.ShakeAction)
	io.Uint8(&enumValue2)
	x.ShakeAction = CameraShakeAction(enumValue2)
	switch int64(enumValue2) {
	case 0, 1:
	default:
		io.InvalidValue(enumValue2, "unknown enum value")
	}
}
