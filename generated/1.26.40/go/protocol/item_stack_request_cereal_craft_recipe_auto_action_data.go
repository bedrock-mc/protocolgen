// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealCraftRecipeAutoActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             TypedServerNetIdStructRecipeNetIdTag
	NumberOfRequestedCrafts uint8
	Ingredients             []ItemStackRequestCerealRecipeIngredientData
}

func (*ItemStackRequestCerealCraftRecipeAutoActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeAutoActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeAutoActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetId.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
	Slice(io, &x.Ingredients)
}
