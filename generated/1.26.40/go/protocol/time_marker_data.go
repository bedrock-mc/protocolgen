// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TimeMarkerData struct {
	Id     uint64
	Name   string
	Time   int32
	Period Optional[int32]
}

// Marshal reads or writes TimeMarkerData using its canonical wire layout.
func (x *TimeMarkerData) Marshal(io IO) {
	io.Varuint64(&x.Id)
	io.String(&x.Name)
	io.Varint32(&x.Time)
	OptionalFunc(io, &x.Period, io.Int32)
}
