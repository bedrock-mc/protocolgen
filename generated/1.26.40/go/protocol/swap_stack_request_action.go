// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SwapStackRequestAction struct {
	ActionType  ItemStackRequestActionType
	Source      StackRequestSlotInfo
	Destination StackRequestSlotInfo
}

func (*SwapStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes SwapStackRequestAction using its canonical wire layout.
func (x *SwapStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}
