// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealDestroyActionData struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestCerealSlotInfoData
}

func (ItemStackRequestCerealDestroyActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealDestroyActionData using its canonical wire layout.
func (x *ItemStackRequestCerealDestroyActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
}
