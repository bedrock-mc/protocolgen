// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"
)

// MoveActorDelta is sent by the server to move an entity. The packet is specifically optimised to save as
// much space as possible, by only writing non-zero fields. As of 1.16.100, this packet no longer actually
// contains any deltas.
type MoveActorDelta struct {
	ActorRuntimeID       uint64
	NewPositionX         protocol.Optional[float32]
	NewPositionY         protocol.Optional[float32]
	NewPositionZ         protocol.Optional[float32]
	RotationX            protocol.Optional[int8]
	RotationY            protocol.Optional[int8]
	RotationYHead        protocol.Optional[int8]
	IsOnGround           bool
	ForceMove            bool
	ForceMoveLocalEntity bool
	ForceCompletion      bool
	Ticks                uint64
}

// ID returns the protocol ID for MoveActorDelta.
func (*MoveActorDelta) ID() uint32 { return IDMoveActorDelta }

// Marshal reads or writes MoveActorDelta using its canonical wire layout.
func (pk *MoveActorDelta) Marshal(io protocol.IO) {
	io.ActorRuntimeID(&pk.ActorRuntimeID)
	protocol.OptionalFunc(io, &pk.NewPositionX, io.Float32)
	protocol.OptionalFunc(io, &pk.NewPositionY, io.Float32)
	protocol.OptionalFunc(io, &pk.NewPositionZ, io.Float32)
	protocol.OptionalFunc(io, &pk.RotationX, io.Int8)
	protocol.OptionalFunc(io, &pk.RotationY, io.Int8)
	protocol.OptionalFunc(io, &pk.RotationYHead, io.Int8)
	io.Bool(&pk.IsOnGround)
	io.Bool(&pk.ForceMove)
	io.Bool(&pk.ForceMoveLocalEntity)
	io.Bool(&pk.ForceCompletion)
	io.Varuint64(&pk.Ticks)
}
