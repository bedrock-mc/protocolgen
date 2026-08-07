// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackResponseInfo struct {
	Result          ItemStackNetResult
	ClientRequestID ItemStackRequestID
	Containers      Optional[[]ItemStackResponseContainerInfo]
}

// Marshal reads or writes ItemStackResponseInfo using its canonical wire layout.
func (x *ItemStackResponseInfo) Marshal(io IO) {
	IntegerFunc(&x.Result, io.Uint8)
	x.ClientRequestID.Marshal(io)
	DoubleOptionalFunc(io, &x.Containers, func(value *[]ItemStackResponseContainerInfo) {
		Slice(io, value)
	})
}
