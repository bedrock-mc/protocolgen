// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type MultiplayerSettings struct {
	PacketType MultiplayerSettingsPacketType
}

// Marshal reads or writes MultiplayerSettings using its canonical wire layout.
func (x *MultiplayerSettings) Marshal(io IO) {
	IntegerFunc(&x.PacketType, io.Varint32)
}
