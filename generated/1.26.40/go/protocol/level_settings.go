// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type LevelSettings struct {
	Seed                                   uint64
	SpawnSettings                          SpawnSettings
	GeneratorType                          GeneratorType
	GameType                               GameType
	IsHardcore                             bool
	GameDifficulty                         LegacyDifficulty
	DefaultSpawnBlockPosition              BlockPos
	AchievementsDisabled                   bool
	EditorWorldType                        EditorWorldType
	IsCreatedInEditor                      bool
	IsExportedFromEditor                   bool
	DayCycleStopTime                       int32
	EducationEditionOffer                  EducationEditionOffer
	EducationFeaturesEnabled               bool
	EducationProductID                     string
	RainLevel                              float32
	LightningLevel                         float32
	HasConfirmedPlatformLockedContent      bool
	MultiplayerGameIntent                  bool
	LANBroadcastIntent                     bool
	XboxLiveBroadcastSetting               SocialGamePublishSetting
	PlatformBroadcastSetting               SocialGamePublishSetting
	CommandsEnabled                        bool
	TexturePacksRequired                   bool
	RuleData                               GameRulesChangedPacketData
	Experiments                            Experiments
	HasBonusChestEnabled                   bool
	StartWithMapEnabled                    bool
	PlayerPermissions                      PlayerPermissionLevel
	ServerChunkTickRange                   int32
	HasLockedBehaviorPack                  bool
	HasLockedResourcePack                  bool
	IsFromLockedTemplate                   bool
	UseMsaGamertagsOnly                    bool
	IsFromWorldTemplate                    bool
	IsWorldTemplateOptionLocked            bool
	OnlySpawnV1Villagers                   bool
	PersonaDisabled                        bool
	CustomSkinsDisabled                    bool
	EmoteChatMuted                         bool
	BaseGameVersion                        string
	LimitedWorldWidth                      int32
	LimitedWorldDepth                      int32
	NetherType                             bool
	EduSharedUriResource                   EduSharedUriResource
	OverrideForceExperimentalGameplay      Optional[bool]
	ChatRestrictionLevel                   ChatRestrictionLevel
	DisablePlayerInteractions              bool
	ServerEditorConnectionPolicy           ServerEditorConnectionPolicy
	AllowAnonymousBlockDropsInEditorWorlds bool
}

// Marshal reads or writes LevelSettings using its canonical wire layout.
func (x *LevelSettings) Marshal(io IO) {
	io.Uint64(&x.Seed)
	x.SpawnSettings.Marshal(io)
	IntegerFunc(&x.GeneratorType, io.Varint32)
	IntegerFunc(&x.GameType, io.Varint32)
	io.Bool(&x.IsHardcore)
	IntegerFunc(&x.GameDifficulty, io.Varint32)
	x.DefaultSpawnBlockPosition.Marshal(io)
	io.Bool(&x.AchievementsDisabled)
	IntegerFunc(&x.EditorWorldType, io.Varint32)
	io.Bool(&x.IsCreatedInEditor)
	io.Bool(&x.IsExportedFromEditor)
	io.Varint32(&x.DayCycleStopTime)
	IntegerFunc(&x.EducationEditionOffer, io.Varuint32)
	io.Bool(&x.EducationFeaturesEnabled)
	io.String(&x.EducationProductID)
	io.Float32(&x.RainLevel)
	io.Float32(&x.LightningLevel)
	io.Bool(&x.HasConfirmedPlatformLockedContent)
	io.Bool(&x.MultiplayerGameIntent)
	io.Bool(&x.LANBroadcastIntent)
	IntegerFunc(&x.XboxLiveBroadcastSetting, io.Varint32)
	IntegerFunc(&x.PlatformBroadcastSetting, io.Varint32)
	io.Bool(&x.CommandsEnabled)
	io.Bool(&x.TexturePacksRequired)
	x.RuleData.Marshal(io)
	x.Experiments.Marshal(io)
	io.Bool(&x.HasBonusChestEnabled)
	io.Bool(&x.StartWithMapEnabled)
	IntegerFunc(&x.PlayerPermissions, io.Int8)
	io.Int32(&x.ServerChunkTickRange)
	io.Bool(&x.HasLockedBehaviorPack)
	io.Bool(&x.HasLockedResourcePack)
	io.Bool(&x.IsFromLockedTemplate)
	io.Bool(&x.UseMsaGamertagsOnly)
	io.Bool(&x.IsFromWorldTemplate)
	io.Bool(&x.IsWorldTemplateOptionLocked)
	io.Bool(&x.OnlySpawnV1Villagers)
	io.Bool(&x.PersonaDisabled)
	io.Bool(&x.CustomSkinsDisabled)
	io.Bool(&x.EmoteChatMuted)
	io.String(&x.BaseGameVersion)
	io.Int32(&x.LimitedWorldWidth)
	io.Int32(&x.LimitedWorldDepth)
	io.Bool(&x.NetherType)
	x.EduSharedUriResource.Marshal(io)
	OptionalFunc(io, &x.OverrideForceExperimentalGameplay, io.Bool)
	IntegerFunc(&x.ChatRestrictionLevel, io.Uint8)
	io.Bool(&x.DisablePlayerInteractions)
	IntegerFunc(&x.ServerEditorConnectionPolicy, io.Varint32)
	io.Bool(&x.AllowAnonymousBlockDropsInEditorWorlds)
}
