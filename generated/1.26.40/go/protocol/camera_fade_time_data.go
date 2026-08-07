// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraFadeTimeData struct {
	FadeInTime  float32
	HoldTime    float32
	FadeOutTime float32
}

// Marshal reads or writes CameraFadeTimeData using its canonical wire layout.
func (x *CameraFadeTimeData) Marshal(io IO) {
	io.Float32(&x.FadeInTime)
	io.Float32(&x.HoldTime)
	io.Float32(&x.FadeOutTime)
}
