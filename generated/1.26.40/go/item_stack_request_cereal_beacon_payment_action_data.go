// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealBeaconPaymentActionData struct {
	ActionType        ItemStackRequestActionType
	PrimaryEffectId   int32
	SecondaryEffectId int32
}

func (ItemStackRequestCerealBeaconPaymentActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealBeaconPaymentActionData using its canonical wire layout.
func (x *ItemStackRequestCerealBeaconPaymentActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Varint32(&x.PrimaryEffectId)
	io.Varint32(&x.SecondaryEffectId)
}
