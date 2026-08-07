// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BookEditActionDeletePage struct {
	PageIndex int32
}

func (BookEditActionDeletePage) isBookEditAction() {}

// Marshal reads or writes BookEditActionDeletePage using its canonical wire layout.
func (x *BookEditActionDeletePage) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
}
