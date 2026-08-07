// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type BoxData struct {
	BoxBound mgl32.Vec3
}

func (*BoxData) isPrimitiveShapeExtraShapeData() {}

// Marshal reads or writes BoxData using its canonical wire layout.
func (x *BoxData) Marshal(io IO) {
	io.Vec3(&x.BoxBound)
}
