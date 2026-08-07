// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetLocalPlayerAsInitialized struct {
	PlayerID ActorRuntimeID
}

// Marshal reads or writes SetLocalPlayerAsInitialized using its canonical wire layout.
func (x *SetLocalPlayerAsInitialized) Marshal(io IO) {
	x.PlayerID.Marshal(io)
}
