// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type EllipsoidData struct {
	Radii           mgl32.Vec3
	SegmentsPerAxis uint8
}

func (*EllipsoidData) isPrimitiveShapeExtraShapeData() {}

// Marshal reads or writes EllipsoidData using its canonical wire layout.
func (x *EllipsoidData) Marshal(io IO) {
	io.Vec3(&x.Radii)
	io.Uint8(&x.SegmentsPerAxis)
}
