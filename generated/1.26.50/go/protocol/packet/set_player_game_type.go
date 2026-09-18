// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// SetPlayerGameType is sent by the server to update the game type, which is otherwise known as the game mode,
// of a player.
type SetPlayerGameType struct {
	PlayerGameType protocol.GameType
}

// ID returns the protocol ID for SetPlayerGameType.
func (*SetPlayerGameType) ID() uint32 { return IDSetPlayerGameType }

// Marshal reads or writes SetPlayerGameType using its canonical wire layout.
func (pk *SetPlayerGameType) Marshal(io protocol.IO) {
	pk.PlayerGameType.Marshal(io)
}
