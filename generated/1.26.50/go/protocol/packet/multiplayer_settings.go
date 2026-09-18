// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

type MultiplayerSettings struct {
	PacketType protocol.MultiplayerSettingsType
}

// ID returns the protocol ID for MultiplayerSettings.
func (*MultiplayerSettings) ID() uint32 { return IDMultiplayerSettings }

// Marshal reads or writes MultiplayerSettings using its canonical wire layout.
func (pk *MultiplayerSettings) Marshal(io protocol.IO) {
	pk.PacketType.Marshal(io)
}
