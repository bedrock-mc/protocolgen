// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequestCerealRequestData struct {
	ClientRequestId       TypedClientNetIdStructItemStackRequestIdTagInt32T0
	Actions               []ItemStackRequestCereal
	StringsToFilter       []string
	StringsToFilterOrigin TextProcessingEventOrigin
}

// Marshal reads or writes ItemStackRequestCerealRequestData using its canonical wire layout.
func (x *ItemStackRequestCerealRequestData) Marshal(io IO) {
	x.ClientRequestId.Marshal(io)
	FuncSlice(io, &x.Actions, io.Varuint32, func(value *ItemStackRequestCereal) {
		item := *value
		marshalItemStackRequestCereal(io, &item)
		*value = item
	})
	FuncSlice(io, &x.StringsToFilter, io.Varuint32, io.String)
	IntegerFunc(&x.StringsToFilterOrigin, io.Int32)
}
