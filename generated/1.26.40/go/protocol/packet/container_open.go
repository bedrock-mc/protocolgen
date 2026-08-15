// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import "protocolgen/generated/1.26.40/go/protocol"

// ContainerOpen is sent by the server to open a container client-side. This container must be
// physically present in the world, for the packet to have any effect. Unlike Java Edition, Bedrock
// Edition requires that chests for example must be present and in range to open its inventory.
type ContainerOpen struct {
	ContainerID uint8
	// ContainerType is the type ID of the container that is being opened when opening the container at
	// the position of the packet. It depends on the block/entity, and could, for example, be the window
	// type of a chest or a hopper, but also a horse inventory.
	ContainerType uint8
	Position      protocol.BlockPos
	TargetActorID int64
}

// Marshal reads or writes ContainerOpen using its canonical wire layout.
func (x *ContainerOpen) Marshal(io protocol.IO) {
	io.Uint8(&x.ContainerID)
	protocol.Minimum(io, &x.ContainerID, 0)
	protocol.Maximum(io, &x.ContainerID, 255)
	io.Uint8(&x.ContainerType)
	protocol.Minimum(io, &x.ContainerType, 0)
	protocol.Maximum(io, &x.ContainerType, 255)
	x.Position.Marshal(io)
	io.ActorUniqueID(&x.TargetActorID)
}

// ID returns the protocol ID for ContainerOpen.
func (*ContainerOpen) ID() uint32 { return IDContainerOpen }
