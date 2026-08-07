// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftRecipeOptionalStackRequestAction struct {
	ActionType          ItemStackRequestActionType
	RecipeNetID         RecipeNetID
	FilteredStringIndex int32
}

func (*CraftRecipeOptionalStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftRecipeOptionalStackRequestAction using its canonical wire layout.
func (x *CraftRecipeOptionalStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetID.Marshal(io)
	io.Int32(&x.FilteredStringIndex)
}
