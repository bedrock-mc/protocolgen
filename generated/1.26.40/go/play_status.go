// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type PlayStatus struct {
	Status PlayStatusType
}

// Marshal reads or writes PlayStatus using its canonical wire layout.
func (x *PlayStatus) Marshal(io IO) {
	IntegerFunc(&x.Status, io.BEInt32)
}
