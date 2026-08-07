// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BookEditActionSwapPages struct {
	PageIndex     int32
	SwapWithIndex int32
}

func (BookEditActionSwapPages) isBookEditAction() {}

// Marshal reads or writes BookEditActionSwapPages using its canonical wire layout.
func (x *BookEditActionSwapPages) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
	io.Varint32(&x.SwapWithIndex)
}
