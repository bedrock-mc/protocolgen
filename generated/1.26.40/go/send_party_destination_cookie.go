// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SendPartyDestinationCookie struct {
	Cookie          string
	Intent          string
	DestinationName string
}

// Marshal reads or writes SendPartyDestinationCookie using its canonical wire layout.
func (x *SendPartyDestinationCookie) Marshal(io IO) {
	io.String(&x.Cookie)
	io.String(&x.Intent)
	io.String(&x.DestinationName)
}
