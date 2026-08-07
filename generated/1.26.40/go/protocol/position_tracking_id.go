// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PositionTrackingId struct {
	Value int32
}

// Marshal reads or writes PositionTrackingId using its canonical wire layout.
func (x *PositionTrackingId) Marshal(io IO) {
	io.Varint32(&x.Value)
}
