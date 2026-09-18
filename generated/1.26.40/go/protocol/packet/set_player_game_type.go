// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"
)

// SetPlayerGameType is sent by the server to update the game type, which is otherwise known as the game mode,
// of a player.
type SetPlayerGameType struct {
	PlayerGameType protocol.GameType
}

// ID ...
func (*SetPlayerGameType) ID() uint32 {
	return IDSetPlayerGameType
}

func (pk *SetPlayerGameType) Marshal(io protocol.IO) {
	pk.PlayerGameType.Marshal(io)
}
