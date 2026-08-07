// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftRepairAndDisenchantStackRequestAction struct {
	ActionType              ItemStackRequestActionType
	RecipeNetID             int32
	NumberOfRequestedCrafts uint8
	RepairCost              int32
}

func (*CraftRepairAndDisenchantStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftRepairAndDisenchantStackRequestAction using its canonical wire layout.
func (x *CraftRepairAndDisenchantStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Int32(&x.RecipeNetID)
	io.Uint8(&x.NumberOfRequestedCrafts)
	io.Varint32(&x.RepairCost)
}
