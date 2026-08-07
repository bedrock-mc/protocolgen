// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftCreativeStackRequestAction struct {
	ActionType              ItemStackRequestActionType
	CreativeItemNetID       uint32
	NumberOfRequestedCrafts uint8
}

func (*CraftCreativeStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftCreativeStackRequestAction using its canonical wire layout.
func (x *CraftCreativeStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Varuint32(&x.CreativeItemNetID)
	io.Uint8(&x.NumberOfRequestedCrafts)
}
