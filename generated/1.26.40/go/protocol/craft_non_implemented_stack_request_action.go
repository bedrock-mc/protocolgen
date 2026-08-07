// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftNonImplementedStackRequestAction struct {
	ActionType ItemStackRequestActionType
}

func (*CraftNonImplementedStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftNonImplementedStackRequestAction using its canonical wire layout.
func (x *CraftNonImplementedStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
}
