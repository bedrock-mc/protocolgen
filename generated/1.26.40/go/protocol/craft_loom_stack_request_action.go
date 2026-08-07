// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftLoomStackRequestAction struct {
	ActionType    ItemStackRequestActionType
	PatternNameID string
	NumCrafts     uint8
}

func (*CraftLoomStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftLoomStackRequestAction using its canonical wire layout.
func (x *CraftLoomStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.String(&x.PatternNameID)
	io.Uint8(&x.NumCrafts)
}
