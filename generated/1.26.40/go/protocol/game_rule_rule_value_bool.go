// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleRuleValueBool struct {
	Value bool
}

func (*GameRuleRuleValueBool) isGameRuleRuleValue() {}

// Marshal reads or writes GameRuleRuleValueBool using its canonical wire layout.
func (x *GameRuleRuleValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}
