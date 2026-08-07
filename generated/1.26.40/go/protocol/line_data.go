// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type LineData struct {
	LineEndLocation mgl32.Vec3
}

func (*LineData) isPrimitiveShapeDataExtraShapeData() {}

// Marshal reads or writes LineData using its canonical wire layout.
func (x *LineData) Marshal(io IO) {
	io.Vec3(&x.LineEndLocation)
}
