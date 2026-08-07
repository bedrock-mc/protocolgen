// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

type AddActor struct {
	TargetActorID     int64
	TargetRuntimeID   uint64
	ActorType         string
	Position          mgl32.Vec3
	Velocity          mgl32.Vec3
	Rotation          mgl32.Vec2
	YHeadRotation     float32
	YBodyRotation     float32
	AttributesList    []protocol.SyncedAttribute
	ActorData         protocol.SynchedActorDataCopyableDataList
	SynchedProperties protocol.PropertySyncData
	ActorLinks        []protocol.ActorLink
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
	protocol.FuncSlice(io, &x.AttributesList, io.Varuint32, func(value *protocol.SyncedAttribute) {
		value.Marshal(io)
	})
	x.ActorData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	protocol.FuncSlice(io, &x.ActorLinks, io.Varuint32, func(value *protocol.ActorLink) {
		value.Marshal(io)
	})
}

// ID returns the protocol ID for AddActor.
func (*AddActor) ID() uint32 { return IDAddActor }
