// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SyncWorldClockStateData struct {
	ClockId  uint64
	Time     int32
	IsPaused bool
}

// Marshal reads or writes SyncWorldClockStateData using its canonical wire layout.
func (x *SyncWorldClockStateData) Marshal(io IO) {
	io.Varuint64(&x.ClockId)
	io.Varint32(&x.Time)
	io.Bool(&x.IsPaused)
}
