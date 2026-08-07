// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleRuleValueFloat struct {
	Value float32
}

func (*GameRuleRuleValueFloat) isGameRuleRuleValue() {}

// Marshal reads or writes GameRuleRuleValueFloat using its canonical wire layout.
func (x *GameRuleRuleValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}
