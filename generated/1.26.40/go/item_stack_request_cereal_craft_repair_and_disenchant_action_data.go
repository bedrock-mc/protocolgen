// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealCraftRepairAndDisenchantActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             int32
	NumberOfRequestedCrafts uint8
	RepairCost              int32
}

func (ItemStackRequestCerealCraftRepairAndDisenchantActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftRepairAndDisenchantActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRepairAndDisenchantActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Int32(&x.RecipeNetId)
	io.Uint8(&x.NumberOfRequestedCrafts)
	io.Varint32(&x.RepairCost)
}
