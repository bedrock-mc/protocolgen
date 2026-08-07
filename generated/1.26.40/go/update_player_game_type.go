// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdatePlayerGameType struct {
	PlayerGameType GameType
	TargetPlayer   ActorUniqueID
	Tick           PlayerInputTick
}

// Marshal reads or writes UpdatePlayerGameType using its canonical wire layout.
func (x *UpdatePlayerGameType) Marshal(io IO) {
	enumValue1 := int32(x.PlayerGameType)
	io.Varint32(&enumValue1)
	x.PlayerGameType = GameType(enumValue1)
	switch int64(enumValue1) {
	case -1, 0, 1, 2, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.TargetPlayer.Marshal(io)
	x.Tick.Marshal(io)
}
