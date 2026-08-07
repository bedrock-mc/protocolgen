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
	io.Bool(&x.DataInstanceId.set)
	if x.DataInstanceId.set {
		io.Uint32(&x.DataInstanceId.val)
	} else if io.Reading() {
		var zero uint32
		x.DataInstanceId.val = zero
	}
}
