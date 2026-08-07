// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

type StartGame struct {
	EntityID                          int64
	RuntimeID                         uint64
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
	io.ActorUniqueID(&x.EntityID)
	io.ActorRuntimeID(&x.RuntimeID)
	IntegerFunc(&x.GameType, io.Varint32)
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
	FuncSlice(io, &x.BlockProperties, io.Varuint32, func(value *ServerBlockProperty) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	io.String(&x.MultiplayerCorrelationId)
	io.Bool(&x.EnableItemStackNetManager)
	io.String(&x.ServerVersion)
	io.NBT(&x.PlayerPropertyData)
	io.Uint64(&x.ServerBlockTypeRegistryChecksum)
	io.UUID(&x.WorldTemplateID)
	io.Bool(&x.ServerEnabledClientSideGeneration)
	io.Bool(&x.BlockNetworkIdsAreHashes)
	x.NetworkPermissions.Marshal(io)
	OptionalFunc(io, &x.ServerConfigurationJoinInfo, func(value *ServerConfigurationServerConfigurationJoinInfo) {
		item := *value
		item.Marshal(io)
		*value = item
	})
	x.ServerTelemetryData.Marshal(io)
}
