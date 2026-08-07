// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealCraftRecipeOptionalActionData struct {
	ActionType          ItemStackRequestActionType
	RecipeNetId         TypedServerNetIdStructRecipeNetIdTag
	FilteredStringIndex int32
}

func (*ItemStackRequestCerealCraftRecipeOptionalActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftRecipeOptionalActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftRecipeOptionalActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	x.RecipeNetId.Marshal(io)
	io.Int32(&x.FilteredStringIndex)
}
