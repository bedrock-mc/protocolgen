// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdatePlayerGameType struct {
	PlayerGameType GameType
	TargetPlayer   int64
	Tick           uint64
}

// Marshal reads or writes UpdatePlayerGameType using its canonical wire layout.
func (x *UpdatePlayerGameType) Marshal(io IO) {
	IntegerFunc(&x.PlayerGameType, io.Varint32)
	io.ActorUniqueID(&x.TargetPlayer)
	io.PlayerInputTick(&x.Tick)
}
