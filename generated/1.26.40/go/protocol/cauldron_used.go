// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CauldronUsed struct {
	ContentsColor uint32
	ContentsType  int32
	FillLevel     int32
}

func (*CauldronUsed) isEventData() {}

// Marshal reads or writes CauldronUsed using its canonical wire layout.
func (x *CauldronUsed) Marshal(io IO) {
	io.Varuint32(&x.ContentsColor)
	io.Varint32(&x.ContentsType)
	io.Varint32(&x.FillLevel)
}
