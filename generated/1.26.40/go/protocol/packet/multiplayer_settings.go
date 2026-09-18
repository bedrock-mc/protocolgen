// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

type MultiplayerSettings struct {
	PacketType protocol.MultiplayerSettingsType
}

// ID ...
func (*MultiplayerSettings) ID() uint32 {
	return IDMultiplayerSettings
}

func (pk *MultiplayerSettings) Marshal(io protocol.IO) {
	pk.PacketType.Marshal(io)
}
