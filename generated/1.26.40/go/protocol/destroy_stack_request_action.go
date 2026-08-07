// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DestroyStackRequestAction struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     StackRequestSlotInfo
}

func (*DestroyStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes DestroyStackRequestAction using its canonical wire layout.
func (x *DestroyStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
}
