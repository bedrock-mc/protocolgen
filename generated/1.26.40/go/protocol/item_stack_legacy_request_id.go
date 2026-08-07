// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackLegacyRequestID struct {
	ID int32
}

// Marshal reads or writes ItemStackLegacyRequestID using its canonical wire layout.
func (x *ItemStackLegacyRequestID) Marshal(io IO) {
	io.Varint32(&x.ID)
}
