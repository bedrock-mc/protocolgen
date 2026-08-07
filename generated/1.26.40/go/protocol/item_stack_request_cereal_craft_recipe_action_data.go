// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealCraftRecipeActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             TypedServerNetIdStructRecipeNetIdTag
	NumberOfRequestedCrafts uint8
}

func (*ItemStackRequestCerealCraftRecipeActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetId.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
}
