// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetDefaultGameType struct {
	DefaultGameType GameType
}

// Marshal reads or writes SetDefaultGameType using its canonical wire layout.
func (x *SetDefaultGameType) Marshal(io IO) {
	enumValue1 := int32(x.DefaultGameType)
	io.Varint32(&enumValue1)
	x.DefaultGameType = GameType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
