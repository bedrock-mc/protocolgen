// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RemoveTimeMarkerData struct {
	ClockID       uint64
	TimeMarkerIds []uint64
}

func (*RemoveTimeMarkerData) isSyncWorldClocksData() {}

// Marshal reads or writes RemoveTimeMarkerData using its canonical wire layout.
func (x *RemoveTimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ClockID)
	FuncSlice(io, &x.TimeMarkerIds, io.Varuint32, io.Varuint64)
}
