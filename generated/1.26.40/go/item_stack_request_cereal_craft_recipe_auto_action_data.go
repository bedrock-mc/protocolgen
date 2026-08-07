// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealCraftRecipeAutoActionData struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             TypedServerNetIdStructRecipeNetIdTag
	NumberOfRequestedCrafts uint8
	Ingredients             []ItemStackRequestCerealRecipeIngredientData
}

func (ItemStackRequestCerealCraftRecipeAutoActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeAutoActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeAutoActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetId.Marshal(io)
	io.Uint8(&x.NumberOfRequestedCrafts)
	FuncSlice(io, &x.Ingredients, io.Varuint32, func(value *ItemStackRequestCerealRecipeIngredientData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
