// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetPlayerGameType struct {
	PlayerGameType GameType
}

// Marshal reads or writes SetPlayerGameType using its canonical wire layout.
func (x *SetPlayerGameType) Marshal(io IO) {
	IntegerFunc(&x.PlayerGameType, io.Varint32)
}
