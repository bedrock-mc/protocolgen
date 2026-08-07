// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PositionTrackingDBClientRequest struct {
	Action PositionTrackingDBClientRequestAction
	Id     PositionTrackingId
}

// Marshal reads or writes PositionTrackingDBClientRequest using its canonical wire layout.
func (x *PositionTrackingDBClientRequest) Marshal(io IO) {
	enumValue1 := uint8(x.Action)
	io.Uint8(&enumValue1)
	x.Action = PositionTrackingDBClientRequestAction(enumValue1)
	switch int64(enumValue1) {
	case 0:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.Id.Marshal(io)
}
