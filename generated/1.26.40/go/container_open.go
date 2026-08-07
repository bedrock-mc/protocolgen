// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ContainerOpen struct {
	ContainerId   uint8
	ContainerType uint8
	Position      BlockPos
	TargetActorID ActorUniqueID
}

// Marshal reads or writes ContainerOpen using its canonical wire layout.
func (x *ContainerOpen) Marshal(io IO) {
	io.Uint8(&x.ContainerId)
	io.Uint8(&x.ContainerType)
	x.Position.Marshal(io)
	x.TargetActorID.Marshal(io)
}
