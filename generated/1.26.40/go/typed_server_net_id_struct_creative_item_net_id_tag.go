// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type TypedServerNetIdStructCreativeItemNetIdTag struct {
	ID uint32
}

// Marshal reads or writes TypedServerNetIdStructCreativeItemNetIdTag using its canonical wire layout.
func (x *TypedServerNetIdStructCreativeItemNetIdTag) Marshal(io IO) {
	io.Varuint32(&x.ID)
}
