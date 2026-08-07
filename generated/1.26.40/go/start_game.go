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
