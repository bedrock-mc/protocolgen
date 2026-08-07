// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

type ContainerOpen struct {
	ContainerID   uint8
	ContainerType uint8
	Position      protocol.BlockPos
	TargetActorID int64
}

// Marshal reads or writes ContainerOpen using its canonical wire layout.
func (x *ContainerOpen) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	io.Uint8(&x.ContainerType)
	x.Position.Marshal(io)
	io.ActorUniqueID(&x.TargetActorID)
}

// ID returns the protocol ID for ContainerOpen.
func (*ContainerOpen) ID() uint32 { return IDContainerOpen }
