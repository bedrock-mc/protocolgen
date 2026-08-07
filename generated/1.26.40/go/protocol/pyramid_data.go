// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PyramidData struct {
	Width  float32
	Depth  Optional[float32]
	Height float32
}

func (*PyramidData) isPrimitiveShapeDataExtraShapeData() {}

// Marshal reads or writes PyramidData using its canonical wire layout.
func (x *PyramidData) Marshal(io IO) {
	io.Float32(&x.Width)
	OptionalFunc(io, &x.Depth, io.Float32)
	io.Float32(&x.Height)
}
