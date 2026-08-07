// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PixelRequest struct {
	Pixel uint32
	Index uint16
}

// Marshal reads or writes PixelRequest using its canonical wire layout.
func (x *PixelRequest) Marshal(io IO) {
	io.Uint32(&x.Pixel)
	io.Uint16(&x.Index)
}
