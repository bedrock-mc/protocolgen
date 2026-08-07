// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PackInstanceId struct {
	PackID      string
	Version     string
	SubPackName string
}

// Marshal reads or writes PackInstanceId using its canonical wire layout.
func (x *PackInstanceId) Marshal(io IO) {
	io.String(&x.PackID)
	io.String(&x.Version)
	io.String(&x.SubPackName)
}
