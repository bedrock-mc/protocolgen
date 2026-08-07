// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SubClientLogin struct {
	SubClientConnectionRequest string
}

// Marshal reads or writes SubClientLogin using its canonical wire layout.
func (x *SubClientLogin) Marshal(io IO) {
	io.String(&x.SubClientConnectionRequest)
}
