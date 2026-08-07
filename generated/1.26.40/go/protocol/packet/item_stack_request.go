// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ItemStackRequest struct {
	Requests []protocol.ItemStackRequestPacketDataRequestData
}

// Marshal reads or writes ItemStackRequest using its canonical wire layout.
func (x *ItemStackRequest) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Requests)
}

// ID returns the protocol ID for ItemStackRequest.
func (*ItemStackRequest) ID() uint32 { return IDItemStackRequest }
