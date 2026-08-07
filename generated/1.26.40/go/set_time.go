// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetTime struct {
	Time int32
}

// Marshal reads or writes SetTime using its canonical wire layout.
func (x *SetTime) Marshal(io IO) {
	io.Varint32(&x.Time)
}
