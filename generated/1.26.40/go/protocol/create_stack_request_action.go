// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CreateStackRequestAction struct {
	ActionType   ItemStackRequestActionType
	ResultsIndex uint8
}

func (*CreateStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CreateStackRequestAction using its canonical wire layout.
func (x *CreateStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.ResultsIndex)
}
