// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRuleValueInt32 struct {
	Value int32
}

func (*GameRuleValueInt32) isGameRuleValue() {}

// Marshal reads or writes GameRuleValueInt32 using its canonical wire layout.
func (x *GameRuleValueInt32) Marshal(io IO) {
	io.Int32(&x.Value)
}
