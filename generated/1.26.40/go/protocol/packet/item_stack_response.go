// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ItemStackResponse is sent by the server in response to an ItemStackRequest packet from the
// client. This packet is used to either approve or reject ItemStackRequests from the client. If a
// request is approved, the client will simply continue as normal. If rejected, the client will undo
// the actions so that the inventory should be in sync with the server again.
type ItemStackResponse struct {
	// Responses is a list of responses to ItemStackRequests sent by the client before. Responses either
	// approve or reject a request from the client. Vanilla limits the size of this slice to 4096.
	Responses []protocol.ItemStackResponseInfo
}

// Marshal reads or writes ItemStackResponse using its canonical wire layout.
func (x *ItemStackResponse) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Responses)
}

// ID returns the protocol ID for ItemStackResponse.
func (*ItemStackResponse) ID() uint32 { return IDItemStackResponse }
