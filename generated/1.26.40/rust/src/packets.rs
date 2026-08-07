// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use crate::enums::*;
use crate::types::*;
use crate::wire;

#[derive(Clone, Debug, Default, PartialEq)]
pub struct Login {
    pub client_network_version: wire::I32BE,
    pub connection_request: String,
}

impl Login {
    pub const ID: u32 = 1;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayStatus {
    pub status: PlayStatusType,
}

impl PlayStatus {
    pub const ID: u32 = 2;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerToClientHandshake {
    pub handshake_web_token: String,
}

impl ServerToClientHandshake {
    pub const ID: u32 = 3;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientToServerHandshake {
}

impl ClientToServerHandshake {
    pub const ID: u32 = 4;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Disconnect {
    pub reason: ConnectionDisconnectFailReason,
    pub messages: DisconnectMessages,
}

impl Disconnect {
    pub const ID: u32 = 5;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePacksInfo {
    pub resource_pack_required: bool,
    pub has_addon_packs: bool,
    pub has_scripts: bool,
    pub force_disable_vibrant_visuals: bool,
    pub world_template_id_and_version: PackIdVersion,
    pub resource_packs: Vec<PackInfoData>,
}

impl ResourcePacksInfo {
    pub const ID: u32 = 6;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackStack {
    pub texture_pack_required: bool,
    pub texture_pack_list: Vec<PackInstanceId>,
    pub base_game_version: String,
    pub experiments: Experiments,
    pub include_editor_packs: bool,
}

impl ResourcePackStack {
    pub const ID: u32 = 7;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackClientResponse {
    pub response: ResourcePackClientResponseData,
}

impl ResourcePackClientResponse {
    pub const ID: u32 = 8;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Text {
    pub localize: bool,
    pub body: TextData,
    pub sender_xuid: String,
    pub platform_id: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub filtered_message: Option<String>,
}

impl Text {
    pub const ID: u32 = 9;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetTime {
    pub time: wire::ZigZag32,
}

impl SetTime {
    pub const ID: u32 = 10;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StartGame {
    pub entity_id: ActorUniqueID,
    pub runtime_id: ActorRuntimeID,
    pub game_type: GameType,
    pub position: glam::Vec3,
    pub rotation: glam::Vec2,
    pub settings: LevelSettings,
    pub level_id: String,
    pub level_name: String,
    pub template_content_identity: String,
    pub is_trial: bool,
    pub movement_settings: SyncedPlayerMovementSettings,
    pub level_current_time: wire::U64LE,
    pub enchantment_seed: wire::ZigZag32,
    pub block_properties: Vec<ServerBlockProperty>,
    pub multiplayer_correlation_id: String,
    pub enable_item_stack_net_manager: bool,
    pub server_version: String,
    pub player_property_data: Nbt,
    pub server_block_type_registry_checksum: wire::U64LE,
    pub world_template_id: uuid::Uuid,
    pub server_enabled_client_side_generation: bool,
    pub block_network_ids_are_hashes: bool,
    pub network_permissions: NetworkPermissions,
    /// Wire presence: optional value is preceded by a presence marker.
    pub server_configuration_join_info: Option<ServerConfigurationServerConfigurationJoinInfo>,
    pub server_telemetry_data: SocialEventsServerTelemetryData,
}

impl StartGame {
    pub const ID: u32 = 11;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddPlayer {
    pub uuid: uuid::Uuid,
    pub player_name: String,
    pub target_runtime_id: ActorRuntimeID,
    pub platform_chat_id: String,
    pub position: glam::Vec3,
    pub velocity: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: wire::F32LE,
    pub carried_item: NetworkItemStackDescriptorSerializedData,
    pub player_game_type: GameType,
    pub entity_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub abilities_data: SerializedAbilitiesData,
    pub actor_links: Vec<EntityLink>,
    pub device_id: String,
    pub build_platform: BuildPlatform,
}

impl AddPlayer {
    pub const ID: u32 = 12;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddActor {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub actor_type: String,
    pub position: glam::Vec3,
    pub velocity: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: wire::F32LE,
    pub y_body_rotation: wire::F32LE,
    pub attributes_list: Vec<SyncedAttribute>,
    pub actor_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub actor_links: Vec<EntityLink>,
}

impl AddActor {
    pub const ID: u32 = 13;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RemoveActor {
    pub target_actor_id: ActorUniqueID,
}

impl RemoveActor {
    pub const ID: u32 = 14;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddItemActor {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub item: NetworkItemStackDescriptorSerializedData,
    pub position: glam::Vec3,
    pub velocity: glam::Vec3,
    pub entity_data: SynchedActorDataCopyableDataList,
    pub is_from_fishing: bool,
}

impl AddItemActor {
    pub const ID: u32 = 15;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerPlayerPostMovePosition {
    pub pos: glam::Vec3,
}

impl ServerPlayerPostMovePosition {
    pub const ID: u32 = 16;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TakeItemActor {
    pub item_runtime_id: ActorRuntimeID,
    pub actor_runtime_id: ActorRuntimeID,
}

impl TakeItemActor {
    pub const ID: u32 = 17;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MoveActorAbsolute {
    pub move_data: MoveActorAbsoluteData,
}

impl MoveActorAbsolute {
    pub const ID: u32 = 18;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MovePlayer {
    pub player_runtime_id: ActorRuntimeID,
    pub position: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: wire::F32LE,
    pub position_mode: PlayerPositionModeComponentPositionMode,
    pub on_ground: bool,
    pub riding_runtime_id: ActorRuntimeID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub teleport_data: Option<MovePlayerTeleportData>,
    pub tick: PlayerInputTick,
}

impl MovePlayer {
    pub const ID: u32 = 19;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateBlock {
    pub block_position: BlockPos,
    pub block_runtime_id: wire::VarUInt,
    pub flags: wire::VarUInt,
    pub layer: wire::VarUInt,
}

impl UpdateBlock {
    pub const ID: u32 = 21;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddPainting {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub position: glam::Vec3,
    pub direction: wire::ZigZag32,
    pub motif: String,
}

impl AddPainting {
    pub const ID: u32 = 22;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelEvent {
    pub event_id: wire::ZigZag32,
    pub position: glam::Vec3,
    pub data: wire::ZigZag32,
}

impl LevelEvent {
    pub const ID: u32 = 25;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockEvent {
    pub block_position: BlockPos,
    pub event_type: wire::ZigZag32,
    pub event_value: wire::ZigZag32,
}

impl BlockEvent {
    pub const ID: u32 = 26;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorEvent {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: ActorEventType,
    pub data: wire::ZigZag32,
    /// Wire presence: optional value is preceded by a presence marker.
    pub fire_at_position: Option<glam::Vec3>,
}

impl ActorEvent {
    pub const ID: u32 = 27;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MobEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: MobEffectEvent,
    pub effect_id: wire::ZigZag32,
    pub effect_amplifier: wire::ZigZag32,
    pub show_particles: bool,
    pub effect_duration_ticks: wire::ZigZag32,
    pub tick: PlayerInputTick,
    pub ambient: bool,
}

impl MobEffect {
    pub const ID: u32 = 28;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateAttributes {
    pub target_runtime_id: ActorRuntimeID,
    pub attribute_list: Vec<AttributeData>,
    pub tick: PlayerInputTick,
}

impl UpdateAttributes {
    pub const ID: u32 = 29;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryTransaction {
    pub legacy_request_id: ItemStackLegacyRequestID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub legacy_set_item_slots: Option<Vec<LegacySetSlot>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub transaction: Option<InventoryTransactionValue>,
}

impl InventoryTransaction {
    pub const ID: u32 = 30;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MobEquipment {
    pub target_runtime_id: ActorRuntimeID,
    pub item: NetworkItemStackDescriptorSerializedData,
    pub slot: wire::U8,
    pub selected_slot: wire::U8,
    pub container_id: wire::U8,
}

impl MobEquipment {
    pub const ID: u32 = 31;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MobArmorEquipment {
    pub target_runtime_id: ActorRuntimeID,
    pub head: NetworkItemStackDescriptorSerializedData,
    pub torso: NetworkItemStackDescriptorSerializedData,
    pub legs: NetworkItemStackDescriptorSerializedData,
    pub feet: NetworkItemStackDescriptorSerializedData,
    pub body: NetworkItemStackDescriptorSerializedData,
}

impl MobArmorEquipment {
    pub const ID: u32 = 32;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Interact {
    pub action: InteractAction,
    pub target_runtime_id: ActorRuntimeID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub position: Option<glam::Vec3>,
}

impl Interact {
    pub const ID: u32 = 33;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockPickRequest {
    pub position: BlockPos,
    pub with_data: bool,
    pub max_slots: wire::U8,
}

impl BlockPickRequest {
    pub const ID: u32 = 34;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorPickRequest {
    pub actor_id: wire::I64LE,
    pub max_slots: wire::U8,
    pub with_data: bool,
}

impl ActorPickRequest {
    pub const ID: u32 = 35;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerAction {
    pub player_runtime_id: ActorRuntimeID,
    pub action: PlayerActionType,
    pub block_position: BlockPos,
    pub result_pos: BlockPos,
    pub face: wire::ZigZag32,
}

impl PlayerAction {
    pub const ID: u32 = 36;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct HurtArmor {
    pub cause: wire::ZigZag32,
    pub damage: wire::ZigZag32,
    pub armor_slots: wire::VarULong,
}

impl HurtArmor {
    pub const ID: u32 = 38;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetActorData {
    pub target_runtime_id: ActorRuntimeID,
    pub actor_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub tick: PlayerInputTick,
}

impl SetActorData {
    pub const ID: u32 = 39;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetActorMotion {
    pub target_runtime_id: ActorRuntimeID,
    pub motion: glam::Vec3,
    pub tick: PlayerInputTick,
}

impl SetActorMotion {
    pub const ID: u32 = 40;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetActorLink {
    pub link: EntityLink,
}

impl SetActorLink {
    pub const ID: u32 = 41;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetHealth {
    pub health: wire::ZigZag32,
}

impl SetHealth {
    pub const ID: u32 = 42;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetSpawnPosition {
    pub spawn_position_type: SpawnPositionType,
    pub block_position: BlockPos,
    pub dimension_type: DimensionType,
    pub spawn_block_pos: BlockPos,
}

impl SetSpawnPosition {
    pub const ID: u32 = 43;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Animate {
    pub action: AnimateAction,
    pub target_actor_runtime_id: ActorRuntimeID,
    pub data: wire::F32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub swing_source: Option<String>,
}

impl Animate {
    pub const ID: u32 = 44;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Respawn {
    pub position: glam::Vec3,
    pub state: PlayerRespawnState,
    pub player_runtime_id: ActorRuntimeID,
}

impl Respawn {
    pub const ID: u32 = 45;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerOpen {
    pub container_id: wire::U8,
    pub container_type: wire::U8,
    pub position: BlockPos,
    pub target_actor_id: ActorUniqueID,
}

impl ContainerOpen {
    pub const ID: u32 = 46;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerClose {
    pub container_id: wire::U8,
    pub container_type: wire::U8,
    pub server_initiated_close: bool,
}

impl ContainerClose {
    pub const ID: u32 = 47;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerHotbar {
    pub selected_slot: wire::VarUInt,
    pub container_id: wire::U8,
    pub should_select_slot: bool,
}

impl PlayerHotbar {
    pub const ID: u32 = 48;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryContent {
    pub container_id: wire::VarUInt,
    pub slots: Vec<NetworkItemStackDescriptorSerializedData>,
    pub full_container_name: FullContainerName,
    pub storage_item: NetworkItemStackDescriptorSerializedData,
}

impl InventoryContent {
    pub const ID: u32 = 49;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventorySlot {
    pub container_id: wire::U8,
    pub slot: wire::VarUInt,
    /// Wire presence: optional value is preceded by a presence marker.
    pub full_container_name: Option<FullContainerName>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub storage_item: Option<NetworkItemStackDescriptorSerializedData>,
    pub item: NetworkItemStackDescriptorSerializedData,
}

impl InventorySlot {
    pub const ID: u32 = 50;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerSetData {
    pub container_id: wire::U8,
    pub id: wire::ZigZag32,
    pub value: wire::ZigZag32,
}

impl ContainerSetData {
    pub const ID: u32 = 51;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CraftingData {
    pub shaped_recipes: Vec<ShapedRecipe>,
    pub shapeless_recipes: Vec<ShapelessRecipe>,
    pub multi_recipes: Vec<MultiRecipe>,
    pub user_data_shapeless_recipes: Vec<ShapelessRecipe>,
    pub shapeless_chemistry_recipes: Vec<ShapelessRecipe>,
    pub shaped_chemistry_recipes: Vec<ShapedRecipe>,
    pub smithing_transform_recipes: Vec<SmithingTransformRecipe>,
    pub smithing_trim_recipes: Vec<SmithingTrimRecipe>,
    pub potion_mixes: Vec<PotionMixDataEntry>,
    pub container_mixes: Vec<ContainerMixDataEntry>,
    pub material_reducers: Vec<MaterialReducerDataEntry>,
    pub clear_recipes: bool,
}

impl CraftingData {
    pub const ID: u32 = 52;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GuiDataPickItem {
    pub item_name: String,
    pub item_effect_name: String,
    pub slot: wire::I32LE,
}

impl GuiDataPickItem {
    pub const ID: u32 = 54;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockActorData {
    pub block_position: BlockPos,
    pub actor_data_tags: Nbt,
}

impl BlockActorData {
    pub const ID: u32 = 56;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelChunk {
    pub chunk_position: ChunkPos,
    pub dimension_id: DimensionType,
    pub sub_chunks_count: wire::VarUInt,
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_request_sub_chunk_limit: Option<wire::ZigZag32>,
    pub cache_enabled: bool,
    pub cache_metadata: Vec<SubChunkMetadata>,
    pub serialized_chunk_data: bytes::Bytes,
}

impl LevelChunk {
    pub const ID: u32 = 58;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetCommandsEnabled {
    pub commands_enabled: bool,
}

impl SetCommandsEnabled {
    pub const ID: u32 = 59;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetDifficulty {
    pub difficulty: wire::VarUInt,
}

impl SetDifficulty {
    pub const ID: u32 = 60;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChangeDimension {
    pub dimension_id: DimensionType,
    pub position: glam::Vec3,
    pub respawn: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub loading_screen_id: Option<wire::U32LE>,
}

impl ChangeDimension {
    pub const ID: u32 = 61;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetPlayerGameType {
    pub player_game_type: GameType,
}

impl SetPlayerGameType {
    pub const ID: u32 = 62;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerList {
    pub entries: Vec<PlayerListData>,
}

impl PlayerList {
    pub const ID: u32 = 63;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SimpleEvent {
    pub type_: Subtype,
}

impl SimpleEvent {
    pub const ID: u32 = 64;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LegacyTelemetryEvent {
    pub target_actor_id: ActorUniqueID,
    pub event_type: LegacyTelemetryType,
    pub use_player_id: bool,
    pub event_data: Event,
}

impl LegacyTelemetryEvent {
    pub const ID: u32 = 65;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SpawnExperienceOrb {
    pub position: glam::Vec3,
    pub xp_value: wire::ZigZag32,
}

impl SpawnExperienceOrb {
    pub const ID: u32 = 66;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundMapItemData {
    pub map_id: ActorUniqueID,
    pub dimension: wire::U8,
    pub is_locked: bool,
    pub map_origin: BlockPos,
    /// Wire presence: optional value is preceded by a presence marker.
    pub creation_map_ids: Option<Vec<ActorUniqueID>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub scale: Option<wire::I8>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub tracked_actor_ids: Option<Vec<MapItemTrackedActorUniqueId>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub decorations: Option<Vec<MapDecoration>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub width: Option<wire::ZigZag32>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub height: Option<wire::ZigZag32>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub start_x: Option<wire::ZigZag32>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub start_y: Option<wire::ZigZag32>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub pixels: Option<Vec<wire::U32LE>>,
}

impl ClientboundMapItemData {
    pub const ID: u32 = 67;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapInfoRequest {
    pub map_unique_id: ActorUniqueID,
    pub client_pixels_list: Vec<PixelRequest>,
}

impl MapInfoRequest {
    pub const ID: u32 = 68;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestChunkRadius {
    pub chunk_radius: wire::ZigZag32,
    pub max_chunk_radius: wire::U8,
}

impl RequestChunkRadius {
    pub const ID: u32 = 69;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChunkRadiusUpdated {
    pub chunk_radius: wire::ZigZag32,
}

impl ChunkRadiusUpdated {
    pub const ID: u32 = 70;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameRulesChanged {
    pub rule_data: GameRulesChangedData,
}

impl GameRulesChanged {
    pub const ID: u32 = 72;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Camera {
    pub camera_id: ActorUniqueID,
    pub target_player_id: ActorUniqueID,
}

impl Camera {
    pub const ID: u32 = 73;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BossEvent {
    pub target_actor_id: ActorUniqueID,
    pub player_id: ActorUniqueID,
    pub event_type: BossEventUpdateType,
    pub name: String,
    pub filtered_name: String,
    pub health_percent: wire::F32LE,
    pub color: BossBarColor,
    pub overlay: BossBarOverlay,
}

impl BossEvent {
    pub const ID: u32 = 74;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShowCredits {
    pub player_runtime_id: ActorRuntimeID,
    pub credits_state: wire::ZigZag32,
}

impl ShowCredits {
    pub const ID: u32 = 75;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommands {
    pub enum_values: Vec<String>,
    pub chained_subcommand_values: Vec<String>,
    pub post_fixes: Vec<String>,
    pub enum_data: Vec<CommandEnum>,
    pub chained_subcommand_data: Vec<ChainedSubcommand>,
    pub commands: Vec<Command>,
    pub soft_enums: Vec<DynamicEnum>,
    pub constraints: Vec<CommandEnumConstraint>,
}

impl AvailableCommands {
    pub const ID: u32 = 76;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandRequest {
    pub command: String,
    pub origin: CommandOriginData,
    pub is_internal: bool,
    pub version: String,
}

impl CommandRequest {
    pub const ID: u32 = 77;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandBlockUpdate {
    pub target: CommandBlockUpdateData,
    pub command: String,
    pub last_output: String,
    pub name: String,
    pub filtered_name: String,
    pub track_output: bool,
    pub tick_delay: wire::I32LE,
    pub execute_on_first_tick: bool,
}

impl CommandBlockUpdate {
    pub const ID: u32 = 78;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOutput {
    pub origin_data: CommandOriginData,
    pub output: CommandOutputData,
}

impl CommandOutput {
    pub const ID: u32 = 79;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateTrade {
    pub container_id: wire::U8,
    pub type_: wire::U8,
    pub size: wire::ZigZag32,
    pub trader_tier: wire::ZigZag32,
    pub entity_unique_id: ActorUniqueID,
    pub last_trading_player: ActorUniqueID,
    pub display_name: String,
    pub use_new_trade_screen: bool,
    pub using_economy_trade: bool,
    pub data: Nbt,
}

impl UpdateTrade {
    pub const ID: u32 = 80;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateEquip {
    pub container_id: wire::U8,
    pub type_: wire::U8,
    pub size: wire::ZigZag32,
    pub entity_unique_id: ActorUniqueID,
    pub data: Nbt,
}

impl UpdateEquip {
    pub const ID: u32 = 81;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackDataInfo {
    pub resource_name: String,
    pub chunk_size: wire::U32LE,
    pub number_of_chunks: wire::U32LE,
    pub file_size: wire::U64LE,
    pub file_hash: String,
    pub is_premium_pack: bool,
    pub pack_type: wire::U8,
}

impl ResourcePackDataInfo {
    pub const ID: u32 = 82;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackChunkData {
    pub resource_name: String,
    pub chunk_id: wire::U32LE,
    pub byte_offset: wire::U64LE,
    pub chunk_data: bytes::Bytes,
}

impl ResourcePackChunkData {
    pub const ID: u32 = 83;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackChunkRequest {
    pub resource_name: String,
    pub chunk: wire::I32LE,
}

impl ResourcePackChunkRequest {
    pub const ID: u32 = 84;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Transfer {
    pub server_address: String,
    pub server_port: wire::U16LE,
    pub reload_world: bool,
    /// Wire presence: optional value is preceded by a presence marker.
    pub gatherings_configuration: Option<ServerConfigurationGatheringsConfigurationJoinInfo>,
}

impl Transfer {
    pub const ID: u32 = 85;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlaySound {
    pub name: String,
    pub position: BlockPos,
    pub volume: wire::F32LE,
    pub pitch: wire::F32LE,
    pub loop_count: wire::ZigZag32,
    /// Wire presence: optional value is preceded by a presence marker.
    pub server_sound_handle: Option<ServerSoundHandle>,
}

impl PlaySound {
    pub const ID: u32 = 86;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StopSound {
    pub sound_name: String,
    pub stop_all_sounds: bool,
    pub stop_music_legacy: bool,
}

impl StopSound {
    pub const ID: u32 = 87;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetTitle {
    pub title_type: TitleType,
    pub title_text: String,
    pub fade_in_time: wire::ZigZag32,
    pub stay_time: wire::ZigZag32,
    pub fade_out_time: wire::ZigZag32,
    pub xuid: String,
    pub platform_online_id: String,
    pub filtered_title_message: String,
}

impl SetTitle {
    pub const ID: u32 = 88;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddBehaviorTree {
    pub behavior_tree_structure_json: String,
}

impl AddBehaviorTree {
    pub const ID: u32 = 89;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureBlockUpdate {
    pub block_position: BlockPos,
    pub structure_data: StructureEditorData,
    pub trigger: bool,
    pub is_waterlogged: bool,
}

impl StructureBlockUpdate {
    pub const ID: u32 = 90;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShowStoreOffer {
    pub offer_id: uuid::Uuid,
    pub redirect_type: ShowStoreOfferRedirectType,
}

impl ShowStoreOffer {
    pub const ID: u32 = 91;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PurchaseReceipt {
    pub purchase_receipts: Vec<String>,
}

impl PurchaseReceipt {
    pub const ID: u32 = 92;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerSkin {
    pub uuid: uuid::Uuid,
    pub serialized_skin: SerializedSkinRef,
    pub localized_new_skin_name: String,
    pub localized_old_skin_name: String,
}

impl PlayerSkin {
    pub const ID: u32 = 93;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubClientLogin {
    pub sub_client_connection_request: String,
}

impl SubClientLogin {
    pub const ID: u32 = 94;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AutomationClientConnect {
    pub web_socket_data: WebSocketData,
}

impl AutomationClientConnect {
    pub const ID: u32 = 95;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetLastHurtBy {
    pub last_hurt_by: ActorType,
}

impl SetLastHurtBy {
    pub const ID: u32 = 96;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BookEdit {
    pub book_slot: wire::ZigZag32,
    pub operation: BookEditAction,
}

impl BookEdit {
    pub const ID: u32 = 97;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NpcRequest {
    pub npc_runtime_id: ActorRuntimeID,
    pub request_type: RequestType,
    pub actions: String,
    pub action_index: wire::U8,
    pub scene_name: String,
}

impl NpcRequest {
    pub const ID: u32 = 98;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PhotoTransfer {
    pub photo_name: String,
    pub photo_data: bytes::Bytes,
    pub book_id: String,
    pub type_: PhotoType,
    pub source_type: PhotoType,
    pub owner_id: wire::I64LE,
    pub new_photo_name: String,
}

impl PhotoTransfer {
    pub const ID: u32 = 99;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ModalFormRequest {
    pub form_id: wire::VarUInt,
    pub form_ui_json: String,
}

impl ModalFormRequest {
    pub const ID: u32 = 100;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ModalFormResponse {
    pub form_id: wire::VarUInt,
    /// Wire presence: optional value is preceded by a presence marker.
    pub json_response: Option<String>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub form_cancel_reason: Option<ModalFormCancelReason>,
}

impl ModalFormResponse {
    pub const ID: u32 = 101;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerSettingsRequest {
}

impl ServerSettingsRequest {
    pub const ID: u32 = 102;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerSettingsResponse {
    pub form_id: wire::VarUInt,
    pub form_ui_json: String,
}

impl ServerSettingsResponse {
    pub const ID: u32 = 103;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShowProfile {
    pub player_xuid: String,
}

impl ShowProfile {
    pub const ID: u32 = 104;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetDefaultGameType {
    pub default_game_type: GameType,
}

impl SetDefaultGameType {
    pub const ID: u32 = 105;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RemoveObjective {
    pub objective_name: String,
}

impl RemoveObjective {
    pub const ID: u32 = 106;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetDisplayObjective {
    pub display_slot_name: String,
    pub objective_name: String,
    pub objective_display_name: String,
    pub criteria_name: String,
    pub sort_order: wire::ZigZag32,
}

impl SetDisplayObjective {
    pub const ID: u32 = 107;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetScore {
    pub score_info: Vec<SetScoreInfoItem>,
}

impl SetScore {
    pub const ID: u32 = 108;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LabTable {
    pub type_: LabTableType,
    pub position: BlockPos,
    pub reaction: LabTableReactionType,
}

impl LabTable {
    pub const ID: u32 = 109;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateBlockSynced {
    pub block_position: BlockPos,
    pub block_runtime_id: wire::VarUInt,
    pub flags: wire::VarUInt,
    pub layer: wire::VarUInt,
    pub unique_actor_id: wire::VarULong,
    pub actor_sync_message: wire::VarULong,
}

impl UpdateBlockSynced {
    pub const ID: u32 = 110;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MoveActorDelta {
    pub move_data: MoveActorDeltaData,
}

impl MoveActorDelta {
    pub const ID: u32 = 111;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetScoreboardIdentity {
    pub scoreboard_identity_packet_type: ScoreboardIdentityPacketType,
    pub scoreboard_identity_info: Vec<ScoreboardIdentityPacketInfo>,
}

impl SetScoreboardIdentity {
    pub const ID: u32 = 112;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetLocalPlayerAsInitialized {
    pub player_id: ActorRuntimeID,
}

impl SetLocalPlayerAsInitialized {
    pub const ID: u32 = 113;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateSoftEnum {
    pub enum_name: String,
    pub values: Vec<String>,
    pub update_type: SoftEnumUpdateType,
}

impl UpdateSoftEnum {
    pub const ID: u32 = 114;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkStackLatency {
    pub creation_time: wire::U64LE,
    pub is_from_server: bool,
}

impl NetworkStackLatency {
    pub const ID: u32 = 115;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SpawnParticleEffect {
    pub dimension_id: wire::U8,
    pub actor_id: ActorUniqueID,
    pub position: glam::Vec3,
    pub effect_name: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub molang_variables: Option<String>,
}

impl SpawnParticleEffect {
    pub const ID: u32 = 118;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableActorIdentifiers {
    pub identifier_list: Nbt,
}

impl AvailableActorIdentifiers {
    pub const ID: u32 = 119;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkChunkPublisherUpdate {
    pub new_position_for_view: BlockPos,
    pub new_radius_for_view: wire::VarUInt,
    pub server_built_chunks_list: Vec<ChunkPos>,
}

impl NetworkChunkPublisherUpdate {
    pub const ID: u32 = 121;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeDefinitionList {
    pub map_of_biome_names_to_data: Vec<(wire::U16LE, BiomeDefinitionData)>,
    pub string_list: BiomeStringList,
}

impl BiomeDefinitionList {
    pub const ID: u32 = 122;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelSoundEvent {
    pub sound_event: String,
    pub position: glam::Vec3,
    pub data: wire::ZigZag32,
    pub actor_identifier: String,
    pub is_baby: bool,
    pub is_global: bool,
    pub actor_unique_id: wire::I64LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub fire_at_position: Option<glam::Vec3>,
}

impl LevelSoundEvent {
    pub const ID: u32 = 123;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelEventGeneric {
    pub event_id: wire::ZigZag32,
    pub ctd: Nbt,
}

impl LevelEventGeneric {
    pub const ID: u32 = 124;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LecternUpdate {
    pub new_page_to_show: wire::U8,
    pub total_pages: wire::U8,
    pub position_of_lectern_to_update: BlockPos,
}

impl LecternUpdate {
    pub const ID: u32 = 125;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCacheStatus {
    pub is_cache_supported: bool,
}

impl ClientCacheStatus {
    pub const ID: u32 = 129;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct OnScreenTextureAnimation {
    pub effect_id: wire::U32LE,
}

impl OnScreenTextureAnimation {
    pub const ID: u32 = 130;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapCreateLockedCopy {
    pub original_map_id: ActorUniqueID,
    pub new_map_id: ActorUniqueID,
}

impl MapCreateLockedCopy {
    pub const ID: u32 = 131;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureTemplateDataRequest {
    pub structure_name: String,
    pub structure_position: BlockPos,
    pub structure_settings: StructureSettings,
    pub requested_operation: StructureTemplateRequestOperation,
}

impl StructureTemplateDataRequest {
    pub const ID: u32 = 132;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureTemplateDataResponse {
    pub structure_name: String,
    pub structure_nbt: Nbt,
    pub response_type: StructureTemplateResponseType,
}

impl StructureTemplateDataResponse {
    pub const ID: u32 = 133;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCacheBlobStatus {
    pub missing_ids: Vec<wire::U64LE>,
    pub found_ids: Vec<wire::U64LE>,
}

impl ClientCacheBlobStatus {
    pub const ID: u32 = 135;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCacheMissResponse {
    pub missing_blobs: Vec<MissingBlobData>,
}

impl ClientCacheMissResponse {
    pub const ID: u32 = 136;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EducationSettings {
    pub education_level_settings: EducationLevelSettings,
}

impl EducationSettings {
    pub const ID: u32 = 137;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Emote {
    pub actor_runtime_id: ActorRuntimeID,
    pub emote_id: String,
    pub emote_length_ticks: wire::VarUInt,
    pub xuid: String,
    pub platform_id: String,
    pub flags: wire::U8,
}

impl Emote {
    pub const ID: u32 = 138;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MultiplayerSettings {
    pub packet_type: MultiplayerSettingsType,
}

impl MultiplayerSettings {
    pub const ID: u32 = 139;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SettingsCommand {
    pub command: String,
    pub suppress_output: bool,
}

impl SettingsCommand {
    pub const ID: u32 = 140;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AnvilDamage {
    pub block_position: BlockPos,
}

impl AnvilDamage {
    pub const ID: u32 = 141;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CompletedUsingItem {
    pub item_id: wire::I16LE,
    pub item_use_method: wire::I32LE,
}

impl CompletedUsingItem {
    pub const ID: u32 = 142;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkSettings {
    pub compression_threshold: wire::U16LE,
    pub compression_algorithm: PacketCompressionAlgorithm,
    pub client_throttle_enabled: bool,
    pub client_throttle_threshold: wire::U8,
    pub client_throttle_scalar: wire::F32LE,
}

impl NetworkSettings {
    pub const ID: u32 = 143;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerAuthInput {
    pub player_rotation: glam::Vec2,
    pub position: glam::Vec3,
    pub move_vector: glam::Vec2,
    pub player_head_rotation: wire::F32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub input_data: Option<Vec<InputData>>,
    pub input_mode: InputMode,
    pub play_mode: ClientPlayMode,
    pub new_interaction_model: NewInteractionModel,
    pub interact_rotation: glam::Vec2,
    pub client_tick: PlayerInputTick,
    pub pos_delta: glam::Vec3,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_use_transaction: Option<PackedItemUseLegacyInventoryTransaction>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_stack_request: Option<ItemStackRequestData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub player_block_actions: Option<Vec<PlayerBlockActionData>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub vehicle_rotation: Option<glam::Vec2>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_predicted_vehicle: Option<ActorUniqueID>,
    pub analog_move_vector: glam::Vec2,
    pub camera_orientation: glam::Vec3,
    pub raw_move_vector: glam::Vec2,
}

impl PlayerAuthInput {
    pub const ID: u32 = 144;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreativeContent {
    pub groups: Vec<CreativeGroupInfo>,
    pub entries: Vec<CreativeItemEntry>,
}

impl CreativeContent {
    pub const ID: u32 = 145;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerEnchantOptions {
    pub options: Vec<ItemEnchantOption>,
}

impl PlayerEnchantOptions {
    pub const ID: u32 = 146;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequest {
    pub requests: Vec<ItemStackRequestPacketData>,
}

impl ItemStackRequest {
    pub const ID: u32 = 147;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponse {
    pub responses: Vec<ItemStackResponseInfo>,
}

impl ItemStackResponse {
    pub const ID: u32 = 148;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerArmorDamage {
    pub armor_slot_and_damage_pairs: Vec<ArmorSlotAndDamagePair>,
}

impl PlayerArmorDamage {
    pub const ID: u32 = 149;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CodeBuilder {
    pub url: String,
    pub should_open_code_builder: bool,
}

impl CodeBuilder {
    pub const ID: u32 = 150;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdatePlayerGameType {
    pub player_game_type: GameType,
    pub target_player: ActorUniqueID,
    pub tick: PlayerInputTick,
}

impl UpdatePlayerGameType {
    pub const ID: u32 = 151;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EmoteList {
    pub runtime_id: ActorRuntimeID,
    pub emote_piece_ids: Vec<uuid::Uuid>,
}

impl EmoteList {
    pub const ID: u32 = 152;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PositionTrackingDBServerBroadcast {
    pub action: PositionTrackingDBServerBroadcastAction,
    pub id: PositionTrackingId,
    pub position_tracking_data: Nbt,
}

impl PositionTrackingDBServerBroadcast {
    pub const ID: u32 = 153;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PositionTrackingDBClientRequest {
    pub action: PositionTrackingDBClientRequestAction,
    pub id: PositionTrackingId,
}

impl PositionTrackingDBClientRequest {
    pub const ID: u32 = 154;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DebugInfo {
    pub actor_id: ActorUniqueID,
    pub data: String,
}

impl DebugInfo {
    pub const ID: u32 = 155;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PacketViolationWarning {
    pub violation_type: PacketViolationType,
    pub violation_severity: PacketViolationSeverity,
    pub violation_packet_id: wire::ZigZag32,
    pub violation_context: String,
}

impl PacketViolationWarning {
    pub const ID: u32 = 156;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MotionPredictionHints {
    pub m_runtime_id: ActorRuntimeID,
    pub m_motion: glam::Vec3,
    pub m_on_ground: bool,
}

impl MotionPredictionHints {
    pub const ID: u32 = 157;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AnimateEntity {
    pub m_animation: String,
    pub m_next_state: String,
    pub m_stop_expression: String,
    pub m_stop_expression_version: wire::I32LE,
    pub m_controller: String,
    pub m_blend_out_time: wire::F32LE,
    pub m_runtime_ids: Vec<ActorRuntimeID>,
}

impl AnimateEntity {
    pub const ID: u32 = 158;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraShake {
    pub intensity: wire::F32LE,
    pub seconds: wire::F32LE,
    pub shake_type: CameraShakeType,
    pub shake_action: CameraShakeAction,
}

impl CameraShake {
    pub const ID: u32 = 159;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerFog {
    pub fog_stack: Vec<String>,
}

impl PlayerFog {
    pub const ID: u32 = 160;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CorrectPlayerMovePrediction {
    pub prediction_type: RewindType,
    pub pos: glam::Vec3,
    pub pos_delta: glam::Vec3,
    pub rotation: glam::Vec2,
    /// Wire presence: optional value is preceded by a presence marker.
    pub vehicle_angular_velocity: Option<wire::F32LE>,
    pub on_ground: bool,
    pub tick: PlayerInputTick,
}

impl CorrectPlayerMovePrediction {
    pub const ID: u32 = 161;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemRegistry {
    pub item_data: Vec<ItemData>,
}

impl ItemRegistry {
    pub const ID: u32 = 162;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundDebugRenderer {
    pub type_: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub debug_marker_data: Option<DebugMarkerData>,
}

impl ClientboundDebugRenderer {
    pub const ID: u32 = 164;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncActorProperty {
    pub property_data: Nbt,
}

impl SyncActorProperty {
    pub const ID: u32 = 165;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddVolumeEntity {
    pub entity_network_id: EntityNetId,
    pub components: Nbt,
    pub json_identifier: String,
    pub instance_name: String,
    pub min_bounds: BlockPos,
    pub max_bounds: BlockPos,
    pub dimension_type: DimensionType,
    pub engine_version: String,
}

impl AddVolumeEntity {
    pub const ID: u32 = 166;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RemoveVolumeEntity {
    pub entity_network_id: EntityNetId,
    pub dimension_type: DimensionType,
}

impl RemoveVolumeEntity {
    pub const ID: u32 = 167;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SimulationType {
    pub sim_type: SimulationTypeEnum,
}

impl SimulationType {
    pub const ID: u32 = 168;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NpcDialogue {
    pub npc_id_raw_id: wire::U64LE,
    pub npc_dialogue_action_type: NpcDialogueActionType,
    pub dialogue: String,
    pub scene_name: String,
    pub npc_name: String,
    pub action_json: String,
}

impl NpcDialogue {
    pub const ID: u32 = 169;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EduUriResource {
    pub edu_shared_uri_resource: EduSharedUriResource,
}

impl EduUriResource {
    pub const ID: u32 = 170;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreatePhoto {
    pub raw_id: wire::U64LE,
    pub photo_name: String,
    pub photo_item_name: String,
}

impl CreatePhoto {
    pub const ID: u32 = 171;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateSubChunkBlocks {
    pub sub_chunk_block_position: BlockPos,
    pub blocks_changed: UpdateSubChunkBlocksChangedInfo,
}

impl UpdateSubChunkBlocks {
    pub const ID: u32 = 172;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunk {
    pub cache_enabled: bool,
    pub dimension_type: DimensionType,
    pub center_pos: SubChunkPos,
    pub sub_chunk_data: Vec<SubChunkData>,
}

impl SubChunk {
    pub const ID: u32 = 174;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunkRequest {
    pub dimension_type: DimensionType,
    pub sub_chunk_position_offset_list: Vec<SubChunkPosOffset>,
    pub center_pos: SubChunkPos,
}

impl SubChunkRequest {
    pub const ID: u32 = 175;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerStartItemCooldown {
    pub item_category: String,
    pub duration_ticks: wire::ZigZag32,
}

impl PlayerStartItemCooldown {
    pub const ID: u32 = 176;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ScriptMessage {
    pub message_id: String,
    pub message_value: String,
}

impl ScriptMessage {
    pub const ID: u32 = 177;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CodeBuilderSource {
    pub operation: CodeBuilderStorageQueryOptionsOperation,
    pub category: CodeBuilderStorageQueryOptionsCategory,
    pub code_status: CodeBuilderExecutionStateCodeStatus,
}

impl CodeBuilderSource {
    pub const ID: u32 = 178;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TickingAreasLoadStatus {
    pub waiting_for_preload: bool,
}

impl TickingAreasLoadStatus {
    pub const ID: u32 = 179;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DimensionData {
    pub definitions: Vec<(String, DimensionDefinition)>,
}

impl DimensionData {
    pub const ID: u32 = 180;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AgentActionEvent {
    pub request_id: String,
    pub action: AgentActionType,
    pub response: String,
}

impl AgentActionEvent {
    pub const ID: u32 = 181;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChangeMobProperty {
    pub actor_id: ActorUniqueID,
    pub property_name: String,
    pub bool_component_value: bool,
    pub string_component_value: String,
    pub int_component_value: wire::ZigZag32,
    pub float_component_value: wire::F32LE,
}

impl ChangeMobProperty {
    pub const ID: u32 = 182;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LessonProgress {
    pub lesson_action: wire::ZigZag32,
    pub score: wire::ZigZag32,
    pub activity_id: String,
}

impl LessonProgress {
    pub const ID: u32 = 183;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestAbility {
    pub ability: wire::ZigZag32,
    pub value_type: RequestAbilityType,
    pub bool: bool,
    pub float: wire::F32LE,
}

impl RequestAbility {
    pub const ID: u32 = 184;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestPermissions {
    pub target_player_id_raw_id: wire::I64LE,
    pub player_permission_level: wire::ZigZag32,
    pub custom_permission_flags: wire::U16LE,
}

impl RequestPermissions {
    pub const ID: u32 = 185;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ToastRequest {
    pub title: String,
    pub content: String,
}

impl ToastRequest {
    pub const ID: u32 = 186;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateAbilities {
    pub data: SerializedAbilitiesData,
}

impl UpdateAbilities {
    pub const ID: u32 = 187;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateAdventureSettings {
    pub adventure_settings: AdventureSettings,
}

impl UpdateAdventureSettings {
    pub const ID: u32 = 188;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DeathInfo {
    pub death_cause_attack_name: String,
    pub death_cause_message_list: Vec<String>,
}

impl DeathInfo {
    pub const ID: u32 = 189;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EditorNetwork {
    pub route_to_manager: bool,
    pub payload: Nbt,
}

impl EditorNetwork {
    pub const ID: u32 = 190;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct FeatureRegistry {
    pub features_data_list: Vec<FeatureRegistryFeatureBinaryJsonFormat>,
}

impl FeatureRegistry {
    pub const ID: u32 = 191;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerStats {
    pub server_time: wire::F32LE,
    pub network_time: wire::F32LE,
}

impl ServerStats {
    pub const ID: u32 = 192;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestNetworkSettings {
    pub client_network_version: wire::I32BE,
}

impl RequestNetworkSettings {
    pub const ID: u32 = 193;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameTestRequest {
    pub max_tests_per_batch: wire::ZigZag32,
    pub repeat_count: wire::ZigZag32,
    pub rotation: Rotation,
    pub stop_on_failure: bool,
    pub test_pos: BlockPos,
    pub tests_per_row: wire::ZigZag32,
    pub test_name: String,
}

impl GameTestRequest {
    pub const ID: u32 = 194;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameTestResults {
    pub succeeded: bool,
    pub error: String,
    pub test_name: String,
}

impl GameTestResults {
    pub const ID: u32 = 195;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateClientInputLocks {
    pub input_lock_component_data: wire::VarUInt,
}

impl UpdateClientInputLocks {
    pub const ID: u32 = 196;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPresets {
    pub camera_presets: CameraPresetList,
}

impl CameraPresets {
    pub const ID: u32 = 198;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UnlockedRecipes {
    pub packet_type: PacketType,
    pub unlocked_recipes_list: Vec<String>,
}

impl UnlockedRecipes {
    pub const ID: u32 = 199;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstruction {
    pub camera_instruction: CameraInstructionData,
}

impl CameraInstruction {
    pub const ID: u32 = 300;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TrimData {
    pub trim_pattern_list: Vec<TrimPattern>,
    pub trim_material_list: Vec<TrimMaterial>,
}

impl TrimData {
    pub const ID: u32 = 302;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct OpenSign {
    pub pos: BlockPos,
    pub is_front_side: bool,
}

impl OpenSign {
    pub const ID: u32 = 303;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AgentAnimation {
    pub agent_animation: AgentAnimationType,
    pub runtime_id: ActorRuntimeID,
}

impl AgentAnimation {
    pub const ID: u32 = 304;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RefreshEntitlements {
}

impl RefreshEntitlements {
    pub const ID: u32 = 305;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerToggleCrafterSlotRequest {
    pub pos_x: wire::I32LE,
    pub pos_y: wire::I32LE,
    pub pos_z: wire::I32LE,
    pub slot_index: wire::U8,
    pub is_disabled: bool,
}

impl PlayerToggleCrafterSlotRequest {
    pub const ID: u32 = 306;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetPlayerInventoryOptions {
    pub inventory_options: InventoryOptions,
}

impl SetPlayerInventoryOptions {
    pub const ID: u32 = 307;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetHud {
    pub hud_element: Vec<HudElement>,
    pub hud_visible: HudVisibility,
}

impl SetHud {
    pub const ID: u32 = 308;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AwardAchievement {
    pub achievement_id: wire::I32LE,
}

impl AwardAchievement {
    pub const ID: u32 = 309;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundCloseForm {
}

impl ClientboundCloseForm {
    pub const ID: u32 = 310;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerboundLoadingScreen {
    pub loading_screen_packet_type: ServerboundLoadingScreenType,
    /// Wire presence: optional value is preceded by a presence marker.
    pub loading_screen_id: Option<wire::U32LE>,
}

impl ServerboundLoadingScreen {
    pub const ID: u32 = 312;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct JigsawStructureData {
    pub jigsaw_structure_data_tag: Nbt,
}

impl JigsawStructureData {
    pub const ID: u32 = 313;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CurrentStructureFeature {
    pub current_structure_feature: String,
}

impl CurrentStructureFeature {
    pub const ID: u32 = 314;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerboundDiagnostics {
    pub avg_fps: wire::F32LE,
    pub avg_server_sim_tick_time_ms: wire::F32LE,
    pub avg_client_sim_tick_time_ms: wire::F32LE,
    pub avg_begin_frame_time_ms: wire::F32LE,
    pub avg_input_time_ms: wire::F32LE,
    pub avg_render_time_ms: wire::F32LE,
    pub avg_end_frame_time_ms: wire::F32LE,
    pub avg_remainder_time_percent: wire::F32LE,
    pub avg_unaccounted_time_percent: wire::F32LE,
    pub memory_category_values: Vec<MemoryCategoryCounter>,
    pub entity_diagnostics: Vec<ECSProfilingDiagnosticsEntityDiagnosticTimingInfo>,
    pub system_diagnostics: Vec<ECSProfilingDiagnosticsSystemDiagnosticTimingInfo>,
    pub system_categories: Vec<ECSProfilingDiagnosticsSystemCategory>,
    pub whisker_scopes: Vec<BedrockProfileWhiskerDiagnosticsScopeDataSummary>,
}

impl ServerboundDiagnostics {
    pub const ID: u32 = 315;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssist {
    pub preset_id: String,
    pub view_angle: glam::Vec2,
    pub distance: wire::F32LE,
    pub target_mode: TargetMode,
    pub action: CameraAimAssistAction,
    pub show_debug_render: bool,
}

impl CameraAimAssist {
    pub const ID: u32 = 316;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerRegistryCleanup {
    pub removed_containers: Vec<FullContainerName>,
}

impl ContainerRegistryCleanup {
    pub const ID: u32 = 317;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MovementEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub effect_id: MovementEffectType,
    pub effect_duration: wire::ZigZag32,
    pub tick: PlayerInputTick,
}

impl MovementEffect {
    pub const ID: u32 = 318;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistPresets {
    pub camera_aim_assist_presets: Vec<CameraAimAssistCategoryDefinition>,
    pub camera_aim_assist_categories: Vec<CameraAimAssistPresetDefinition>,
    pub operation: CameraAimAssistPresetOperation,
}

impl CameraAimAssistPresets {
    pub const ID: u32 = 320;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCameraAimAssist {
    pub camera_preset_id: String,
    pub action: ClientCameraAimAssistAction,
    pub allow_aim_assist: bool,
}

impl ClientCameraAimAssist {
    pub const ID: u32 = 321;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientMovementPredictionSync {
    pub actor_data_flag: ActorDataFlagComponent,
    pub actor_bounding_box: ActorDataBoundingBoxComponent,
    pub movement_attributes: [wire::F32LE; 9],
    pub actor_unique_id: ActorUniqueID,
    pub actor_flying_state: bool,
}

impl ClientMovementPredictionSync {
    pub const ID: u32 = 322;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateClientOptions {
    /// Wire presence: optional value is preceded by a presence marker.
    pub graphics_mode_change: Option<GraphicsMode>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub filter_profanity_change: Option<bool>,
}

impl UpdateClientOptions {
    pub const ID: u32 = 323;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerVideoCapture {
    pub action: PlayerVideoCaptureData,
}

impl PlayerVideoCapture {
    pub const ID: u32 = 324;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerUpdateEntityOverrides {
    pub target_id: ActorUniqueID,
    pub property_index: wire::VarUInt,
    pub update: PlayerUpdateEntityOverridesData,
}

impl PlayerUpdateEntityOverrides {
    pub const ID: u32 = 325;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerLocation {
    pub target_actor_id: ActorUniqueID,
    pub location: PlayerLocationData,
}

impl PlayerLocation {
    pub const ID: u32 = 326;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundControlSchemeSet {
    pub control_scheme: ControlScheme,
}

impl ClientboundControlSchemeSet {
    pub const ID: u32 = 327;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PrimitiveShapes {
    pub array_of_primitive_shapes_can_be_a_mix_of_new_updated_or_removed: Vec<PrimitiveShape>,
}

impl PrimitiveShapes {
    pub const ID: u32 = 328;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerboundPackSettingChange {
    pub pack_id: uuid::Uuid,
    pub pack_setting_name: String,
    pub pack_setting_value: ServerboundPackSettingChangePackSettingValue,
}

impl ServerboundPackSettingChange {
    pub const ID: u32 = 329;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundDataStore {
    pub updates: Vec<BedrockDDUI>,
}

impl ClientboundDataStore {
    pub const ID: u32 = 330;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GraphicsOverrideParameter {
    pub parameter_keyframe_values: Vec<(wire::F32LE, glam::Vec3)>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub float_value: Option<wire::F32LE>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub vec3_value: Option<glam::Vec3>,
    pub biome_identifier: String,
    /// Wire presence: optional value is preceded by a presence marker.
    pub player_identifier: Option<String>,
    pub identifier_for_parameter: GraphicsOverrideParameterType,
    pub reset_parameter: bool,
}

impl GraphicsOverrideParameter {
    pub const ID: u32 = 331;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerboundDataStore {
    pub update: BedrockDDUIDataStoreUpdate,
}

impl ServerboundDataStore {
    pub const ID: u32 = 332;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundDataDrivenUIShowScreen {
    pub screen_id: String,
    pub form_id: wire::U32LE,
    /// Wire presence: optional value is preceded by a presence marker.
    pub data_instance_id: Option<wire::U32LE>,
}

impl ClientboundDataDrivenUIShowScreen {
    pub const ID: u32 = 333;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundDataDrivenUICloseScreen {
    /// Wire presence: optional value is preceded by a presence marker.
    pub form_id: Option<wire::U32LE>,
}

impl ClientboundDataDrivenUICloseScreen {
    pub const ID: u32 = 334;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundDataDrivenUIReload {
}

impl ClientboundDataDrivenUIReload {
    pub const ID: u32 = 335;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundTextureShift {
    pub action_id: ClientboundTextureShiftAction,
    pub collection_name: String,
    pub from_step: String,
    pub to_step: String,
    pub all_steps: Vec<String>,
    pub current_length_in_ticks: wire::VarULong,
    pub total_length_in_ticks: wire::VarULong,
    pub enabled: bool,
}

impl ClientboundTextureShift {
    pub const ID: u32 = 336;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct VoxelShapes {
    pub shapes: Vec<VoxelShapesSerializableVoxelShape>,
    pub name_map: Vec<(String, VoxelShapesRegistryHandle)>,
    pub custom_shape_count: wire::U16LE,
}

impl VoxelShapes {
    pub const ID: u32 = 337;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSpline {
    pub camera_data_splines: Vec<CameraSplineDefinition>,
}

impl CameraSpline {
    pub const ID: u32 = 338;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistActorPriority {
    pub camera_aim_assist_actor_priority_list: Vec<CameraAimAssistActorPriorityData>,
}

impl CameraAimAssistActorPriority {
    pub const ID: u32 = 339;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePacksReadyForValidation {
}

impl ResourcePacksReadyForValidation {
    pub const ID: u32 = 340;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LocatorBar {
    pub waypoints: Vec<LocatorBarWaypoint>,
}

impl LocatorBar {
    pub const ID: u32 = 341;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PartyChanged {
    /// Wire presence: optional value is preceded by a presence marker.
    pub party_info: Option<PlayerPartyInfo>,
}

impl PartyChanged {
    pub const ID: u32 = 342;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerboundDataDrivenScreenClosed {
    pub form_id: wire::U32LE,
    pub close_reason: String,
}

impl ServerboundDataDrivenScreenClosed {
    pub const ID: u32 = 343;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncWorldClocks {
    pub data: SyncWorldClocksData,
}

impl SyncWorldClocks {
    pub const ID: u32 = 344;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundAttributeLayerSync {
    pub data: AttributeLayerSyncData,
}

impl ClientboundAttributeLayerSync {
    pub const ID: u32 = 345;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerStoreInfo {
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_store_entry_point_configuration: Option<ServerConfigurationClientStoreEntryPointConfiguration>,
}

impl ServerStoreInfo {
    pub const ID: u32 = 346;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerPresenceInfo {
    /// Wire presence: optional value is preceded by a presence marker.
    pub presence_configuration: Option<ServerConfigurationPresenceConfiguration>,
}

impl ServerPresenceInfo {
    pub const ID: u32 = 347;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundUpdateSoundData {
    pub server_sound_handle: ServerSoundHandle,
    /// Wire presence: optional value is preceded by a presence marker.
    pub stop: Option<SoundDataEvent>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub set_volume: Option<SoundDataEvent>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub set_pitch: Option<SoundDataEvent>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub fade: Option<SoundDataEvent>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub seek_to: Option<SoundDataEvent>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub pause: Option<SoundDataEvent>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub resume: Option<SoundDataEvent>,
}

impl ClientboundUpdateSoundData {
    pub const ID: u32 = 348;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SendPartyDestinationCookie {
    pub cookie: String,
    pub intent: String,
    pub destination_name: String,
}

impl SendPartyDestinationCookie {
    pub const ID: u32 = 349;
}
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PartyDestinationCookieResponse {
    pub cookie: String,
    pub accepted: bool,
}

impl PartyDestinationCookieResponse {
    pub const ID: u32 = 350;
}

#[derive(Clone, Copy, Debug, Eq, Hash, PartialEq)]
#[repr(u32)]
pub enum PacketId {
    Login = 1,
    PlayStatus = 2,
    ServerToClientHandshake = 3,
    ClientToServerHandshake = 4,
    Disconnect = 5,
    ResourcePacksInfo = 6,
    ResourcePackStack = 7,
    ResourcePackClientResponse = 8,
    Text = 9,
    SetTime = 10,
    StartGame = 11,
    AddPlayer = 12,
    AddActor = 13,
    RemoveActor = 14,
    AddItemActor = 15,
    ServerPlayerPostMovePosition = 16,
    TakeItemActor = 17,
    MoveActorAbsolute = 18,
    MovePlayer = 19,
    UpdateBlock = 21,
    AddPainting = 22,
    LevelEvent = 25,
    BlockEvent = 26,
    ActorEvent = 27,
    MobEffect = 28,
    UpdateAttributes = 29,
    InventoryTransaction = 30,
    MobEquipment = 31,
    MobArmorEquipment = 32,
    Interact = 33,
    BlockPickRequest = 34,
    ActorPickRequest = 35,
    PlayerAction = 36,
    HurtArmor = 38,
    SetActorData = 39,
    SetActorMotion = 40,
    SetActorLink = 41,
    SetHealth = 42,
    SetSpawnPosition = 43,
    Animate = 44,
    Respawn = 45,
    ContainerOpen = 46,
    ContainerClose = 47,
    PlayerHotbar = 48,
    InventoryContent = 49,
    InventorySlot = 50,
    ContainerSetData = 51,
    CraftingData = 52,
    GuiDataPickItem = 54,
    BlockActorData = 56,
    LevelChunk = 58,
    SetCommandsEnabled = 59,
    SetDifficulty = 60,
    ChangeDimension = 61,
    SetPlayerGameType = 62,
    PlayerList = 63,
    SimpleEvent = 64,
    LegacyTelemetryEvent = 65,
    SpawnExperienceOrb = 66,
    ClientboundMapItemData = 67,
    MapInfoRequest = 68,
    RequestChunkRadius = 69,
    ChunkRadiusUpdated = 70,
    GameRulesChanged = 72,
    Camera = 73,
    BossEvent = 74,
    ShowCredits = 75,
    AvailableCommands = 76,
    CommandRequest = 77,
    CommandBlockUpdate = 78,
    CommandOutput = 79,
    UpdateTrade = 80,
    UpdateEquip = 81,
    ResourcePackDataInfo = 82,
    ResourcePackChunkData = 83,
    ResourcePackChunkRequest = 84,
    Transfer = 85,
    PlaySound = 86,
    StopSound = 87,
    SetTitle = 88,
    AddBehaviorTree = 89,
    StructureBlockUpdate = 90,
    ShowStoreOffer = 91,
    PurchaseReceipt = 92,
    PlayerSkin = 93,
    SubClientLogin = 94,
    AutomationClientConnect = 95,
    SetLastHurtBy = 96,
    BookEdit = 97,
    NpcRequest = 98,
    PhotoTransfer = 99,
    ModalFormRequest = 100,
    ModalFormResponse = 101,
    ServerSettingsRequest = 102,
    ServerSettingsResponse = 103,
    ShowProfile = 104,
    SetDefaultGameType = 105,
    RemoveObjective = 106,
    SetDisplayObjective = 107,
    SetScore = 108,
    LabTable = 109,
    UpdateBlockSynced = 110,
    MoveActorDelta = 111,
    SetScoreboardIdentity = 112,
    SetLocalPlayerAsInitialized = 113,
    UpdateSoftEnum = 114,
    NetworkStackLatency = 115,
    SpawnParticleEffect = 118,
    AvailableActorIdentifiers = 119,
    NetworkChunkPublisherUpdate = 121,
    BiomeDefinitionList = 122,
    LevelSoundEvent = 123,
    LevelEventGeneric = 124,
    LecternUpdate = 125,
    ClientCacheStatus = 129,
    OnScreenTextureAnimation = 130,
    MapCreateLockedCopy = 131,
    StructureTemplateDataRequest = 132,
    StructureTemplateDataResponse = 133,
    ClientCacheBlobStatus = 135,
    ClientCacheMissResponse = 136,
    EducationSettings = 137,
    Emote = 138,
    MultiplayerSettings = 139,
    SettingsCommand = 140,
    AnvilDamage = 141,
    CompletedUsingItem = 142,
    NetworkSettings = 143,
    PlayerAuthInput = 144,
    CreativeContent = 145,
    PlayerEnchantOptions = 146,
    ItemStackRequest = 147,
    ItemStackResponse = 148,
    PlayerArmorDamage = 149,
    CodeBuilder = 150,
    UpdatePlayerGameType = 151,
    EmoteList = 152,
    PositionTrackingDBServerBroadcast = 153,
    PositionTrackingDBClientRequest = 154,
    DebugInfo = 155,
    PacketViolationWarning = 156,
    MotionPredictionHints = 157,
    AnimateEntity = 158,
    CameraShake = 159,
    PlayerFog = 160,
    CorrectPlayerMovePrediction = 161,
    ItemRegistry = 162,
    ClientboundDebugRenderer = 164,
    SyncActorProperty = 165,
    AddVolumeEntity = 166,
    RemoveVolumeEntity = 167,
    SimulationType = 168,
    NpcDialogue = 169,
    EduUriResource = 170,
    CreatePhoto = 171,
    UpdateSubChunkBlocks = 172,
    SubChunk = 174,
    SubChunkRequest = 175,
    PlayerStartItemCooldown = 176,
    ScriptMessage = 177,
    CodeBuilderSource = 178,
    TickingAreasLoadStatus = 179,
    DimensionData = 180,
    AgentActionEvent = 181,
    ChangeMobProperty = 182,
    LessonProgress = 183,
    RequestAbility = 184,
    RequestPermissions = 185,
    ToastRequest = 186,
    UpdateAbilities = 187,
    UpdateAdventureSettings = 188,
    DeathInfo = 189,
    EditorNetwork = 190,
    FeatureRegistry = 191,
    ServerStats = 192,
    RequestNetworkSettings = 193,
    GameTestRequest = 194,
    GameTestResults = 195,
    UpdateClientInputLocks = 196,
    CameraPresets = 198,
    UnlockedRecipes = 199,
    CameraInstruction = 300,
    TrimData = 302,
    OpenSign = 303,
    AgentAnimation = 304,
    RefreshEntitlements = 305,
    PlayerToggleCrafterSlotRequest = 306,
    SetPlayerInventoryOptions = 307,
    SetHud = 308,
    AwardAchievement = 309,
    ClientboundCloseForm = 310,
    ServerboundLoadingScreen = 312,
    JigsawStructureData = 313,
    CurrentStructureFeature = 314,
    ServerboundDiagnostics = 315,
    CameraAimAssist = 316,
    ContainerRegistryCleanup = 317,
    MovementEffect = 318,
    CameraAimAssistPresets = 320,
    ClientCameraAimAssist = 321,
    ClientMovementPredictionSync = 322,
    UpdateClientOptions = 323,
    PlayerVideoCapture = 324,
    PlayerUpdateEntityOverrides = 325,
    PlayerLocation = 326,
    ClientboundControlSchemeSet = 327,
    PrimitiveShapes = 328,
    ServerboundPackSettingChange = 329,
    ClientboundDataStore = 330,
    GraphicsOverrideParameter = 331,
    ServerboundDataStore = 332,
    ClientboundDataDrivenUIShowScreen = 333,
    ClientboundDataDrivenUICloseScreen = 334,
    ClientboundDataDrivenUIReload = 335,
    ClientboundTextureShift = 336,
    VoxelShapes = 337,
    CameraSpline = 338,
    CameraAimAssistActorPriority = 339,
    ResourcePacksReadyForValidation = 340,
    LocatorBar = 341,
    PartyChanged = 342,
    ServerboundDataDrivenScreenClosed = 343,
    SyncWorldClocks = 344,
    ClientboundAttributeLayerSync = 345,
    ServerStoreInfo = 346,
    ServerPresenceInfo = 347,
    ClientboundUpdateSoundData = 348,
    SendPartyDestinationCookie = 349,
    PartyDestinationCookieResponse = 350,
}

impl PacketId {
    pub fn from_raw(raw: u32) -> Option<Self> {
        match raw {
            1 => Some(Self::Login),
            2 => Some(Self::PlayStatus),
            3 => Some(Self::ServerToClientHandshake),
            4 => Some(Self::ClientToServerHandshake),
            5 => Some(Self::Disconnect),
            6 => Some(Self::ResourcePacksInfo),
            7 => Some(Self::ResourcePackStack),
            8 => Some(Self::ResourcePackClientResponse),
            9 => Some(Self::Text),
            10 => Some(Self::SetTime),
            11 => Some(Self::StartGame),
            12 => Some(Self::AddPlayer),
            13 => Some(Self::AddActor),
            14 => Some(Self::RemoveActor),
            15 => Some(Self::AddItemActor),
            16 => Some(Self::ServerPlayerPostMovePosition),
            17 => Some(Self::TakeItemActor),
            18 => Some(Self::MoveActorAbsolute),
            19 => Some(Self::MovePlayer),
            21 => Some(Self::UpdateBlock),
            22 => Some(Self::AddPainting),
            25 => Some(Self::LevelEvent),
            26 => Some(Self::BlockEvent),
            27 => Some(Self::ActorEvent),
            28 => Some(Self::MobEffect),
            29 => Some(Self::UpdateAttributes),
            30 => Some(Self::InventoryTransaction),
            31 => Some(Self::MobEquipment),
            32 => Some(Self::MobArmorEquipment),
            33 => Some(Self::Interact),
            34 => Some(Self::BlockPickRequest),
            35 => Some(Self::ActorPickRequest),
            36 => Some(Self::PlayerAction),
            38 => Some(Self::HurtArmor),
            39 => Some(Self::SetActorData),
            40 => Some(Self::SetActorMotion),
            41 => Some(Self::SetActorLink),
            42 => Some(Self::SetHealth),
            43 => Some(Self::SetSpawnPosition),
            44 => Some(Self::Animate),
            45 => Some(Self::Respawn),
            46 => Some(Self::ContainerOpen),
            47 => Some(Self::ContainerClose),
            48 => Some(Self::PlayerHotbar),
            49 => Some(Self::InventoryContent),
            50 => Some(Self::InventorySlot),
            51 => Some(Self::ContainerSetData),
            52 => Some(Self::CraftingData),
            54 => Some(Self::GuiDataPickItem),
            56 => Some(Self::BlockActorData),
            58 => Some(Self::LevelChunk),
            59 => Some(Self::SetCommandsEnabled),
            60 => Some(Self::SetDifficulty),
            61 => Some(Self::ChangeDimension),
            62 => Some(Self::SetPlayerGameType),
            63 => Some(Self::PlayerList),
            64 => Some(Self::SimpleEvent),
            65 => Some(Self::LegacyTelemetryEvent),
            66 => Some(Self::SpawnExperienceOrb),
            67 => Some(Self::ClientboundMapItemData),
            68 => Some(Self::MapInfoRequest),
            69 => Some(Self::RequestChunkRadius),
            70 => Some(Self::ChunkRadiusUpdated),
            72 => Some(Self::GameRulesChanged),
            73 => Some(Self::Camera),
            74 => Some(Self::BossEvent),
            75 => Some(Self::ShowCredits),
            76 => Some(Self::AvailableCommands),
            77 => Some(Self::CommandRequest),
            78 => Some(Self::CommandBlockUpdate),
            79 => Some(Self::CommandOutput),
            80 => Some(Self::UpdateTrade),
            81 => Some(Self::UpdateEquip),
            82 => Some(Self::ResourcePackDataInfo),
            83 => Some(Self::ResourcePackChunkData),
            84 => Some(Self::ResourcePackChunkRequest),
            85 => Some(Self::Transfer),
            86 => Some(Self::PlaySound),
            87 => Some(Self::StopSound),
            88 => Some(Self::SetTitle),
            89 => Some(Self::AddBehaviorTree),
            90 => Some(Self::StructureBlockUpdate),
            91 => Some(Self::ShowStoreOffer),
            92 => Some(Self::PurchaseReceipt),
            93 => Some(Self::PlayerSkin),
            94 => Some(Self::SubClientLogin),
            95 => Some(Self::AutomationClientConnect),
            96 => Some(Self::SetLastHurtBy),
            97 => Some(Self::BookEdit),
            98 => Some(Self::NpcRequest),
            99 => Some(Self::PhotoTransfer),
            100 => Some(Self::ModalFormRequest),
            101 => Some(Self::ModalFormResponse),
            102 => Some(Self::ServerSettingsRequest),
            103 => Some(Self::ServerSettingsResponse),
            104 => Some(Self::ShowProfile),
            105 => Some(Self::SetDefaultGameType),
            106 => Some(Self::RemoveObjective),
            107 => Some(Self::SetDisplayObjective),
            108 => Some(Self::SetScore),
            109 => Some(Self::LabTable),
            110 => Some(Self::UpdateBlockSynced),
            111 => Some(Self::MoveActorDelta),
            112 => Some(Self::SetScoreboardIdentity),
            113 => Some(Self::SetLocalPlayerAsInitialized),
            114 => Some(Self::UpdateSoftEnum),
            115 => Some(Self::NetworkStackLatency),
            118 => Some(Self::SpawnParticleEffect),
            119 => Some(Self::AvailableActorIdentifiers),
            121 => Some(Self::NetworkChunkPublisherUpdate),
            122 => Some(Self::BiomeDefinitionList),
            123 => Some(Self::LevelSoundEvent),
            124 => Some(Self::LevelEventGeneric),
            125 => Some(Self::LecternUpdate),
            129 => Some(Self::ClientCacheStatus),
            130 => Some(Self::OnScreenTextureAnimation),
            131 => Some(Self::MapCreateLockedCopy),
            132 => Some(Self::StructureTemplateDataRequest),
            133 => Some(Self::StructureTemplateDataResponse),
            135 => Some(Self::ClientCacheBlobStatus),
            136 => Some(Self::ClientCacheMissResponse),
            137 => Some(Self::EducationSettings),
            138 => Some(Self::Emote),
            139 => Some(Self::MultiplayerSettings),
            140 => Some(Self::SettingsCommand),
            141 => Some(Self::AnvilDamage),
            142 => Some(Self::CompletedUsingItem),
            143 => Some(Self::NetworkSettings),
            144 => Some(Self::PlayerAuthInput),
            145 => Some(Self::CreativeContent),
            146 => Some(Self::PlayerEnchantOptions),
            147 => Some(Self::ItemStackRequest),
            148 => Some(Self::ItemStackResponse),
            149 => Some(Self::PlayerArmorDamage),
            150 => Some(Self::CodeBuilder),
            151 => Some(Self::UpdatePlayerGameType),
            152 => Some(Self::EmoteList),
            153 => Some(Self::PositionTrackingDBServerBroadcast),
            154 => Some(Self::PositionTrackingDBClientRequest),
            155 => Some(Self::DebugInfo),
            156 => Some(Self::PacketViolationWarning),
            157 => Some(Self::MotionPredictionHints),
            158 => Some(Self::AnimateEntity),
            159 => Some(Self::CameraShake),
            160 => Some(Self::PlayerFog),
            161 => Some(Self::CorrectPlayerMovePrediction),
            162 => Some(Self::ItemRegistry),
            164 => Some(Self::ClientboundDebugRenderer),
            165 => Some(Self::SyncActorProperty),
            166 => Some(Self::AddVolumeEntity),
            167 => Some(Self::RemoveVolumeEntity),
            168 => Some(Self::SimulationType),
            169 => Some(Self::NpcDialogue),
            170 => Some(Self::EduUriResource),
            171 => Some(Self::CreatePhoto),
            172 => Some(Self::UpdateSubChunkBlocks),
            174 => Some(Self::SubChunk),
            175 => Some(Self::SubChunkRequest),
            176 => Some(Self::PlayerStartItemCooldown),
            177 => Some(Self::ScriptMessage),
            178 => Some(Self::CodeBuilderSource),
            179 => Some(Self::TickingAreasLoadStatus),
            180 => Some(Self::DimensionData),
            181 => Some(Self::AgentActionEvent),
            182 => Some(Self::ChangeMobProperty),
            183 => Some(Self::LessonProgress),
            184 => Some(Self::RequestAbility),
            185 => Some(Self::RequestPermissions),
            186 => Some(Self::ToastRequest),
            187 => Some(Self::UpdateAbilities),
            188 => Some(Self::UpdateAdventureSettings),
            189 => Some(Self::DeathInfo),
            190 => Some(Self::EditorNetwork),
            191 => Some(Self::FeatureRegistry),
            192 => Some(Self::ServerStats),
            193 => Some(Self::RequestNetworkSettings),
            194 => Some(Self::GameTestRequest),
            195 => Some(Self::GameTestResults),
            196 => Some(Self::UpdateClientInputLocks),
            198 => Some(Self::CameraPresets),
            199 => Some(Self::UnlockedRecipes),
            300 => Some(Self::CameraInstruction),
            302 => Some(Self::TrimData),
            303 => Some(Self::OpenSign),
            304 => Some(Self::AgentAnimation),
            305 => Some(Self::RefreshEntitlements),
            306 => Some(Self::PlayerToggleCrafterSlotRequest),
            307 => Some(Self::SetPlayerInventoryOptions),
            308 => Some(Self::SetHud),
            309 => Some(Self::AwardAchievement),
            310 => Some(Self::ClientboundCloseForm),
            312 => Some(Self::ServerboundLoadingScreen),
            313 => Some(Self::JigsawStructureData),
            314 => Some(Self::CurrentStructureFeature),
            315 => Some(Self::ServerboundDiagnostics),
            316 => Some(Self::CameraAimAssist),
            317 => Some(Self::ContainerRegistryCleanup),
            318 => Some(Self::MovementEffect),
            320 => Some(Self::CameraAimAssistPresets),
            321 => Some(Self::ClientCameraAimAssist),
            322 => Some(Self::ClientMovementPredictionSync),
            323 => Some(Self::UpdateClientOptions),
            324 => Some(Self::PlayerVideoCapture),
            325 => Some(Self::PlayerUpdateEntityOverrides),
            326 => Some(Self::PlayerLocation),
            327 => Some(Self::ClientboundControlSchemeSet),
            328 => Some(Self::PrimitiveShapes),
            329 => Some(Self::ServerboundPackSettingChange),
            330 => Some(Self::ClientboundDataStore),
            331 => Some(Self::GraphicsOverrideParameter),
            332 => Some(Self::ServerboundDataStore),
            333 => Some(Self::ClientboundDataDrivenUIShowScreen),
            334 => Some(Self::ClientboundDataDrivenUICloseScreen),
            335 => Some(Self::ClientboundDataDrivenUIReload),
            336 => Some(Self::ClientboundTextureShift),
            337 => Some(Self::VoxelShapes),
            338 => Some(Self::CameraSpline),
            339 => Some(Self::CameraAimAssistActorPriority),
            340 => Some(Self::ResourcePacksReadyForValidation),
            341 => Some(Self::LocatorBar),
            342 => Some(Self::PartyChanged),
            343 => Some(Self::ServerboundDataDrivenScreenClosed),
            344 => Some(Self::SyncWorldClocks),
            345 => Some(Self::ClientboundAttributeLayerSync),
            346 => Some(Self::ServerStoreInfo),
            347 => Some(Self::ServerPresenceInfo),
            348 => Some(Self::ClientboundUpdateSoundData),
            349 => Some(Self::SendPartyDestinationCookie),
            350 => Some(Self::PartyDestinationCookieResponse),
            _ => None,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum Packet {
    Login(Login),
    PlayStatus(PlayStatus),
    ServerToClientHandshake(ServerToClientHandshake),
    ClientToServerHandshake(ClientToServerHandshake),
    Disconnect(Disconnect),
    ResourcePacksInfo(ResourcePacksInfo),
    ResourcePackStack(ResourcePackStack),
    ResourcePackClientResponse(ResourcePackClientResponse),
    Text(Text),
    SetTime(SetTime),
    StartGame(Box<StartGame>),
    AddPlayer(Box<AddPlayer>),
    AddActor(Box<AddActor>),
    RemoveActor(RemoveActor),
    AddItemActor(AddItemActor),
    ServerPlayerPostMovePosition(ServerPlayerPostMovePosition),
    TakeItemActor(TakeItemActor),
    MoveActorAbsolute(MoveActorAbsolute),
    MovePlayer(Box<MovePlayer>),
    UpdateBlock(UpdateBlock),
    AddPainting(AddPainting),
    LevelEvent(LevelEvent),
    BlockEvent(BlockEvent),
    ActorEvent(ActorEvent),
    MobEffect(Box<MobEffect>),
    UpdateAttributes(UpdateAttributes),
    InventoryTransaction(InventoryTransaction),
    MobEquipment(MobEquipment),
    MobArmorEquipment(MobArmorEquipment),
    Interact(Interact),
    BlockPickRequest(BlockPickRequest),
    ActorPickRequest(ActorPickRequest),
    PlayerAction(PlayerAction),
    HurtArmor(HurtArmor),
    SetActorData(SetActorData),
    SetActorMotion(SetActorMotion),
    SetActorLink(SetActorLink),
    SetHealth(SetHealth),
    SetSpawnPosition(SetSpawnPosition),
    Animate(Animate),
    Respawn(Respawn),
    ContainerOpen(ContainerOpen),
    ContainerClose(ContainerClose),
    PlayerHotbar(PlayerHotbar),
    InventoryContent(InventoryContent),
    InventorySlot(InventorySlot),
    ContainerSetData(ContainerSetData),
    CraftingData(Box<CraftingData>),
    GuiDataPickItem(GuiDataPickItem),
    BlockActorData(BlockActorData),
    LevelChunk(LevelChunk),
    SetCommandsEnabled(SetCommandsEnabled),
    SetDifficulty(SetDifficulty),
    ChangeDimension(ChangeDimension),
    SetPlayerGameType(SetPlayerGameType),
    PlayerList(PlayerList),
    SimpleEvent(SimpleEvent),
    LegacyTelemetryEvent(LegacyTelemetryEvent),
    SpawnExperienceOrb(SpawnExperienceOrb),
    ClientboundMapItemData(Box<ClientboundMapItemData>),
    MapInfoRequest(MapInfoRequest),
    RequestChunkRadius(RequestChunkRadius),
    ChunkRadiusUpdated(ChunkRadiusUpdated),
    GameRulesChanged(GameRulesChanged),
    Camera(Camera),
    BossEvent(Box<BossEvent>),
    ShowCredits(ShowCredits),
    AvailableCommands(Box<AvailableCommands>),
    CommandRequest(CommandRequest),
    CommandBlockUpdate(Box<CommandBlockUpdate>),
    CommandOutput(CommandOutput),
    UpdateTrade(Box<UpdateTrade>),
    UpdateEquip(UpdateEquip),
    ResourcePackDataInfo(ResourcePackDataInfo),
    ResourcePackChunkData(ResourcePackChunkData),
    ResourcePackChunkRequest(ResourcePackChunkRequest),
    Transfer(Transfer),
    PlaySound(PlaySound),
    StopSound(StopSound),
    SetTitle(Box<SetTitle>),
    AddBehaviorTree(AddBehaviorTree),
    StructureBlockUpdate(StructureBlockUpdate),
    ShowStoreOffer(ShowStoreOffer),
    PurchaseReceipt(PurchaseReceipt),
    PlayerSkin(PlayerSkin),
    SubClientLogin(SubClientLogin),
    AutomationClientConnect(AutomationClientConnect),
    SetLastHurtBy(SetLastHurtBy),
    BookEdit(BookEdit),
    NpcRequest(NpcRequest),
    PhotoTransfer(PhotoTransfer),
    ModalFormRequest(ModalFormRequest),
    ModalFormResponse(ModalFormResponse),
    ServerSettingsRequest(ServerSettingsRequest),
    ServerSettingsResponse(ServerSettingsResponse),
    ShowProfile(ShowProfile),
    SetDefaultGameType(SetDefaultGameType),
    RemoveObjective(RemoveObjective),
    SetDisplayObjective(SetDisplayObjective),
    SetScore(SetScore),
    LabTable(LabTable),
    UpdateBlockSynced(UpdateBlockSynced),
    MoveActorDelta(MoveActorDelta),
    SetScoreboardIdentity(SetScoreboardIdentity),
    SetLocalPlayerAsInitialized(SetLocalPlayerAsInitialized),
    UpdateSoftEnum(UpdateSoftEnum),
    NetworkStackLatency(NetworkStackLatency),
    SpawnParticleEffect(SpawnParticleEffect),
    AvailableActorIdentifiers(AvailableActorIdentifiers),
    NetworkChunkPublisherUpdate(NetworkChunkPublisherUpdate),
    BiomeDefinitionList(BiomeDefinitionList),
    LevelSoundEvent(Box<LevelSoundEvent>),
    LevelEventGeneric(LevelEventGeneric),
    LecternUpdate(LecternUpdate),
    ClientCacheStatus(ClientCacheStatus),
    OnScreenTextureAnimation(OnScreenTextureAnimation),
    MapCreateLockedCopy(MapCreateLockedCopy),
    StructureTemplateDataRequest(StructureTemplateDataRequest),
    StructureTemplateDataResponse(StructureTemplateDataResponse),
    ClientCacheBlobStatus(ClientCacheBlobStatus),
    ClientCacheMissResponse(ClientCacheMissResponse),
    EducationSettings(EducationSettings),
    Emote(Emote),
    MultiplayerSettings(MultiplayerSettings),
    SettingsCommand(SettingsCommand),
    AnvilDamage(AnvilDamage),
    CompletedUsingItem(CompletedUsingItem),
    NetworkSettings(NetworkSettings),
    PlayerAuthInput(Box<PlayerAuthInput>),
    CreativeContent(CreativeContent),
    PlayerEnchantOptions(PlayerEnchantOptions),
    ItemStackRequest(ItemStackRequest),
    ItemStackResponse(ItemStackResponse),
    PlayerArmorDamage(PlayerArmorDamage),
    CodeBuilder(CodeBuilder),
    UpdatePlayerGameType(UpdatePlayerGameType),
    EmoteList(EmoteList),
    PositionTrackingDBServerBroadcast(PositionTrackingDBServerBroadcast),
    PositionTrackingDBClientRequest(PositionTrackingDBClientRequest),
    DebugInfo(DebugInfo),
    PacketViolationWarning(PacketViolationWarning),
    MotionPredictionHints(MotionPredictionHints),
    AnimateEntity(AnimateEntity),
    CameraShake(CameraShake),
    PlayerFog(PlayerFog),
    CorrectPlayerMovePrediction(CorrectPlayerMovePrediction),
    ItemRegistry(ItemRegistry),
    ClientboundDebugRenderer(ClientboundDebugRenderer),
    SyncActorProperty(SyncActorProperty),
    AddVolumeEntity(Box<AddVolumeEntity>),
    RemoveVolumeEntity(RemoveVolumeEntity),
    SimulationType(SimulationType),
    NpcDialogue(NpcDialogue),
    EduUriResource(EduUriResource),
    CreatePhoto(CreatePhoto),
    UpdateSubChunkBlocks(UpdateSubChunkBlocks),
    SubChunk(SubChunk),
    SubChunkRequest(SubChunkRequest),
    PlayerStartItemCooldown(PlayerStartItemCooldown),
    ScriptMessage(ScriptMessage),
    CodeBuilderSource(CodeBuilderSource),
    TickingAreasLoadStatus(TickingAreasLoadStatus),
    DimensionData(DimensionData),
    AgentActionEvent(AgentActionEvent),
    ChangeMobProperty(ChangeMobProperty),
    LessonProgress(LessonProgress),
    RequestAbility(RequestAbility),
    RequestPermissions(RequestPermissions),
    ToastRequest(ToastRequest),
    UpdateAbilities(UpdateAbilities),
    UpdateAdventureSettings(UpdateAdventureSettings),
    DeathInfo(DeathInfo),
    EditorNetwork(EditorNetwork),
    FeatureRegistry(FeatureRegistry),
    ServerStats(ServerStats),
    RequestNetworkSettings(RequestNetworkSettings),
    GameTestRequest(GameTestRequest),
    GameTestResults(GameTestResults),
    UpdateClientInputLocks(UpdateClientInputLocks),
    CameraPresets(CameraPresets),
    UnlockedRecipes(UnlockedRecipes),
    CameraInstruction(CameraInstruction),
    TrimData(TrimData),
    OpenSign(OpenSign),
    AgentAnimation(AgentAnimation),
    RefreshEntitlements(RefreshEntitlements),
    PlayerToggleCrafterSlotRequest(PlayerToggleCrafterSlotRequest),
    SetPlayerInventoryOptions(SetPlayerInventoryOptions),
    SetHud(SetHud),
    AwardAchievement(AwardAchievement),
    ClientboundCloseForm(ClientboundCloseForm),
    ServerboundLoadingScreen(ServerboundLoadingScreen),
    JigsawStructureData(JigsawStructureData),
    CurrentStructureFeature(CurrentStructureFeature),
    ServerboundDiagnostics(Box<ServerboundDiagnostics>),
    CameraAimAssist(CameraAimAssist),
    ContainerRegistryCleanup(ContainerRegistryCleanup),
    MovementEffect(MovementEffect),
    CameraAimAssistPresets(CameraAimAssistPresets),
    ClientCameraAimAssist(ClientCameraAimAssist),
    ClientMovementPredictionSync(ClientMovementPredictionSync),
    UpdateClientOptions(UpdateClientOptions),
    PlayerVideoCapture(PlayerVideoCapture),
    PlayerUpdateEntityOverrides(PlayerUpdateEntityOverrides),
    PlayerLocation(PlayerLocation),
    ClientboundControlSchemeSet(ClientboundControlSchemeSet),
    PrimitiveShapes(PrimitiveShapes),
    ServerboundPackSettingChange(ServerboundPackSettingChange),
    ClientboundDataStore(ClientboundDataStore),
    GraphicsOverrideParameter(GraphicsOverrideParameter),
    ServerboundDataStore(ServerboundDataStore),
    ClientboundDataDrivenUIShowScreen(ClientboundDataDrivenUIShowScreen),
    ClientboundDataDrivenUICloseScreen(ClientboundDataDrivenUICloseScreen),
    ClientboundDataDrivenUIReload(ClientboundDataDrivenUIReload),
    ClientboundTextureShift(Box<ClientboundTextureShift>),
    VoxelShapes(VoxelShapes),
    CameraSpline(CameraSpline),
    CameraAimAssistActorPriority(CameraAimAssistActorPriority),
    ResourcePacksReadyForValidation(ResourcePacksReadyForValidation),
    LocatorBar(LocatorBar),
    PartyChanged(PartyChanged),
    ServerboundDataDrivenScreenClosed(ServerboundDataDrivenScreenClosed),
    SyncWorldClocks(SyncWorldClocks),
    ClientboundAttributeLayerSync(ClientboundAttributeLayerSync),
    ServerStoreInfo(ServerStoreInfo),
    ServerPresenceInfo(ServerPresenceInfo),
    ClientboundUpdateSoundData(Box<ClientboundUpdateSoundData>),
    SendPartyDestinationCookie(SendPartyDestinationCookie),
    PartyDestinationCookieResponse(PartyDestinationCookieResponse),
}

impl From<Login> for Packet {
    fn from(value: Login) -> Self {
        Self::Login(value)
    }
}

impl From<PlayStatus> for Packet {
    fn from(value: PlayStatus) -> Self {
        Self::PlayStatus(value)
    }
}

impl From<ServerToClientHandshake> for Packet {
    fn from(value: ServerToClientHandshake) -> Self {
        Self::ServerToClientHandshake(value)
    }
}

impl From<ClientToServerHandshake> for Packet {
    fn from(value: ClientToServerHandshake) -> Self {
        Self::ClientToServerHandshake(value)
    }
}

impl From<Disconnect> for Packet {
    fn from(value: Disconnect) -> Self {
        Self::Disconnect(value)
    }
}

impl From<ResourcePacksInfo> for Packet {
    fn from(value: ResourcePacksInfo) -> Self {
        Self::ResourcePacksInfo(value)
    }
}

impl From<ResourcePackStack> for Packet {
    fn from(value: ResourcePackStack) -> Self {
        Self::ResourcePackStack(value)
    }
}

impl From<ResourcePackClientResponse> for Packet {
    fn from(value: ResourcePackClientResponse) -> Self {
        Self::ResourcePackClientResponse(value)
    }
}

impl From<Text> for Packet {
    fn from(value: Text) -> Self {
        Self::Text(value)
    }
}

impl From<SetTime> for Packet {
    fn from(value: SetTime) -> Self {
        Self::SetTime(value)
    }
}

impl From<StartGame> for Packet {
    fn from(value: StartGame) -> Self {
        Self::StartGame(Box::new(value))
    }
}

impl From<AddPlayer> for Packet {
    fn from(value: AddPlayer) -> Self {
        Self::AddPlayer(Box::new(value))
    }
}

impl From<AddActor> for Packet {
    fn from(value: AddActor) -> Self {
        Self::AddActor(Box::new(value))
    }
}

impl From<RemoveActor> for Packet {
    fn from(value: RemoveActor) -> Self {
        Self::RemoveActor(value)
    }
}

impl From<AddItemActor> for Packet {
    fn from(value: AddItemActor) -> Self {
        Self::AddItemActor(value)
    }
}

impl From<ServerPlayerPostMovePosition> for Packet {
    fn from(value: ServerPlayerPostMovePosition) -> Self {
        Self::ServerPlayerPostMovePosition(value)
    }
}

impl From<TakeItemActor> for Packet {
    fn from(value: TakeItemActor) -> Self {
        Self::TakeItemActor(value)
    }
}

impl From<MoveActorAbsolute> for Packet {
    fn from(value: MoveActorAbsolute) -> Self {
        Self::MoveActorAbsolute(value)
    }
}

impl From<MovePlayer> for Packet {
    fn from(value: MovePlayer) -> Self {
        Self::MovePlayer(Box::new(value))
    }
}

impl From<UpdateBlock> for Packet {
    fn from(value: UpdateBlock) -> Self {
        Self::UpdateBlock(value)
    }
}

impl From<AddPainting> for Packet {
    fn from(value: AddPainting) -> Self {
        Self::AddPainting(value)
    }
}

impl From<LevelEvent> for Packet {
    fn from(value: LevelEvent) -> Self {
        Self::LevelEvent(value)
    }
}

impl From<BlockEvent> for Packet {
    fn from(value: BlockEvent) -> Self {
        Self::BlockEvent(value)
    }
}

impl From<ActorEvent> for Packet {
    fn from(value: ActorEvent) -> Self {
        Self::ActorEvent(value)
    }
}

impl From<MobEffect> for Packet {
    fn from(value: MobEffect) -> Self {
        Self::MobEffect(Box::new(value))
    }
}

impl From<UpdateAttributes> for Packet {
    fn from(value: UpdateAttributes) -> Self {
        Self::UpdateAttributes(value)
    }
}

impl From<InventoryTransaction> for Packet {
    fn from(value: InventoryTransaction) -> Self {
        Self::InventoryTransaction(value)
    }
}

impl From<MobEquipment> for Packet {
    fn from(value: MobEquipment) -> Self {
        Self::MobEquipment(value)
    }
}

impl From<MobArmorEquipment> for Packet {
    fn from(value: MobArmorEquipment) -> Self {
        Self::MobArmorEquipment(value)
    }
}

impl From<Interact> for Packet {
    fn from(value: Interact) -> Self {
        Self::Interact(value)
    }
}

impl From<BlockPickRequest> for Packet {
    fn from(value: BlockPickRequest) -> Self {
        Self::BlockPickRequest(value)
    }
}

impl From<ActorPickRequest> for Packet {
    fn from(value: ActorPickRequest) -> Self {
        Self::ActorPickRequest(value)
    }
}

impl From<PlayerAction> for Packet {
    fn from(value: PlayerAction) -> Self {
        Self::PlayerAction(value)
    }
}

impl From<HurtArmor> for Packet {
    fn from(value: HurtArmor) -> Self {
        Self::HurtArmor(value)
    }
}

impl From<SetActorData> for Packet {
    fn from(value: SetActorData) -> Self {
        Self::SetActorData(value)
    }
}

impl From<SetActorMotion> for Packet {
    fn from(value: SetActorMotion) -> Self {
        Self::SetActorMotion(value)
    }
}

impl From<SetActorLink> for Packet {
    fn from(value: SetActorLink) -> Self {
        Self::SetActorLink(value)
    }
}

impl From<SetHealth> for Packet {
    fn from(value: SetHealth) -> Self {
        Self::SetHealth(value)
    }
}

impl From<SetSpawnPosition> for Packet {
    fn from(value: SetSpawnPosition) -> Self {
        Self::SetSpawnPosition(value)
    }
}

impl From<Animate> for Packet {
    fn from(value: Animate) -> Self {
        Self::Animate(value)
    }
}

impl From<Respawn> for Packet {
    fn from(value: Respawn) -> Self {
        Self::Respawn(value)
    }
}

impl From<ContainerOpen> for Packet {
    fn from(value: ContainerOpen) -> Self {
        Self::ContainerOpen(value)
    }
}

impl From<ContainerClose> for Packet {
    fn from(value: ContainerClose) -> Self {
        Self::ContainerClose(value)
    }
}

impl From<PlayerHotbar> for Packet {
    fn from(value: PlayerHotbar) -> Self {
        Self::PlayerHotbar(value)
    }
}

impl From<InventoryContent> for Packet {
    fn from(value: InventoryContent) -> Self {
        Self::InventoryContent(value)
    }
}

impl From<InventorySlot> for Packet {
    fn from(value: InventorySlot) -> Self {
        Self::InventorySlot(value)
    }
}

impl From<ContainerSetData> for Packet {
    fn from(value: ContainerSetData) -> Self {
        Self::ContainerSetData(value)
    }
}

impl From<CraftingData> for Packet {
    fn from(value: CraftingData) -> Self {
        Self::CraftingData(Box::new(value))
    }
}

impl From<GuiDataPickItem> for Packet {
    fn from(value: GuiDataPickItem) -> Self {
        Self::GuiDataPickItem(value)
    }
}

impl From<BlockActorData> for Packet {
    fn from(value: BlockActorData) -> Self {
        Self::BlockActorData(value)
    }
}

impl From<LevelChunk> for Packet {
    fn from(value: LevelChunk) -> Self {
        Self::LevelChunk(value)
    }
}

impl From<SetCommandsEnabled> for Packet {
    fn from(value: SetCommandsEnabled) -> Self {
        Self::SetCommandsEnabled(value)
    }
}

impl From<SetDifficulty> for Packet {
    fn from(value: SetDifficulty) -> Self {
        Self::SetDifficulty(value)
    }
}

impl From<ChangeDimension> for Packet {
    fn from(value: ChangeDimension) -> Self {
        Self::ChangeDimension(value)
    }
}

impl From<SetPlayerGameType> for Packet {
    fn from(value: SetPlayerGameType) -> Self {
        Self::SetPlayerGameType(value)
    }
}

impl From<PlayerList> for Packet {
    fn from(value: PlayerList) -> Self {
        Self::PlayerList(value)
    }
}

impl From<SimpleEvent> for Packet {
    fn from(value: SimpleEvent) -> Self {
        Self::SimpleEvent(value)
    }
}

impl From<LegacyTelemetryEvent> for Packet {
    fn from(value: LegacyTelemetryEvent) -> Self {
        Self::LegacyTelemetryEvent(value)
    }
}

impl From<SpawnExperienceOrb> for Packet {
    fn from(value: SpawnExperienceOrb) -> Self {
        Self::SpawnExperienceOrb(value)
    }
}

impl From<ClientboundMapItemData> for Packet {
    fn from(value: ClientboundMapItemData) -> Self {
        Self::ClientboundMapItemData(Box::new(value))
    }
}

impl From<MapInfoRequest> for Packet {
    fn from(value: MapInfoRequest) -> Self {
        Self::MapInfoRequest(value)
    }
}

impl From<RequestChunkRadius> for Packet {
    fn from(value: RequestChunkRadius) -> Self {
        Self::RequestChunkRadius(value)
    }
}

impl From<ChunkRadiusUpdated> for Packet {
    fn from(value: ChunkRadiusUpdated) -> Self {
        Self::ChunkRadiusUpdated(value)
    }
}

impl From<GameRulesChanged> for Packet {
    fn from(value: GameRulesChanged) -> Self {
        Self::GameRulesChanged(value)
    }
}

impl From<Camera> for Packet {
    fn from(value: Camera) -> Self {
        Self::Camera(value)
    }
}

impl From<BossEvent> for Packet {
    fn from(value: BossEvent) -> Self {
        Self::BossEvent(Box::new(value))
    }
}

impl From<ShowCredits> for Packet {
    fn from(value: ShowCredits) -> Self {
        Self::ShowCredits(value)
    }
}

impl From<AvailableCommands> for Packet {
    fn from(value: AvailableCommands) -> Self {
        Self::AvailableCommands(Box::new(value))
    }
}

impl From<CommandRequest> for Packet {
    fn from(value: CommandRequest) -> Self {
        Self::CommandRequest(value)
    }
}

impl From<CommandBlockUpdate> for Packet {
    fn from(value: CommandBlockUpdate) -> Self {
        Self::CommandBlockUpdate(Box::new(value))
    }
}

impl From<CommandOutput> for Packet {
    fn from(value: CommandOutput) -> Self {
        Self::CommandOutput(value)
    }
}

impl From<UpdateTrade> for Packet {
    fn from(value: UpdateTrade) -> Self {
        Self::UpdateTrade(Box::new(value))
    }
}

impl From<UpdateEquip> for Packet {
    fn from(value: UpdateEquip) -> Self {
        Self::UpdateEquip(value)
    }
}

impl From<ResourcePackDataInfo> for Packet {
    fn from(value: ResourcePackDataInfo) -> Self {
        Self::ResourcePackDataInfo(value)
    }
}

impl From<ResourcePackChunkData> for Packet {
    fn from(value: ResourcePackChunkData) -> Self {
        Self::ResourcePackChunkData(value)
    }
}

impl From<ResourcePackChunkRequest> for Packet {
    fn from(value: ResourcePackChunkRequest) -> Self {
        Self::ResourcePackChunkRequest(value)
    }
}

impl From<Transfer> for Packet {
    fn from(value: Transfer) -> Self {
        Self::Transfer(value)
    }
}

impl From<PlaySound> for Packet {
    fn from(value: PlaySound) -> Self {
        Self::PlaySound(value)
    }
}

impl From<StopSound> for Packet {
    fn from(value: StopSound) -> Self {
        Self::StopSound(value)
    }
}

impl From<SetTitle> for Packet {
    fn from(value: SetTitle) -> Self {
        Self::SetTitle(Box::new(value))
    }
}

impl From<AddBehaviorTree> for Packet {
    fn from(value: AddBehaviorTree) -> Self {
        Self::AddBehaviorTree(value)
    }
}

impl From<StructureBlockUpdate> for Packet {
    fn from(value: StructureBlockUpdate) -> Self {
        Self::StructureBlockUpdate(value)
    }
}

impl From<ShowStoreOffer> for Packet {
    fn from(value: ShowStoreOffer) -> Self {
        Self::ShowStoreOffer(value)
    }
}

impl From<PurchaseReceipt> for Packet {
    fn from(value: PurchaseReceipt) -> Self {
        Self::PurchaseReceipt(value)
    }
}

impl From<PlayerSkin> for Packet {
    fn from(value: PlayerSkin) -> Self {
        Self::PlayerSkin(value)
    }
}

impl From<SubClientLogin> for Packet {
    fn from(value: SubClientLogin) -> Self {
        Self::SubClientLogin(value)
    }
}

impl From<AutomationClientConnect> for Packet {
    fn from(value: AutomationClientConnect) -> Self {
        Self::AutomationClientConnect(value)
    }
}

impl From<SetLastHurtBy> for Packet {
    fn from(value: SetLastHurtBy) -> Self {
        Self::SetLastHurtBy(value)
    }
}

impl From<BookEdit> for Packet {
    fn from(value: BookEdit) -> Self {
        Self::BookEdit(value)
    }
}

impl From<NpcRequest> for Packet {
    fn from(value: NpcRequest) -> Self {
        Self::NpcRequest(value)
    }
}

impl From<PhotoTransfer> for Packet {
    fn from(value: PhotoTransfer) -> Self {
        Self::PhotoTransfer(value)
    }
}

impl From<ModalFormRequest> for Packet {
    fn from(value: ModalFormRequest) -> Self {
        Self::ModalFormRequest(value)
    }
}

impl From<ModalFormResponse> for Packet {
    fn from(value: ModalFormResponse) -> Self {
        Self::ModalFormResponse(value)
    }
}

impl From<ServerSettingsRequest> for Packet {
    fn from(value: ServerSettingsRequest) -> Self {
        Self::ServerSettingsRequest(value)
    }
}

impl From<ServerSettingsResponse> for Packet {
    fn from(value: ServerSettingsResponse) -> Self {
        Self::ServerSettingsResponse(value)
    }
}

impl From<ShowProfile> for Packet {
    fn from(value: ShowProfile) -> Self {
        Self::ShowProfile(value)
    }
}

impl From<SetDefaultGameType> for Packet {
    fn from(value: SetDefaultGameType) -> Self {
        Self::SetDefaultGameType(value)
    }
}

impl From<RemoveObjective> for Packet {
    fn from(value: RemoveObjective) -> Self {
        Self::RemoveObjective(value)
    }
}

impl From<SetDisplayObjective> for Packet {
    fn from(value: SetDisplayObjective) -> Self {
        Self::SetDisplayObjective(value)
    }
}

impl From<SetScore> for Packet {
    fn from(value: SetScore) -> Self {
        Self::SetScore(value)
    }
}

impl From<LabTable> for Packet {
    fn from(value: LabTable) -> Self {
        Self::LabTable(value)
    }
}

impl From<UpdateBlockSynced> for Packet {
    fn from(value: UpdateBlockSynced) -> Self {
        Self::UpdateBlockSynced(value)
    }
}

impl From<MoveActorDelta> for Packet {
    fn from(value: MoveActorDelta) -> Self {
        Self::MoveActorDelta(value)
    }
}

impl From<SetScoreboardIdentity> for Packet {
    fn from(value: SetScoreboardIdentity) -> Self {
        Self::SetScoreboardIdentity(value)
    }
}

impl From<SetLocalPlayerAsInitialized> for Packet {
    fn from(value: SetLocalPlayerAsInitialized) -> Self {
        Self::SetLocalPlayerAsInitialized(value)
    }
}

impl From<UpdateSoftEnum> for Packet {
    fn from(value: UpdateSoftEnum) -> Self {
        Self::UpdateSoftEnum(value)
    }
}

impl From<NetworkStackLatency> for Packet {
    fn from(value: NetworkStackLatency) -> Self {
        Self::NetworkStackLatency(value)
    }
}

impl From<SpawnParticleEffect> for Packet {
    fn from(value: SpawnParticleEffect) -> Self {
        Self::SpawnParticleEffect(value)
    }
}

impl From<AvailableActorIdentifiers> for Packet {
    fn from(value: AvailableActorIdentifiers) -> Self {
        Self::AvailableActorIdentifiers(value)
    }
}

impl From<NetworkChunkPublisherUpdate> for Packet {
    fn from(value: NetworkChunkPublisherUpdate) -> Self {
        Self::NetworkChunkPublisherUpdate(value)
    }
}

impl From<BiomeDefinitionList> for Packet {
    fn from(value: BiomeDefinitionList) -> Self {
        Self::BiomeDefinitionList(value)
    }
}

impl From<LevelSoundEvent> for Packet {
    fn from(value: LevelSoundEvent) -> Self {
        Self::LevelSoundEvent(Box::new(value))
    }
}

impl From<LevelEventGeneric> for Packet {
    fn from(value: LevelEventGeneric) -> Self {
        Self::LevelEventGeneric(value)
    }
}

impl From<LecternUpdate> for Packet {
    fn from(value: LecternUpdate) -> Self {
        Self::LecternUpdate(value)
    }
}

impl From<ClientCacheStatus> for Packet {
    fn from(value: ClientCacheStatus) -> Self {
        Self::ClientCacheStatus(value)
    }
}

impl From<OnScreenTextureAnimation> for Packet {
    fn from(value: OnScreenTextureAnimation) -> Self {
        Self::OnScreenTextureAnimation(value)
    }
}

impl From<MapCreateLockedCopy> for Packet {
    fn from(value: MapCreateLockedCopy) -> Self {
        Self::MapCreateLockedCopy(value)
    }
}

impl From<StructureTemplateDataRequest> for Packet {
    fn from(value: StructureTemplateDataRequest) -> Self {
        Self::StructureTemplateDataRequest(value)
    }
}

impl From<StructureTemplateDataResponse> for Packet {
    fn from(value: StructureTemplateDataResponse) -> Self {
        Self::StructureTemplateDataResponse(value)
    }
}

impl From<ClientCacheBlobStatus> for Packet {
    fn from(value: ClientCacheBlobStatus) -> Self {
        Self::ClientCacheBlobStatus(value)
    }
}

impl From<ClientCacheMissResponse> for Packet {
    fn from(value: ClientCacheMissResponse) -> Self {
        Self::ClientCacheMissResponse(value)
    }
}

impl From<EducationSettings> for Packet {
    fn from(value: EducationSettings) -> Self {
        Self::EducationSettings(value)
    }
}

impl From<Emote> for Packet {
    fn from(value: Emote) -> Self {
        Self::Emote(value)
    }
}

impl From<MultiplayerSettings> for Packet {
    fn from(value: MultiplayerSettings) -> Self {
        Self::MultiplayerSettings(value)
    }
}

impl From<SettingsCommand> for Packet {
    fn from(value: SettingsCommand) -> Self {
        Self::SettingsCommand(value)
    }
}

impl From<AnvilDamage> for Packet {
    fn from(value: AnvilDamage) -> Self {
        Self::AnvilDamage(value)
    }
}

impl From<CompletedUsingItem> for Packet {
    fn from(value: CompletedUsingItem) -> Self {
        Self::CompletedUsingItem(value)
    }
}

impl From<NetworkSettings> for Packet {
    fn from(value: NetworkSettings) -> Self {
        Self::NetworkSettings(value)
    }
}

impl From<PlayerAuthInput> for Packet {
    fn from(value: PlayerAuthInput) -> Self {
        Self::PlayerAuthInput(Box::new(value))
    }
}

impl From<CreativeContent> for Packet {
    fn from(value: CreativeContent) -> Self {
        Self::CreativeContent(value)
    }
}

impl From<PlayerEnchantOptions> for Packet {
    fn from(value: PlayerEnchantOptions) -> Self {
        Self::PlayerEnchantOptions(value)
    }
}

impl From<ItemStackRequest> for Packet {
    fn from(value: ItemStackRequest) -> Self {
        Self::ItemStackRequest(value)
    }
}

impl From<ItemStackResponse> for Packet {
    fn from(value: ItemStackResponse) -> Self {
        Self::ItemStackResponse(value)
    }
}

impl From<PlayerArmorDamage> for Packet {
    fn from(value: PlayerArmorDamage) -> Self {
        Self::PlayerArmorDamage(value)
    }
}

impl From<CodeBuilder> for Packet {
    fn from(value: CodeBuilder) -> Self {
        Self::CodeBuilder(value)
    }
}

impl From<UpdatePlayerGameType> for Packet {
    fn from(value: UpdatePlayerGameType) -> Self {
        Self::UpdatePlayerGameType(value)
    }
}

impl From<EmoteList> for Packet {
    fn from(value: EmoteList) -> Self {
        Self::EmoteList(value)
    }
}

impl From<PositionTrackingDBServerBroadcast> for Packet {
    fn from(value: PositionTrackingDBServerBroadcast) -> Self {
        Self::PositionTrackingDBServerBroadcast(value)
    }
}

impl From<PositionTrackingDBClientRequest> for Packet {
    fn from(value: PositionTrackingDBClientRequest) -> Self {
        Self::PositionTrackingDBClientRequest(value)
    }
}

impl From<DebugInfo> for Packet {
    fn from(value: DebugInfo) -> Self {
        Self::DebugInfo(value)
    }
}

impl From<PacketViolationWarning> for Packet {
    fn from(value: PacketViolationWarning) -> Self {
        Self::PacketViolationWarning(value)
    }
}

impl From<MotionPredictionHints> for Packet {
    fn from(value: MotionPredictionHints) -> Self {
        Self::MotionPredictionHints(value)
    }
}

impl From<AnimateEntity> for Packet {
    fn from(value: AnimateEntity) -> Self {
        Self::AnimateEntity(value)
    }
}

impl From<CameraShake> for Packet {
    fn from(value: CameraShake) -> Self {
        Self::CameraShake(value)
    }
}

impl From<PlayerFog> for Packet {
    fn from(value: PlayerFog) -> Self {
        Self::PlayerFog(value)
    }
}

impl From<CorrectPlayerMovePrediction> for Packet {
    fn from(value: CorrectPlayerMovePrediction) -> Self {
        Self::CorrectPlayerMovePrediction(value)
    }
}

impl From<ItemRegistry> for Packet {
    fn from(value: ItemRegistry) -> Self {
        Self::ItemRegistry(value)
    }
}

impl From<ClientboundDebugRenderer> for Packet {
    fn from(value: ClientboundDebugRenderer) -> Self {
        Self::ClientboundDebugRenderer(value)
    }
}

impl From<SyncActorProperty> for Packet {
    fn from(value: SyncActorProperty) -> Self {
        Self::SyncActorProperty(value)
    }
}

impl From<AddVolumeEntity> for Packet {
    fn from(value: AddVolumeEntity) -> Self {
        Self::AddVolumeEntity(Box::new(value))
    }
}

impl From<RemoveVolumeEntity> for Packet {
    fn from(value: RemoveVolumeEntity) -> Self {
        Self::RemoveVolumeEntity(value)
    }
}

impl From<SimulationType> for Packet {
    fn from(value: SimulationType) -> Self {
        Self::SimulationType(value)
    }
}

impl From<NpcDialogue> for Packet {
    fn from(value: NpcDialogue) -> Self {
        Self::NpcDialogue(value)
    }
}

impl From<EduUriResource> for Packet {
    fn from(value: EduUriResource) -> Self {
        Self::EduUriResource(value)
    }
}

impl From<CreatePhoto> for Packet {
    fn from(value: CreatePhoto) -> Self {
        Self::CreatePhoto(value)
    }
}

impl From<UpdateSubChunkBlocks> for Packet {
    fn from(value: UpdateSubChunkBlocks) -> Self {
        Self::UpdateSubChunkBlocks(value)
    }
}

impl From<SubChunk> for Packet {
    fn from(value: SubChunk) -> Self {
        Self::SubChunk(value)
    }
}

impl From<SubChunkRequest> for Packet {
    fn from(value: SubChunkRequest) -> Self {
        Self::SubChunkRequest(value)
    }
}

impl From<PlayerStartItemCooldown> for Packet {
    fn from(value: PlayerStartItemCooldown) -> Self {
        Self::PlayerStartItemCooldown(value)
    }
}

impl From<ScriptMessage> for Packet {
    fn from(value: ScriptMessage) -> Self {
        Self::ScriptMessage(value)
    }
}

impl From<CodeBuilderSource> for Packet {
    fn from(value: CodeBuilderSource) -> Self {
        Self::CodeBuilderSource(value)
    }
}

impl From<TickingAreasLoadStatus> for Packet {
    fn from(value: TickingAreasLoadStatus) -> Self {
        Self::TickingAreasLoadStatus(value)
    }
}

impl From<DimensionData> for Packet {
    fn from(value: DimensionData) -> Self {
        Self::DimensionData(value)
    }
}

impl From<AgentActionEvent> for Packet {
    fn from(value: AgentActionEvent) -> Self {
        Self::AgentActionEvent(value)
    }
}

impl From<ChangeMobProperty> for Packet {
    fn from(value: ChangeMobProperty) -> Self {
        Self::ChangeMobProperty(value)
    }
}

impl From<LessonProgress> for Packet {
    fn from(value: LessonProgress) -> Self {
        Self::LessonProgress(value)
    }
}

impl From<RequestAbility> for Packet {
    fn from(value: RequestAbility) -> Self {
        Self::RequestAbility(value)
    }
}

impl From<RequestPermissions> for Packet {
    fn from(value: RequestPermissions) -> Self {
        Self::RequestPermissions(value)
    }
}

impl From<ToastRequest> for Packet {
    fn from(value: ToastRequest) -> Self {
        Self::ToastRequest(value)
    }
}

impl From<UpdateAbilities> for Packet {
    fn from(value: UpdateAbilities) -> Self {
        Self::UpdateAbilities(value)
    }
}

impl From<UpdateAdventureSettings> for Packet {
    fn from(value: UpdateAdventureSettings) -> Self {
        Self::UpdateAdventureSettings(value)
    }
}

impl From<DeathInfo> for Packet {
    fn from(value: DeathInfo) -> Self {
        Self::DeathInfo(value)
    }
}

impl From<EditorNetwork> for Packet {
    fn from(value: EditorNetwork) -> Self {
        Self::EditorNetwork(value)
    }
}

impl From<FeatureRegistry> for Packet {
    fn from(value: FeatureRegistry) -> Self {
        Self::FeatureRegistry(value)
    }
}

impl From<ServerStats> for Packet {
    fn from(value: ServerStats) -> Self {
        Self::ServerStats(value)
    }
}

impl From<RequestNetworkSettings> for Packet {
    fn from(value: RequestNetworkSettings) -> Self {
        Self::RequestNetworkSettings(value)
    }
}

impl From<GameTestRequest> for Packet {
    fn from(value: GameTestRequest) -> Self {
        Self::GameTestRequest(value)
    }
}

impl From<GameTestResults> for Packet {
    fn from(value: GameTestResults) -> Self {
        Self::GameTestResults(value)
    }
}

impl From<UpdateClientInputLocks> for Packet {
    fn from(value: UpdateClientInputLocks) -> Self {
        Self::UpdateClientInputLocks(value)
    }
}

impl From<CameraPresets> for Packet {
    fn from(value: CameraPresets) -> Self {
        Self::CameraPresets(value)
    }
}

impl From<UnlockedRecipes> for Packet {
    fn from(value: UnlockedRecipes) -> Self {
        Self::UnlockedRecipes(value)
    }
}

impl From<CameraInstruction> for Packet {
    fn from(value: CameraInstruction) -> Self {
        Self::CameraInstruction(value)
    }
}

impl From<TrimData> for Packet {
    fn from(value: TrimData) -> Self {
        Self::TrimData(value)
    }
}

impl From<OpenSign> for Packet {
    fn from(value: OpenSign) -> Self {
        Self::OpenSign(value)
    }
}

impl From<AgentAnimation> for Packet {
    fn from(value: AgentAnimation) -> Self {
        Self::AgentAnimation(value)
    }
}

impl From<RefreshEntitlements> for Packet {
    fn from(value: RefreshEntitlements) -> Self {
        Self::RefreshEntitlements(value)
    }
}

impl From<PlayerToggleCrafterSlotRequest> for Packet {
    fn from(value: PlayerToggleCrafterSlotRequest) -> Self {
        Self::PlayerToggleCrafterSlotRequest(value)
    }
}

impl From<SetPlayerInventoryOptions> for Packet {
    fn from(value: SetPlayerInventoryOptions) -> Self {
        Self::SetPlayerInventoryOptions(value)
    }
}

impl From<SetHud> for Packet {
    fn from(value: SetHud) -> Self {
        Self::SetHud(value)
    }
}

impl From<AwardAchievement> for Packet {
    fn from(value: AwardAchievement) -> Self {
        Self::AwardAchievement(value)
    }
}

impl From<ClientboundCloseForm> for Packet {
    fn from(value: ClientboundCloseForm) -> Self {
        Self::ClientboundCloseForm(value)
    }
}

impl From<ServerboundLoadingScreen> for Packet {
    fn from(value: ServerboundLoadingScreen) -> Self {
        Self::ServerboundLoadingScreen(value)
    }
}

impl From<JigsawStructureData> for Packet {
    fn from(value: JigsawStructureData) -> Self {
        Self::JigsawStructureData(value)
    }
}

impl From<CurrentStructureFeature> for Packet {
    fn from(value: CurrentStructureFeature) -> Self {
        Self::CurrentStructureFeature(value)
    }
}

impl From<ServerboundDiagnostics> for Packet {
    fn from(value: ServerboundDiagnostics) -> Self {
        Self::ServerboundDiagnostics(Box::new(value))
    }
}

impl From<CameraAimAssist> for Packet {
    fn from(value: CameraAimAssist) -> Self {
        Self::CameraAimAssist(value)
    }
}

impl From<ContainerRegistryCleanup> for Packet {
    fn from(value: ContainerRegistryCleanup) -> Self {
        Self::ContainerRegistryCleanup(value)
    }
}

impl From<MovementEffect> for Packet {
    fn from(value: MovementEffect) -> Self {
        Self::MovementEffect(value)
    }
}

impl From<CameraAimAssistPresets> for Packet {
    fn from(value: CameraAimAssistPresets) -> Self {
        Self::CameraAimAssistPresets(value)
    }
}

impl From<ClientCameraAimAssist> for Packet {
    fn from(value: ClientCameraAimAssist) -> Self {
        Self::ClientCameraAimAssist(value)
    }
}

impl From<ClientMovementPredictionSync> for Packet {
    fn from(value: ClientMovementPredictionSync) -> Self {
        Self::ClientMovementPredictionSync(value)
    }
}

impl From<UpdateClientOptions> for Packet {
    fn from(value: UpdateClientOptions) -> Self {
        Self::UpdateClientOptions(value)
    }
}

impl From<PlayerVideoCapture> for Packet {
    fn from(value: PlayerVideoCapture) -> Self {
        Self::PlayerVideoCapture(value)
    }
}

impl From<PlayerUpdateEntityOverrides> for Packet {
    fn from(value: PlayerUpdateEntityOverrides) -> Self {
        Self::PlayerUpdateEntityOverrides(value)
    }
}

impl From<PlayerLocation> for Packet {
    fn from(value: PlayerLocation) -> Self {
        Self::PlayerLocation(value)
    }
}

impl From<ClientboundControlSchemeSet> for Packet {
    fn from(value: ClientboundControlSchemeSet) -> Self {
        Self::ClientboundControlSchemeSet(value)
    }
}

impl From<PrimitiveShapes> for Packet {
    fn from(value: PrimitiveShapes) -> Self {
        Self::PrimitiveShapes(value)
    }
}

impl From<ServerboundPackSettingChange> for Packet {
    fn from(value: ServerboundPackSettingChange) -> Self {
        Self::ServerboundPackSettingChange(value)
    }
}

impl From<ClientboundDataStore> for Packet {
    fn from(value: ClientboundDataStore) -> Self {
        Self::ClientboundDataStore(value)
    }
}

impl From<GraphicsOverrideParameter> for Packet {
    fn from(value: GraphicsOverrideParameter) -> Self {
        Self::GraphicsOverrideParameter(value)
    }
}

impl From<ServerboundDataStore> for Packet {
    fn from(value: ServerboundDataStore) -> Self {
        Self::ServerboundDataStore(value)
    }
}

impl From<ClientboundDataDrivenUIShowScreen> for Packet {
    fn from(value: ClientboundDataDrivenUIShowScreen) -> Self {
        Self::ClientboundDataDrivenUIShowScreen(value)
    }
}

impl From<ClientboundDataDrivenUICloseScreen> for Packet {
    fn from(value: ClientboundDataDrivenUICloseScreen) -> Self {
        Self::ClientboundDataDrivenUICloseScreen(value)
    }
}

impl From<ClientboundDataDrivenUIReload> for Packet {
    fn from(value: ClientboundDataDrivenUIReload) -> Self {
        Self::ClientboundDataDrivenUIReload(value)
    }
}

impl From<ClientboundTextureShift> for Packet {
    fn from(value: ClientboundTextureShift) -> Self {
        Self::ClientboundTextureShift(Box::new(value))
    }
}

impl From<VoxelShapes> for Packet {
    fn from(value: VoxelShapes) -> Self {
        Self::VoxelShapes(value)
    }
}

impl From<CameraSpline> for Packet {
    fn from(value: CameraSpline) -> Self {
        Self::CameraSpline(value)
    }
}

impl From<CameraAimAssistActorPriority> for Packet {
    fn from(value: CameraAimAssistActorPriority) -> Self {
        Self::CameraAimAssistActorPriority(value)
    }
}

impl From<ResourcePacksReadyForValidation> for Packet {
    fn from(value: ResourcePacksReadyForValidation) -> Self {
        Self::ResourcePacksReadyForValidation(value)
    }
}

impl From<LocatorBar> for Packet {
    fn from(value: LocatorBar) -> Self {
        Self::LocatorBar(value)
    }
}

impl From<PartyChanged> for Packet {
    fn from(value: PartyChanged) -> Self {
        Self::PartyChanged(value)
    }
}

impl From<ServerboundDataDrivenScreenClosed> for Packet {
    fn from(value: ServerboundDataDrivenScreenClosed) -> Self {
        Self::ServerboundDataDrivenScreenClosed(value)
    }
}

impl From<SyncWorldClocks> for Packet {
    fn from(value: SyncWorldClocks) -> Self {
        Self::SyncWorldClocks(value)
    }
}

impl From<ClientboundAttributeLayerSync> for Packet {
    fn from(value: ClientboundAttributeLayerSync) -> Self {
        Self::ClientboundAttributeLayerSync(value)
    }
}

impl From<ServerStoreInfo> for Packet {
    fn from(value: ServerStoreInfo) -> Self {
        Self::ServerStoreInfo(value)
    }
}

impl From<ServerPresenceInfo> for Packet {
    fn from(value: ServerPresenceInfo) -> Self {
        Self::ServerPresenceInfo(value)
    }
}

impl From<ClientboundUpdateSoundData> for Packet {
    fn from(value: ClientboundUpdateSoundData) -> Self {
        Self::ClientboundUpdateSoundData(Box::new(value))
    }
}

impl From<SendPartyDestinationCookie> for Packet {
    fn from(value: SendPartyDestinationCookie) -> Self {
        Self::SendPartyDestinationCookie(value)
    }
}

impl From<PartyDestinationCookieResponse> for Packet {
    fn from(value: PartyDestinationCookieResponse) -> Self {
        Self::PartyDestinationCookieResponse(value)
    }
}
