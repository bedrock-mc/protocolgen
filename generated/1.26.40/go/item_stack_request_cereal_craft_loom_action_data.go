// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealCraftLoomActionData struct {
	ActionType    ItemStackRequestActionType
	PatternNameId string
	NumCrafts     uint8
}

func (ItemStackRequestCerealCraftLoomActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftLoomActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftLoomActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.String(&x.PatternNameId)
	io.Uint8(&x.NumCrafts)
}
