// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealLabTableCombineActionData struct {
	ActionType ItemStackRequestActionType
}

func (ItemStackRequestCerealLabTableCombineActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealLabTableCombineActionData using its canonical wire layout.
func (x *ItemStackRequestCerealLabTableCombineActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
}
