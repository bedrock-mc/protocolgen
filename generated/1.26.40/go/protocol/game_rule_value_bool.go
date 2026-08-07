// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleValueBool struct {
	Value bool
}

func (*GameRuleValueBool) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueBool using its canonical wire layout.
func (x *GameRuleValueBool) Marshal(io IO) {
	io.Bool(&x.Value)
}
