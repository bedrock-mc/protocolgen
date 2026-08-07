// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealCraftCreativeActionData struct {
	ActionType              ItemStackRequestActionType
	CreativeItemNetId       uint32
	NumberOfRequestedCrafts uint8
}

func (ItemStackRequestCerealCraftCreativeActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftCreativeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftCreativeActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Varuint32(&x.CreativeItemNetId)
	io.Uint8(&x.NumberOfRequestedCrafts)
}
