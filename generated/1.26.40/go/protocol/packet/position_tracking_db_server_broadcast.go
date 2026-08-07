// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type PositionTrackingDBServerBroadcast struct {
	Action               protocol.PositionTrackingDBServerBroadcastAction
	Id                   protocol.PositionTrackingId
	PositionTrackingData []byte
}

// Marshal reads or writes PositionTrackingDBServerBroadcast using its canonical wire layout.
func (x *PositionTrackingDBServerBroadcast) Marshal(io protocol.IO) {
	protocol.IntegerFunc(&x.Action, io.Uint8)
	x.Id.Marshal(io)
	io.NBT(&x.PositionTrackingData)
}
