// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type AddItemActor struct {
	TargetActorID   int64
	TargetRuntimeID uint64
	Item            protocol.CerealizerNetworkItemStackDescriptorSerializedData
	Position        mgl32.Vec3
	Velocity        mgl32.Vec3
	EntityData      protocol.SynchedActorDataCopyableDataList
	IsFromFishing   bool
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
