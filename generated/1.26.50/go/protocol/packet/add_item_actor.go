// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// AddItemActor is sent by the server to the client to make an item entity show up. It is one of the
// few entities that cannot be sent using the AddActor packet
type AddItemActor struct {
	TargetActorID   int64
	TargetRuntimeID uint64
	// Item is the item that is spawned. It must have a valid ID for it to show up client-side. If it is
	// not a valid item, the client will crash when coming near.
	Item protocol.NetworkItemStackDescriptorSerializedData
	// Position is the position to spawn the entity on. If the entity is on a distance that the player
	// cannot see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Velocity is the initial velocity the entity spawns with. This velocity will initiate client side
	// movement of the entity.
	Velocity      mgl32.Vec3
	EntityData    protocol.SynchedActorDataCopyableDataList
	IsFromFishing bool
}

// Marshal reads or writes AddItemActor using its canonical wire layout.
func (x *AddItemActor) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	x.Item.Marshal(io)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	x.EntityData.Marshal(io)
	io.Bool(&x.IsFromFishing)
}

// ID returns the protocol ID for AddItemActor.
func (*AddItemActor) ID() uint32 { return IDAddItemActor }
