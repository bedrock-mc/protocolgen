// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ConsumeStackRequestAction struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     StackRequestSlotInfo
}

func (*ConsumeStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes ConsumeStackRequestAction using its canonical wire layout.
func (x *ConsumeStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
}
