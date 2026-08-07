// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type SetHealth struct {
	Health int32
}

// Marshal reads or writes SetHealth using its canonical wire layout.
func (x *SetHealth) Marshal(io IO) {
	io.Varint32(&x.Health)
}
