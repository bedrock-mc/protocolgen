// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type WorldClockData struct {
	Id          uint64
	Name        string
	Time        int32
	IsPaused    bool
	TimeMarkers []TimeMarkerData
}

// Marshal reads or writes WorldClockData using its canonical wire layout.
func (x *WorldClockData) Marshal(io IO) {
	io.Varuint64(&x.Id)
	io.String(&x.Name)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
	FuncSlice(io, &x.TimeMarkers, io.Varuint32, func(value *TimeMarkerData) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
