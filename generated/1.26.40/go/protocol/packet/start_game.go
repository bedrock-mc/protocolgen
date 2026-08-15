// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.40/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// StartGame is sent by the server to send information about the world the player will be spawned
// in. It contains information about the position the player spawns in, and information about the
// world in general such as its game rules.
type StartGame struct {
	EntityID  int64
	RuntimeID uint64
	GameType  protocol.GameType
	Position  mgl32.Vec3
	Rotation  mgl32.Vec2
	Settings  protocol.LevelSettings
	// LevelID is a base64 encoded world ID that is used to identify the world.
	LevelID   string
	LevelName string
	// TemplateContentIdentity is a UUID specific to the premium world template that might have been
	// used to generate the world. Servers should always fill out an empty string for this.
	TemplateContentIdentity string
	IsTrial                 bool
	MovementSettings        protocol.SyncedPlayerMovementSettings
	LevelCurrentTime        uint64
	// EnchantmentSeed is the seed used to seed the random used to produce enchantments in the
	// enchantment table. Note that the exact correct random implementation must be used to produce the
	// correct results both client- and server-side.
	EnchantmentSeed int32
	BlockProperties []protocol.ServerBlockProperty
	// MultiplayerCorrelationID is a unique ID specifying the multi-player session of the player. A
	// random UUID should be filled out for this field.
	MultiplayerCorrelationID        string
	EnableItemStackNetManager       bool
	ServerVersion                   string
	PlayerPropertyData              []byte
	ServerBlockTypeRegistryChecksum uint64
	// WorldTemplateID is a UUID that identifies the template that was used to generate the world.
	// Servers that do not use a world based off of a template can set this to an empty UUID.
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
	protocol.Minimum(io, &x.LevelCurrentTime, 0)
	io.Varint32(&x.EnchantmentSeed)
	protocol.Slice(io, &x.BlockProperties)
	io.String(&x.MultiplayerCorrelationID)
	io.Bool(&x.EnableItemStackNetManager)
	io.String(&x.ServerVersion)
	io.NBT(&x.PlayerPropertyData, protocol.NBTNetwork)
	io.Uint64(&x.ServerBlockTypeRegistryChecksum)
	protocol.Minimum(io, &x.ServerBlockTypeRegistryChecksum, 0)
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
