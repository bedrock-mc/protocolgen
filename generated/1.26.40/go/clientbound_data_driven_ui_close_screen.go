// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundDataDrivenUICloseScreen struct {
	FormId Optional[uint32]
}

// Marshal reads or writes ClientboundDataDrivenUICloseScreen using its canonical wire layout.
func (x *ClientboundDataDrivenUICloseScreen) Marshal(io IO) {
	io.Bool(&x.FormId.set)
	if x.FormId.set {
		io.Uint32(&x.FormId.val)
	} else if io.Reading() {
		var zero uint32
		x.FormId.val = zero
	}
}
