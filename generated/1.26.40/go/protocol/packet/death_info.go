// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// DeathInfo is a packet sent from the server to the client expected to be sent when a player dies.
// It contains messages related to the player's death, which are shown on the death screen as of
// v1.19.10.
type DeathInfo struct {
	// DeathCauseAttackName is the cause of the player's death, such as "suffocation" or "suicide".
	DeathCauseAttackName string
	// DeathCauseMessageList is a list of death messages to be shown on the death screen.
	DeathCauseMessageList []string
}

// Marshal reads or writes DeathInfo using its canonical wire layout.
func (x *DeathInfo) Marshal(io protocol.IO) {
	io.String(&x.DeathCauseAttackName)
	protocol.FuncSlice(io, &x.DeathCauseMessageList, io.Varuint32, io.String)
}

// ID returns the protocol ID for DeathInfo.
func (*DeathInfo) ID() uint32 { return IDDeathInfo }
