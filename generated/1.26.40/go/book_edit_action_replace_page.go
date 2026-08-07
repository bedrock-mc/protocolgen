// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BookEditActionReplacePage struct {
	PageIndex int32
	PageText  string
	PhotoName string
}

func (BookEditActionReplacePage) isBookEditAction() {}

// Marshal reads or writes BookEditActionReplacePage using its canonical wire layout.
func (x *BookEditActionReplacePage) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
	io.String(&x.PageText)
	io.String(&x.PhotoName)
}
