// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealDropActionData struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestCerealSlotInfoData
	Randomly   bool
}

func (*ItemStackRequestCerealDropActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealDropActionData using its canonical wire layout.
func (x *ItemStackRequestCerealDropActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	io.Bool(&x.Randomly)
}
