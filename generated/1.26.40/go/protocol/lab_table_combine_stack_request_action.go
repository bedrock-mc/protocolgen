// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LabTableCombineStackRequestAction struct {
	ActionType ItemStackRequestActionType
}

func (*LabTableCombineStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes LabTableCombineStackRequestAction using its canonical wire layout.
func (x *LabTableCombineStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
}
