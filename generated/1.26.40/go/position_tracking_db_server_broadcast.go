// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PositionTrackingDBServerBroadcast struct {
	Action               PositionTrackingDBServerBroadcastAction
	Id                   PositionTrackingId
	PositionTrackingData []byte
}

// Marshal reads or writes PositionTrackingDBServerBroadcast using its canonical wire layout.
func (x *PositionTrackingDBServerBroadcast) Marshal(io IO) {
	enumValue1 := uint8(x.Action)
	io.Uint8(&enumValue1)
	x.Action = PositionTrackingDBServerBroadcastAction(enumValue1)
	switch int64(enumValue1) {
	case 0, 1, 2:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.Id.Marshal(io)
	io.NBT(&x.PositionTrackingData)
}
