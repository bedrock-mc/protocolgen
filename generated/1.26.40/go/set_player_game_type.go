// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetPlayerGameType struct {
	PlayerGameType GameType
}

// Marshal reads or writes SetPlayerGameType using its canonical wire layout.
func (x *SetPlayerGameType) Marshal(io IO) {
	enumValue1 := int32(x.PlayerGameType)
	io.Varint32(&enumValue1)
	x.PlayerGameType = GameType(enumValue1)
	switch int64(enumValue1) {
	case -1, 0, 1, 2, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
