// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EduSharedURIResource struct {
	ButtonName string
	LinkURI    string
}

// Marshal reads or writes EduSharedURIResource using its canonical wire layout.
func (x *EduSharedURIResource) Marshal(io IO) {
	io.String(&x.ButtonName)
	io.String(&x.LinkURI)
}
