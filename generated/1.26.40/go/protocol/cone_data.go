// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type ConeData struct {
	Radii       mgl32.Vec2
	Height      float32
	NumSegments uint8
}

func (*ConeData) isPrimitiveShapeExtraShapeData() {}

// Marshal reads or writes ConeData using its canonical wire layout.
func (x *ConeData) Marshal(io IO) {
	io.Vec2(&x.Radii)
	io.Float32(&x.Height)
	io.Uint8(&x.NumSegments)
}
