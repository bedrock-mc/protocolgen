// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackResponseInfo struct {
	Result          ItemStackNetResult
	ClientRequestId TypedClientNetIdStructItemStackRequestIdTagInt32T0
	Containers      Optional[[]ItemStackResponseContainerInfo]
}

// Marshal reads or writes ItemStackResponseInfo using its canonical wire layout.
func (x *ItemStackResponseInfo) Marshal(io IO) {
	IntegerFunc(&x.Result, io.Uint8)
	x.ClientRequestId.Marshal(io)
	DoubleOptionalFunc(io, &x.Containers, func(value *[]ItemStackResponseContainerInfo) {
		FuncSlice(io, value, io.Varuint32, func(value *ItemStackResponseContainerInfo) {
			value.Marshal(io)
		})
	})
}
