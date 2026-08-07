// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackNetID struct {
	ID int32
}

// Marshal reads or writes ItemStackNetID using its canonical wire layout.
func (x *ItemStackNetID) Marshal(io IO) {
	io.Varint32(&x.ID)
}
