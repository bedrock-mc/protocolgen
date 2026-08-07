// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleRuleValueInt32 struct {
	Value int32
}

func (*GameRuleRuleValueInt32) isGameRuleRuleValue() {}

// Marshal reads or writes GameRuleRuleValueInt32 using its canonical wire layout.
func (x *GameRuleRuleValueInt32) Marshal(io IO) {
	io.Int32(&x.Value)
}
