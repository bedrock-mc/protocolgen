// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundDataDrivenUIShowScreen struct {
	ScreenId       string
	FormId         uint32
	DataInstanceId Optional[uint32]
}

// Marshal reads or writes ClientboundDataDrivenUIShowScreen using its canonical wire layout.
func (x *ClientboundDataDrivenUIShowScreen) Marshal(io IO) {
	io.String(&x.ScreenId)
	io.Uint32(&x.FormId)
	OptionalFunc(io, &x.DataInstanceId, io.Uint32)
}
