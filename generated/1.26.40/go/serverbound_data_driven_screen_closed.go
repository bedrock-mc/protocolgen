// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerboundDataDrivenScreenClosed struct {
	FormId      uint32
	CloseReason string
}

// Marshal reads or writes ServerboundDataDrivenScreenClosed using its canonical wire layout.
func (x *ServerboundDataDrivenScreenClosed) Marshal(io IO) {
	io.Uint32(&x.FormId)
	io.String(&x.CloseReason)
}
