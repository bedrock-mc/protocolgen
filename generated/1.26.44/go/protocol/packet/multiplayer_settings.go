// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

type MultiplayerSettings struct {
	PacketType protocol.MultiplayerSettingsType
}

// Marshal reads or writes MultiplayerSettings using its canonical wire layout.
func (x *MultiplayerSettings) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.PacketType, io.Varint32)
}

// ID returns the protocol ID for MultiplayerSettings.
func (*MultiplayerSettings) ID() uint32 { return IDMultiplayerSettings }
