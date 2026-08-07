// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BookEditActionFinalize struct {
	Title  string
	Author string
	XUID   string
}

func (BookEditActionFinalize) isBookEditAction() {}

// Marshal reads or writes BookEditActionFinalize using its canonical wire layout.
func (x *BookEditActionFinalize) Marshal(io IO) {
	io.String(&x.Title)
	io.String(&x.Author)
	io.String(&x.XUID)
}
