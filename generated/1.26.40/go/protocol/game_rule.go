// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// GameRule contains game rule data.
type GameRule struct {
	// Name is the name of the game rule.
	RuleName string
	// CanBeModifiedByPlayer specifies if the game rule can be modified by the player through the
	// in-game UI.
	RuleCanBeModified bool
	// Value is the new value of the game rule. This is either a bool, uint32 or float32, or nil for the
	// null variant, which carries no value at all.
	RuleValue GameRuleValue
}

// Marshal reads or writes GameRule using its canonical wire layout.
func (x *GameRule) Marshal(io IO) {
	io.String(&x.RuleName)
	io.Bool(&x.RuleCanBeModified)
	MarshalGameRuleValue(io, &x.RuleValue)
}
