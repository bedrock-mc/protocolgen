// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealPlaceActionData struct {
	ActionType  ItemStackRequestActionType
	Amount      uint8
	Source      ItemStackRequestCerealSlotInfoData
	Destination ItemStackRequestCerealSlotInfoData
}

func (*ItemStackRequestCerealPlaceActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealPlaceActionData using its canonical wire layout.
func (x *ItemStackRequestCerealPlaceActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.Amount)
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}
