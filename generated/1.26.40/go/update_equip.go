// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type UpdateEquip struct {
	ContainerId    uint8
	Type           uint8
	Size           int32
	EntityUniqueId int64
	Data           []byte
}

// Marshal reads or writes UpdateEquip using its canonical wire layout.
func (x *UpdateEquip) Marshal(io IO) {
	io.Uint8(&x.ContainerId)
	io.Uint8(&x.Type)
	io.Varint32(&x.Size)
	io.ActorUniqueID(&x.EntityUniqueId)
	io.NBT(&x.Data)
}
