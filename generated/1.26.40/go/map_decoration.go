// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "image/color"

type MapDecoration struct {
	ImageType MapDecorationType
	Rotation  uint8
	X         uint8
	Y         uint8
	Label     string
	Color     color.RGBA
}

// Marshal reads or writes MapDecoration using its canonical wire layout.
func (x *MapDecoration) Marshal(io IO) {
	IntegerFunc(&x.ImageType, io.Int8)
	io.Uint8(&x.Rotation)
	io.Uint8(&x.X)
	io.Uint8(&x.Y)
	io.String(&x.Label)
	io.RGBA(&x.Color)
}
