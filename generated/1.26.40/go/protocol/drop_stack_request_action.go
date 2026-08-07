// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type DropStackRequestAction struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     StackRequestSlotInfo
	Randomly   bool
}

func (*DropStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes DropStackRequestAction using its canonical wire layout.
func (x *DropStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	io.Bool(&x.Randomly)
}
