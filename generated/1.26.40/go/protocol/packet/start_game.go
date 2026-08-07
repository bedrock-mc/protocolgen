// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

type StartGame struct {
	EntityID                          int64
	RuntimeID                         uint64
	GameType                          protocol.GameType
	Position                          mgl32.Vec3
	Rotation                          mgl32.Vec2
	Settings                          protocol.LevelSettings
	LevelID                           string
	LevelName                         string
	TemplateContentIdentity           string
	IsTrial                           bool
	MovementSettings                  protocol.SyncedPlayerMovementSettings
	LevelCurrentTime                  uint64
	EnchantmentSeed                   int32
	BlockProperties                   []protocol.ServerBlockProperty
	MultiplayerCorrelationID          string
	EnableItemStackNetManager         bool
	ServerVersion                     string
	PlayerPropertyData                []byte
	ServerBlockTypeRegistryChecksum   uint64
	WorldTemplateID                   uuid.UUID
	ServerEnabledClientSideGeneration bool
	BlockNetworkIdsAreHashes          bool
	NetworkPermissions                protocol.NetworkPermissions
	ServerConfigurationJoinInfo       protocol.Optional[protocol.ServerConfigurationServerConfigurationJoinInfo]
	ServerTelemetryData               protocol.SocialEventsServerTelemetryData
}

// Marshal reads or writes StartGame using its canonical wire layout.
func (x *StartGame) Marshal(io protocol.IO) {
	io.ActorUniqueID(&x.EntityID)
	io.ActorRuntimeID(&x.RuntimeID)
	protocol.IntegerFunc(&x.GameType, io.Varint32)
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
	protocol.Slice(io, &x.BlockProperties)
	io.String(&x.MultiplayerCorrelationID)
	io.Bool(&x.EnableItemStackNetManager)
	io.String(&x.ServerVersion)
	io.NBT(&x.PlayerPropertyData)
	io.Uint64(&x.ServerBlockTypeRegistryChecksum)
	io.UUID(&x.WorldTemplateID)
	io.Bool(&x.ServerEnabledClientSideGeneration)
	io.Bool(&x.BlockNetworkIdsAreHashes)
	x.NetworkPermissions.Marshal(io)
	protocol.OptionalFunc(io, &x.ServerConfigurationJoinInfo, func(value *protocol.ServerConfigurationServerConfigurationJoinInfo) {
		value.Marshal(io)
	})
	x.ServerTelemetryData.Marshal(io)
}

// ID returns the protocol ID for StartGame.
func (*StartGame) ID() uint32 { return IDStartGame }
