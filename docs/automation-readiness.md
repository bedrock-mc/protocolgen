# Automation readiness — September 16, 2026

CI verifies reproducible generation, tests, and an exact reviewed comparison baseline. Passing CI is not proof that every packet agrees with vanilla. This inventory separates decisions already resolved for a pinned release, new U6 schema gaps, and independent-verification limits.

## Checked-in snapshots

| Snapshot | Protocol | Packets | Adjudications | Correction operations in current files | Recorded manifest overrides |
|---|---:|---:|---:|---:|---:|
| 1.26.40 | 2168 | 229 | 168 | 81 | 81 |
| 1.26.44 | 2168 | 229 | 168 | 0 | 81 |
| 1.26.50 | 2187 | 231 | 173 | 83 | 83 |

The .44 snapshot inherits .40 and applies an evidenced RemoveScore.ObjectiveName optional-wrapper hotfix. These manifests have no unresolved/opaque nodes; their decisions are resolved for their pinned targets. The baseline .50 manifest retained eight obsolete overrides after commit `bbaa61a` deliberately removed their source corrections (TextData.LineGapHeight and seven sound-slot optional markers). Regeneration now removes those stale wrappers and records 83 overrides, matching 83 correction operations.

## U6 raw Mojang gaps

Pinned source `19e25de129227373e97dfa340af93ccf12c0feb9`, Minecraft 1.26.60-beta.23/protocol 2208: 235 packets, 745 top-level field claims, **79 reachable unresolved node occurrences**. This is raw ingestion, not a reconciled U6 release. All 235 packet directions are unknown in this raw source.

The previous 44 selector gaps are resolved by the zero-based VarUInt32 rule for wire variants. No new discriminator field is needed. Entering those alternatives reveals 50 additional unresolved nodes, so the total changes from 73 to 79 rather than falling to 29. See [the rule and PlayerList example](mojang-u6-metadata.md#variant-selectors).

### 50 ItemStackNetIdVariant occurrences (16 fields)

- `AddPlayerPacket.Carried Item`
- `AddItemActorPacket.Item`
- `InventoryTransactionPacket.Transaction` (13 occurrences)
- `MobEquipmentPacket.Item`
- `MobArmorEquipmentPacket.Head`
- `MobArmorEquipmentPacket.Torso`
- `MobArmorEquipmentPacket.Legs`
- `MobArmorEquipmentPacket.Feet`
- `MobArmorEquipmentPacket.Body`
- `InventoryContentPacket.Slots`
- `InventoryContentPacket.Storage Item`
- `InventorySlotPacket.Storage Item`
- `InventorySlotPacket.Item`
- `PlayerAuthInputPacket.Item Use Transaction` (3 occurrences)
- `PlayerAuthInputPacket.Item Stack Request` (11 occurrences)
- `ItemStackRequestPacket.Requests` (11 occurrences)

### 18 untyped payload occurrences

- `StartGamePacket.Block Properties`
- `StartGamePacket.Player Property Data`
- `AddPlayerPacket.Entity Data`
- `AddActorPacket.Actor Data`
- `AddItemActorPacket.Entity Data`
- `SetActorDataPacket.Actor Data`
- `BlockActorDataPacket.Actor Data Tags`
- `UpdateTradePacket.Data`
- `UpdateEquipPacket.Data`
- `AvailableActorIdentifiersPacket.Identifier List`
- `LevelEventGenericPacket.__[[CTD]]__`
- `StructureTemplateDataResponsePacket.Structure's NBT`
- `PositionTrackingDBServerBroadcastPacket.Position tracking data`
- `ItemRegistryPacket.Item Data`
- `SyncActorPropertyPacket.Property Data`
- `AddVolumeEntityPacket.Components`
- `JigsawStructureDataPacket.Jigsaw Structure Data Tag`
- `ClientboundDataStorePacket.Updates`

### 10 colour alternatives

All ten occurrences are inside `ClientboundAttributeLayerSyncPacket.Data`.
The colour schemas describe alternate JSON representations without a wire
selector marker. Their actual binary representation still needs evidence.

### 1 WebToken self-reference

- `ServerToClientHandshakePacket.Handshake WebToken`

The largest remaining metadata win is the precise wire representation of
[ItemStackNetIdVariant](https://github.com/Mojang/bedrock-protocol-docs/blob/19e25de129227373e97dfa340af93ccf12c0feb9/json/ItemStackNetIdVariant.json).
Next are concrete encodings for untyped payloads, including which NBT format
and framing each field uses, followed by colour and WebToken representations.
Byte examples for representative variants would help independently verify
these rules. Also reconcile optionals/defaults, bytes versus text, signedness,
length prefixes, and encode/decode asymmetry against the exact same release.
Zero unresolved nodes alone does not prove wire correctness.

## Existing manual evidence categories

- Source corrections: arbitrary byte buffers versus UTF-8 (login, chunks, item user data, features, hashes, script/debug/photo payloads); default/required/conditional fields; flattened packet structure; discriminator widths and values; fixed-u8 inventory/equipment fields; NBT editor payload; camera target mode; optional booleans; diagnostics arrays; Text category; PrimitiveShapes attached actor runtime ID.

- Adjudications: .40 has 168 complete-field selections; .50 has 173, of which 128 carry prior decisions after byte-equivalence checks and 45 are reviewed new selections. Commands `carry-adjudications` and `adjudicate-claims` already automate safe fingerprinting; humans review changes and evidence, not manually recompute hashes.

- Direction overlays:229/.40 and 231/.50 packet entries, pinned to gophertunnel pool evidence. NBT overlays:18 exact field-path encoding entries per snapshot. Add/update machine extraction with exhaustiveness guards for new versions.

- Consumer metadata overlays: naming 121/.40 and 120/.50 entries; domains 504/.40 and 510/.50; documentation 859/.40 and 834/.50. These are ergonomics/semantic mappings, not unresolved wire layouts. Seeding exists under tools/seed-gophertunnel-overlays; changed/new names and hand-written integration still need review.

- Same-protocol hotfix derivation currently supports only wrap_optional and needs pinned exact-version evidence/base hashes. Other hotfix shapes require implementing an explicit safe operation, not arbitrary replacement.

## Independent oracle verification (protocol 2168)

The regenerated `.40` manifest compared with gophertunnel `be6713da4dc051a4197f897d04835e89e9c54321` has **190 agreements, 23 reviewed divergences, 15 unresolved comparisons, and one packet without an oracle**. The reviewed divergences include extractor limitations; they are not 23 proved vanilla bugs. No disputed codec was changed merely to satisfy this oracle.

CI fails on a new or changed divergence. Each reviewed entry pins the complete manifest packet, source operation tree, compared paths, and oracle revision, excluding diagnostic checkout paths. If analysis or input changes, its acceptance fingerprint expires. Entries that stop diverging must also be removed after review. Unresolved comparisons remain visible and non-fatal; a fully automatic promotion gate must eventually require coverage rather than only this divergence baseline.

The four newly distinguished selector-binding limits (8, 72, 324, 329) are **UNRESOLVED**, not agreements: byte-operation paths match but the extractor cannot prove the branch-to-discriminant relationship.

### All reviewed divergences

The authoritative per-packet evidence and settlement criteria are in [accepted-divergences.json](../tools/gophertunnel-oracle/accepted-divergences.json).

| ID | Packet | Remaining question or tooling limitation |
|---:|---|---|
| 5 | `DisconnectPacket` | Wire conflict: the canonical Messages union uses var_u32; the pinned Disconnect marshal uses a HideDisconnectionScreen bool controlling the message strings. |
| 12 | `AddPlayerPacket` | Mixed comparison limits and wire conflict: item extra-data parsing is counted as outer data and stack-ID bool branches are not normalized to optionals. Entity metadata selector placement differs. Ability Layers additionally has a real var_u32 versus u8 prefix conflict shared with packet 187. |
| 13 | `AddActorPacket` | Extractor/control-layout conflict: the pinned EntityMetadata writer emits a cereal selector and a legacy type byte before its switch; the canonical union places its variant annotation earlier. The report cannot establish equivalent selector binding. |
| 15 | `AddItemActorPacket` | Extractor limitation: ItemInstance extra-data is read into one byte buffer, then bufReader parses canBePlacedOn/canBreak inside that buffer. The extractor incorrectly adds those inner operations to the outer packet and does not normalize the stack-ID bool branch. Entity metadata also has unproved selector binding. |
| 31 | `MobEquipmentPacket` | Extractor limitation: ItemInstance extra-data buffer parsing appears as additional outer canBePlacedOn/canBreak arrays, and the stack-ID bool-plus-conditional is not represented as the canonical optional. |
| 38 | `HurtArmorPacket` | The canonical manifest writes Armor Slots as unsigned var_u64, while the pinned gophertunnel Marshal writes ArmourSlots with Varint64 (zigzag_i64). This is an oracle conflict, not a normalization equivalence: the independent protocol documentation and the existing gophertunnel drift adjudication both identify the unsigned shape, so the hand-written oracle is the likely defect. |
| 39 | `SetActorDataPacket` | Extractor/control-layout conflict: EntityMetadata writes its cereal selector and legacy type byte before dispatch, while the canonical union annotates its variant earlier. This is not sufficient evidence of a wire width defect. |
| 49 | `InventoryContentPacket` | Extractor limitation: both Content items and StorageItem include itemUserData inner-buffer parsing in the extracted outer stream; StackNetworkID conditional presence is also not normalized. |
| 50 | `InventorySlotPacket` | Mixed comparison limits and wire conflict: the canonical Container Id is u8 while pinned InventorySlot.WindowID uses var_u32. ItemInstance also has inner-buffer and conditional-presence extraction limitations. |
| 63 | `PlayerListPacket` | Extractor limitation: PlayerListEntry uses an early return for removal entries, but the extractor also includes the subsequent addition payload. Selector placement and skin-colour endianness require separate verification after early-return handling. |
| 65 | `LegacyTelemetryEventPacket` | Wire/control conflict: the LegacyTelemetry event union and the pinned Event marshaler disagree on event discriminants and payload fields, including Contents Color. A packet-level exception does not establish any event as correct. |
| 78 | `CommandBlockUpdatePacket` | Wire conflict: the canonical Target union uses var_u32, while pinned CommandBlockUpdate uses a bool to distinguish block and actor targets. |
| 93 | `PlayerSkinPacket` | Wire conflict: Serialized Skin colour and piece tint colour values are i32le in the manifest but use BEARGB in the pinned skin serializer. |
| 108 | `SetScorePacket` | Wire/control-layout conflict: the canonical Score Info union control is u8, while pinned ScoreboardEntry writes a var_u32 selector and type-name string before switching. Variant annotations also appear at different positions. Small current discriminator values may produce the same bytes, but the comparator deliberately preserves the primitive distinction. |
| 133 | `StructureTemplateDataResponsePacket` | Wire conflict: the pinned StructureTemplateDataResponse writes Success bool and conditionally NBT, while the manifest has different presence framing. |
| 145 | `CreativeContentPacket` | Extractor limitation: CreativeContent item icons/items contain canBePlacedOn/canBreak parsing through a bufReader after the outer extraData byte slice. These inner arrays are currently counted twice in the extracted wire sequence. |
| 147 | `ItemStackRequestPacket` | Extractor/control-layout conflict: ItemStackRequest cereal action selector, legacy action byte and switch annotation appear at different positions; action-ID mappings also need exact evidence. Removing marker tokens would not prove the cases correspond. |
| 164 | `ClientboundDebugRendererPacket` | Wire conflict: the manifest adds optional bool presence around DebugMarkerData, while pinned DebugRenderer includes marker data under its string-type condition without that presence byte. |
| 187 | `UpdateAbilitiesPacket` | The canonical manifest prefixes Layers with unsigned var_u32, while the pinned gophertunnel AbilityData Marshal uses SliceUint8Length and therefore an unsigned fixed u8 count. The array-prefix distinction is intentionally preserved. |
| 325 | `PlayerUpdateEntityOverridesPacket` | Wire conflict: pinned PlayerUpdateEntityOverrides emits an extra var_u32/legacy dispatch layout while the canonical update payload contains a Type string not represented by that marshal. |
| 326 | `PlayerLocationPacket` | Extractor limitation: PlayerLocation writes the control, then a reserved zigzag_i32, then dispatches. The canonical annotation is immediately after the control; identical wire fields do not prove which value selects the branch. |
| 338 | `CameraSplinePacket` | Wire conflict: canonical spline_type is unconditional but pinned CameraSpline makes it optional; canonical progress/rotation easing is optional but pinned code writes strings unconditionally. |
| 347 | `ServerPresenceInfoPacket` | The canonical manifest contains only the optional rich-presence value, while the pinned gophertunnel PresenceInfo type adds optional ExperienceName and WorldName and an unconditional RichPresenceID inside the outer optional. This can change packet length rather than merely field grouping. |

### All unresolved comparisons

| ID | Packet | Limitation |
|---:|---|---|
| 8 | `ResourcePackClientResponsePacket` | gophertunnel: byte operations match but union selector binding is unproved because conditional paths lack variant metadata |
| 11 | `StartGamePacket` | control-flow path expansion exceeds limit 256 |
| 30 | `InventoryTransactionPacket` | control-flow path expansion exceeds limit 256 |
| 32 | `MobArmorEquipmentPacket` | control-flow path expansion exceeds limit 256 |
| 52 | `CraftingDataPacket` | control-flow path expansion exceeds limit 256 |
| 67 | `ClientboundMapItemDataPacket` | control-flow path expansion exceeds limit 256 |
| 72 | `GameRulesChangedPacket` | gophertunnel: byte operations match but union selector binding is unproved because conditional paths lack variant metadata |
| 144 | `PlayerAuthInputPacket` | control-flow path expansion exceeds limit 256 |
| 300 | `CameraInstructionPacket` | control-flow path expansion exceeds limit 256 |
| 324 | `PlayerVideoCapturePacket` | gophertunnel: byte operations match but union selector binding is unproved because conditional paths lack variant metadata |
| 328 | `PrimitiveShapesPacket` | control-flow path expansion exceeds limit 256 |
| 329 | `ServerboundPackSettingChangePacket` | gophertunnel: byte operations match but union selector binding is unproved because conditional paths lack variant metadata |
| 330 | `ClientboundDataStorePacket` | manifest: recursive node at Updates[].variant.The New Property Value.variant[] is not statically finite; manifest: recursive node at Updates[].variant.The New Property Value.variant.<value> is not statically finite; gophertunnel: recursive local helper call at ClientBoundDataStore.Updates[].Change.NewValue.ListValue[] (/Users/hashim/Library/Caches/protocolgen/gophertunnel/be6713da4dc051a4197f897d04835e89e9c54321/minecraft/protocol/data_store.go:139); gophertunnel: recursive local helper call at ClientBoundDataStore.Updates[].Change.NewValue.MapValue[] (/Users/hashim/Library/Caches/protocolgen/gophertunnel/be6713da4dc051a4197f897d04835e89e9c54321/minecraft/protocol/data_store.go:144) |
| 345 | `ClientboundAttributeLayerSyncPacket` | control-flow path expansion exceeds limit 256 |
| 348 | `ClientboundUpdateSoundDataPacket` | control-flow path expansion exceeds limit 256 |

`ServerPlayerPostMovePositionPacket` (ID 16) has no oracle implementation. Five deprecated oracle-only IDs are intentionally excluded from the generated pools: 55 `AdventureSettings`, 117 `ScriptCustomEvent`, 163 `FilterText`, 173 `PhotoInfoRequest`, 197 `ClientCheatAbility`. This is a coverage boundary, not a request to restore obsolete packets.

### Differential byte-test allowances

The separate executable differential harness targets protocol 2168. Its committed corpus covers shared packet pools, but this does not prove every branch or value. `differential/accepted-divergences.json` has **40 allowances across 30 packet IDs**: 33 re-encoded-byte allowances and 7 decode-success allowances. All use empty `input_hex`, so they apply to a packet/direction/error class rather than one exact sample. These are broader than the new AST comparison fingerprints and should be narrowed to semantic predicates or exact reproductions before unattended promotion.

Affected IDs: 8, 9, 11, 12, 15, 30, 31, 32, 49, 50, 52, 56, 65, 80, 81, 86, 93, 108, 119, 124, 148, 153, 162, 165, 166, 184, 190, 313, 326, 347. The file records each rationale. Many concern oracle canonicalization, derived values, reserved bytes, or malformed-input tolerance; they are not automatically generated-code defects.

## What remains for an automated release pipeline

1. Automatically detect matching Mojang/Endstone releases, verify full version/protocol/tree locks, prepare a new version directory, and run raw-claim/reconciliation reports. Current Make targets and CI pins are version-specific. U6 still needs a separately pinned matching runtime graph and all evidence/overlays; the enum improvement alone is not a U6 release.

2. Group changed disagreements by reusable type and present exact evidence to review. Safely carry unchanged adjudications using the existing command; never blindly carry corrections or let one source silently win. Add official metadata to remove manual decisions where it actually supplies equivalent evidence.

3. Resolve the remaining raw primitive, untyped payload, colour and WebToken gaps above and verify source optionality and directional asymmetry. Some require better metadata/frontends, some a wire capture or binary serializer trace. Human adjudication remains appropriate whenever the upstream sources contradict.

4. Complete independent verification: comparator control-flow/recursive/encapsulated-buffer handling; per-version oracle locks and coverage; live decode/encode corpus of rare branches. Axolotl parity is currently a synthetic small check in make regen; the parity adapter rejects asymmetric fields, and the CI parity command exercises only a small synthetic fixture. Recursive and encapsulated markers already have adapter support.

5. Complete downstream integration: Go already emits concrete Reader/Writer implementations and packet pools. The remaining integration work is adopting generated packet APIs in consumers and validating their actual network behavior, not writing a byte runtime from scratch. The update guide is transcription assistance. Rust borrowed views are optional ergonomics, not a wire automation blocker.

6. Capture BDS data after corrections with exact binary/codec pins; the capture workflow requires explicit EULA acceptance and does not automatically commit captures. New versions need a matching generated-codec binding: the current build tags select `.40` or `.44`, not `.50`. Matching Endstone/BDS pins are required for registry export; .44 has no matching Endstone pin. PMMP creative/recipe outputs still lack canonical block-item/legacy mapping; runtime registry cannot invent it. New versions require generated decoder binding and gophertunnel pin updates.

7. CI now regenerates `.40`, its `.44` hotfix, and `.50`, checks both languages, and uploads the protocol 2168 oracle report. The remaining release-orchestration work is a generic version matrix, per-version independent-oracle/differential coverage, and an evidence gate for automatic promotion.

A realistically fully automated flow can prepare, regenerate, compare and validate an update; it should stop for human evidence only on genuinely new ambiguities. Removing that stop without new evidence would make the generator less trustworthy.
