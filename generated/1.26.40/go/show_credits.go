// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ShowCredits struct {
	PlayerRuntimeID ActorRuntimeID
	CreditsState    int32
}

// Marshal reads or writes ShowCredits using its canonical wire layout.
func (x *ShowCredits) Marshal(io IO) {
	x.PlayerRuntimeID.Marshal(io)
	io.Varint32(&x.CreditsState)
}
