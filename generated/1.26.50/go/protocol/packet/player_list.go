// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// PlayerList is sent by the server to update the client-side player list in the in-game menu
// screen. It shows the icon of each player if the correct XUID is written in the packet. Sending
// the PlayerList packet is obligatory when sending an AddPlayer packet. The added player will not
// show up to a client if it has not been added to the player list, because several properties of
// the player are obtained from the player list, such as the skin.
type PlayerList struct {
	// Entries is a list of all player list entries that should be added/removed from the player list,
	// depending on the ActionType set.
	Entries []protocol.PlayerListData
}

// Marshal reads or writes PlayerList using its canonical wire layout.
func (x *PlayerList) Marshal(io protocol.IO) {
	protocol.FuncSliceLimits(io, &x.Entries, io.Varuint32, 0, 1000, func(value *protocol.PlayerListData) {
		protocol.MarshalPlayerListData(io, value)
	})
}

// ID returns the protocol ID for PlayerList.
func (*PlayerList) ID() uint32 { return IDPlayerList }
