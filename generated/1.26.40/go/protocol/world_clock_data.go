// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type WorldClockData struct {
	ID          uint64
	Name        string
	Time        int32
	IsPaused    bool
	TimeMarkers []TimeMarkerData
}

// Marshal reads or writes WorldClockData using its canonical wire layout.
func (x *WorldClockData) Marshal(io IO) {
	io.Varuint64(&x.ID)
	io.String(&x.Name)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
	Slice(io, &x.TimeMarkers)
}
