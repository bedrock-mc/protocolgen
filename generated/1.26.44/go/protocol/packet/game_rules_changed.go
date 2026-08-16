// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.44/go/protocol"

// GameRulesChanged is sent by the server to the client to update client-side game rules, such as
// game rules like the 'showCoordinates' game rule.
type GameRulesChanged struct {
	RuleData protocol.GameRulesChangedData
}

// Marshal reads or writes GameRulesChanged using its canonical wire layout.
func (x *GameRulesChanged) Marshal(io protocol.IO) {
	x.RuleData.Marshal(io)
}

// ID returns the protocol ID for GameRulesChanged.
func (*GameRulesChanged) ID() uint32 { return IDGameRulesChanged }
