// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ItemStackRequest is sent by the client to change item stacks in an inventory. It is essentially a
// replacement of the InventoryTransaction packet added in 1.16 for inventory specific actions, such
// as moving items around or crafting. The InventoryTransaction packet is still used for actions
// such as placing blocks and interacting with entities.
type ItemStackRequest struct {
	// Requests holds a list of item stack requests. These requests are all separate, but the client
	// buffers the requests, so you might find multiple unrelated requests in this packet.
	Requests []protocol.ItemStackRequestPacketData
}

// Marshal reads or writes ItemStackRequest using its canonical wire layout.
func (x *ItemStackRequest) Marshal(io protocol.IO) {
	protocol.Slice(io, &x.Requests)
}

// ID returns the protocol ID for ItemStackRequest.
func (*ItemStackRequest) ID() uint32 { return IDItemStackRequest }
