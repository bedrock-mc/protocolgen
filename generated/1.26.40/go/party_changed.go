// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PartyChanged struct {
	PartyInfo Optional[PlayerPartyInfo]
}

// Marshal reads or writes PartyChanged using its canonical wire layout.
func (x *PartyChanged) Marshal(io IO) {
	io.Bool(&x.PartyInfo.set)
	if x.PartyInfo.set {
		x.PartyInfo.val.Marshal(io)
	} else if io.Reading() {
		var zero PlayerPartyInfo
		x.PartyInfo.val = zero
	}
}
