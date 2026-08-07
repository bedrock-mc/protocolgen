// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TypedServerNetIdStructItemStackNetIdTagInt32T0 struct {
	ID int32
}

// Marshal reads or writes TypedServerNetIdStructItemStackNetIdTagInt32T0 using its canonical wire layout.
func (x *TypedServerNetIdStructItemStackNetIdTagInt32T0) Marshal(io IO) {
	io.Varint32(&x.ID)
}
