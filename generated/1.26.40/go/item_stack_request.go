// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackRequest struct {
	Requests []ItemStackRequestPacketDataRequestData
}

// Marshal reads or writes ItemStackRequest using its canonical wire layout.
func (x *ItemStackRequest) Marshal(io IO) {
	FuncSlice(io, &x.Requests, io.Varuint32, func(value *ItemStackRequestPacketDataRequestData) {
		value.Marshal(io)
	})
}
