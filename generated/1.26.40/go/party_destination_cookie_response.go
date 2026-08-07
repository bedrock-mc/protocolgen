// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PartyDestinationCookieResponse struct {
	Cookie   string
	Accepted bool
}

// Marshal reads or writes PartyDestinationCookieResponse using its canonical wire layout.
func (x *PartyDestinationCookieResponse) Marshal(io IO) {
	io.String(&x.Cookie)
	io.Bool(&x.Accepted)
}
