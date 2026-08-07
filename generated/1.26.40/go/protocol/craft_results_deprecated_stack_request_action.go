// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CraftResultsDeprecatedStackRequestAction struct {
	ActionType   ItemStackRequestActionType
	CraftResults []ItemInstance
	NumCrafts    uint8
}

func (*CraftResultsDeprecatedStackRequestAction) isStackRequestAction() {}

// Marshal reads or writes CraftResultsDeprecatedStackRequestAction using its canonical wire layout.
func (x *CraftResultsDeprecatedStackRequestAction) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	Slice(io, &x.CraftResults)
	io.Uint8(&x.NumCrafts)
}
