// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type StartGame struct {
	EntityID                          ActorUniqueID
	RuntimeID                         ActorRuntimeID
	GameType                          GameType
	Position                          Vec3
	Rotation                          Vec2
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
	WorldTemplateID                   [16]byte
	ServerEnabledClientSideGeneration bool
	BlockNetworkIdsAreHashes          bool
	NetworkPermissions                NetworkPermissions
	ServerConfigurationJoinInfo       *ServerConfigurationServerConfigurationJoinInfo
	ServerTelemetryData               SocialEventsServerTelemetryData
}

func (p *StartGame) Encode(w Encoder) error {
	if err := w.Write("StartGamePacket.Entity ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}}, p.EntityID); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}}, p.RuntimeID); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}, p.GameType); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Position); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}, p.Rotation); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Settings", Shape{Kind: "struct", Semantic: "LevelSettings", TypeID: "LevelSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Seed", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 1, Name: "Spawn Settings", Shape: Shape{Kind: "struct", Semantic: "SpawnSettings", TypeID: "SpawnSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Spawn Biome Type", Shape: Shape{Kind: "enum", Semantic: "SpawnBiomeType", TypeID: "enums/SpawnBiomeType", PrimitiveCode: "i16le", Variants: []ShapeVariant{{Value: 0, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "UserDefined", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "User Defined Biome Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Dimension", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 2, Name: "Generator Type", Shape: Shape{Kind: "enum", Semantic: "GeneratorType", TypeID: "enums/GeneratorType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Legacy", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Overworld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Flat", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Nether", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "TheEnd", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Void", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Undefined", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Game Type", Shape: Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 4, Name: "Is Hardcore", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Game Difficulty", Shape: Shape{Kind: "enum", Semantic: "SharedTypes::Legacy::Difficulty", TypeID: "enums/SharedTypes::Legacy::Difficulty", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Peaceful", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Easy", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Normal", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Hard", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Count", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Unknown", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 6, Name: "Default Spawn Block Position", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 7, Name: "Achievements Disabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 8, Name: "Editor World Type", Shape: Shape{Kind: "enum", Semantic: "Editor::WorldType", TypeID: "enums/Editor::WorldType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "NonEditor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "EditorProject", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "EditorTestLevel", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "EditorRealmsUpload", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 9, Name: "Is Created In Editor", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 10, Name: "Is Exported From Editor", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 11, Name: "Day Cycle Stop Time", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 12, Name: "Education Edition Offer", Shape: Shape{Kind: "enum", Semantic: "EducationEditionOffer", TypeID: "enums/EducationEditionOffer", PrimitiveCode: "var_u32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "RestOfWorld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "China_Deprecated", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 13, Name: "Education Features Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 14, Name: "Education Product ID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 15, Name: "Rain Level", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 16, Name: "Lightning Level", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 17, Name: "Has Confirmed Platform Locked Content", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 18, Name: "Multiplayer Game Intent", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 19, Name: "LAN Broadcast Intent", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 20, Name: "Xbox Live Broadcast Setting", Shape: Shape{Kind: "enum", Semantic: "Social::GamePublishSetting", TypeID: "enums/Social::GamePublishSetting", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "NoMultiPlay", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InviteOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "FriendsOnly", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "FriendsOfFriends", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Public", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 21, Name: "Platform Broadcast Setting", Shape: Shape{Kind: "enum", Semantic: "Social::GamePublishSetting", TypeID: "enums/Social::GamePublishSetting", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "NoMultiPlay", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InviteOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "FriendsOnly", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "FriendsOfFriends", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Public", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 22, Name: "Commands Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 23, Name: "Texture Packs Required", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 24, Name: "Rule Data", Shape: Shape{Kind: "struct", Semantic: "GameRulesChangedPacketData", TypeID: "GameRulesChangedPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "Rules List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "GameRule", TypeID: "GameRule", Fields: []ShapeField{{Ordinal: 0, Name: "Rule Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Rule Can Be Modified", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Rule Value", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "Empty0", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "int32", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Value: 3, Name: "float", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}}, {Ordinal: 25, Name: "Experiments", Shape: Shape{Kind: "struct", Semantic: "Experiments", TypeID: "Experiments", Fields: []ShapeField{{Ordinal: 0, Name: "Toggles", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "cerealizer_ExperimentsAnon::ExperimentToggle", TypeID: "cerealizer_ExperimentsAnon::ExperimentToggle", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}, {Ordinal: 1, Name: "ExperimentsEverToggled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}, {Ordinal: 26, Name: "Has Bonus Chest Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 27, Name: "Start With Map Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 28, Name: "Player Permissions", Shape: Shape{Kind: "enum", Semantic: "PlayerPermissionLevel", TypeID: "enums/PlayerPermissionLevel", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "Visitor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Member", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Operator", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Custom", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 29, Name: "Server Chunk Tick Range", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 30, Name: "Has Locked Behavior Pack", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 31, Name: "Has Locked Resource Pack", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 32, Name: "Is From Locked Template", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 33, Name: "Use Msa Gamertags Only", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 34, Name: "Is From World Template", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 35, Name: "Is World Template Option Locked", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 36, Name: "Only Spawn V1 Villagers", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 37, Name: "Persona Disabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 38, Name: "Custom Skins Disabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 39, Name: "Emote Chat Muted", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 40, Name: "Base Game Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 41, Name: "Limited World Width", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 42, Name: "Limited World Depth", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 43, Name: "Nether Type", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 44, Name: "Edu Shared Uri Resource", Shape: Shape{Kind: "struct", Semantic: "EduSharedUriResource", TypeID: "EduSharedUriResource", Fields: []ShapeField{{Ordinal: 0, Name: "Button Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Link Uri", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 45, Name: "Override Force Experimental Gameplay", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "bool"}}}, {Ordinal: 46, Name: "Chat Restriction Level", Shape: Shape{Kind: "enum", Semantic: "ChatRestrictionLevel", TypeID: "enums/ChatRestrictionLevel", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Dropped", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Disabled", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 47, Name: "Disable Player Interactions", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 48, Name: "Server Editor Connection Policy", Shape: Shape{Kind: "enum", Semantic: "ServerEditorConnectionPolicy", TypeID: "enums/ServerEditorConnectionPolicy", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "MatchWorldType", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "EditorOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "VanillaOnly", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Mixed", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 49, Name: "Allow Anonymous Block Drops In Editor Worlds", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}, p.Settings); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Level ID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.LevelID); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Level Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.LevelName); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Template Content Identity", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.TemplateContentIdentity); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Is Trial", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.IsTrial); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Movement Settings", Shape{Kind: "struct", Semantic: "SyncedPlayerMovementSettings", TypeID: "SyncedPlayerMovementSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Rewind History Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Server Authoritative Block Breaking", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}, p.MovementSettings); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Level Current Time", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.LevelCurrentTime); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Enchantment Seed", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.EnchantmentSeed); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Block Properties", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ServerBlockProperty", TypeID: "ServerBlockProperty", Fields: []ShapeField{{Ordinal: 0, Name: "Block Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Block Definition", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}}, p.BlockProperties); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Multiplayer Correlation Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.MultiplayerCorrelationId); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Enable Item Stack Net Manager", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.EnableItemStackNetManager); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Server Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}, p.ServerVersion); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Player Property Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}, p.PlayerPropertyData); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Server Block Type Registry Checksum", Shape{Kind: "primitive", PrimitiveCode: "u64le"}, p.ServerBlockTypeRegistryChecksum); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.World Template ID", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"}, p.WorldTemplateID); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Server Enabled ClientSide Generation", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ServerEnabledClientSideGeneration); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.BlockNetworkIds Are Hashes", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.BlockNetworkIdsAreHashes); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.NetworkPermissions", Shape{Kind: "struct", Semantic: "NetworkPermissions", TypeID: "NetworkPermissions", Fields: []ShapeField{{Ordinal: 0, Name: "Server Auth Sound Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}, p.NetworkPermissions); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Server Configuration Join Info", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::ServerConfigurationJoinInfo", TypeID: "ServerConfiguration::ServerConfigurationJoinInfo", Fields: []ShapeField{{Ordinal: 0, Name: "gathering", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::GatheringsConfigurationJoinInfo", TypeID: "ServerConfiguration::GatheringsConfigurationJoinInfo", Fields: []ShapeField{{Ordinal: 0, Name: "experienceId", Shape: Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "experienceName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "worldId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 3, Name: "worldName", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 4, Name: "creatorId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "targetId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 6, Name: "scenarioId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 7, Name: "serverId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, {Ordinal: 1, Name: "clientStoreEntryPoint", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::ClientStoreEntryPointConfiguration", TypeID: "ServerConfiguration::ClientStoreEntryPointConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "storeId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "storeName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Ordinal: 2, Name: "presence", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::PresenceConfiguration", TypeID: "ServerConfiguration::PresenceConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "richPresenceId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}}}, p.ServerConfigurationJoinInfo); err != nil {
		return err
	}
	if err := w.Write("StartGamePacket.Server Telemetry Data", Shape{Kind: "struct", Semantic: "Social::Events::ServerTelemetryData", TypeID: "Social::Events::ServerTelemetryData", Fields: []ShapeField{{Ordinal: 0, Name: "Server Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Scenario Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "World Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "Owner Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}, p.ServerTelemetryData); err != nil {
		return err
	}
	return nil
}

func DecodeStartGame(r Decoder) (StartGame, error) {
	var p StartGame
	{
		raw, err := r.Read("StartGamePacket.Entity ID", Shape{Kind: "struct", Semantic: "ActorUniqueID", TypeID: "ActorUniqueID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Unique ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorUniqueID)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Entity ID has unexpected decoded type %T", raw)
		}
		p.EntityID = value
	}
	{
		raw, err := r.Read("StartGamePacket.Runtime ID", Shape{Kind: "struct", Semantic: "ActorRuntimeID", TypeID: "ActorRuntimeID", Fields: []ShapeField{{Ordinal: 0, Name: "Actor Runtime ID", Shape: Shape{Kind: "primitive", PrimitiveCode: "var_u64"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(ActorRuntimeID)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Runtime ID has unexpected decoded type %T", raw)
		}
		p.RuntimeID = value
	}
	{
		raw, err := r.Read("StartGamePacket.Game Type", Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(GameType)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Game Type has unexpected decoded type %T", raw)
		}
		p.GameType = value
	}
	{
		raw, err := r.Read("StartGamePacket.Position", Shape{Kind: "struct", Semantic: "Vec3", TypeID: "Vec3", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec3)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Position has unexpected decoded type %T", raw)
		}
		p.Position = value
	}
	{
		raw, err := r.Read("StartGamePacket.Rotation", Shape{Kind: "struct", Semantic: "Vec2", TypeID: "Vec2", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(Vec2)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Rotation has unexpected decoded type %T", raw)
		}
		p.Rotation = value
	}
	{
		raw, err := r.Read("StartGamePacket.Settings", Shape{Kind: "struct", Semantic: "LevelSettings", TypeID: "LevelSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Seed", Shape: Shape{Kind: "primitive", PrimitiveCode: "u64le"}}, {Ordinal: 1, Name: "Spawn Settings", Shape: Shape{Kind: "struct", Semantic: "SpawnSettings", TypeID: "SpawnSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Spawn Biome Type", Shape: Shape{Kind: "enum", Semantic: "SpawnBiomeType", TypeID: "enums/SpawnBiomeType", PrimitiveCode: "i16le", Variants: []ShapeVariant{{Value: 0, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "UserDefined", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 1, Name: "User Defined Biome Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "Dimension", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 2, Name: "Generator Type", Shape: Shape{Kind: "enum", Semantic: "GeneratorType", TypeID: "enums/GeneratorType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Legacy", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Overworld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Flat", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Nether", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "TheEnd", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Void", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Undefined", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 3, Name: "Game Type", Shape: Shape{Kind: "enum", Semantic: "GameType", TypeID: "enums/GameType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: -1, Name: "Undefined", Shape: Shape{Kind: "void"}}, {Value: 0, Name: "Survival", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Creative", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Adventure", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Default", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "Spectator", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 4, Name: "Is Hardcore", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 5, Name: "Game Difficulty", Shape: Shape{Kind: "enum", Semantic: "SharedTypes::Legacy::Difficulty", TypeID: "enums/SharedTypes::Legacy::Difficulty", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "Peaceful", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Easy", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Normal", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Hard", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Count", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "Unknown", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 6, Name: "Default Spawn Block Position", Shape: Shape{Kind: "struct", Semantic: "BlockPos", TypeID: "BlockPos", Fields: []ShapeField{{Ordinal: 0, Name: "X", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Y", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 2, Name: "Z", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}}}}, {Ordinal: 7, Name: "Achievements Disabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 8, Name: "Editor World Type", Shape: Shape{Kind: "enum", Semantic: "Editor::WorldType", TypeID: "enums/Editor::WorldType", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "NonEditor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "EditorProject", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "EditorTestLevel", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "EditorRealmsUpload", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 9, Name: "Is Created In Editor", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 10, Name: "Is Exported From Editor", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 11, Name: "Day Cycle Stop Time", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 12, Name: "Education Edition Offer", Shape: Shape{Kind: "enum", Semantic: "EducationEditionOffer", TypeID: "enums/EducationEditionOffer", PrimitiveCode: "var_u32", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "RestOfWorld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "China_Deprecated", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 13, Name: "Education Features Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 14, Name: "Education Product ID", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 15, Name: "Rain Level", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 16, Name: "Lightning Level", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}, {Ordinal: 17, Name: "Has Confirmed Platform Locked Content", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 18, Name: "Multiplayer Game Intent", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 19, Name: "LAN Broadcast Intent", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 20, Name: "Xbox Live Broadcast Setting", Shape: Shape{Kind: "enum", Semantic: "Social::GamePublishSetting", TypeID: "enums/Social::GamePublishSetting", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "NoMultiPlay", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InviteOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "FriendsOnly", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "FriendsOfFriends", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Public", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 21, Name: "Platform Broadcast Setting", Shape: Shape{Kind: "enum", Semantic: "Social::GamePublishSetting", TypeID: "enums/Social::GamePublishSetting", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "NoMultiPlay", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "InviteOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "FriendsOnly", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "FriendsOfFriends", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "Public", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 22, Name: "Commands Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 23, Name: "Texture Packs Required", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 24, Name: "Rule Data", Shape: Shape{Kind: "struct", Semantic: "GameRulesChangedPacketData", TypeID: "GameRulesChangedPacketData", Fields: []ShapeField{{Ordinal: 0, Name: "Rules List", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "GameRule", TypeID: "GameRule", Fields: []ShapeField{{Ordinal: 0, Name: "Rule Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Rule Can Be Modified", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 2, Name: "Rule Value", Shape: Shape{Kind: "union", Control: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Variants: []ShapeVariant{{Value: 0, Name: "Empty0", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "bool", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Value: 2, Name: "int32", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Value: 3, Name: "float", Shape: Shape{Kind: "primitive", PrimitiveCode: "f32le"}}}}}}}}}}}}, {Ordinal: 25, Name: "Experiments", Shape: Shape{Kind: "struct", Semantic: "Experiments", TypeID: "Experiments", Fields: []ShapeField{{Ordinal: 0, Name: "Toggles", Shape: Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "u32le"}, Element: &Shape{Kind: "struct", Semantic: "cerealizer_ExperimentsAnon::ExperimentToggle", TypeID: "cerealizer_ExperimentsAnon::ExperimentToggle", Fields: []ShapeField{{Ordinal: 0, Name: "Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}}, {Ordinal: 1, Name: "ExperimentsEverToggled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}}}, {Ordinal: 26, Name: "Has Bonus Chest Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 27, Name: "Start With Map Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 28, Name: "Player Permissions", Shape: Shape{Kind: "enum", Semantic: "PlayerPermissionLevel", TypeID: "enums/PlayerPermissionLevel", PrimitiveCode: "i8", Variants: []ShapeVariant{{Value: 0, Name: "Visitor", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Member", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Operator", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Custom", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 29, Name: "Server Chunk Tick Range", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 30, Name: "Has Locked Behavior Pack", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 31, Name: "Has Locked Resource Pack", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 32, Name: "Is From Locked Template", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 33, Name: "Use Msa Gamertags Only", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 34, Name: "Is From World Template", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 35, Name: "Is World Template Option Locked", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 36, Name: "Only Spawn V1 Villagers", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 37, Name: "Persona Disabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 38, Name: "Custom Skins Disabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 39, Name: "Emote Chat Muted", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 40, Name: "Base Game Version", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 41, Name: "Limited World Width", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 42, Name: "Limited World Depth", Shape: Shape{Kind: "primitive", PrimitiveCode: "i32le"}}, {Ordinal: 43, Name: "Nether Type", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 44, Name: "Edu Shared Uri Resource", Shape: Shape{Kind: "struct", Semantic: "EduSharedUriResource", TypeID: "EduSharedUriResource", Fields: []ShapeField{{Ordinal: 0, Name: "Button Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Link Uri", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}, {Ordinal: 45, Name: "Override Force Experimental Gameplay", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "bool"}}}, {Ordinal: 46, Name: "Chat Restriction Level", Shape: Shape{Kind: "enum", Semantic: "ChatRestrictionLevel", TypeID: "enums/ChatRestrictionLevel", PrimitiveCode: "u8", Variants: []ShapeVariant{{Value: 0, Name: "None", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "Dropped", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "Disabled", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 47, Name: "Disable Player Interactions", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}, {Ordinal: 48, Name: "Server Editor Connection Policy", Shape: Shape{Kind: "enum", Semantic: "ServerEditorConnectionPolicy", TypeID: "enums/ServerEditorConnectionPolicy", PrimitiveCode: "zigzag_i32", Variants: []ShapeVariant{{Value: 0, Name: "MatchWorldType", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "EditorOnly", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "VanillaOnly", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "Mixed", Shape: Shape{Kind: "void"}}}}}, {Ordinal: 49, Name: "Allow Anonymous Block Drops In Editor Worlds", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(LevelSettings)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Settings has unexpected decoded type %T", raw)
		}
		p.Settings = value
	}
	{
		raw, err := r.Read("StartGamePacket.Level ID", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Level ID has unexpected decoded type %T", raw)
		}
		p.LevelID = value
	}
	{
		raw, err := r.Read("StartGamePacket.Level Name", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Level Name has unexpected decoded type %T", raw)
		}
		p.LevelName = value
	}
	{
		raw, err := r.Read("StartGamePacket.Template Content Identity", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Template Content Identity has unexpected decoded type %T", raw)
		}
		p.TemplateContentIdentity = value
	}
	{
		raw, err := r.Read("StartGamePacket.Is Trial", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Is Trial has unexpected decoded type %T", raw)
		}
		p.IsTrial = value
	}
	{
		raw, err := r.Read("StartGamePacket.Movement Settings", Shape{Kind: "struct", Semantic: "SyncedPlayerMovementSettings", TypeID: "SyncedPlayerMovementSettings", Fields: []ShapeField{{Ordinal: 0, Name: "Rewind History Size", Shape: Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}}, {Ordinal: 1, Name: "Server Authoritative Block Breaking", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SyncedPlayerMovementSettings)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Movement Settings has unexpected decoded type %T", raw)
		}
		p.MovementSettings = value
	}
	{
		raw, err := r.Read("StartGamePacket.Level Current Time", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Level Current Time has unexpected decoded type %T", raw)
		}
		p.LevelCurrentTime = value
	}
	{
		raw, err := r.Read("StartGamePacket.Enchantment Seed", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Enchantment Seed has unexpected decoded type %T", raw)
		}
		p.EnchantmentSeed = value
	}
	{
		raw, err := r.Read("StartGamePacket.Block Properties", Shape{Kind: "array", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}, Element: &Shape{Kind: "struct", Semantic: "ServerBlockProperty", TypeID: "ServerBlockProperty", Fields: []ShapeField{{Ordinal: 0, Name: "Block Name", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Block Definition", Shape: Shape{Kind: "primitive", PrimitiveCode: "nbt_le"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]ServerBlockProperty)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Block Properties has unexpected decoded type %T", raw)
		}
		p.BlockProperties = value
	}
	{
		raw, err := r.Read("StartGamePacket.Multiplayer Correlation Id", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Multiplayer Correlation Id has unexpected decoded type %T", raw)
		}
		p.MultiplayerCorrelationId = value
	}
	{
		raw, err := r.Read("StartGamePacket.Enable Item Stack Net Manager", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Enable Item Stack Net Manager has unexpected decoded type %T", raw)
		}
		p.EnableItemStackNetManager = value
	}
	{
		raw, err := r.Read("StartGamePacket.Server Version", Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(string)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Server Version has unexpected decoded type %T", raw)
		}
		p.ServerVersion = value
	}
	{
		raw, err := r.Read("StartGamePacket.Player Property Data", Shape{Kind: "primitive", PrimitiveCode: "nbt_le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([]byte)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Player Property Data has unexpected decoded type %T", raw)
		}
		p.PlayerPropertyData = value
	}
	{
		raw, err := r.Read("StartGamePacket.Server Block Type Registry Checksum", Shape{Kind: "primitive", PrimitiveCode: "u64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint64)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Server Block Type Registry Checksum has unexpected decoded type %T", raw)
		}
		p.ServerBlockTypeRegistryChecksum = value
	}
	{
		raw, err := r.Read("StartGamePacket.World Template ID", Shape{Kind: "primitive", Semantic: "mce::UUID", TypeID: "mce__UUID.json#", PrimitiveCode: "uuid"})
		if err != nil {
			return p, err
		}
		value, ok := raw.([16]byte)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.World Template ID has unexpected decoded type %T", raw)
		}
		p.WorldTemplateID = value
	}
	{
		raw, err := r.Read("StartGamePacket.Server Enabled ClientSide Generation", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Server Enabled ClientSide Generation has unexpected decoded type %T", raw)
		}
		p.ServerEnabledClientSideGeneration = value
	}
	{
		raw, err := r.Read("StartGamePacket.BlockNetworkIds Are Hashes", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.BlockNetworkIds Are Hashes has unexpected decoded type %T", raw)
		}
		p.BlockNetworkIdsAreHashes = value
	}
	{
		raw, err := r.Read("StartGamePacket.NetworkPermissions", Shape{Kind: "struct", Semantic: "NetworkPermissions", TypeID: "NetworkPermissions", Fields: []ShapeField{{Ordinal: 0, Name: "Server Auth Sound Enabled", Shape: Shape{Kind: "primitive", PrimitiveCode: "bool"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(NetworkPermissions)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.NetworkPermissions has unexpected decoded type %T", raw)
		}
		p.NetworkPermissions = value
	}
	{
		raw, err := r.Read("StartGamePacket.Server Configuration Join Info", Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::ServerConfigurationJoinInfo", TypeID: "ServerConfiguration::ServerConfigurationJoinInfo", Fields: []ShapeField{{Ordinal: 0, Name: "gathering", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::GatheringsConfigurationJoinInfo", TypeID: "ServerConfiguration::GatheringsConfigurationJoinInfo", Fields: []ShapeField{{Ordinal: 0, Name: "experienceId", Shape: Shape{Kind: "primitive", PrimitiveCode: "uuid"}}, {Ordinal: 1, Name: "experienceName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "worldId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 3, Name: "worldName", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 4, Name: "creatorId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 5, Name: "targetId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "primitive", PrimitiveCode: "uuid"}}}, {Ordinal: 6, Name: "scenarioId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}, {Ordinal: 7, Name: "serverId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}, {Ordinal: 1, Name: "clientStoreEntryPoint", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::ClientStoreEntryPointConfiguration", TypeID: "ServerConfiguration::ClientStoreEntryPointConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "storeId", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "storeName", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}, {Ordinal: 2, Name: "presence", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "struct", Semantic: "ServerConfiguration::PresenceConfiguration", TypeID: "ServerConfiguration::PresenceConfiguration", Fields: []ShapeField{{Ordinal: 0, Name: "richPresenceId", Shape: Shape{Kind: "optional", Value: &Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}}}}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(*ServerConfigurationServerConfigurationJoinInfo)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Server Configuration Join Info has unexpected decoded type %T", raw)
		}
		p.ServerConfigurationJoinInfo = value
	}
	{
		raw, err := r.Read("StartGamePacket.Server Telemetry Data", Shape{Kind: "struct", Semantic: "Social::Events::ServerTelemetryData", TypeID: "Social::Events::ServerTelemetryData", Fields: []ShapeField{{Ordinal: 0, Name: "Server Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 1, Name: "Scenario Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 2, Name: "World Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}, {Ordinal: 3, Name: "Owner Id", Shape: Shape{Kind: "string", Encoding: "utf8", Representation: "text", Prefix: &Shape{Kind: "primitive", PrimitiveCode: "var_u32"}}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(SocialEventsServerTelemetryData)
		if !ok {
			return p, fmt.Errorf("field StartGamePacket.Server Telemetry Data has unexpected decoded type %T", raw)
		}
		p.ServerTelemetryData = value
	}
	return p, nil
}
