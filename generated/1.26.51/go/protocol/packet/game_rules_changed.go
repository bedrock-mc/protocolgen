// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.51/go/protocol"
)

// GameRulesChanged is sent by the server to the client to update client-side game rules, such as game rules
// like the 'showCoordinates' game rule.
type GameRulesChanged struct {
	RuleData protocol.GameRulesChangedData
}

// ID ...
func (*GameRulesChanged) ID() uint32 {
	return IDGameRulesChanged
}

func (pk *GameRulesChanged) Marshal(io protocol.IO) {
	pk.RuleData.Marshal(io)
}
