// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PositionTrackingID struct {
	Value int32
}

// Marshal reads or writes PositionTrackingID using its canonical wire layout.
func (x *PositionTrackingID) Marshal(io IO) {
	io.Varint32(&x.Value)
}
