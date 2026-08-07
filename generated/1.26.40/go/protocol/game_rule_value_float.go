// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleValueFloat struct {
	Value float32
}

func (*GameRuleValueFloat) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueFloat using its canonical wire layout.
func (x *GameRuleValueFloat) Marshal(io IO) {
	io.Float32(&x.Value)
}
