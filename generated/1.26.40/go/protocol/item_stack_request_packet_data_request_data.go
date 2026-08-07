// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestPacketDataRequestData struct {
	ClientRequestId       TypedClientNetIdStructItemStackRequestIdTagInt32T0
	Actions               []ItemStackRequestCereal
	StringsToFilter       []string
	StringsToFilterOrigin TextProcessingEventOrigin
}

// Marshal reads or writes ItemStackRequestPacketDataRequestData using its canonical wire layout.
func (x *ItemStackRequestPacketDataRequestData) Marshal(io IO) {
	x.ClientRequestId.Marshal(io)
	FuncSlice(io, &x.Actions, io.Varuint32, func(value *ItemStackRequestCereal) {
		MarshalItemStackRequestCereal(io, value)
	})
	FuncSlice(io, &x.StringsToFilter, io.Varuint32, io.String)
	IntegerFunc(&x.StringsToFilterOrigin, io.Int32)
}
