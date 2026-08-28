// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// SyncWorldClockStateData represents the state data for synchronising a world clock.
type SyncWorldClockStateData struct {
	// ClockID is the unique identifier for the clock.
	ClockID uint64
	// Time is the current time of the clock.
	Time int32
	// IsPaused indicates if the clock is paused.
	IsPaused bool
}

// Marshal reads or writes SyncWorldClockStateData using its canonical wire layout.
func (x *SyncWorldClockStateData) Marshal(io IO) {
	io.Varuint64(&x.ClockID)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
}

type SyncWorldClocksData interface {
	isSyncWorldClocksData()
}

// MarshalSyncWorldClocksData reads or writes the SyncWorldClocksData union using its canonical wire layout.
func MarshalSyncWorldClocksData(io IO, x *SyncWorldClocksData) {
	UnionFunc(io,
		func() {
			var tag uint32
			io.Varuint32(&tag)
			switch int64(tag) {
			case 0:
				value := new(SyncStateData)
				value.Marshal(io)
				*x = value
			case 1:
				value := new(InitializeRegistryData)
				value.Marshal(io)
				*x = value
			case 2:
				value := new(AddTimeMarkerData)
				value.Marshal(io)
				*x = value
			case 3:
				value := new(RemoveTimeMarkerData)
				value.Marshal(io)
				*x = value
			default:
				io.InvalidValue(tag, "unknown union tag")
			}
		},
		func() {
			switch value := (*x).(type) {
			case *SyncStateData:
				tag := uint32(0)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *InitializeRegistryData:
				tag := uint32(1)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *AddTimeMarkerData:
				tag := uint32(2)
				io.Varuint32(&tag)
				value.Marshal(io)
			case *RemoveTimeMarkerData:
				tag := uint32(3)
				io.Varuint32(&tag)
				value.Marshal(io)
			default:
				io.InvalidValue(*x, "unknown union value")
			}
		},
	)
}

// TimeMarkerData represents a time marker within a world clock.
type TimeMarkerData struct {
	// ID is the unique identifier for the time marker.
	ID uint64
	// Name is the name of the time marker.
	Name string
	// Time is the time at which the marker is set.
	Time int32
	// Period is the optional period for the time marker.
	Period Optional[int32]
}

// Marshal reads or writes TimeMarkerData using its canonical wire layout.
func (x *TimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.ID)
	io.StringLimits(&x.Name, 0, 128)
	io.Varint32(&x.Time)
	OptionalFunc(io, &x.Period, io.Int32)
}

// WorldClockData represents a complete world clock with its time markers.
type WorldClockData struct {
	// ID is the unique identifier for the clock.
	ID uint64
	// Name is the name of the clock.
	Name string
	// Time is the current time of the clock.
	Time int32
	// IsPaused indicates if the clock is paused.
	IsPaused bool
	// TimeMarkers is a list of time markers for this clock.
	TimeMarkers []TimeMarkerData
}

// Marshal reads or writes WorldClockData using its canonical wire layout.
func (x *WorldClockData) Marshal(io IO) {
	io.Varuint64(&x.ID)
	io.StringLimits(&x.Name, 0, 128)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
	SliceLimits(io, &x.TimeMarkers, 0, 256)
}
