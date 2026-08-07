// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type TypedClientNetIdStructItemStackRequestIdTagInt32T0 struct {
	ID int32
}

// Marshal reads or writes TypedClientNetIdStructItemStackRequestIdTagInt32T0 using its canonical wire layout.
func (x *TypedClientNetIdStructItemStackRequestIdTagInt32T0) Marshal(io IO) {
	io.Varint32(&x.ID)
}
