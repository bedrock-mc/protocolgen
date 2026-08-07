// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlaceStackRequestAction struct {
	ActionType  ItemStackRequestActionType
	Amount      uint8
	Source      StackRequestSlotInfo
	Destination StackRequestSlotInfo
}

func (*PlaceStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes PlaceStackRequestAction using its canonical wire layout.
func (x *PlaceStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}
