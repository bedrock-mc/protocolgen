// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

type AddPlayer struct {
	UUID              uuid.UUID
	PlayerName        string
	TargetRuntimeID   ActorRuntimeID
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
	x.TargetRuntimeID.Marshal(io)
	io.String(&x.PlatformChatId)
	io.Vec3(&x.Position)
	io.Vec3(&x.Velocity)
	io.Vec2(&x.Rotation)
	io.Float32(&x.YHeadRotation)
	x.CarriedItem.Marshal(io)
	enumValue1 := int32(x.PlayerGameType)
	io.Varint32(&enumValue1)
	x.PlayerGameType = GameType(enumValue1)
	switch int64(enumValue1) {
	case -1, 0, 1, 2, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	x.EntityData.Marshal(io)
	x.SynchedProperties.Marshal(io)
	x.AbilitiesData.Marshal(io)
	if !io.Reading() && uint64(len(x.ActorLinks)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.ActorLinks), "collection length overflows uint32")
		return
	}
	count2 := uint32(len(x.ActorLinks))
	io.Varuint32(&count2)
	if io.Reading() {
		if uint64(count2) > uint64(^uint(0)>>1) {
			io.InvalidValue(count2, "collection length overflows int")
			return
		}
		x.ActorLinks = make([]ActorLink, int(count2))
	}
	for index3 := range x.ActorLinks {
		x.ActorLinks[index3].Marshal(io)
	}
	io.String(&x.DeviceId)
	enumValue4 := int32(x.BuildPlatform)
	io.Int32(&enumValue4)
	x.BuildPlatform = BuildPlatform(enumValue4)
	switch int64(enumValue4) {
	case -1, 1, 2, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15:
	default:
		io.InvalidValue(enumValue4, "unknown enum value")
	}
}
