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

// Marshal reads or writes AddActor using its canonical wire layout.
func (x *AddActor) Marshal(io IO) {
	x.TargetActorID.Marshal(io)
	x.TargetRuntimeID.Marshal(io)
	io.String(&x.ActorType)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	io.Float32(&x.YBodyRotation)
	if !io.Reading() && uint64(len(x.AttributesList)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.AttributesList), "collection length overflows uint32")
		return
	}
	count1 := uint32(len(x.AttributesList))
	io.Varuint32(&count1)
	if io.Reading() {
		if uint64(count1) > uint64(^uint(0)>>1) {
			io.InvalidValue(count1, "collection length overflows int")
			return
		}
		x.AttributesList = make([]SyncedAttribute, int(count1))
	}
	for index2 := range x.AttributesList {
		x.AttributesList[index2].Marshal(io)
	}
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	if !io.Reading() && uint64(len(x.ActorLinks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ActorLinks), "collection length overflows uint32")
		return
	}
	count3 := uint32(len(x.ActorLinks))
	io.Varuint32(&count3)
	if io.Reading() {
		if uint64(count3) > uint64(^uint(0)>>1) {
			io.InvalidValue(count3, "collection length overflows int")
			return
		}
		x.ActorLinks = make([]ActorLink, int(count3))
	}
	for index4 := range x.ActorLinks {
		x.ActorLinks[index4].Marshal(io)
	}
}
