// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ActorDataFlagComponent struct {
	ActorFlagBitsetData Bitset131
}

// Marshal reads or writes ActorDataFlagComponent using its canonical wire layout.
func (x *ActorDataFlagComponent) Marshal(io IO) {
	io.Bitset(x.ActorFlagBitsetData[:], 131)
}
