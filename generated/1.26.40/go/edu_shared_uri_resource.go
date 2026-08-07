// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EduSharedUriResource struct {
	ButtonName string
	LinkUri    string
}

// Marshal reads or writes EduSharedUriResource using its canonical wire layout.
func (x *EduSharedUriResource) Marshal(io IO) {
	io.String(&x.ButtonName)
	io.String(&x.LinkUri)
}
