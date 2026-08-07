// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SimpleEvent struct {
	Type SimpleEventSubtype
}

// Marshal reads or writes SimpleEvent using its canonical wire layout.
func (x *SimpleEvent) Marshal(io IO) {
	IntegerFunc(&x.Type, io.Uint16)
}
