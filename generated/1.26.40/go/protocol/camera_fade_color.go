// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraFadeColor struct {
	Red   float32
	Green float32
	Blue  float32
}

// Marshal reads or writes CameraFadeColor using its canonical wire layout.
func (x *CameraFadeColor) Marshal(io IO) {
	io.Float32(&x.Red)
	io.Float32(&x.Green)
	io.Float32(&x.Blue)
}
