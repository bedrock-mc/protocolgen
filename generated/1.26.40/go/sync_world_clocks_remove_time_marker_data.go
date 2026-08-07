// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SyncWorldClocksRemoveTimeMarkerData struct {
	ClockId       uint64
	TimeMarkerIds []uint64
}

func (SyncWorldClocksRemoveTimeMarkerData) isSyncWorldClocksData() {}

// Marshal reads or writes SyncWorldClocksRemoveTimeMarkerData using its canonical wire layout.
func (x *SyncWorldClocksRemoveTimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ClockId)
	FuncSlice(io, &x.TimeMarkerIds, io.Varuint32, io.Varuint64)
}
