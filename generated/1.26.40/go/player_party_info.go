// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayerPartyInfo struct {
	PartyId       string
	IsPartyLeader bool
}

// Marshal reads or writes PlayerPartyInfo using its canonical wire layout.
func (x *PlayerPartyInfo) Marshal(io IO) {
	io.String(&x.PartyId)
	io.Bool(&x.IsPartyLeader)
}
