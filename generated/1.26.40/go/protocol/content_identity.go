// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ContentIdentity struct {
	Identity string
}

// Marshal reads or writes ContentIdentity using its canonical wire layout.
func (x *ContentIdentity) Marshal(io IO) {
	io.String(&x.Identity)
}
