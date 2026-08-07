// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type BookEdit struct {
	BookSlot  int32
	Operation BookEditAction
}

// Marshal reads or writes BookEdit using its canonical wire layout.
func (x *BookEdit) Marshal(io IO) {
	io.Varint32(&x.BookSlot)
	marshalBookEditAction(io, &x.Operation)
}
