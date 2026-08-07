// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type CylinderData struct {
	RadiusX     mgl32.Vec2
	RadiusZ     mgl32.Vec2
	Height      float32
	NumSegments uint8
}

func (CylinderData) isPrimitiveShapeDataExtraShapeData() {}

// Marshal reads or writes CylinderData using its canonical wire layout.
func (x *CylinderData) Marshal(io IO) {
	io.Vec2(&x.RadiusX)
	io.Vec2(&x.RadiusZ)
	io.Float32(&x.Height)
	io.Uint8(&x.NumSegments)
}
