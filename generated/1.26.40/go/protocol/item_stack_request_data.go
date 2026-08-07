// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestData struct {
	ClientRequestID       ItemStackRequestID
	Actions               []StackRequestAction
	StringsToFilter       []string
	StringsToFilterOrigin TextProcessingEventOrigin
}

// Marshal reads or writes ItemStackRequestData using its canonical wire layout.
func (x *ItemStackRequestData) Marshal(io IO) {
	x.ClientRequestID.Marshal(io)
	FuncSlice(io, &x.Actions, io.Varuint32, func(value *StackRequestAction) {
		MarshalStackRequestAction(io, value)
	})
	FuncSlice(io, &x.StringsToFilter, io.Varuint32, io.String)
	IntegerFunc(&x.StringsToFilterOrigin, io.Int32)
}
