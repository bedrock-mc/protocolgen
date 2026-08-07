// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SkinImage struct {
	Width      uint32
	Height     uint32
	ImageBytes []uint8
}

// Marshal reads or writes SkinImage using its canonical wire layout.
func (x *SkinImage) Marshal(io IO) {
	io.Uint32(&x.Width)
	io.Uint32(&x.Height)
	FuncSlice(io, &x.ImageBytes, io.Varuint32, io.Uint8)
}
