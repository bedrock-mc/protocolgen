// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

type StartGame struct {
	EntityID                          ActorUniqueID
	RuntimeID                         ActorRuntimeID
	GameType                          GameType
	Position                          mgl32.Vec3
	Rotation                          mgl32.Vec2
	Settings                          LevelSettings
	LevelID                           string
	LevelName                         string
	TemplateContentIdentity           string
	IsTrial                           bool
	MovementSettings                  SyncedPlayerMovementSettings
	LevelCurrentTime                  uint64
	EnchantmentSeed                   int32
	BlockProperties                   []ServerBlockProperty
	MultiplayerCorrelationId          string
	EnableItemStackNetManager         bool
	ServerVersion                     string
	PlayerPropertyData                []byte
	ServerBlockTypeRegistryChecksum   uint64
	WorldTemplateID                   uuid.UUID
	ServerEnabledClientSideGeneration bool
	BlockNetworkIdsAreHashes          bool
	NetworkPermissions                NetworkPermissions
	ServerConfigurationJoinInfo       Optional[ServerConfigurationServerConfigurationJoinInfo]
	ServerTelemetryData               SocialEventsServerTelemetryData
}

// Marshal reads or writes StartGame using its canonical wire layout.
func (x *StartGame) Marshal(io IO) {
	x.EntityID.Marshal(io)
	x.RuntimeID.Marshal(io)
	enumValue1 := int32(x.GameType)
	io.Varint32(&enumValue1)
	x.GameType = GameType(enumValue1)
	switch int64(enumValue1) {
	case -1, 0, 1, 2, 5, 6:
	default:
		io.InvalidValue(enumValue1, "unknown enum value")
	}
	io.Vec3(&x.Position)
	io.Vec2(&x.Rotation)
	x.Settings.Marshal(io)
	io.String(&x.LevelID)
	io.String(&x.LevelName)
	io.String(&x.TemplateContentIdentity)
	io.Bool(&x.IsTrial)
	x.MovementSettings.Marshal(io)
	io.Uint64(&x.LevelCurrentTime)
	io.Varint32(&x.EnchantmentSeed)
	if !io.Reading() && uint64(len(x.BlockProperties)) > uint64(^uint32(0)) {
		io.InvalidValue(len(x.BlockProperties), "collection length overflows uint32")
		return
	}
	count2 := uint32(len(x.BlockProperties))
	io.Varuint32(&count2)
	if io.Reading() {
		if uint64(count2) > uint64(^uint(0)>>1) {
			io.InvalidValue(count2, "collection length overflows int")
			return
		}
		x.BlockProperties = make([]ServerBlockProperty, int(count2))
	}
	for index3 := range x.BlockProperties {
		x.BlockProperties[index3].Marshal(io)
	}
	io.String(&x.MultiplayerCorrelationId)
	io.Bool(&x.EnableItemStackNetManager)
	io.String(&x.ServerVersion)
	io.NBT(&x.PlayerPropertyData)
	io.Uint64(&x.ServerBlockTypeRegistryChecksum)
	io.UUID(&x.WorldTemplateID)
	io.Bool(&x.ServerEnabledClientSideGeneration)
	io.Bool(&x.BlockNetworkIdsAreHashes)
	x.NetworkPermissions.Marshal(io)
	io.Bool(&x.ServerConfigurationJoinInfo.set)
	if x.ServerConfigurationJoinInfo.set {
		x.ServerConfigurationJoinInfo.val.Marshal(io)
	} else if io.Reading() {
		var zero ServerConfigurationServerConfigurationJoinInfo
		x.ServerConfigurationJoinInfo.val = zero
	}
	x.ServerTelemetryData.Marshal(io)
}
