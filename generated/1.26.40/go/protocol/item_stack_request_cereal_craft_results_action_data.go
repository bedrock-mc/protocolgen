// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestCerealCraftResultsActionData struct {
	ActionType   ItemStackRequestActionType
	CraftResults []ItemStackRequestCerealNetworkItemInstanceDescriptorData
	NumCrafts    uint8
}

func (*ItemStackRequestCerealCraftResultsActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCraftResultsActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCraftResultsActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	Slice(io, &x.CraftResults)
	io.Uint8(&x.NumCrafts)
}
