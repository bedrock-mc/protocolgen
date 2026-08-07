// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ItemStackResponse struct {
	Responses []ItemStackResponseInfo
}

// Marshal reads or writes ItemStackResponse using its canonical wire layout.
func (x *ItemStackResponse) Marshal(io IO) {
	FuncSlice(io, &x.Responses, io.Varuint32, func(value *ItemStackResponseInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
