// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type AddActor struct {
	TargetActorID     ActorUniqueID
	TargetRuntimeID   ActorRuntimeID
	ActorType         string
	Position          mgl32.Vec3
	Velocity          mgl32.Vec3
	Rotation          mgl32.Vec2
	YHeadRotation     float32
	YBodyRotation     float32
	AttributesList    []SyncedAttribute
	ActorData         SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	ActorLinks        []ActorLink
}
