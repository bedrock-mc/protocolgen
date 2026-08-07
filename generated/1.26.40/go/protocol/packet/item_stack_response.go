// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ItemStackResponse struct {
	Responses []protocol.ItemStackResponseInfo
}

// Marshal reads or writes ItemStackResponse using its canonical wire layout.
func (x *ItemStackResponse) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &x.Responses, io.Varuint32, func(value *protocol.ItemStackResponseInfo) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for ItemStackResponse.
func (*ItemStackResponse) ID() uint32 { return IDItemStackResponse }
