// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SemVersion struct {
	Version string
}

// Marshal reads or writes SemVersion using its canonical wire layout.
func (x *SemVersion) Marshal(io IO) {
	io.String(&x.Version)
}
