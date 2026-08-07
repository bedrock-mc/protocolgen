// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PositionTrackingDBServerBroadcast struct {
	Action               PositionTrackingDBServerBroadcastAction
	Id                   PositionTrackingId
	PositionTrackingData []byte
}

// Marshal reads or writes PositionTrackingDBServerBroadcast using its canonical wire layout.
func (x *PositionTrackingDBServerBroadcast) Marshal(io IO) {
	IntegerFunc(&x.Action, io.Uint8)
	x.Id.Marshal(io)
	io.NBT(&x.PositionTrackingData)
}
