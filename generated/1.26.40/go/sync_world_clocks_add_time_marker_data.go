// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SyncWorldClocksAddTimeMarkerData struct {
	ClockId     uint64
	TimeMarkers []TimeMarkerData
}

func (SyncWorldClocksAddTimeMarkerData) isSyncWorldClocksData() {}

// Marshal reads or writes SyncWorldClocksAddTimeMarkerData using its canonical wire layout.
func (x *SyncWorldClocksAddTimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ClockId)
	FuncSlice(io, &x.TimeMarkers, io.Varuint32, func(value *TimeMarkerData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
