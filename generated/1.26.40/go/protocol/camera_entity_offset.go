// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CameraEntityOffset struct {
	EntityOffsetX float32
	EntityOffsetY float32
	EntityOffsetZ float32
}

// Marshal reads or writes CameraEntityOffset using its canonical wire layout.
func (x *CameraEntityOffset) Marshal(io IO) {
	io.Float32(&x.EntityOffsetX)
	io.Float32(&x.EntityOffsetY)
	io.Float32(&x.EntityOffsetZ)
}
