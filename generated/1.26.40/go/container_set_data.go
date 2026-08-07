// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ContainerSetData struct {
	ContainerID uint8
	ID          int32
	Value       int32
}

// Marshal reads or writes ContainerSetData using its canonical wire layout.
func (x *ContainerSetData) Marshal(io IO) {
	io.Uint8(&x.ContainerID)
	io.Varint32(&x.ID)
	io.Varint32(&x.Value)
}
