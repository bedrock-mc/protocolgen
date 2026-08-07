// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRulesChangedData struct {
	RulesList []GameRule
}

// Marshal reads or writes GameRulesChangedData using its canonical wire layout.
func (x *GameRulesChangedData) Marshal(io IO) {
	Slice(io, &x.RulesList)
}
