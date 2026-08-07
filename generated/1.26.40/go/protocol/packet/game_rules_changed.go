// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type GameRulesChanged struct {
	RuleData protocol.GameRulesChangedPacketData
}

// Marshal reads or writes GameRulesChanged using its canonical wire layout.
func (x *GameRulesChanged) Marshal(io protocol.IO) {
	x.RuleData.Marshal(io)
}

// ID returns the protocol ID for GameRulesChanged.
func (*GameRulesChanged) ID() uint32 { return IDGameRulesChanged }
