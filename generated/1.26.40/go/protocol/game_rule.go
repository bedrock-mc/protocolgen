// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type GameRule struct {
	RuleName          string
	RuleCanBeModified bool
	RuleValue         GameRuleRuleValue
}

// Marshal reads or writes GameRule using its canonical wire layout.
func (x *GameRule) Marshal(io IO) {
	io.String(&x.RuleName)
	io.Bool(&x.RuleCanBeModified)
	MarshalGameRuleRuleValue(io, &x.RuleValue)
}
