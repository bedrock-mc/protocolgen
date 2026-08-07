// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type GameRulesChanged struct {
	RuleData GameRulesChangedPacketData
}

// Marshal reads or writes GameRulesChanged using its canonical wire layout.
func (x *GameRulesChanged) Marshal(io IO) {
	x.RuleData.Marshal(io)
}
