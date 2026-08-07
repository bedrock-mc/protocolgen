// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// AddActor is sent by the server to the client to spawn an entity to the player. It is used for
// every entity except other players, for which the AddPlayer packet is used.
type AddActor struct {
	TargetActorID   int64
	TargetRuntimeID uint64
	ActorType       string
	// Position is the position to spawn the entity on. If the entity is on a distance that the player
	// cannot see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Velocity is the initial velocity the entity spawns with. This velocity will initiate client side
	// movement of the entity.
	Velocity          mgl32.Vec3
	Rotation          mgl32.Vec2
	YHeadRotation     float32
	YBodyRotation     float32
	AttributesList    []protocol.SyncedAttribute
	ActorData         protocol.SynchedActorDataCopyableDataList
	SynchedProperties protocol.PropertySyncData
	ActorLinks        []protocol.EntityLink
}

// Marshal reads or writes AddActor using its canonical wire layout.
func (x *AddActor) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.String(&x.ActorType)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	io.Float32(&x.YBodyRotation)
	protocol.Slice(io, &x.AttributesList)
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	protocol.Slice(io, &x.ActorLinks)
}

// ID returns the protocol ID for AddActor.
func (*AddActor) ID() uint32 { return IDAddActor }
