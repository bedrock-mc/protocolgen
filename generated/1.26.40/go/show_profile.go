// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ShowProfile struct {
	PlayerXUID string
}

// Marshal reads or writes ShowProfile using its canonical wire layout.
func (x *ShowProfile) Marshal(io IO) {
	io.String(&x.PlayerXUID)
}
