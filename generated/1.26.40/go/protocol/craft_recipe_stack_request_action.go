// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftRecipeStackRequestAction struct {
	ActionType              ItemStackRequestActionType
	RecipeNetID             RecipeNetID
	NumberOfRequestedCrafts uint8
}

func (*CraftRecipeStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftRecipeStackRequestAction using its canonical wire layout.
func (x *CraftRecipeStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetID.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
}
