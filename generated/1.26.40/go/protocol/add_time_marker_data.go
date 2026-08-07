// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type AddTimeMarkerData struct {
	ClockID     uint64
	TimeMarkers []TimeMarkerData
}

func (*AddTimeMarkerData) isSyncWorldClocksData() {}

// Marshal reads or writes AddTimeMarkerData using its canonical wire layout.
func (x *AddTimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ClockID)
	Slice(io, &x.TimeMarkers)
}
