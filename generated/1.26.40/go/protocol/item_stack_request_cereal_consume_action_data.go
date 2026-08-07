// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealConsumeActionData struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestCerealSlotInfoData
}

func (*ItemStackRequestCerealConsumeActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealConsumeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealConsumeActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
}
