// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BookEditActionAddPage struct {
	PageIndex int32
	PageText  string
	PhotoName string
}

func (BookEditActionAddPage) isBookEditAction() {}

// Marshal reads or writes BookEditActionAddPage using its canonical wire layout.
func (x *BookEditActionAddPage) Marshal(io IO) {
	io.Varint32(&x.PageIndex)
	io.String(&x.PageText)
	io.String(&x.PhotoName)
}
