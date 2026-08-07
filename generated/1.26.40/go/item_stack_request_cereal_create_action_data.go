// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealCreateActionData struct {
	ActionType   ItemStackRequestActionType
	ResultsIndex uint8
}

func (ItemStackRequestCerealCreateActionData) isItemStackRequestCereal() {}

// Marshal reads or writes ItemStackRequestCerealCreateActionData using its canonical wire layout.
func (x *ItemStackRequestCerealCreateActionData) Marshal(io IO) {
	IntegerFunc(&x.ActionType, io.Uint8)
	io.Uint8(&x.ResultsIndex)
}
