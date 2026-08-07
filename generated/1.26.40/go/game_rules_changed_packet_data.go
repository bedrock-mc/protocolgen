// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type GameRulesChangedPacketData struct {
	RulesList []GameRule
}

// Marshal reads or writes GameRulesChangedPacketData using its canonical wire layout.
func (x *GameRulesChangedPacketData) Marshal(io IO) {
	FuncSlice(io, &x.RulesList, io.Varuint32, func(value *GameRule) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
