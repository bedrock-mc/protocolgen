// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SemVersionData struct {
	Version string
}

// Marshal reads or writes SemVersionData using its canonical wire layout.
func (x *SemVersionData) Marshal(io IO) {
	io.String(&x.Version)
}
