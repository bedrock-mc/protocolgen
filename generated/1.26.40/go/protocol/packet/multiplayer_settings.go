// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type MultiplayerSettings struct {
	PacketType protocol.MultiplayerSettingsType
}

// Marshal reads or writes MultiplayerSettings using its canonical wire layout.
func (x *MultiplayerSettings) Marshal(io protocol.IO) {
	x.PacketType.Marshal(io)
}

// ID returns the protocol ID for MultiplayerSettings.
func (*MultiplayerSettings) ID() uint32 { return IDMultiplayerSettings }
