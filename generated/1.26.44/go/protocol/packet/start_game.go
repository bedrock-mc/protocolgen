// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package packet

import (
	"protocolgen/generated/1.26.44/go/protocol"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// StartGame is sent by the server to send information about the world the player will be spawned in. It
// contains information about the position the player spawns in, and information about the world in general
// such as its game rules.
type StartGame struct {
	EntityID  int64
	RuntimeID uint64
	// EntityUniqueID is the unique ID of the player. The unique ID is a value that remains consistent across
	// different sessions of the same world, but most servers simply fill the runtime ID of the entity out for
	// this field.
	GameType protocol.GameType
	// PlayerPosition is the spawn position of the player in the world. In servers this is often the same as the
	// world's spawn position found below.
	Position mgl32.Vec3
	// WorldSpawn is the block on which the world spawn of the world. This coordinate has no effect on the place
	// that the client spawns, but it does have an effect on the direction that a compass points.
	Rotation mgl32.Vec2
	// EducationSharedResourceURI is an education edition feature that transmits education resource settings to
	// clients.
	Settings protocol.LevelSettings
	// LevelID is a base64 encoded world ID that is used to identify the world.
	LevelID string
	// WorldName is the name of the world that the player is joining. Note that this field shows up above the
	// player list for the rest of the game session, and cannot be changed. Setting the server name to this field
	// is recommended.
	LevelName string
	// TemplateContentIdentity is a UUID specific to the premium world template that might have been used to
	// generate the world. Servers should always fill out an empty string for this.
	TemplateContentIdentity string
	// Trial specifies if the world was a trial world, meaning features are limited and there is a time limit on
	// the world.
	IsTrial bool
	// PlayerMovementSettings ...
	MovementSettings protocol.SyncedPlayerMovementSettings
	// Time is the total time that has elapsed since the start of the world.
	LevelCurrentTime uint64
	// EnchantmentSeed is the seed used to seed the random used to produce enchantments in the enchantment table.
	// Note that the exact correct random implementation must be used to produce the correct results both client-
	// and server-side.
	EnchantmentSeed int32
	// Blocks is a list of all custom blocks registered on the server.
	BlockProperties []protocol.ServerBlockProperty
	// MultiplayerCorrelationID is a unique ID specifying the multi-player session of the player. A random UUID
	// should be filled out for this field.
	MultiplayerCorrelationID string
	// ServerAuthoritativeInventory specifies if the server authoritative inventory system is enabled. This is a
	// new system introduced in 1.16. Backwards compatibility with the inventory transactions has to some extent
	// been preserved, but will eventually be removed.
	EnableItemStackNetManager bool
	// GameVersion is the version of the game the server is running. The exact function of this field isn't clear.
	ServerVersion      string
	PlayerPropertyData []byte
	// ServerBlockStateChecksum is a checksum to ensure block states between the server and client match. This can
	// simply be left empty, and the client will avoid trying to verify it.
	ServerBlockTypeRegistryChecksum uint64
	// WorldTemplateID is a UUID that identifies the template that was used to generate the world. Servers that do
	// not use a world based off of a template can set this to an empty UUID.
	WorldTemplateID uuid.UUID
	// ClientSideGeneration is true if the client should use the features registered in the FeatureRegistry packet
	// to generate terrain client-side to save on bandwidth.
	ServerEnabledClientSideGeneration bool
	// DisablePlayerInteractions is true if the client should ignore other players when interacting with the
	// world.
	BlockNetworkIdsAreHashes bool
	NetworkPermissions       protocol.NetworkPermissions
	// ServerJoinInformation contains optional information about the server the player is joining.
	ServerConfigurationJoinInfo protocol.Optional[protocol.ServerConfigurationServerConfigurationJoinInfo]
	ServerTelemetryData         protocol.SocialEventsServerTelemetryData
}

// ID ...
func (*StartGame) ID() uint32 {
	return IDStartGame
}

func (pk *StartGame) Marshal(io protocol.IO) {
	io.ActorUniqueID(&pk.EntityID)
	io.ActorRuntimeID(&pk.RuntimeID)
	pk.GameType.Marshal(io)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.Rotation)
	pk.Settings.Marshal(io)
	io.String(&pk.LevelID)
	io.String(&pk.LevelName)
	io.String(&pk.TemplateContentIdentity)
	io.Bool(&pk.IsTrial)
	pk.MovementSettings.Marshal(io)
	io.Uint64(&pk.LevelCurrentTime)
	io.Varint32(&pk.EnchantmentSeed)
	protocol.Slice(io, &pk.BlockProperties)
	io.String(&pk.MultiplayerCorrelationID)
	io.Bool(&pk.EnableItemStackNetManager)
	io.String(&pk.ServerVersion)
	io.NBT(&pk.PlayerPropertyData, protocol.NBTNetwork)
	io.Uint64(&pk.ServerBlockTypeRegistryChecksum)
	io.UUID(&pk.WorldTemplateID)
	io.Bool(&pk.ServerEnabledClientSideGeneration)
	io.Bool(&pk.BlockNetworkIdsAreHashes)
	pk.NetworkPermissions.Marshal(io)
	protocol.OptionalMarshaler(io, &pk.ServerConfigurationJoinInfo)
	pk.ServerTelemetryData.Marshal(io)
}
