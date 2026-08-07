// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraEase struct {
	Type uint8
	Time float32
}

// Marshal reads or writes CameraEase using its canonical wire layout.
func (x *CameraEase) Marshal(io IO) {
	io.Uint8(&x.Type)
	io.Float32(&x.Time)
}
