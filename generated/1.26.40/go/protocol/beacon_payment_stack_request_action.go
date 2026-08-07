// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type BeaconPaymentStackRequestAction struct {
	ActionType        ItemStackRequestActionType
	PrimaryEffectID   int32
	SecondaryEffectID int32
}

func (*BeaconPaymentStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes BeaconPaymentStackRequestAction using its canonical wire layout.
func (x *BeaconPaymentStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Varint32(&x.PrimaryEffectID)
	io.Varint32(&x.SecondaryEffectID)
}
