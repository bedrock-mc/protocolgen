// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type EntityNetID struct {
	RawID uint32
}

// Marshal reads or writes EntityNetID using its canonical wire layout.
func (x *EntityNetID) Marshal(io IO) {
	io.Varuint32(&x.RawID)
}
