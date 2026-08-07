// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "github.com/go-gl/mathgl/mgl32"

type ArrowData struct {
	ArrowEndLocation Optional[mgl32.Vec3]
	ArrowHeadLength  Optional[float32]
	ArrowHeadRadius  Optional[float32]
	NumSegments      Optional[uint8]
}

func (*ArrowData) isPrimitiveShapeDataExtraShapeData() {}

// Marshal reads or writes ArrowData using its canonical wire layout.
func (x *ArrowData) Marshal(io IO) {
	OptionalFunc(io, &x.ArrowEndLocation, io.Vec3)
	OptionalFunc(io, &x.ArrowHeadLength, io.Float32)
	OptionalFunc(io, &x.ArrowHeadRadius, io.Float32)
	OptionalFunc(io, &x.NumSegments, io.Uint8)
}
