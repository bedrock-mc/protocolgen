# gophertunnel update guide — protocol 2168 to 2187

1.26.40-beta.0 → 1.26.50-beta.25

| | From | To |
| --- | --- | --- |
| Branch | `automated/1.26.40` | `r/26_u5` |
| Protocol | 2168 | 2187 |
| Game version | 1.26.40-beta.0 | 1.26.50-beta.25 |
| Upstream commit | `0e00fe80f4f3c71572ff6429de40146d1f4412fc` | `e0a1dad9341b85d3191f998c3d84b7d2c96e6c81` |
| Fixer commit | `0399c985c7bd18e050c71b5eee8150e98fcc19fb` | `0399c985c7bd18e050c71b5eee8150e98fcc19fb` |

Read `changelog.md` for the human-readable diff. This guide gives a target-version `Marshal` transcription for each changed definition.

**The Go below is a transcription aid, not a patch.** Names come from the schema, so they may not match gophertunnel's existing names. Field comments are emitted only when Mojang provides a description.

## Added Packets

### SetPlayerFurnaceOptionsPacket
packet id `351` · struct

```go
const IDSetPlayerFurnaceOptionsPacket uint32 = 351

type SetPlayerFurnaceOptionsPacket struct {
	FurnaceType    uint8
	FurnaceOptions FurnaceOptions
}

func (*SetPlayerFurnaceOptionsPacket) ID() uint32 {
	return IDSetPlayerFurnaceOptionsPacket
}

func (pk *SetPlayerFurnaceOptionsPacket) Marshal(io protocol.IO) {
	io.Uint8(&pk.FurnaceType)
	protocol.Single(io, &pk.FurnaceOptions)
}
```

### RecordStartedPacket
packet id `352` · struct

```go
const IDRecordStartedPacket uint32 = 352

type RecordStartedPacket struct {
	BlockPosition     BlockPos
	ServerSoundHandle ServerSoundHandle
}

func (*RecordStartedPacket) ID() uint32 {
	return IDRecordStartedPacket
}

func (pk *RecordStartedPacket) Marshal(io protocol.IO) {
	protocol.Single(io, &pk.BlockPosition)
	protocol.Single(io, &pk.ServerSoundHandle)
}
```

## Added Types

### FurnaceOptions
struct

```go
type FurnaceOptions struct {
	LeftFurnaceTab int32
	Filtering      bool
	Layout         int32
}

func (pk *FurnaceOptions) Marshal(io IO) {
	io.Int32(&pk.LeftFurnaceTab)
	io.Bool(&pk.Filtering)
	io.Int32(&pk.Layout)
}
```

### NoiseAlignment
struct

```go
type NoiseAlignment struct {
	Type  uint8
	Value uint32
}

func (pk *NoiseAlignment) Marshal(io IO) {
	io.Uint8(&pk.Type)
	io.Varuint32(&pk.Value)
	if pk.Value < 0 {
		io.InvalidValue(pk.Value, "value", "value below minimum")
	}
}
```

## Added Enums

### FurnaceLayout
enum

```go
const (
	FurnaceLayoutNone          = 0
	FurnaceLayoutInventoryOnly = 1
	FurnaceLayoutDefault       = 2
)
```

### FurnaceLeftTabIndex
enum

```go
const (
	FurnaceLeftTabIndexNone         = 0
	FurnaceLeftTabIndexRecipeFood   = 1
	FurnaceLeftTabIndexRecipeItems  = 2
	FurnaceLeftTabIndexRecipeBlocks = 3
	FurnaceLeftTabIndexRecipeSearch = 4
	FurnaceLeftTabIndexInventory    = 5
)
```

### NoiseAlignmentType
enum

```go
const (
	NoiseAlignmentTypeMinLocalTransitionEnd = 0
)
```

### SetPlayerFurnaceOptionsPacketPayload_FurnaceType
enum

```go
const (
	SetPlayerFurnaceOptionsFurnaceTypeNone         = 0
	SetPlayerFurnaceOptionsFurnaceTypeFurnace      = 1
	SetPlayerFurnaceOptionsFurnaceTypeBlastFurnace = 2
	SetPlayerFurnaceOptionsFurnaceTypeSmoker       = 3
)
```

## Modified Packets

### PlaySoundPacket
packet id `86` · struct · **wire break**

- added `Bypass Listener Range Check` at ordinal 5 (boolean)
- `Server Sound Handle` moved from ordinal 5 to 6
- added `Playback Position Seconds` at ordinal 7 (float, optional)

```go
type PlaySoundPacket struct {
	Name                     string
	Position                 BlockPos
	Volume                   float32
	Pitch                    float32
	LoopCount                int32
	BypassListenerRangeCheck bool
	ServerSoundHandle        protocol.Optional[ServerSoundHandle]
	PlaybackPositionSeconds  protocol.Optional[float32]
}

func (*PlaySoundPacket) ID() uint32 {
	return IDPlaySoundPacket
}

func (pk *PlaySoundPacket) Marshal(io protocol.IO) {
	io.String(&pk.Name)
	protocol.Single(io, &pk.Position)
	io.Float32(&pk.Volume)
	io.Float32(&pk.Pitch)
	io.Varint32(&pk.LoopCount)
	io.Bool(&pk.BypassListenerRangeCheck)
	protocol.OptionalMarshaler(io, &pk.ServerSoundHandle)
	protocol.OptionalFunc(io, &pk.PlaybackPositionSeconds, io.Float32)
}
```

### RequestNetworkSettingsPacket
packet id `193` · struct

- `ClientNetworkVersion` constraints changed from maximum=2168, minimum=2168 to maximum=2187, minimum=2187

```go
type RequestNetworkSettingsPacket struct {
	ClientNetworkVersion int32
}

func (*RequestNetworkSettingsPacket) ID() uint32 {
	return IDRequestNetworkSettingsPacket
}

func (pk *RequestNetworkSettingsPacket) Marshal(io protocol.IO) {
	io.BEInt32(&pk.ClientNetworkVersion)
	if pk.ClientNetworkVersion < 2187 {
		io.InvalidValue(pk.ClientNetworkVersion, "clientNetworkVersion", "value below minimum")
	}
	if pk.ClientNetworkVersion > 2187 {
		io.InvalidValue(pk.ClientNetworkVersion, "clientNetworkVersion", "value above maximum")
	}
}
```

### ServerboundPackSettingChangePacket
packet id `329` · struct

- schema changed

```go
type ServerboundPackSettingChangePacket struct {
	// PackID The UUID of the pack whose setting is being changed.
	PackID MceUUID
	// PackSettingName The name of the setting being changed.
	PackSettingName string
	// PackSettingValue The value of the setting being changed.
	PackSettingValue any
}

func (*ServerboundPackSettingChangePacket) ID() uint32 {
	return IDServerboundPackSettingChangePacket
}

func (pk *ServerboundPackSettingChangePacket) Marshal(io protocol.IO) {
	protocol.Single(io, &pk.PackID)
	io.String(&pk.PackSettingName)
	if len(pk.PackSettingName) > 128 {
		io.InvalidValue(pk.PackSettingName, "packSettingName", "string too long")
	}
	// PackSettingValue is untyped in the target schema; resolve its wire shape from exact-version evidence.
}
```

## Modified Types

### CameraPresets
struct · **wire break**

- added `Apply Inherited Starting Rotation` at ordinal 22 (boolean)
- added `Starting Rotation` at ordinal 23 (Vec2)

**Also update:** `CameraPresetsPacket` — it embeds this type.

```go
type CameraPresets struct {
	Name                    IdentifierSharedTypesAssetTypeCameraPreset
	InheritFrom             ReferenceSharedTypesAssetTypeCameraPreset
	PosX                    Optional[float32]
	PosY                    Optional[float32]
	PosZ                    Optional[float32]
	RotX                    Optional[float32]
	RotY                    Optional[float32]
	RotationSpeed           Optional[float32]
	SnapToTarget            Optional[bool]
	HorizontalRotationLimit Optional[Vec2]
	VerticalRotationLimit   Optional[Vec2]
	ContinueTargeting       Optional[bool]
	BlockListeningRadius    Optional[float32]
	ViewOffset              Optional[Vec2]
	// EntityOffset Changing the camera's pivot point from the center of the entity
	EntityOffset                   Optional[Vec3]
	Radius                         Optional[float32]
	YawLimitMin                    Optional[float32]
	YawLimitMax                    Optional[float32]
	Listener                       Optional[uint8]
	PlayerEffects                  Optional[bool]
	AimAssist                      Optional[ComprehensiveCameraAimAssistCommandDefinition]
	ControlScheme                  Optional[uint8]
	ApplyInheritedStartingRotation bool
	StartingRotation               Vec2
}

func (pk *CameraPresets) Marshal(io IO) {
	Single(io, &pk.Name)
	Single(io, &pk.InheritFrom)
	OptionalFunc(io, &pk.PosX, io.Float32)
	OptionalFunc(io, &pk.PosY, io.Float32)
	OptionalFunc(io, &pk.PosZ, io.Float32)
	OptionalFunc(io, &pk.RotX, io.Float32)
	OptionalFunc(io, &pk.RotY, io.Float32)
	OptionalFunc(io, &pk.RotationSpeed, io.Float32)
	OptionalFunc(io, &pk.SnapToTarget, io.Bool)
	OptionalMarshaler(io, &pk.HorizontalRotationLimit)
	OptionalMarshaler(io, &pk.VerticalRotationLimit)
	OptionalFunc(io, &pk.ContinueTargeting, io.Bool)
	OptionalFunc(io, &pk.BlockListeningRadius, io.Float32)
	OptionalMarshaler(io, &pk.ViewOffset)
	OptionalMarshaler(io, &pk.EntityOffset)
	OptionalFunc(io, &pk.Radius, io.Float32)
	OptionalFunc(io, &pk.YawLimitMin, io.Float32)
	OptionalFunc(io, &pk.YawLimitMax, io.Float32)
	OptionalFunc(io, &pk.Listener, io.Uint8)
	OptionalFunc(io, &pk.PlayerEffects, io.Bool)
	OptionalMarshaler(io, &pk.AimAssist)
	OptionalFunc(io, &pk.ControlScheme, io.Uint8)
	io.Bool(&pk.ApplyInheritedStartingRotation)
	Single(io, &pk.StartingRotation)
}
```

### DimensionDefinition
struct · **wire break**

- added `Minimum Y` at ordinal 0 (int32 (Compression))
- added `Height Range` at ordinal 1 (int32 (Compression))
- added `Default Biome` at ordinal 5 (string)
- removed `Height Maximum` from ordinal 0
- removed `Height Minimum` from ordinal 1

**Also update:** `DimensionDataPacket` — it embeds this type.

```go
type DimensionDefinition struct {
	MinimumY      int32
	HeightRange   int32
	GeneratorType string
	// DimensionType Dimension runtime ID.
	DimensionType DimensionType
	// PackID UUID of the pack that registered this dimension.
	PackID MceUUID
	// DefaultBiome Identifier of the biome a custom dimension generates with.
	DefaultBiome string
}

func (pk *DimensionDefinition) Marshal(io IO) {
	io.Varint32(&pk.MinimumY)
	io.Varint32(&pk.HeightRange)
	io.String(&pk.GeneratorType)
	Single(io, &pk.DimensionType)
	Single(io, &pk.PackID)
	io.String(&pk.DefaultBiome)
	if len(pk.DefaultBiome) > 256 {
		io.InvalidValue(pk.DefaultBiome, "defaultBiome", "string too long")
	}
}
```

### EntityDiagnosticTimingInfo
struct · **wire break**

- added `Position` at ordinal 4 (Vec3)
- added `Dimension` at ordinal 5 (string)

**Also update:** `ServerboundDiagnosticsPacket` — it embeds this type.

```go
type EntityDiagnosticTimingInfo struct {
	DisplayName    string
	Entity         string
	TimeInNS       uint64
	PercentOfTotal uint8
	Position       Vec3
	Dimension      string
}

func (pk *EntityDiagnosticTimingInfo) Marshal(io IO) {
	io.String(&pk.DisplayName)
	io.String(&pk.Entity)
	io.Uint64(&pk.TimeInNS)
	if pk.TimeInNS < 0 {
		io.InvalidValue(pk.TimeInNS, "timeInNS", "value below minimum")
	}
	io.Uint8(&pk.PercentOfTotal)
	if pk.PercentOfTotal < 0 {
		io.InvalidValue(pk.PercentOfTotal, "percentOfTotal", "value below minimum")
	}
	if pk.PercentOfTotal > 255 {
		io.InvalidValue(pk.PercentOfTotal, "percentOfTotal", "value above maximum")
	}
	Single(io, &pk.Position)
	io.String(&pk.Dimension)
}
```

### EnvironmentAttributeData
struct · **wire break**

- added `NoiseAlignment` at ordinal 9 (NoiseAlignment)

**Also update:** `ClientboundAttributeLayerSyncPacket` — it embeds this type.

```go
type EnvironmentAttributeData struct {
	AttributeName          string
	FromAttribute          Optional[any]
	Attribute              any
	ToAttribute            Optional[any]
	CurrentTransitionTicks uint32
	TotalTransitionTicks   uint32
	Easing                 int32
	LocalTransitionTicks   uint32
	NoiseTransition        bool
	NoiseAlignment         NoiseAlignment
}

func (pk *EnvironmentAttributeData) Marshal(io IO) {
	io.String(&pk.AttributeName)
	if len(pk.AttributeName) > 128 {
		io.InvalidValue(pk.AttributeName, "attributeName", "string too long")
	}
	// FromAttribute is untyped in the target schema; resolve its wire shape from exact-version evidence.
	// Attribute is untyped in the target schema; resolve its wire shape from exact-version evidence.
	// ToAttribute is untyped in the target schema; resolve its wire shape from exact-version evidence.
	io.Uint32(&pk.CurrentTransitionTicks)
	if pk.CurrentTransitionTicks < 0 {
		io.InvalidValue(pk.CurrentTransitionTicks, "currentTransitionTicks", "value below minimum")
	}
	io.Uint32(&pk.TotalTransitionTicks)
	if pk.TotalTransitionTicks < 0 {
		io.InvalidValue(pk.TotalTransitionTicks, "totalTransitionTicks", "value below minimum")
	}
	io.Int32(&pk.Easing)
	io.Uint32(&pk.LocalTransitionTicks)
	if pk.LocalTransitionTicks < 0 {
		io.InvalidValue(pk.LocalTransitionTicks, "localTransitionTicks", "value below minimum")
	}
	io.Bool(&pk.NoiseTransition)
	Single(io, &pk.NoiseAlignment)
}
```

### MoveActorDeltaData
struct · **wire break**

- added `Ticks` at ordinal 11 (uint64 (Compression))

**Also update:** `MoveActorDeltaPacket` — it embeds this type.

```go
type MoveActorDeltaData struct {
	// ActorRuntimeID The runtime id of the actor being moved.
	ActorRuntimeID ActorRuntimeID
	// NewPositionX New X coordinate of the actor; absent when X position has not changed.
	NewPositionX Optional[float32]
	// NewPositionY New Y coordinate of the actor; absent when Y position has not changed.
	NewPositionY Optional[float32]
	// NewPositionZ New Z coordinate of the actor; absent when Z position has not changed.
	NewPositionZ Optional[float32]
	// RotationX New X rotation packed as int8; absent when X rotation has not changed.
	RotationX Optional[int8]
	// RotationY New Y rotation packed as int8; absent when Y rotation has not changed.
	RotationY Optional[int8]
	// RotationYHead New head Y rotation packed as int8; absent when head Y rotation has not changed or the actor is not a Mob.
	RotationYHead Optional[int8]
	// IsOnGround Whether the actor is on the ground after applying this update.
	IsOnGround bool
	// ForceMove Whether the receiver should snap the actor to the new position without interpolation.
	ForceMove bool
	// ForceMoveLocalEntity Whether the receiver should snap a locally-owned entity to the new position.
	ForceMoveLocalEntity bool
	// ForceCompletion Whether the receiver should complete any in-progress local movement before applying this update.
	ForceCompletion bool
	// Ticks Expected number of ticks before the next movement update, used for position interpolation duration on the client.
	Ticks uint64
}

func (pk *MoveActorDeltaData) Marshal(io IO) {
	Single(io, &pk.ActorRuntimeID)
	OptionalFunc(io, &pk.NewPositionX, io.Float32)
	OptionalFunc(io, &pk.NewPositionY, io.Float32)
	OptionalFunc(io, &pk.NewPositionZ, io.Float32)
	OptionalFunc(io, &pk.RotationX, func(value *int8) {
		io.Int8(value)
		if *value < -128 {
			io.InvalidValue(*value, "rotationX", "value below minimum")
		}
		if *value > 127 {
			io.InvalidValue(*value, "rotationX", "value above maximum")
		}
	})
	OptionalFunc(io, &pk.RotationY, func(value *int8) {
		io.Int8(value)
		if *value < -128 {
			io.InvalidValue(*value, "rotationY", "value below minimum")
		}
		if *value > 127 {
			io.InvalidValue(*value, "rotationY", "value above maximum")
		}
	})
	OptionalFunc(io, &pk.RotationYHead, func(value *int8) {
		io.Int8(value)
		if *value < -128 {
			io.InvalidValue(*value, "rotationYHead", "value below minimum")
		}
		if *value > 127 {
			io.InvalidValue(*value, "rotationYHead", "value above maximum")
		}
	})
	io.Bool(&pk.IsOnGround)
	io.Bool(&pk.ForceMove)
	io.Bool(&pk.ForceMoveLocalEntity)
	io.Bool(&pk.ForceCompletion)
	io.Varuint64(&pk.Ticks)
	if pk.Ticks < 0 {
		io.InvalidValue(pk.Ticks, "ticks", "value below minimum")
	}
}
```

### TextDataPayload
struct · **wire break**

- added `LineGapHeight` at ordinal 3 (float, optional)
- `DepthTest` moved from ordinal 3 to 4
- `ShowBackface` moved from ordinal 4 to 5
- `ShowTextBackface` moved from ordinal 5 to 6

**Also update:** `PrimitiveShapesPacket` — it embeds this type.

```go
type TextData struct {
	// Text Text (string) of the debug text shape.
	Text string
	// UseRotation Whether to use specified rotation, otherwise faces camera.
	UseRotation bool
	// BackgroundColor Color to draw background.
	BackgroundColor Optional[Color]
	// LineGapHeight Line gap height for multiline text rendering.
	LineGapHeight Optional[float32]
	// DepthTest Whether other objects block seeing the text (false=like nametags,true=like signs).
	DepthTest bool
	// ShowBackface Whether the backface of the background should be rendered.
	ShowBackface bool
	// ShowTextBackface Whether the backface of the text should be rendered.
	ShowTextBackface bool
}

func (pk *TextData) Marshal(io IO) {
	io.String(&pk.Text)
	io.Bool(&pk.UseRotation)
	OptionalMarshaler(io, &pk.BackgroundColor)
	OptionalFunc(io, &pk.LineGapHeight, io.Float32)
	io.Bool(&pk.DepthTest)
	io.Bool(&pk.ShowBackface)
	io.Bool(&pk.ShowTextBackface)
}
```

## Modified Enums

### BuildPlatform
enum

- added value `Nintendo` = 12
- removed value `Nx` = 12

**Also update:** `AddPlayerPacket`, `PlayerListPacket` — each embeds this type.

```go
const (
	BuildPlatformGoogle       = 1
	BuildPlatformIOS          = 2
	BuildPlatformOSX          = 3
	BuildPlatformAmazon       = 4
	BuildPlatformGearVR       = 5
	BuildPlatformUWP          = 7
	BuildPlatformWin32        = 8
	BuildPlatformDedicated    = 9
	BuildPlatformTvOS         = 10
	BuildPlatformSony         = 11
	BuildPlatformNintendo     = 12
	BuildPlatformXbox         = 13
	BuildPlatformWindowsPhone = 14
	BuildPlatformLinux        = 15
	BuildPlatformUnknown      = -1
)
```

### CurrentCmdVersion
enum · **wire break**

- value `Count` changed from 51 to 52
- value `Latest` changed from 50 to 51

**Also update:** `CommandRequestPacket` — it embeds this type.

```go
const (
	CurrentCmdVersionInvalid                                                                                                               = -1
	CurrentCmdVersionInitial                                                                                                               = 1
	CurrentCmdVersionTpRotationClamping                                                                                                    = 2
	CurrentCmdVersionNewBedrockCmdSystem                                                                                                   = 3
	CurrentCmdVersionExecuteUsesVec3                                                                                                       = 4
	CurrentCmdVersionCloneFixes                                                                                                            = 5
	CurrentCmdVersionUpdateAquatic                                                                                                         = 6
	CurrentCmdVersionEntitySelectorUsesVec3                                                                                                = 7
	CurrentCmdVersionContainersDontDropItemsAnymore                                                                                        = 8
	CurrentCmdVersionFiltersObeyDimensions                                                                                                 = 9
	CurrentCmdVersionExecuteAndBlockCommandAndSelfSelectorFixes                                                                            = 10
	CurrentCmdVersionInstantEffectsUseTicks                                                                                                = 11
	CurrentCmdVersionDontRegisterBrokenFunctionCommands                                                                                    = 12
	CurrentCmdVersionClearSpawnPointCommand                                                                                                = 13
	CurrentCmdVersionCloneAndTeleportRotationFixes                                                                                         = 14
	CurrentCmdVersionTeleportDimensionFixes                                                                                                = 15
	CurrentCmdVersionCloneUpdateBlockAndTimeFixes                                                                                          = 16
	CurrentCmdVersionCloneIntersectFix                                                                                                     = 17
	CurrentCmdVersionFunctionExecuteOrderAndChestSlotFix                                                                                   = 18
	CurrentCmdVersionNonTickingAreasNoLongerConsideredLoaded                                                                               = 19
	CurrentCmdVersionSpreadplayersHazardAndResolvePlayerByNameFix                                                                          = 20
	CurrentCmdVersionNewExecuteCommandSyntaxExperimentAndChestLootTableFixAndTeleportFacingVerticalUnclampedAndLocateBiomeAndFeatureMerged = 21
	CurrentCmdVersionWaterloggingAddedToStructureCommand                                                                                   = 22
	CurrentCmdVersionSelectorDistanceFilteredAndRelativeRotationFix                                                                        = 23
	CurrentCmdVersionNewSummonCommandAddedRotationOptionsAndBubbleColumnCloneFixAndExecuteInDimensionTeleportFixAndNewExecuteRotationFix   = 24
	CurrentCmdVersionNewExecuteCommandReleaseEnchantCommandLevelFixAndHasItemDataFixAndCommandDeferral                                     = 25
	CurrentCmdVersionExecuteIfScoreFixes                                                                                                   = 26
	CurrentCmdVersionReplaceItemAndLootReplaceBlockCommandsDoNotPlaceItemsIntoCauldronsFix                                                 = 27
	CurrentCmdVersionChangesToCommandOriginRotation                                                                                        = 28
	CurrentCmdVersionRemoveAuxValueParameterFromBlockCommands                                                                              = 29
	CurrentCmdVersionVolumeSelectorFixes                                                                                                   = 30
	CurrentCmdVersionEnableSummonRotation                                                                                                  = 31
	CurrentCmdVersionSummonCommandDefaultRotation                                                                                          = 32
	CurrentCmdVersionPositionalDimensionFiltering                                                                                          = 33
	CurrentCmdVersionCommandSelectorHasItemFilterNoLongerCallsSameItemFunction                                                             = 34
	CurrentCmdVersionAgentSweepingBlockTest                                                                                                = 34
	CurrentCmdVersionBlockStateEquals                                                                                                      = 35
	CurrentCmdVersionCommandPositionFix                                                                                                    = 35
	CurrentCmdVersionCommandSelectorHasItemFilterUsesDataAsDamageForSelectingDamageableItems                                               = 36
	CurrentCmdVersionExecuteDetectConditionSubcommandNotAllowNonLoadedBlocks                                                               = 37
	CurrentCmdVersionRemoveSuicideKeyword                                                                                                  = 38
	CurrentCmdVersionCloneContainerBlockEntityRemovalFix                                                                                   = 39
	CurrentCmdVersionStopSoundMusicFix                                                                                                     = 40
	CurrentCmdVersionSpreadPlayersStuckInGroundFixAndMaxHeightParameter                                                                    = 41
	CurrentCmdVersionLocateStructureOutput                                                                                                 = 42
	CurrentCmdVersionPostBlockFlattening                                                                                                   = 43
	CurrentCmdVersionTestForBlockCommandDoesNotIgnoreBlockState                                                                            = 44
	CurrentCmdVersionCount                                                                                                                 = 52
	CurrentCmdVersionLatest                                                                                                                = 51
)
```

### DisconnectFailReason
enum

- added value `MissingStructureData` = 148
- added value `UnsupportedTransport` = 149

**Also update:** `DisconnectPacket` — it embeds this type.

```go
const (
	DisconnectFailReasonUnknown                                       = 0
	DisconnectFailReasonCantConnectNoInternet                         = 1
	DisconnectFailReasonNoPermissions                                 = 2
	DisconnectFailReasonUnrecoverableError                            = 3
	DisconnectFailReasonThirdPartyBlocked                             = 4
	DisconnectFailReasonThirdPartyNoInternet                          = 5
	DisconnectFailReasonThirdPartyBadIP                               = 6
	DisconnectFailReasonThirdPartyNoServerOrServerLocked              = 7
	DisconnectFailReasonVersionMismatch                               = 8
	DisconnectFailReasonSkinIssue                                     = 9
	DisconnectFailReasonInviteSessionNotFound                         = 10
	DisconnectFailReasonEduLevelSettingsMissing                       = 11
	DisconnectFailReasonLocalServerNotFound                           = 12
	DisconnectFailReasonLegacyDisconnect                              = 13
	DisconnectFailReasonINTERNALUserLeaveGameAttempted                = 14
	DisconnectFailReasonPlatformLockedSkinsError                      = 15
	DisconnectFailReasonRealmsWorldUnassigned                         = 16
	DisconnectFailReasonRealmsServerCantConnect                       = 17
	DisconnectFailReasonRealmsServerHidden                            = 18
	DisconnectFailReasonRealmsServerDisabledBeta                      = 19
	DisconnectFailReasonRealmsServerDisabled                          = 20
	DisconnectFailReasonCrossPlatformDisabled                         = 21
	DisconnectFailReasonTESTONLYCantConnect                           = 22
	DisconnectFailReasonSessionNotFound                               = 23
	DisconnectFailReasonClientSettingsIncompatibleWithServer          = 24
	DisconnectFailReasonServerFull                                    = 25
	DisconnectFailReasonInvalidPlatformSkin                           = 26
	DisconnectFailReasonEditionVersionMismatch                        = 27
	DisconnectFailReasonEditionMismatch                               = 28
	DisconnectFailReasonLevelNewerThanExeVersion                      = 29
	DisconnectFailReasonINTERNALNoFailOccurred                        = 30
	DisconnectFailReasonBannedSkin                                    = 31
	DisconnectFailReasonTimeout                                       = 32
	DisconnectFailReasonServerNotFound                                = 33
	DisconnectFailReasonOutdatedServer                                = 34
	DisconnectFailReasonOutdatedClient                                = 35
	DisconnectFailReasonNoPremiumPlatform                             = 36
	DisconnectFailReasonMultiplayerDisabled                           = 37
	DisconnectFailReasonNoWiFi                                        = 38
	DisconnectFailReasonWorldCorruption                               = 39
	DisconnectFailReasonNoReason                                      = 40
	DisconnectFailReasonDisconnected                                  = 41
	DisconnectFailReasonInvalidPlayer                                 = 42
	DisconnectFailReasonLoggedInOtherLocation                         = 43
	DisconnectFailReasonServerIDConflict                              = 44
	DisconnectFailReasonNotAllowed                                    = 45
	DisconnectFailReasonNotAuthenticated                              = 46
	DisconnectFailReasonInvalidTenant                                 = 47
	DisconnectFailReasonUnknownPacket                                 = 48
	DisconnectFailReasonUnexpectedPacket                              = 49
	DisconnectFailReasonInvalidCommandRequestPacket                   = 50
	DisconnectFailReasonHostSuspended                                 = 51
	DisconnectFailReasonLoginPacketNoRequest                          = 52
	DisconnectFailReasonLoginPacketNoCert                             = 53
	DisconnectFailReasonMissingClient                                 = 54
	DisconnectFailReasonKicked                                        = 55
	DisconnectFailReasonKickedForExploit                              = 56
	DisconnectFailReasonKickedForIdle                                 = 57
	DisconnectFailReasonResourcePackProblem                           = 58
	DisconnectFailReasonIncompatiblePack                              = 59
	DisconnectFailReasonOutOfStorage                                  = 60
	DisconnectFailReasonInvalidLevel                                  = 61
	DisconnectFailReasonDisconnectPacket                              = 62
	DisconnectFailReasonBlockMismatch                                 = 63
	DisconnectFailReasonInvalidHeights                                = 64
	DisconnectFailReasonInvalidWidths                                 = 65
	DisconnectFailReasonConnectionLost                                = 66
	DisconnectFailReasonZombieConnection                              = 67
	DisconnectFailReasonShutdown                                      = 68
	DisconnectFailReasonReasonNotSet                                  = 69
	DisconnectFailReasonLoadingStateTimeout                           = 70
	DisconnectFailReasonResourcePackLoadingFailed                     = 71
	DisconnectFailReasonSearchingForSessionLoadingScreenFailed        = 72
	DisconnectFailReasonNetherNetProtocolVersion                      = 73
	DisconnectFailReasonSubsystemStatusError                          = 74
	DisconnectFailReasonEmptyAuthFromDiscovery                        = 75
	DisconnectFailReasonEmptyURLFromDiscovery                         = 76
	DisconnectFailReasonExpiredAuthFromDiscovery                      = 77
	DisconnectFailReasonUnknownSignalServiceSignInFailure             = 78
	DisconnectFailReasonXBLJoinLobbyFailure                           = 79
	DisconnectFailReasonUnspecifiedClientInstanceDisconnection        = 80
	DisconnectFailReasonNetherNetSessionNotFound                      = 81
	DisconnectFailReasonNetherNetCreatePeerConnection                 = 82
	DisconnectFailReasonNetherNetICE                                  = 83
	DisconnectFailReasonNetherNetConnectRequest                       = 84
	DisconnectFailReasonNetherNetConnectResponse                      = 85
	DisconnectFailReasonNetherNetNegotiationTimeout                   = 86
	DisconnectFailReasonNetherNetInactivityTimeout                    = 87
	DisconnectFailReasonStaleConnectionBeingReplaced                  = 88
	DisconnectFailReasonRealmsSessionNotFound                         = 89
	DisconnectFailReasonBadPacket                                     = 90
	DisconnectFailReasonNetherNetFailedToCreateOffer                  = 91
	DisconnectFailReasonNetherNetFailedToCreateAnswer                 = 92
	DisconnectFailReasonNetherNetFailedToSetLocalDescription          = 93
	DisconnectFailReasonNetherNetFailedToSetRemoteDescription         = 94
	DisconnectFailReasonNetherNetNegotiationTimeoutWaitingForResponse = 95
	DisconnectFailReasonNetherNetNegotiationTimeoutWaitingForAccept   = 96
	DisconnectFailReasonNetherNetIncomingConnectionIgnored            = 97
	DisconnectFailReasonNetherNetSignalingParsingFailure              = 98
	DisconnectFailReasonNetherNetSignalingUnknownError                = 99
	DisconnectFailReasonNetherNetSignalingUnicastDeliveryFailed       = 100
	DisconnectFailReasonNetherNetSignalingBroadcastDeliveryFailed     = 101
	DisconnectFailReasonNetherNetSignalingGenericDeliveryFailed       = 102
	DisconnectFailReasonEditorMismatchEditorWorld                     = 103
	DisconnectFailReasonEditorMismatchVanillaWorld                    = 104
	DisconnectFailReasonWorldTransferNotPrimaryClient                 = 105
	DisconnectFailReasonINTERNALRequestServerShutdown                 = 106
	DisconnectFailReasonClientGameSetupCancelled                      = 107
	DisconnectFailReasonClientGameSetupFailed                         = 108
	DisconnectFailReasonNoVenue                                       = 109
	DisconnectFailReasonNetherNetSignalingSigninFailed                = 110
	DisconnectFailReasonSessionAccessDenied                           = 111
	DisconnectFailReasonServiceSigninIssue                            = 112
	DisconnectFailReasonNetherNetNoSignalingChannel                   = 113
	DisconnectFailReasonNetherNetNotLoggedIn                          = 114
	DisconnectFailReasonNetherNetClientSignalingError                 = 115
	DisconnectFailReasonSubClientLoginDisabled                        = 116
	DisconnectFailReasonDeepLinkTryingToOpenDemoWorldWhileSignedIn    = 117
	DisconnectFailReasonAsyncJoinTaskDenied                           = 118
	DisconnectFailReasonRealmsTimelineRequired                        = 119
	DisconnectFailReasonGuestWithoutHost                              = 120
	DisconnectFailReasonFailedToJoinExperience                        = 121
	DisconnectFailReasonNetherNetDataChannelClosed                    = 122
	DisconnectFailReasonDiscoveryEnvironmentMismatch                  = 123
	DisconnectFailReasonHostWithoutKeys                               = 124
	DisconnectFailReasonHostSignedOut                                 = 125
	DisconnectFailReasonScriptWatchdogException                       = 126
	DisconnectFailReasonScriptMemoryLimitExceeded                     = 127
	DisconnectFailReasonStorageLowDuringGameplay                      = 128
	DisconnectFailReasonStorageFullDuringGameplay                     = 129
	DisconnectFailReasonLevelStorageCorruption                        = 130
	DisconnectFailReasonEditionMismatchVanillaToEdu                   = 131
	DisconnectFailReasonEditionMismatchEduToVanilla                   = 132
	DisconnectFailReasonEditorMismatchEditorToVanilla                 = 133
	DisconnectFailReasonEditorMismatchVanillaToEditor                 = 134
	DisconnectFailReasonDenyListed                                    = 135
	DisconnectFailReasonNonceMissing                                  = 136
	DisconnectFailReasonNonceNotFound                                 = 137
	DisconnectFailReasonNonceExpired                                  = 138
	DisconnectFailReasonNonceNotValid                                 = 139
	DisconnectFailReasonHostDisconnected                              = 140
	DisconnectFailReasonEditorJoinIntentPolicyFailure                 = 141
	DisconnectFailReasonNetherNetIdentityNotAllowed                   = 142
	DisconnectFailReasonInvalidName                                   = 143
	DisconnectFailReasonExpiredToken                                  = 144
	DisconnectFailReasonHostAcceptsNoTypeOfAuth                       = 145
	DisconnectFailReasonNotAuthenticatedFastFail                      = 146
	DisconnectFailReasonEditorNotAllowed                              = 147
	DisconnectFailReasonMissingStructureData                          = 148
	DisconnectFailReasonUnsupportedTransport                          = 149
)
```

### MapDecoration__Type
enum

- added value `AbandonedCamp` = 25
- added value `BuriedAncientCity` = 26
- added value `BuriedMineshaft` = 27
- added value `DesertPyramid` = 28
- added value `WarmOceanRuins` = 29

**Also update:** `ClientboundMapItemDataPacket` — it embeds this type.

```go
const (
	MapDecorationTypeMarkerWhite       = 0
	MapDecorationTypeMarkerGreen       = 1
	MapDecorationTypeMarkerRed         = 2
	MapDecorationTypeMarkerBlue        = 3
	MapDecorationTypeXWhite            = 4
	MapDecorationTypeTriangleRed       = 5
	MapDecorationTypeSquareWhite       = 6
	MapDecorationTypeMarkerSign        = 7
	MapDecorationTypeMarkerPink        = 8
	MapDecorationTypeMarkerOrange      = 9
	MapDecorationTypeMarkerYellow      = 10
	MapDecorationTypeMarkerTeal        = 11
	MapDecorationTypeTriangleGreen     = 12
	MapDecorationTypeSmallSquareWhite  = 13
	MapDecorationTypeMansion           = 14
	MapDecorationTypeMonument          = 15
	MapDecorationTypeNoDraw            = 16
	MapDecorationTypeVillageDesert     = 17
	MapDecorationTypeVillagePlains     = 18
	MapDecorationTypeVillageSavanna    = 19
	MapDecorationTypeVillageSnowy      = 20
	MapDecorationTypeVillageTaiga      = 21
	MapDecorationTypeJungleTemple      = 22
	MapDecorationTypeWitchHut          = 23
	MapDecorationTypeTrialChambers     = 24
	MapDecorationTypeAbandonedCamp     = 25
	MapDecorationTypeBuriedAncientCity = 26
	MapDecorationTypeBuriedMineshaft   = 27
	MapDecorationTypeDesertPyramid     = 28
	MapDecorationTypeWarmOceanRuins    = 29
	MapDecorationTypeCount             = 30
)
```

### Memory__MemoryCategory
enum

- removed value `Persona_Textures` = 58

**Also update:** `ServerboundDiagnosticsPacket` — it embeds this type.

```go
const (
	MemoryCategoryUnknown                               = 0
	MemoryCategoryInvalidSizeUnknown                    = 1
	MemoryCategoryActor                                 = 2
	MemoryCategoryActorAnimation                        = 3
	MemoryCategoryActorRendering                        = 4
	MemoryCategoryBlockTickingQueues                    = 5
	MemoryCategoryBiomeStorage                          = 6
	MemoryCategoryBlobs                                 = 7
	MemoryCategoryCereal                                = 8
	MemoryCategoryCircuitSystem                         = 9
	MemoryCategoryClient                                = 10
	MemoryCategoryCommands                              = 11
	MemoryCategoryDBStorage                             = 12
	MemoryCategoryDebug                                 = 13
	MemoryCategoryDocumentation                         = 14
	MemoryCategoryECSSystems                            = 15
	MemoryCategoryFMOD                                  = 16
	MemoryCategoryFonts                                 = 17
	MemoryCategoryImGui                                 = 18
	MemoryCategoryInput                                 = 19
	MemoryCategoryJSONUI                                = 20
	MemoryCategoryJSONUIControlFactoryJSON              = 21
	MemoryCategoryJSONUIControlTree                     = 22
	MemoryCategoryJSONUIControlTreeControlElement       = 23
	MemoryCategoryJSONUIControlTreePopulateDataBinding  = 24
	MemoryCategoryJSONUIControlTreePopulateFocus        = 25
	MemoryCategoryJSONUIControlTreePopulateLayout       = 26
	MemoryCategoryJSONUIControlTreePopulateOther        = 27
	MemoryCategoryJSONUIControlTreePopulateSprite       = 28
	MemoryCategoryJSONUIControlTreePopulateText         = 29
	MemoryCategoryJSONUIControlTreePopulateTTS          = 30
	MemoryCategoryJSONUIControlTreeVisibility           = 31
	MemoryCategoryJSONUICreateUI                        = 32
	MemoryCategoryJSONUIDefs                            = 33
	MemoryCategoryJSONUILayoutManager                   = 34
	MemoryCategoryJSONUILayoutManagerRemoveDependencies = 35
	MemoryCategoryJSONUILayoutManagerInitVariable       = 36
	MemoryCategoryLanguages                             = 37
	MemoryCategoryLevel                                 = 38
	MemoryCategoryLevelStructures                       = 39
	MemoryCategoryLevelChunk                            = 40
	MemoryCategoryLevelChunkGen                         = 41
	MemoryCategoryLevelChunkGenThreadLocal              = 42
	MemoryCategoryLightVolumeManager                    = 43
	MemoryCategoryNetwork                               = 44
	MemoryCategoryMarketplace                           = 45
	MemoryCategoryMaterialDragonCompiledDefinition      = 46
	MemoryCategoryMaterialDragonMaterial                = 47
	MemoryCategoryMaterialDragonResource                = 48
	MemoryCategoryMaterialDragonUniformMap              = 49
	MemoryCategoryMaterialRenderMaterial                = 50
	MemoryCategoryMaterialRenderMaterialGroup           = 51
	MemoryCategoryMaterialVariationManager              = 52
	MemoryCategoryMoLang                                = 53
	MemoryCategoryOreUI                                 = 54
	MemoryCategoryOreUIClient                           = 55
	MemoryCategoryPersonaPieces                         = 56
	MemoryCategoryPersonaAnimations                     = 57
	MemoryCategoryPersonaCharacters                     = 58
	MemoryCategoryPersonaSkinPacks                      = 59
	MemoryCategoryPersonaRepo                           = 60
	MemoryCategoryPlayer                                = 61
	MemoryCategoryRenderChunk                           = 62
	MemoryCategoryRenderChunkIndexBuffer                = 63
	MemoryCategoryRenderChunkVertexBuffer               = 64
	MemoryCategoryRendering                             = 65
	MemoryCategoryRenderingBgfxInit                     = 66
	MemoryCategoryRenderingBgfxStartFrame               = 67
	MemoryCategoryRenderingBlockTessellator             = 68
	MemoryCategoryRenderingEndFrame                     = 69
	MemoryCategoryRenderingGraphicsTasksInit            = 70
	MemoryCategoryRenderingLibrary                      = 71
	MemoryCategoryRenderingPolygonOperatorPool          = 72
	MemoryCategoryRenderingPBRTextureData               = 73
	MemoryCategoryRenderingRenderRegistry               = 74
	MemoryCategoryRenderingSetup                        = 75
	MemoryCategoryRenderingVertices                     = 76
	MemoryCategoryRequestLog                            = 77
	MemoryCategoryResourcePacks                         = 78
	MemoryCategorySound                                 = 79
	MemoryCategorySubChunkBiomeData                     = 80
	MemoryCategorySubChunkBlockData                     = 81
	MemoryCategorySubChunkLightData                     = 82
	MemoryCategoryTextures                              = 83
	MemoryCategoryWeatherRenderer                       = 84
	MemoryCategoryWorldGenerator                        = 85
	MemoryCategoryTasks                                 = 86
	MemoryCategoryTest                                  = 87
	MemoryCategoryTestLoadTestTags                      = 88
	MemoryCategoryScripting                             = 89
	MemoryCategoryScriptingRuntime                      = 90
	MemoryCategoryScriptingContext                      = 91
	MemoryCategoryScriptingContextBindingsMC            = 92
	MemoryCategoryScriptingContextBindingsGT            = 93
	MemoryCategoryScriptingContextRun                   = 94
	MemoryCategoryDataDrivenUI                          = 95
	MemoryCategoryDataDrivenUIDefs                      = 96
	MemoryCategoryGameface                              = 97
	MemoryCategoryGamefaceSystem                        = 98
	MemoryCategoryGamefaceDOM                           = 99
	MemoryCategoryGamefaceCSS                           = 100
	MemoryCategoryGamefaceDisplay                       = 101
	MemoryCategoryGamefaceTempAllocator                 = 102
	MemoryCategoryGamefacePoolAllocator                 = 103
	MemoryCategoryGamefaceDump                          = 104
	MemoryCategoryGamefaceMedia                         = 105
	MemoryCategoryGamefaceJSON                          = 106
	MemoryCategoryGamefaceScriptEngine                  = 107
	MemoryCategoryGamefaceScript                        = 108
	MemoryCategoryGamefaceLayout                        = 109
)
```

### MinecraftPacketIds
enum

- added value `SetPlayerFurnaceOptions` = 351
- added value `RecordStarted` = 352

```go
const (
	MinecraftPacketIdsKeepAlive                            = 0
	MinecraftPacketIdsLogin                                = 1
	MinecraftPacketIdsPlayStatus                           = 2
	MinecraftPacketIdsServerToClientHandshake              = 3
	MinecraftPacketIdsClientToServerHandshake              = 4
	MinecraftPacketIdsDisconnect                           = 5
	MinecraftPacketIdsResourcePacksInfo                    = 6
	MinecraftPacketIdsResourcePackStack                    = 7
	MinecraftPacketIdsResourcePackClientResponse           = 8
	MinecraftPacketIdsText                                 = 9
	MinecraftPacketIdsSetTime                              = 10
	MinecraftPacketIdsStartGame                            = 11
	MinecraftPacketIdsAddPlayer                            = 12
	MinecraftPacketIdsAddActor                             = 13
	MinecraftPacketIdsRemoveActor                          = 14
	MinecraftPacketIdsAddItemActor                         = 15
	MinecraftPacketIdsServerPlayerPostMovePosition         = 16
	MinecraftPacketIdsTakeItemActor                        = 17
	MinecraftPacketIdsMoveAbsoluteActor                    = 18
	MinecraftPacketIdsMovePlayer                           = 19
	MinecraftPacketIdsPassengerJump                        = 20
	MinecraftPacketIdsUpdateBlock                          = 21
	MinecraftPacketIdsAddPainting                          = 22
	MinecraftPacketIdsTickSync                             = 23
	MinecraftPacketIdsLevelSoundEventV1                    = 24
	MinecraftPacketIdsLevelEvent                           = 25
	MinecraftPacketIdsTileEvent                            = 26
	MinecraftPacketIdsActorEvent                           = 27
	MinecraftPacketIdsMobEffect                            = 28
	MinecraftPacketIdsUpdateAttributes                     = 29
	MinecraftPacketIdsInventoryTransaction                 = 30
	MinecraftPacketIdsPlayerEquipment                      = 31
	MinecraftPacketIdsMobArmorEquipment                    = 32
	MinecraftPacketIdsInteract                             = 33
	MinecraftPacketIdsBlockPickRequest                     = 34
	MinecraftPacketIdsActorPickRequest                     = 35
	MinecraftPacketIdsPlayerAction                         = 36
	MinecraftPacketIdsActorFall                            = 37
	MinecraftPacketIdsHurtArmor                            = 38
	MinecraftPacketIdsSetActorData                         = 39
	MinecraftPacketIdsSetActorMotion                       = 40
	MinecraftPacketIdsSetActorLink                         = 41
	MinecraftPacketIdsSetHealth                            = 42
	MinecraftPacketIdsSetSpawnPosition                     = 43
	MinecraftPacketIdsAnimate                              = 44
	MinecraftPacketIdsRespawn                              = 45
	MinecraftPacketIdsContainerOpen                        = 46
	MinecraftPacketIdsContainerClose                       = 47
	MinecraftPacketIdsPlayerHotbar                         = 48
	MinecraftPacketIdsInventoryContent                     = 49
	MinecraftPacketIdsInventorySlot                        = 50
	MinecraftPacketIdsContainerSetData                     = 51
	MinecraftPacketIdsCraftingData                         = 52
	MinecraftPacketIdsCraftingEvent                        = 53
	MinecraftPacketIdsGuiDataPickItem                      = 54
	MinecraftPacketIdsAdventureSettings                    = 55
	MinecraftPacketIdsBlockActorData                       = 56
	MinecraftPacketIdsPlayerInput                          = 57
	MinecraftPacketIdsFullChunkData                        = 58
	MinecraftPacketIdsSetCommandsEnabled                   = 59
	MinecraftPacketIdsSetDifficulty                        = 60
	MinecraftPacketIdsChangeDimension                      = 61
	MinecraftPacketIdsSetPlayerGameType                    = 62
	MinecraftPacketIdsPlayerList                           = 63
	MinecraftPacketIdsSimpleEvent                          = 64
	MinecraftPacketIdsLegacyTelemetryEvent                 = 65
	MinecraftPacketIdsSpawnExperienceOrb                   = 66
	MinecraftPacketIdsMapData                              = 67
	MinecraftPacketIdsMapInfoRequest                       = 68
	MinecraftPacketIdsRequestChunkRadius                   = 69
	MinecraftPacketIdsChunkRadiusUpdated                   = 70
	MinecraftPacketIdsItemFrameDropItem                    = 71
	MinecraftPacketIdsGameRulesChanged                     = 72
	MinecraftPacketIdsCamera                               = 73
	MinecraftPacketIdsBossEvent                            = 74
	MinecraftPacketIdsShowCredits                          = 75
	MinecraftPacketIdsAvailableCommands                    = 76
	MinecraftPacketIdsCommandRequest                       = 77
	MinecraftPacketIdsCommandBlockUpdate                   = 78
	MinecraftPacketIdsCommandOutput                        = 79
	MinecraftPacketIdsUpdateTrade                          = 80
	MinecraftPacketIdsUpdateEquip                          = 81
	MinecraftPacketIdsResourcePackDataInfo                 = 82
	MinecraftPacketIdsResourcePackChunkData                = 83
	MinecraftPacketIdsResourcePackChunkRequest             = 84
	MinecraftPacketIdsTransfer                             = 85
	MinecraftPacketIdsPlaySound                            = 86
	MinecraftPacketIdsStopSound                            = 87
	MinecraftPacketIdsSetTitle                             = 88
	MinecraftPacketIdsAddBehaviorTree                      = 89
	MinecraftPacketIdsStructureBlockUpdate                 = 90
	MinecraftPacketIdsShowStoreOffer                       = 91
	MinecraftPacketIdsPurchaseReceipt                      = 92
	MinecraftPacketIdsPlayerSkin                           = 93
	MinecraftPacketIdsSubclientLogin                       = 94
	MinecraftPacketIdsAutomationClientConnect              = 95
	MinecraftPacketIdsSetLastHurtBy                        = 96
	MinecraftPacketIdsBookEdit                             = 97
	MinecraftPacketIdsNPCRequest                           = 98
	MinecraftPacketIdsPhotoTransfer                        = 99
	MinecraftPacketIdsShowModalForm                        = 100
	MinecraftPacketIdsModalFormResponse                    = 101
	MinecraftPacketIdsServerSettingsRequest                = 102
	MinecraftPacketIdsServerSettingsResponse               = 103
	MinecraftPacketIdsShowProfile                          = 104
	MinecraftPacketIdsSetDefaultGameType                   = 105
	MinecraftPacketIdsRemoveObjective                      = 106
	MinecraftPacketIdsSetDisplayObjective                  = 107
	MinecraftPacketIdsSetScore                             = 108
	MinecraftPacketIdsLabTable                             = 109
	MinecraftPacketIdsUpdateBlockSynced                    = 110
	MinecraftPacketIdsMoveDeltaActor                       = 111
	MinecraftPacketIdsSetScoreboardIdentity                = 112
	MinecraftPacketIdsSetLocalPlayerAsInit                 = 113
	MinecraftPacketIdsUpdateSoftEnum                       = 114
	MinecraftPacketIdsPing                                 = 115
	MinecraftPacketIdsBlockPalette                         = 116
	MinecraftPacketIdsScriptCustomEvent                    = 117
	MinecraftPacketIdsSpawnParticleEffect                  = 118
	MinecraftPacketIdsAvailableActorIDList                 = 119
	MinecraftPacketIdsLevelSoundEventV2                    = 120
	MinecraftPacketIdsNetworkChunkPublisherUpdate          = 121
	MinecraftPacketIdsBiomeDefinitionList                  = 122
	MinecraftPacketIdsLevelSoundEvent                      = 123
	MinecraftPacketIdsLevelEventGeneric                    = 124
	MinecraftPacketIdsLecternUpdate                        = 125
	MinecraftPacketIdsVideoStreamConnect                   = 126
	MinecraftPacketIdsAddEntity                            = 127
	MinecraftPacketIdsRemoveEntity                         = 128
	MinecraftPacketIdsClientCacheStatus                    = 129
	MinecraftPacketIdsOnScreenTextureAnimation             = 130
	MinecraftPacketIdsMapCreateLockedCopy                  = 131
	MinecraftPacketIdsStructureTemplateDataExportRequest   = 132
	MinecraftPacketIdsStructureTemplateDataExportResponse  = 133
	MinecraftPacketIdsClientCacheBlobStatusPacket          = 135
	MinecraftPacketIdsClientCacheMissResponsePacket        = 136
	MinecraftPacketIdsEducationSettingsPacket              = 137
	MinecraftPacketIdsEmote                                = 138
	MinecraftPacketIdsMultiplayerSettingsPacket            = 139
	MinecraftPacketIdsSettingsCommandPacket                = 140
	MinecraftPacketIdsAnvilDamage                          = 141
	MinecraftPacketIdsCompletedUsingItem                   = 142
	MinecraftPacketIdsNetworkSettings                      = 143
	MinecraftPacketIdsPlayerAuthInputPacket                = 144
	MinecraftPacketIdsCreativeContent                      = 145
	MinecraftPacketIdsPlayerEnchantOptions                 = 146
	MinecraftPacketIdsItemStackRequest                     = 147
	MinecraftPacketIdsItemStackResponse                    = 148
	MinecraftPacketIdsPlayerArmorDamage                    = 149
	MinecraftPacketIdsCodeBuilderPacket                    = 150
	MinecraftPacketIdsUpdatePlayerGameType                 = 151
	MinecraftPacketIdsEmoteList                            = 152
	MinecraftPacketIdsPositionTrackingDBServerBroadcast    = 153
	MinecraftPacketIdsPositionTrackingDBClientRequest      = 154
	MinecraftPacketIdsDebugInfoPacket                      = 155
	MinecraftPacketIdsPacketViolationWarning               = 156
	MinecraftPacketIdsMotionPredictionHints                = 157
	MinecraftPacketIdsTriggerAnimation                     = 158
	MinecraftPacketIdsCameraShake                          = 159
	MinecraftPacketIdsPlayerFogSetting                     = 160
	MinecraftPacketIdsCorrectPlayerMovePredictionPacket    = 161
	MinecraftPacketIdsItemRegistryPacket                   = 162
	MinecraftPacketIdsFilterTextPacket                     = 163
	MinecraftPacketIdsClientBoundDebugRendererPacket       = 164
	MinecraftPacketIdsSyncActorProperty                    = 165
	MinecraftPacketIdsAddVolumeEntityPacket                = 166
	MinecraftPacketIdsRemoveVolumeEntityPacket             = 167
	MinecraftPacketIdsSimulationTypePacket                 = 168
	MinecraftPacketIdsNpcDialoguePacket                    = 169
	MinecraftPacketIdsEduURIResourcePacket                 = 170
	MinecraftPacketIdsCreatePhotoPacket                    = 171
	MinecraftPacketIdsUpdateSubChunkBlocks                 = 172
	MinecraftPacketIdsPhotoInfoRequest                     = 173
	MinecraftPacketIdsSubChunkPacket                       = 174
	MinecraftPacketIdsSubChunkRequestPacket                = 175
	MinecraftPacketIdsPlayerStartItemCooldown              = 176
	MinecraftPacketIdsScriptMessagePacket                  = 177
	MinecraftPacketIdsCodeBuilderSourcePacket              = 178
	MinecraftPacketIdsTickingAreasLoadStatus               = 179
	MinecraftPacketIdsDimensionDataPacket                  = 180
	MinecraftPacketIdsAgentAction                          = 181
	MinecraftPacketIdsChangeMobProperty                    = 182
	MinecraftPacketIdsLessonProgressPacket                 = 183
	MinecraftPacketIdsRequestAbilityPacket                 = 184
	MinecraftPacketIdsRequestPermissionsPacket             = 185
	MinecraftPacketIdsToastRequest                         = 186
	MinecraftPacketIdsUpdateAbilitiesPacket                = 187
	MinecraftPacketIdsUpdateAdventureSettingsPacket        = 188
	MinecraftPacketIdsDeathInfo                            = 189
	MinecraftPacketIdsEditorNetworkPacket                  = 190
	MinecraftPacketIdsFeatureRegistryPacket                = 191
	MinecraftPacketIdsServerStats                          = 192
	MinecraftPacketIdsRequestNetworkSettings               = 193
	MinecraftPacketIdsGameTestRequestPacket                = 194
	MinecraftPacketIdsGameTestResultsPacket                = 195
	MinecraftPacketIdsPlayerClientInputPermissions         = 196
	MinecraftPacketIdsClientCheatAbilityPacket             = 197
	MinecraftPacketIdsCameraPresets                        = 198
	MinecraftPacketIdsUnlockedRecipes                      = 199
	MinecraftPacketIdsTitleSpecificPacketsStart            = 200
	MinecraftPacketIdsTitleSpecificPacketsEnd              = 299
	MinecraftPacketIdsCameraInstruction                    = 300
	MinecraftPacketIdsCompressedBiomeDefinitionList        = 301
	MinecraftPacketIdsTrimData                             = 302
	MinecraftPacketIdsOpenSign                             = 303
	MinecraftPacketIdsAgentAnimation                       = 304
	MinecraftPacketIdsRefreshEntitlementsPacket            = 305
	MinecraftPacketIdsPlayerToggleCrafterSlotRequestPacket = 306
	MinecraftPacketIdsSetPlayerInventoryOptions            = 307
	MinecraftPacketIdsSetHudPacket                         = 308
	MinecraftPacketIdsAwardAchievementPacket               = 309
	MinecraftPacketIdsClientboundCloseScreen               = 310
	MinecraftPacketIdsClientboundLoadingScreenPacket       = 311
	MinecraftPacketIdsServerboundLoadingScreenPacket       = 312
	MinecraftPacketIdsJigsawStructureDataPacket            = 313
	MinecraftPacketIdsCurrentStructureFeaturePacket        = 314
	MinecraftPacketIdsServerboundDiagnosticsPacket         = 315
	MinecraftPacketIdsCameraAimAssist                      = 316
	MinecraftPacketIdsContainerRegistryCleanup             = 317
	MinecraftPacketIdsMovementEffect                       = 318
	MinecraftPacketIdsSetMovementAuthorityMode             = 319
	MinecraftPacketIdsCameraAimAssistActorPriority         = 339
	MinecraftPacketIdsCameraAimAssistPresets               = 320
	MinecraftPacketIdsClientCameraAimAssist                = 321
	MinecraftPacketIdsClientMovementPredictionSyncPacket   = 322
	MinecraftPacketIdsUpdateClientOptions                  = 323
	MinecraftPacketIdsPlayerVideoCapturePacket             = 324
	MinecraftPacketIdsPlayerUpdateEntityOverridesPacket    = 325
	MinecraftPacketIdsPlayerLocation                       = 326
	MinecraftPacketIdsSyncWorldClocks                      = 344
	MinecraftPacketIdsSendPartyDestinationCookie           = 349
	MinecraftPacketIdsPartyDestinationCookieResponse       = 350
	MinecraftPacketIdsSetPlayerFurnaceOptions              = 351
	MinecraftPacketIdsRecordStarted                        = 352
)
```

### persona__AnimatedTextureType
enum

- added value `None` = 0

**Also update:** `PlayerListPacket`, `PlayerSkinPacket` — each embeds this type.

```go
const (
	PersonaAnimatedTextureTypeNone        = 0
	PersonaAnimatedTextureTypeFace        = 1
	PersonaAnimatedTextureTypeBody32x32   = 2
	PersonaAnimatedTextureTypeBody128x128 = 3
)
```
