// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AutoCraftRecipeStackRequestAction struct {
	ActionType              ItemStackRequestActionType
	RecipeNetID             RecipeNetID
	NumberOfRequestedCrafts uint8
	Ingredients             []RecipeIngredient
}

func (*AutoCraftRecipeStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes AutoCraftRecipeStackRequestAction using its canonical wire layout.
func (x *AutoCraftRecipeStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetID.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
	Slice(io, &x.Ingredients)
}
