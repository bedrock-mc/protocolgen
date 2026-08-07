// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type CreativeItemNetID struct {
	ID uint32
}

// Marshal reads or writes CreativeItemNetID using its canonical wire layout.
func (x *CreativeItemNetID) Marshal(io IO) {
	io.Varuint32(&x.ID)
}
