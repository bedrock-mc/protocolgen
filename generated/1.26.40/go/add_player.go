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
