// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "github.com/go-gl/mathgl/mgl32"

type AddActor struct {
	TargetActorID     int64
	TargetRuntimeID   uint64
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

// Marshal reads or writes AddActor using its canonical wire layout.
func (x *AddActor) Marshal(io IO) {
	io.ActorUniqueID(&x.TargetActorID)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.String(&x.ActorType)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	io.Float32(&x.YBodyRotation)
	FuncSlice(io, &x.AttributesList, io.Varuint32, func(value *SyncedAttribute) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	FuncSlice(io, &x.ActorLinks, io.Varuint32, func(value *ActorLink) {
		item := *value
		item.Marshal(io)
		*value = item
	})
}
