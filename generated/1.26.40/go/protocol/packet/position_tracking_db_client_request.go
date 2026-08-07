// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// PositionTrackingDBClientRequest is a packet sent by the client to request the position and
// dimension of a 'tracking ID'. These IDs are tracked in a database by the server. In 1.16, this is
// used for lodestones. The client will send this request to find the position a lodestone compass
// needs to point to. If found, it will point to the lodestone. If not, it will start spinning
// around. A PositionTrackingDBServerBroadcast packet should be sent in response to this packet.
type PositionTrackingDBClientRequest struct {
	Action  protocol.PositionTrackingDBClientRequestAction
	IDValue protocol.PositionTrackingID
}

// Marshal reads or writes PositionTrackingDBClientRequest using its canonical wire layout.
func (x *PositionTrackingDBClientRequest) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Action, io.Uint8)
	x.IDValue.Marshal(io)
}

// ID returns the protocol ID for PositionTrackingDBClientRequest.
func (*PositionTrackingDBClientRequest) ID() uint32 { return IDPositionTrackingDBClientRequest }
