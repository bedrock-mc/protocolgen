// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PackInstanceID struct {
	PackID      string
	Version     string
	SubPackName string
}

// Marshal reads or writes PackInstanceID using its canonical wire layout.
func (x *PackInstanceID) Marshal(io IO) {
	io.String(&x.PackID)
	io.String(&x.Version)
	io.String(&x.SubPackName)
}
