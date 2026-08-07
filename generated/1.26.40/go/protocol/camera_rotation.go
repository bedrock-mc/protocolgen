// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraRotation struct {
	X float32
	Y float32
}

// Marshal reads or writes CameraRotation using its canonical wire layout.
func (x *CameraRotation) Marshal(io IO) {
	io.Float32(&x.X)
	io.Float32(&x.Y)
}
