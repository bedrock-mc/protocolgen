// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type EntityNetId struct {
	RawId uint32
}

// Marshal reads or writes EntityNetId using its canonical wire layout.
func (x *EntityNetId) Marshal(io IO) {
	io.Varuint32(&x.RawId)
}
