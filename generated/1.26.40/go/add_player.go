// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

type AddPlayer struct {
	UUID              uuid.UUID
	PlayerName        string
	TargetRuntimeID   uint64
	PlatformChatId    string
	Position          mgl32.Vec3
	Velocity          mgl32.Vec3
	Rotation          mgl32.Vec2
	YHeadRotation     float32
	CarriedItem       CerealizerNetworkItemStackDescriptorSerializedData
	PlayerGameType    GameType
	EntityData        SynchedActorDataCopyableDataList
	SynchedProperties PropertySyncData
	AbilitiesData     SerializedAbilitiesData
	ActorLinks        []ActorLink
	DeviceId          string
	BuildPlatform     BuildPlatform
}

// Marshal reads or writes AddPlayer using its canonical wire layout.
func (x *AddPlayer) Marshal(io IO) {
	io.UUID(&x.UUID)
	io.String(&x.PlayerName)
	io.ActorRuntimeID(&x.TargetRuntimeID)
	io.String(&x.PlatformChatId)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	x.CarriedItem.Marshal(io)
	IntegerFunc(&x.PlayerGameType, io.Varint32)
	x.EntityData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	x.AbilitiesData.Marshal(io)
	FuncSlice(io, &x.ActorLinks, io.Varuint32, func(value *ActorLink) {
		value.Marshal(io)
	})
	io.String(&x.DeviceId)
	IntegerFunc(&x.BuildPlatform, io.Int32)
}
