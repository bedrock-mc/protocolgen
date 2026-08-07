// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealSwapActionData struct {
	ActionType  ItemStackRequestActionType
	Source      ItemStackRequestCerealSlotInfoData
	Destination ItemStackRequestCerealSlotInfoData
}

func (*ItemStackRequestCerealSwapActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealSwapActionData using its canonical wire layout.
func (x *ItemStackRequestCerealSwapActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.Source.Marshal(io)
	x.Destination.Marshal(io)
}
