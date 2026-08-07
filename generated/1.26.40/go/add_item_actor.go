// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type AddItemActor struct {
	TargetActorID   ActorUniqueID
	TargetRuntimeID ActorRuntimeID
	Item            CerealizerNetworkItemStackDescriptorSerializedData
	Position        mgl32.Vec3
	Velocity        mgl32.Vec3
	EntityData      SynchedActorDataCopyableDataList
	IsFromFishing   bool
}

// Marshal reads or writes AddItemActor using its canonical wire layout.
func (x *AddItemActor) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
	x.TargetRuntimeID.Marshal(io)
	x.Item.Marshal(io)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	x.EntityData.Marshal(io)
	io.Bool(&x.IsFromFishing)
}
