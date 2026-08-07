// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

import "image/color"

type TextShape struct {
	Text             string
	UseRotation      bool
	BackgroundColor  Optional[color.RGBA]
	DepthTest        bool
	ShowBackface     bool
	ShowTextBackface bool
}

func (*TextShape) isPrimitiveShapeExtraShapeData() {}

// Marshal reads or writes TextShape using its canonical wire layout.
func (x *TextShape) Marshal(io IO) {
	io.String(&x.Text)
	io.Bool(&x.UseRotation)
	OptionalFunc(io, &x.BackgroundColor, io.RGBA)
	io.Bool(&x.DepthTest)
	io.Bool(&x.ShowBackface)
	io.Bool(&x.ShowTextBackface)
}
