// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MultiplayerSettings struct {
	PacketType MultiplayerSettingsPacketType
}

// Marshal reads or writes MultiplayerSettings using its canonical wire layout.
func (x *MultiplayerSettings) Marshal(io IO) {
	enumValue1 := int32(x.PacketType)
	io.Varint32(&enumValue1)
	x.PacketType = MultiplayerSettingsPacketType(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
}
