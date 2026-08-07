// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MapInfoRequestPacketAnonClientPixelsProxy struct {
	Pixel uint32
	Index uint16
}

// Marshal reads or writes MapInfoRequestPacketAnonClientPixelsProxy using its canonical wire layout.
func (x *MapInfoRequestPacketAnonClientPixelsProxy) Marshal(io IO) {
	io.Uint32(&x.Pixel)
	io.Uint16(&x.Index)
}
