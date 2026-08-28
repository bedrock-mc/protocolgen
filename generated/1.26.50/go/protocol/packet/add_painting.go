// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.50/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// AddPainting is sent by the server to the client to make a painting entity show up. It is one of
// the few entities that cannot be sent using the AddActor packet.
type AddPainting struct {
	TargetActorID   int64
	TargetRuntimeID uint64
	// Position is the position to spawn the entity on. If the entity is on a distance that the player
	// cannot see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Direction is the facing direction of the painting.
	Direction int32
	Motif     string
}

// Marshal reads or writes AddPainting using its canonical wire layout.
func (x *AddPainting) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.Vec3(&x.Position)
	io.Varint32(&x.Direction)
	io.String(&x.Motif)
}

// ID returns the protocol ID for AddPainting.
func (*AddPainting) ID() uint32 { return IDAddPainting }
