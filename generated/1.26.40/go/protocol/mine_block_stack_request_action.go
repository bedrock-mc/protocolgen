// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type MineBlockStackRequestAction struct {
	ActionType          ItemStackRequestActionType
	Slot                int32
	PredictedDurability int32
	NetIDVariant        int32
}

func (*MineBlockStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes MineBlockStackRequestAction using its canonical wire layout.
func (x *MineBlockStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Varint32(&x.Slot)
	io.Varint32(&x.PredictedDurability)
	io.Int32(&x.NetIDVariant)
}
