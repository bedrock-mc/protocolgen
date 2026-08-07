// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type ItemStackRequestID struct {
	ID int32
}

// Marshal reads or writes ItemStackRequestID using its canonical wire layout.
func (x *ItemStackRequestID) Marshal(io IO) {
	io.Varint32(&x.ID)
}
