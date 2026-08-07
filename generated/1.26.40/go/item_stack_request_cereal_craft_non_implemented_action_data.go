// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealCraftNonImplementedActionData struct {
	ActionType ItemStackRequestActionType
}

func (ItemStackRequestCerealCraftNonImplementedActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftNonImplementedActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftNonImplementedActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
}
