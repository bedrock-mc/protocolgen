// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayerPartyInfo struct {
	PartyID       string
	IsPartyLeader bool
}

// Marshal reads or writes PlayerPartyInfo using its canonical wire layout.
func (x *PlayerPartyInfo) Marshal(io IO) {
	io.String(&x.PartyID)
	io.Bool(&x.IsPartyLeader)
}
