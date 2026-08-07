// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealMineBlockActionData struct {
	ActionType          ItemStackRequestActionType
	Slot                int32
	PredictedDurability int32
	NetIdVariant        int32
}

func (ItemStackRequestCerealMineBlockActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealMineBlockActionData using its canonical wire layout.
func (x *ItemStackRequestCerealMineBlockActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Varint32(&x.Slot)
	io.Varint32(&x.PredictedDurability)
	io.Int32(&x.NetIdVariant)
}
