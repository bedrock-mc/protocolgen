# Bedrock protocol changes — 2168 to 2187

1.26.40-beta.0 → 1.26.50-beta.25

| | From | To |
| --- | --- | --- |
| Branch | `automated/1.26.40` | `r/26_u5` |
| Protocol | 2168 | 2187 |
| Game version | 1.26.40-beta.0 | 1.26.50-beta.25 |
| Upstream commit | `0e00fe80f4f3c71572ff6429de40146d1f4412fc` | `e0a1dad9341b85d3191f998c3d84b7d2c96e6c81` |
| Fixer commit | `0399c985c7bd18e050c71b5eee8150e98fcc19fb` | `0399c985c7bd18e050c71b5eee8150e98fcc19fb` |

- **8** new, **0** removed, **0** renamed, **16** modified
- **8** type(s) change the bytes on the wire
- **17** packet(s) touched, counting those that only embed a changed type
- **43** individual changes

_Derived by diffing bpd-fixer's corrected schemas for the two versions. "Wire break" means the byte layout moves; it says nothing about whether your code needs a version gate — see the per-project guides for that._

## Added Packets

### SetPlayerFurnaceOptionsPacket
packet id `351` · struct

### RecordStartedPacket
packet id `352` · struct

## Added Types

### FurnaceOptions
struct

### NoiseAlignment
struct

## Added Enums

### FurnaceLayout
enum

### FurnaceLeftTabIndex
enum

### NoiseAlignmentType
enum

### SetPlayerFurnaceOptionsPacketPayload_FurnaceType
enum

## Modified Packets

### PlaySoundPacket
packet id `86` · struct · **wire break**

- added `Bypass Listener Range Check` at ordinal 5 (boolean)
- `Server Sound Handle` moved from ordinal 5 to 6
- added `Playback Position Seconds` at ordinal 7 (float, optional)

| # | Before (PlaySoundPacket) | After |  |
| --- | --- | --- | --- |
| 0 | `Name` — string | `Name` — string |   |
| 1 | `Position` — BlockPos | `Position` — BlockPos |   |
| 2 | `Volume` — float | `Volume` — float |   |
| 3 | `Pitch` — float | `Pitch` — float |   |
| 4 | `Loop Count` — int32 (Compression) | `Loop Count` — int32 (Compression) |   |
| 5 | `Server Sound Handle` — ServerSoundHandle, optional | `Bypass Listener Range Check` — boolean | **changed** |
| 6 | — | `Server Sound Handle` — ServerSoundHandle, optional | **changed** |
| 7 | — | `Playback Position Seconds` — float, optional | **changed** |

### RequestNetworkSettingsPacket
packet id `193` · struct

- `ClientNetworkVersion` constraints changed from maximum=2168, minimum=2168 to maximum=2187, minimum=2187

| # | Before (RequestNetworkSettingsPacket) | After |  |
| --- | --- | --- | --- |
| 0 | `ClientNetworkVersion` — int32 (Big Endian) | `ClientNetworkVersion` — int32 (Big Endian) |   |

### ServerboundPackSettingChangePacket
packet id `329` · struct

- schema changed

| # | Before (ServerboundPackSettingChangePacket) | After |  |
| --- | --- | --- | --- |
| 0 | `PackId` — mce__UUID | `PackId` — mce__UUID |   |
| 1 | `PackSettingName` — string | `PackSettingName` — string |   |
| 2 | `PackSettingValue` — untyped (Compression) | `PackSettingValue` — untyped (Compression) |   |

## Modified Types

### CameraPresets
struct · **wire break**

- added `Apply Inherited Starting Rotation` at ordinal 22 (boolean)
- added `Starting Rotation` at ordinal 23 (Vec2)

| # | Before (CameraPresets) | After |  |
| --- | --- | --- | --- |
| 0 | `Name` — SharedTypes__Identifier_SharedTypes__AssetType__CameraPreset_ | `Name` — SharedTypes__Identifier_SharedTypes__AssetType__CameraPreset_ |   |
| 1 | `Inherit From` — SharedTypes__Reference_SharedTypes__AssetType__CameraPreset_ | `Inherit From` — SharedTypes__Reference_SharedTypes__AssetType__CameraPreset_ |   |
| 2 | `Pos X` — float, optional | `Pos X` — float, optional |   |
| 3 | `Pos Y` — float, optional | `Pos Y` — float, optional |   |
| 4 | `Pos Z` — float, optional | `Pos Z` — float, optional |   |
| 5 | `Rot X` — float, optional | `Rot X` — float, optional |   |
| 6 | `Rot Y` — float, optional | `Rot Y` — float, optional |   |
| 7 | `Rotation Speed` — float, optional | `Rotation Speed` — float, optional |   |
| 8 | `Snap to Target` — boolean, optional | `Snap to Target` — boolean, optional |   |
| 9 | `Horizontal Rotation Limit` — Vec2, optional | `Horizontal Rotation Limit` — Vec2, optional |   |
| 10 | `Vertical Rotation Limit` — Vec2, optional | `Vertical Rotation Limit` — Vec2, optional |   |
| 11 | `Continue Targeting` — boolean, optional | `Continue Targeting` — boolean, optional |   |
| 12 | `Block Listening Radius` — float, optional | `Block Listening Radius` — float, optional |   |
| 13 | `View Offset` — Vec2, optional | `View Offset` — Vec2, optional |   |
| 14 | `Entity Offset` — Vec3, optional | `Entity Offset` — Vec3, optional |   |
| 15 | `Radius` — float, optional | `Radius` — float, optional |   |
| 16 | `Yaw Limit Min` — float, optional | `Yaw Limit Min` — float, optional |   |
| 17 | `Yaw Limit Max` — float, optional | `Yaw Limit Max` — float, optional |   |
| 18 | `Listener` — SharedTypes__Comprehensive__CameraPreset__AudioListener (Enum-as-Value), optional | `Listener` — SharedTypes__Comprehensive__CameraPreset__AudioListener (Enum-as-Value), optional |   |
| 19 | `Player Effects` — boolean, optional | `Player Effects` — boolean, optional |   |
| 20 | `Aim Assist` — SharedTypes__Comprehensive__CameraAimAssistCommandDefinition, optional | `Aim Assist` — SharedTypes__Comprehensive__CameraAimAssistCommandDefinition, optional |   |
| 21 | `Control Scheme` — Control_Scheme (Enum-as-Value), optional | `Control Scheme` — Control_Scheme (Enum-as-Value), optional |   |
| 22 | — | `Apply Inherited Starting Rotation` — boolean | **changed** |
| 23 | — | `Starting Rotation` — Vec2 | **changed** |

**Also affects:** `CameraPresetsPacket`

### DimensionDefinition
struct · **wire break**

- added `Minimum Y` at ordinal 0 (int32 (Compression))
- added `Height Range` at ordinal 1 (int32 (Compression))
- added `Default Biome` at ordinal 5 (string)
- removed `Height Maximum` from ordinal 0
- removed `Height Minimum` from ordinal 1

| # | Before (DimensionDefinition) | After |  |
| --- | --- | --- | --- |
| 0 | `Height Maximum` — int32 (Compression) | `Minimum Y` — int32 (Compression) | **changed** |
| 1 | `Height Minimum` — int32 (Compression) | `Height Range` — int32 (Compression) | **changed** |
| 2 | `Generator Type` — int32 (Compression, Enum-as-Value) | `Generator Type` — int32 (Compression, Enum-as-Value) |   |
| 3 | `Dimension Type` — DimensionType | `Dimension Type` — DimensionType |   |
| 4 | `Pack Id` — mce__UUID | `Pack Id` — mce__UUID |   |
| 5 | — | `Default Biome` — string | **changed** |

**Also affects:** `DimensionDataPacket`

### EntityDiagnosticTimingInfo
struct · **wire break**

- added `Position` at ordinal 4 (Vec3)
- added `Dimension` at ordinal 5 (string)

| # | Before (EntityDiagnosticTimingInfo) | After |  |
| --- | --- | --- | --- |
| 0 | `Display Name` — string | `Display Name` — string |   |
| 1 | `Entity` — string | `Entity` — string |   |
| 2 | `Time in NS` — uint64 | `Time in NS` — uint64 |   |
| 3 | `Percent of Total` — uint8 | `Percent of Total` — uint8 |   |
| 4 | — | `Position` — Vec3 | **changed** |
| 5 | — | `Dimension` — string | **changed** |

**Also affects:** `ServerboundDiagnosticsPacket`

### EnvironmentAttributeData
struct · **wire break**

- added `NoiseAlignment` at ordinal 9 (NoiseAlignment)

| # | Before (EnvironmentAttributeData) | After |  |
| --- | --- | --- | --- |
| 0 | `AttributeName` — string | `AttributeName` — string |   |
| 1 | `FromAttribute` — untyped (Compression), optional | `FromAttribute` — untyped (Compression), optional |   |
| 2 | `Attribute` — untyped (Compression) | `Attribute` — untyped (Compression) |   |
| 3 | `ToAttribute` — untyped (Compression), optional | `ToAttribute` — untyped (Compression), optional |   |
| 4 | `CurrentTransitionTicks` — uint32 | `CurrentTransitionTicks` — uint32 |   |
| 5 | `TotalTransitionTicks` — uint32 | `TotalTransitionTicks` — uint32 |   |
| 6 | `Easing` — easing_function | `Easing` — easing_function |   |
| 7 | `LocalTransitionTicks` — uint32 | `LocalTransitionTicks` — uint32 |   |
| 8 | `NoiseTransition` — boolean | `NoiseTransition` — boolean |   |
| 9 | — | `NoiseAlignment` — NoiseAlignment | **changed** |

**Also affects:** `ClientboundAttributeLayerSyncPacket`

### MoveActorDeltaData
struct · **wire break**

- added `Ticks` at ordinal 11 (uint64 (Compression))

| # | Before (MoveActorDeltaData) | After |  |
| --- | --- | --- | --- |
| 0 | `Actor Runtime ID` — ActorRuntimeID | `Actor Runtime ID` — ActorRuntimeID |   |
| 1 | `New Position X` — float, optional | `New Position X` — float, optional |   |
| 2 | `New Position Y` — float, optional | `New Position Y` — float, optional |   |
| 3 | `New Position Z` — float, optional | `New Position Z` — float, optional |   |
| 4 | `Rotation X` — int8, optional | `Rotation X` — int8, optional |   |
| 5 | `Rotation Y` — int8, optional | `Rotation Y` — int8, optional |   |
| 6 | `Rotation Y Head` — int8, optional | `Rotation Y Head` — int8, optional |   |
| 7 | `Is On Ground` — boolean | `Is On Ground` — boolean |   |
| 8 | `Force Move` — boolean | `Force Move` — boolean |   |
| 9 | `Force Move Local Entity` — boolean | `Force Move Local Entity` — boolean |   |
| 10 | `Force Completion` — boolean | `Force Completion` — boolean |   |
| 11 | — | `Ticks` — uint64 (Compression) | **changed** |

**Also affects:** `MoveActorDeltaPacket`

### TextDataPayload
struct · **wire break**

- added `LineGapHeight` at ordinal 3 (float, optional)
- `DepthTest` moved from ordinal 3 to 4
- `ShowBackface` moved from ordinal 4 to 5
- `ShowTextBackface` moved from ordinal 5 to 6

| # | Before (TextDataPayload) | After |  |
| --- | --- | --- | --- |
| 0 | `Text` — string | `Text` — string |   |
| 1 | `UseRotation` — boolean | `UseRotation` — boolean |   |
| 2 | `BackgroundColor` — Color, optional | `BackgroundColor` — Color, optional |   |
| 3 | `DepthTest` — boolean | `LineGapHeight` — float, optional | **changed** |
| 4 | `ShowBackface` — boolean | `DepthTest` — boolean | **changed** |
| 5 | `ShowTextBackface` — boolean | `ShowBackface` — boolean | **changed** |
| 6 | — | `ShowTextBackface` — boolean | **changed** |

**Also affects:** `PrimitiveShapesPacket`

## Modified Enums

### BuildPlatform
enum

- added value `Nintendo` = 12
- removed value `Nx` = 12

**Also affects:** `AddPlayerPacket`, `PlayerListPacket`

### CurrentCmdVersion
enum · **wire break**

- value `Count` changed from 51 to 52
- value `Latest` changed from 50 to 51

**Also affects:** `CommandRequestPacket`

### DisconnectFailReason
enum

- added value `MissingStructureData` = 148
- added value `UnsupportedTransport` = 149

**Also affects:** `DisconnectPacket`

### MapDecoration__Type
enum

- added value `AbandonedCamp` = 25
- added value `BuriedAncientCity` = 26
- added value `BuriedMineshaft` = 27
- added value `DesertPyramid` = 28
- added value `WarmOceanRuins` = 29

**Also affects:** `ClientboundMapItemDataPacket`

### Memory__MemoryCategory
enum

- removed value `Persona_Textures` = 58

**Also affects:** `ServerboundDiagnosticsPacket`

### MinecraftPacketIds
enum

- added value `SetPlayerFurnaceOptions` = 351
- added value `RecordStarted` = 352

### persona__AnimatedTextureType
enum

- added value `None` = 0

**Also affects:** `PlayerListPacket`, `PlayerSkinPacket`

## Fields the docs leave untyped

These carry a raw NBT compound with no schema type. Listed for completeness — their contents are not tracked by this diff.

- `AddVolumeEntityPacketPayload` · `Components`
- `AvailableActorIdentifiersPacketPayload` · `Identifier List`
- `BlockActorDataPacketPayload` · `Actor Data Tags`
- `DataItemCompoundTagPayload` · `Value`
- `Data_Store_Change` · `The New Property Value`
- `ItemData` · `Item Component Data`
- `JigsawStructureDataPacketPayload` · `Jigsaw Structure Data Tag`
- `LevelEventGenericPacketPayload` · `__[[CTD]]__`
- `PositionTrackingDBServerBroadcastPacketPayload` · `Position tracking data`
- `ServerBlockProperty` · `Block Definition`
- `StartGamePacketPayload` · `Player Property Data`
- `StructureTemplateDataResponsePacketPayload` · `Structure's NBT`
- `SyncActorPropertyPacketPayload` · `Property Data`
- `UpdateEquipPacketPayload` · `Data`
- `UpdateTradePacketPayload` · `Data`
