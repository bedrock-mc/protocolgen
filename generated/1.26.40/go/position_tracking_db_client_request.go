// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PositionTrackingDBClientRequest struct {
	Action PositionTrackingDBClientRequestAction
	Id     PositionTrackingId
}

// Marshal reads or writes PositionTrackingDBClientRequest using its canonical wire layout.
func (x *PositionTrackingDBClientRequest) Marshal(io IO) {
	IntegerFunc(&x.Action, io.Uint8)
	x.Id.Marshal(io)
}
