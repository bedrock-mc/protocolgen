// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRulesChangedPacketData struct {
	RulesList []GameRule
}

// Marshal reads or writes GameRulesChangedPacketData using its canonical wire layout.
func (x *GameRulesChangedPacketData) Marshal(io IO) {
	Slice(io, &x.RulesList)
}
