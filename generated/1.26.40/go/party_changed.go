// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PartyChanged struct {
	PartyInfo Optional[PlayerPartyInfo]
}

// Marshal reads or writes PartyChanged using its canonical wire layout.
func (x *PartyChanged) Marshal(io IO) {
	OptionalFunc(io, &x.PartyInfo, func(value *PlayerPartyInfo) {
		value.Marshal(io)
	})
}
