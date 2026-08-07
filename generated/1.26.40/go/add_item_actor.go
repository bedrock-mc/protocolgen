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
