# Bedrock protocol changes — 2168 to 2169

1.26.40-beta.0 → 1.26.50

| | From | To |
| --- | --- | --- |
| Branch | `automated/1.26.40` | `r/26_u4` |
| Protocol | 2168 | 2169 |
| Game version | 1.26.40-beta.0 | 1.26.50 |
| Upstream commit | `0e00fe80f4` | `0f6a0bff19` |
| Fixer commit | `9fa8eb75f9` | `9fa8eb75f9` |

- **0** new, **0** removed, **0** renamed, **3** modified
- **1** type(s) change the bytes on the wire
- **4** packet(s) touched, counting those that only embed a changed type
- **6** individual changes

_Derived by diffing bpd-fixer's corrected schemas for the two versions. "Wire break" means the byte layout moves; it says nothing about whether your code needs a version gate — see the per-project guides for that._

## Modified Packets

### RequestNetworkSettingsPacket
packet id `193` · struct

- `ClientNetworkVersion` constraints changed from maximum=2168, minimum=2168 to maximum=2169, minimum=2169

| # | Before (RequestNetworkSettingsPacket) | After |  |
| --- | --- | --- | --- |
| 0 | `ClientNetworkVersion` — int32 (Big Endian) | `ClientNetworkVersion` — int32 (Big Endian) |   |

## Modified Types

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
