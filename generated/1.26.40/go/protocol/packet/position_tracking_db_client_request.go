// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

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
