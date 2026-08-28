// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.50/go/protocol"

// SetScore is sent by the server to send the contents of a scoreboard to the player. It may be used
// to either add, remove or edit entries on the scoreboard.
type SetScore struct {
	// ScoreInfo is a list of all entries that the client should operate on. Each entry's IdentityType
	// specifies whether it is added, modified or removed.
	ScoreInfo []protocol.SetScoreInfoItem
}

// Marshal reads or writes SetScore using its canonical wire layout.
func (x *SetScore) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.ScoreInfo, io.Varuint32, func(value *protocol.SetScoreInfoItem) {
		protocol.MarshalSetScoreInfoItem(io, value)
	})
}

// ID returns the protocol ID for SetScore.
func (*SetScore) ID() uint32 { return IDSetScore }
