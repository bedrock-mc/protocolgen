// Code generated from canonical protocol manifest v2. DO NOT EDIT.

use crate::enums::*;
use crate::types::*;
use crate::wire;

/// Login is sent when the client initially tries to join the server. It is the first packet sent
/// and contains information specific to the player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Login {
    pub client_network_version: wire::I32BE,
    /// `connection_request` is a string containing information about the player and JWTs that may be
    /// used to verify if the player is connected to XBOX Live. The connection request also contains the
    /// necessary client public key to initiate encryption.
    pub connection_request: String,
}

impl Login {
    pub const ID: u32 = 1;
}
/// PlayStatus is sent by the server to update a player on the play status. This includes failed
/// statuses due to a mismatched version, but also success statuses.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayStatus {
    /// `status` is the status of the packet. It is one of the constants found above.
    pub status: PlayStatusType,
}

impl PlayStatus {
    pub const ID: u32 = 2;
}
/// ServerToClientHandshake is sent by the server to the client to complete the key exchange in
/// order to initialise encryption on client and server side. It is followed up by a
/// ClientToServerHandshake packet from the client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerToClientHandshake {
    pub handshake_web_token: String,
}

impl ServerToClientHandshake {
    pub const ID: u32 = 3;
}
/// ClientToServerHandshake is sent by the client in response to a ServerToClientHandshake packet
/// sent by the server. It is the first encrypted packet in the login handshake and serves as a
/// confirmation that encryption is correctly initialised client side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientToServerHandshake {
}

impl ClientToServerHandshake {
    pub const ID: u32 = 4;
}
/// Disconnect may be sent by the server to disconnect the client using an optional message to send
/// as the disconnect screen.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Disconnect {
    /// `reason` is the reason for the disconnection. This affects the error code displayed on the Ore
    /// UI disconnection screen and is one of the constants above.
    pub reason: ConnectionDisconnectFailReason,
    pub messages: DisconnectMessages,
}

impl Disconnect {
    pub const ID: u32 = 5;
}
/// ResourcePacksInfo is sent by the server to inform the client on what resource packs the server
/// has. It sends a list of the resource packs it has and basic information on them like the version
/// and description.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePacksInfo {
    pub resource_pack_required: bool,
    pub has_addon_packs: bool,
    /// `has_scripts` specifies if any of the resource packs contain scripts in them. If set to true,
    /// only clients that support scripts will be able to download them.
    pub has_scripts: bool,
    /// `force_disable_vibrant_visuals` specifies if the vibrant visuals feature should be forcibly
    /// disabled on the server. If set to true, the server will ensure that vibrant visuals are not
    /// enabled, regardless of the client's settings.
    pub force_disable_vibrant_visuals: bool,
    pub world_template_id_and_version: PackIdVersion,
    pub resource_packs: Vec<PackInfoData>,
}

impl ResourcePacksInfo {
    pub const ID: u32 = 6;
}
/// ResourcePackStack is sent by the server to send the order in which resource packs and behaviour
/// packs should be applied (and downloaded) by the client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackStack {
    /// `texture_pack_required` specifies if the client must accept the texture packs the server has in
    /// order to join the server. If set to true, the client gets the option to either download the
    /// resource packs and join, or quit entirely. Behaviour packs never have to be downloaded.
    pub texture_pack_required: bool,
    pub texture_pack_list: Vec<PackInstanceId>,
    /// `base_game_version` is the vanilla version that the client should set its resource pack stack
    /// to.
    pub base_game_version: String,
    /// `experiments` holds a list of experiments that are either enabled or disabled in the world that
    /// the player spawns in. It is not clear why experiments are sent both here and in the StartGame
    /// packet.
    pub experiments: Experiments,
    /// `include_editor_packs` specifies if vanilla editor packs should be included in the resource pack
    /// stack when connecting to an editor world.
    pub include_editor_packs: bool,
}

impl ResourcePackStack {
    pub const ID: u32 = 7;
}
/// ResourcePackClientResponse is sent by the client in response to resource packets sent by the
/// server. It is used to let the server know what action needs to be taken for the client to have
/// all resource packs ready and set.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackClientResponse {
    /// `response` is the response type of the response. It is one of the constants found above.
    pub response: ResourcePackClientResponseData,
}

impl ResourcePackClientResponse {
    pub const ID: u32 = 8;
}
/// Text is sent by the client to the server to send chat messages, and by the server to the client
/// to forward or send messages, which may be chat, popups, tips etc.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Text {
    pub localize: bool,
    pub body: TextData,
    pub sender_xuid: String,
    pub platform_id: String,
    /// `filtered_message` is a filtered version of Message with all the profanity removed. The client
    /// will use this over Message if this field is not empty and they have the "Filter Profanity"
    /// setting enabled.
    /// Wire presence: optional value is preceded by a presence marker.
    pub filtered_message: Option<String>,
}

impl Text {
    pub const ID: u32 = 9;
}
/// SetTime is sent by the server to update the current time client-side. The client actually
/// advances time client-side by itself, so this packet does not need to be sent each tick. It is
/// merely a means of synchronising time between server and client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetTime {
    /// `time` is the current time. The time is not limited to 24000 (time of day), but continues
    /// progressing after that.
    pub time: wire::ZigZag32,
}

impl SetTime {
    pub const ID: u32 = 10;
}
/// StartGame is sent by the server to send information about the world the player will be spawned
/// in. It contains information about the position the player spawns in, and information about the
/// world in general such as its game rules.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StartGame {
    pub entity_id: ActorUniqueID,
    pub runtime_id: ActorRuntimeID,
    pub game_type: GameType,
    pub position: glam::Vec3,
    pub rotation: glam::Vec2,
    pub settings: LevelSettings,
    /// `level_id` is a base64 encoded world ID that is used to identify the world.
    pub level_id: String,
    pub level_name: String,
    /// `template_content_identity` is a UUID specific to the premium world template that might have
    /// been used to generate the world. Servers should always fill out an empty string for this.
    pub template_content_identity: String,
    pub is_trial: bool,
    pub movement_settings: SyncedPlayerMovementSettings,
    pub level_current_time: wire::U64LE,
    /// `enchantment_seed` is the seed used to seed the random used to produce enchantments in the
    /// enchantment table. Note that the exact correct random implementation must be used to produce the
    /// correct results both client- and server-side.
    pub enchantment_seed: wire::ZigZag32,
    pub block_properties: Vec<ServerBlockProperty>,
    /// `multiplayer_correlation_id` is a unique ID specifying the multi-player session of the player. A
    /// random UUID should be filled out for this field.
    pub multiplayer_correlation_id: String,
    pub enable_item_stack_net_manager: bool,
    pub server_version: String,
    pub player_property_data: wire::NetworkNbt,
    pub server_block_type_registry_checksum: wire::U64LE,
    /// `world_template_id` is a UUID that identifies the template that was used to generate the world.
    /// Servers that do not use a world based off of a template can set this to an empty UUID.
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
/// AddPlayer is sent by the server to the client to make a player entity show up client-side. It is
/// one of the few entities that cannot be sent using the AddActor packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddPlayer {
    /// `uuid` is the UUID of the player. It is the same UUID that the client sent in the Login packet
    /// at the start of the session. A player with this UUID must exist in the player list (built up
    /// using the PlayerList packet), for it to show up in-game.
    pub uuid: uuid::Uuid,
    pub player_name: String,
    pub target_runtime_id: ActorRuntimeID,
    /// `platform_chat_id` is an identifier only set for particular platforms when chatting (presumably
    /// only for Nintendo Switch). It is otherwise an empty string, and is used to decide which players
    /// are able to chat with each other.
    pub platform_chat_id: String,
    /// `position` is the position to spawn the player on. If the player is on a distance that the
    /// viewer cannot see it, the player will still show up if the viewer moves closer.
    pub position: glam::Vec3,
    /// `velocity` is the initial velocity the player spawns with. This velocity will initiate client
    /// side movement of the player.
    pub velocity: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: wire::F32LE,
    pub carried_item: NetworkItemStackDescriptorSerializedData,
    pub player_game_type: GameType,
    pub entity_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub abilities_data: SerializedAbilitiesData,
    pub actor_links: Vec<EntityLink>,
    /// `device_id` is the device ID set in one of the files found in the storage of the device of the
    /// player. It may be changed freely, so it should not be relied on for anything.
    pub device_id: String,
    /// `build_platform` is the build platform/device OS of the player that is about to be added, as it
    /// sent in the Login packet when joining.
    pub build_platform: BuildPlatform,
}

impl AddPlayer {
    pub const ID: u32 = 12;
}
/// AddActor is sent by the server to the client to spawn an entity to the player. It is used for
/// every entity except other players, for which the AddPlayer packet is used.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddActor {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub actor_type: String,
    /// `position` is the position to spawn the entity on. If the entity is on a distance that the
    /// player cannot see it, the entity will still show up if the player moves closer.
    pub position: glam::Vec3,
    /// `velocity` is the initial velocity the entity spawns with. This velocity will initiate client
    /// side movement of the entity.
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
/// RemoveActor is sent by the server to remove an entity that currently exists in the world from
/// the client- side. Sending this packet if the client cannot already see this entity will have no
/// effect.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RemoveActor {
    pub target_actor_id: ActorUniqueID,
}

impl RemoveActor {
    pub const ID: u32 = 14;
}
/// AddItemActor is sent by the server to the client to make an item entity show up. It is one of
/// the few entities that cannot be sent using the AddActor packet
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddItemActor {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    /// `item` is the item that is spawned. It must have a valid ID for it to show up client-side. If it
    /// is not a valid item, the client will crash when coming near.
    pub item: NetworkItemStackDescriptorSerializedData,
    /// `position` is the position to spawn the entity on. If the entity is on a distance that the
    /// player cannot see it, the entity will still show up if the player moves closer.
    pub position: glam::Vec3,
    /// `velocity` is the initial velocity the entity spawns with. This velocity will initiate client
    /// side movement of the entity.
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
/// TakeItemActor is sent by the server when a player picks up an item entity. It makes the item
/// entity disappear to viewers and shows the pick-up animation.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TakeItemActor {
    pub item_runtime_id: ActorRuntimeID,
    pub actor_runtime_id: ActorRuntimeID,
}

impl TakeItemActor {
    pub const ID: u32 = 17;
}
/// MoveActorAbsolute is sent by the server to move an entity to an absolute position. It is
/// typically used for movements where high accuracy isn't needed, such as for long range
/// teleporting.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MoveActorAbsolute {
    pub move_data: MoveActorAbsoluteData,
}

impl MoveActorAbsolute {
    pub const ID: u32 = 18;
}
/// MovePlayer is sent by players to send their movement to the server, and by the server to update
/// the movement of player entities to other players.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MovePlayer {
    pub player_runtime_id: ActorRuntimeID,
    /// `position` is the position to spawn the player on. If the player is on a distance that the
    /// viewer cannot see it, the player will still show up if the viewer moves closer.
    pub position: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: wire::F32LE,
    pub position_mode: PlayerPositionModeComponentPositionMode,
    /// `on_ground` specifies if the player is considered on the ground. Note that proxies or hacked
    /// clients could fake this to always be true, so it should not be taken for granted.
    pub on_ground: bool,
    pub riding_runtime_id: ActorRuntimeID,
    /// Wire presence: optional value is preceded by a presence marker.
    pub teleport_data: Option<MovePlayerTeleportData>,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
}

impl MovePlayer {
    pub const ID: u32 = 19;
}
/// UpdateBlock is sent by the server to update a block client-side, without resending the entire
/// chunk that the block is located in. It is particularly useful for small modifications like block
/// breaking/placing.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateBlock {
    /// `block_position` is the block position at which a block is updated.
    pub block_position: BlockPos,
    /// `block_runtime_id` is the runtime ID of the block that is placed at Position after sending the
    /// packet to the client.
    pub block_runtime_id: wire::VarUInt,
    /// `flags` is a combination of flags that specify the way the block is updated client-side. It is a
    /// combination of the flags above, but typically sending only the BlockUpdateNetwork flag is
    /// sufficient.
    pub flags: wire::VarUInt,
    /// `layer` is the world layer on which the block is updated. For most blocks, this is the first
    /// layer, as that layer is the default layer to place blocks on, but for blocks inside of each
    /// other, this differs.
    pub layer: wire::VarUInt,
}

impl UpdateBlock {
    pub const ID: u32 = 21;
}
/// AddPainting is sent by the server to the client to make a painting entity show up. It is one of
/// the few entities that cannot be sent using the AddActor packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddPainting {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    /// `position` is the position to spawn the entity on. If the entity is on a distance that the
    /// player cannot see it, the entity will still show up if the player moves closer.
    pub position: glam::Vec3,
    /// `direction` is the facing direction of the painting.
    pub direction: wire::ZigZag32,
    pub motif: String,
}

impl AddPainting {
    pub const ID: u32 = 22;
}
/// LevelEvent is sent by the server to make a certain event in the level occur. It ranges from
/// particles, to sounds, and other events such as starting rain and block breaking.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelEvent {
    /// `event_id` is the ID of the event that is being 'called'. It is one of the events found in the
    /// constants above.
    pub event_id: wire::ZigZag32,
    /// `position` is the position of the level event. Practically every event requires this Vec3 set
    /// for it, as particles, sounds and block editing relies on it.
    pub position: glam::Vec3,
    /// `data` is an integer holding additional data of the event. The type of data held depends on the
    /// EventType.
    pub data: wire::ZigZag32,
}

impl LevelEvent {
    pub const ID: u32 = 25;
}
/// BlockEvent is sent by the server to initiate a certain event that has something to do with
/// blocks in specific, for example opening a chest.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockEvent {
    /// `block_position` is the position of the block that an event occurred at.
    pub block_position: BlockPos,
    /// `event_type` is the type of the block event. The event type decides the way the event data that
    /// follows is used. It is one of the constants found above.
    pub event_type: wire::ZigZag32,
    /// `event_value` holds event type specific data. For chests for example, opening the chest means
    /// the data must hold 1, whereas closing it should hold 0.
    pub event_value: wire::ZigZag32,
}

impl BlockEvent {
    pub const ID: u32 = 26;
}
/// ActorEvent is sent by the server when a particular event happens that has to do with an entity.
/// Some of these events are entity-specific, for example a wolf shaking itself dry, but others are
/// used for each entity, such as dying.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorEvent {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: ActorEventType,
    pub data: wire::ZigZag32,
    /// `fire_at_position` is the position in the same world at which the event should fire. If this is
    /// not present, the position entity will be used instead.
    /// Wire presence: optional value is preceded by a presence marker.
    pub fire_at_position: Option<glam::Vec3>,
}

impl ActorEvent {
    pub const ID: u32 = 27;
}
/// MobEffect is sent by the server to apply an effect to the player, for example an effect like
/// poison. It may also be used to modify existing effects, or removing them completely.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MobEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: MobEffectEvent,
    pub effect_id: wire::ZigZag32,
    pub effect_amplifier: wire::ZigZag32,
    pub show_particles: bool,
    pub effect_duration_ticks: wire::ZigZag32,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
    /// `ambient` specifies if the effect is ambient. If set to false, it will not get treated as an
    /// ambient effect.
    pub ambient: bool,
}

impl MobEffect {
    pub const ID: u32 = 28;
}
/// UpdateAttributes is sent by the server to update an amount of attributes of any entity in the
/// world. These attributes include ones such as the health or the movement speed of the entity.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateAttributes {
    pub target_runtime_id: ActorRuntimeID,
    pub attribute_list: Vec<AttributeData>,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
}

impl UpdateAttributes {
    pub const ID: u32 = 29;
}
/// InventoryTransaction is a packet sent by the client. It essentially exists out of multiple
/// sub-packets, each of which have something to do with the inventory in one way or another. Some
/// of these sub-packets directly relate to the inventory, others relate to interaction with the
/// world, that could potentially result in a change in the inventory.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryTransaction {
    /// `legacy_request_id` is an ID that is only non-zero at times when sent by the client. The server
    /// should always send 0 for this. When this field is not 0, the LegacySetItemSlots slice below will
    /// have values in it. LegacyRequestID ties in with the ItemStackResponse packet. If this field is
    /// non-0, the server should respond with an ItemStackResponse packet. Some inventory actions such
    /// as dropping an item out of the hotbar are still one using this packet, and the ItemStackResponse
    /// packet needs to tie in with it.
    pub legacy_request_id: ItemStackLegacyRequestID,
    /// `legacy_set_item_slots` are only present if the LegacyRequestID is non-zero. These item slots
    /// inform the server of the slots that were changed during the inventory transaction, and the
    /// server should send back an ItemStackResponse packet with these slots present in it. (Or false
    /// with no slots, if rejected.)
    /// Wire presence: optional value is preceded by a presence marker.
    pub legacy_set_item_slots: Option<Vec<LegacySetSlot>>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub transaction: Option<InventoryTransactionValue>,
}

impl InventoryTransaction {
    pub const ID: u32 = 30;
}
/// MobEquipment is sent by the client to the server and the server to the client to make the other
/// side aware of the new item that an entity is holding. It is used to show the item in the hand of
/// entities such as zombies too.
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
/// Interact is sent by the client when it interacts with another entity in some way. It used to be
/// used for normal entity and block interaction, but this is no longer the case now.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Interact {
    pub action: InteractAction,
    pub target_runtime_id: ActorRuntimeID,
    /// `position` associated with the ActionType above. For the InteractActionMouseOverEntity, this is
    /// the position relative to the entity moused over over which the player hovered with its
    /// mouse/touch. For the InteractActionLeaveVehicle, this is the position that the player spawns at
    /// after leaving the vehicle.
    /// Wire presence: optional value is preceded by a presence marker.
    pub position: Option<glam::Vec3>,
}

impl Interact {
    pub const ID: u32 = 33;
}
/// BlockPickRequest is sent by the client when it requests to pick a block in the world and place
/// its item in their inventory.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockPickRequest {
    /// `position` is the position at which the client requested to pick the block. The block at that
    /// position should have its item put in HotBarSlot if it is empty.
    pub position: BlockPos,
    pub with_data: bool,
    pub max_slots: wire::U8,
}

impl BlockPickRequest {
    pub const ID: u32 = 34;
}
/// ActorPickRequest is sent by the client when it tries to pick an entity, so that it gets a spawn
/// egg which can spawn that entity.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ActorPickRequest {
    pub actor_id: wire::I64LE,
    pub max_slots: wire::U8,
    /// `with_data` is true if the pick request requests the entity metadata.
    pub with_data: bool,
}

impl ActorPickRequest {
    pub const ID: u32 = 35;
}
/// PlayerAction is sent by the client when it executes any action, for example starting to sprint,
/// swim, starting the breaking of a block, dropping an item, etc.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerAction {
    pub player_runtime_id: ActorRuntimeID,
    pub action: PlayerActionType,
    /// `block_position` is the position of the target block, if the action with the ActionType set
    /// concerned a block. If that is not the case, the block position will be zero.
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
/// SetActorData is sent by the server to update the entity metadata of an entity. It includes flags
/// such as if the entity is on fire, but also properties such as the air it has left until it
/// starts drowning.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetActorData {
    pub target_runtime_id: ActorRuntimeID,
    pub actor_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
}

impl SetActorData {
    pub const ID: u32 = 39;
}
/// SetActorMotion is sent by the server to change the client-side velocity of an entity. It is
/// usually used in combination with server-side movement calculation.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetActorMotion {
    pub target_runtime_id: ActorRuntimeID,
    pub motion: glam::Vec3,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
}

impl SetActorMotion {
    pub const ID: u32 = 40;
}
/// SetActorLink is sent by the server to initiate an entity link client-side, meaning one entity
/// will start riding another.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetActorLink {
    /// `link` is the link to be set client-side. It links two entities together, so that one entity
    /// rides another. Note that players that see those entities later will not see the link, unless it
    /// is also sent in the AddActor and AddPlayer packets.
    pub link: EntityLink,
}

impl SetActorLink {
    pub const ID: u32 = 41;
}
/// SetHealth is sent by the server. It sets the health of the player it is sent to. The SetHealth
/// packet should no longer be used. Instead, the health attribute should be used so that the health
/// and maximum health may be changed directly.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetHealth {
    /// `health` is the new health of the player.
    pub health: wire::ZigZag32,
}

impl SetHealth {
    pub const ID: u32 = 42;
}
/// SetSpawnPosition is sent by the server to update the spawn position of a player, for example
/// when sleeping in a bed.
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
/// Animate is sent by the server to send a player animation from one player to all viewers of that
/// player. It is used for a couple of actions, such as arm swimming and critical hits.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Animate {
    pub action: AnimateAction,
    pub target_actor_runtime_id: ActorRuntimeID,
    /// `data` ...
    pub data: wire::F32LE,
    /// `swing_source` is the source for swing actions. It is one of the action type constants that may
    /// be found above.
    /// Wire presence: optional value is preceded by a presence marker.
    pub swing_source: Option<String>,
}

impl Animate {
    pub const ID: u32 = 44;
}
/// Respawn is sent by the server to make a player respawn client-side. It is sent in response to a
/// PlayerAction packet with ActionType PlayerActionRespawn. As of 1.13, the server sends two of
/// these packets with different states, and the client sends one of these back in order to complete
/// the respawn.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Respawn {
    /// `position` is the position on which the player should be respawned. The position might be in a
    /// different dimension, in which case the client should first be sent a ChangeDimension packet.
    pub position: glam::Vec3,
    /// `state` is the 'state' of the respawn. It is one of the constants that may be found above, and
    /// the value the packet contains depends on whether the server or client sends it.
    pub state: PlayerRespawnState,
    pub player_runtime_id: ActorRuntimeID,
}

impl Respawn {
    pub const ID: u32 = 45;
}
/// ContainerOpen is sent by the server to open a container client-side. This container must be
/// physically present in the world, for the packet to have any effect. Unlike Java Edition, Bedrock
/// Edition requires that chests for example must be present and in range to open its inventory.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerOpen {
    pub container_id: wire::U8,
    /// `container_type` is the type ID of the container that is being opened when opening the container
    /// at the position of the packet. It depends on the block/entity, and could, for example, be the
    /// window type of a chest or a hopper, but also a horse inventory.
    pub container_type: wire::U8,
    pub position: BlockPos,
    pub target_actor_id: ActorUniqueID,
}

impl ContainerOpen {
    pub const ID: u32 = 46;
}
/// ContainerClose is sent by the server to close a container the player currently has opened, which
/// was opened using the ContainerOpen packet, or by the client to tell the server it closed a
/// particular container, such as the crafting grid.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerClose {
    pub container_id: wire::U8,
    /// `container_type` is the type of container that the server is trying to close. This is used to
    /// validate on the client side whether or not the server's close request is valid.
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
/// InventoryContent is sent by the server to update the full content of a particular inventory. It
/// is usually sent for the main inventory of the player, but also works for other inventories that
/// are currently opened by the player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventoryContent {
    /// `container_id` is the ID that identifies one of the windows that the client currently has
    /// opened, or one of the consistent windows such as the main inventory.
    pub container_id: wire::VarUInt,
    /// `slots` is the new content of the inventory. The length of this slice must be equal to the full
    /// size of the inventory window updated.
    pub slots: Vec<NetworkItemStackDescriptorSerializedData>,
    /// `full_container_name` is the protocol.FullContainerName that describes the container that the
    /// content is for.
    pub full_container_name: FullContainerName,
    /// `storage_item` is the item that is acting as the storage container for the inventory. If the
    /// inventory is not a dynamic container then this field should be left empty. When set, only the
    /// item type is used by the client and none of the other stack info.
    pub storage_item: NetworkItemStackDescriptorSerializedData,
}

impl InventoryContent {
    pub const ID: u32 = 49;
}
/// InventorySlot is sent by the server to update a single slot in one of the inventory windows that
/// the client currently has opened. Usually this is the main inventory, but it may also be the off
/// hand or, for example, a chest inventory.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct InventorySlot {
    /// `container_id` is the ID of the window that the packet modifies. It must point to one of the
    /// windows that the client currently has opened.
    pub container_id: wire::U8,
    /// `slot` is the index of the slot that the packet modifies. The new item will be set to the slot
    /// at this index.
    pub slot: wire::VarUInt,
    /// `full_container_name` is the protocol.FullContainerName that describes the container that the
    /// content is for.
    /// Wire presence: optional value is preceded by a presence marker.
    pub full_container_name: Option<FullContainerName>,
    /// `storage_item` is the item that is acting as the storage container for the inventory. If the
    /// inventory is not a dynamic container then this field should be left empty. When set, only the
    /// item type is used by the client and none of the other stack info.
    /// Wire presence: optional value is preceded by a presence marker.
    pub storage_item: Option<NetworkItemStackDescriptorSerializedData>,
    /// `item` is the item to be put in the slot at Slot. It will overwrite any item that may currently
    /// be present in that slot.
    pub item: NetworkItemStackDescriptorSerializedData,
}

impl InventorySlot {
    pub const ID: u32 = 50;
}
/// ContainerSetData is sent by the server to update specific data of a single container, meaning a
/// block such as a furnace or a brewing stand. This data is usually used by the client to display
/// certain features client-side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerSetData {
    pub container_id: wire::U8,
    pub id: wire::ZigZag32,
    /// `value` is the value of the property. Its use differs per property.
    pub value: wire::ZigZag32,
}

impl ContainerSetData {
    pub const ID: u32 = 51;
}
/// CraftingData is sent by the server to let the client know all crafting data that the server
/// maintains. This includes shapeless crafting, crafting table recipes, furnace recipes etc. Each
/// crafting station's recipes are included in it.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CraftingData {
    /// `shaped_recipes` through SmithingTrimRecipes are the typed recipe vectors used by 1.26.40.
    pub shaped_recipes: Vec<ShapedRecipe>,
    pub shapeless_recipes: Vec<ShapelessRecipe>,
    pub multi_recipes: Vec<MultiRecipe>,
    pub user_data_shapeless_recipes: Vec<ShapelessRecipe>,
    pub shapeless_chemistry_recipes: Vec<ShapelessRecipe>,
    pub shaped_chemistry_recipes: Vec<ShapedRecipe>,
    pub smithing_transform_recipes: Vec<SmithingTransformRecipe>,
    pub smithing_trim_recipes: Vec<SmithingTrimRecipe>,
    /// `potion_mixes` is a list of all potion mixing recipes which may be used in the brewing stand.
    pub potion_mixes: Vec<PotionMixDataEntry>,
    /// `container_mixes` is a list of all recipes to convert a potion from one type to another, such as
    /// from a drinkable potion to a splash potion, or from a splash potion to a lingering potion.
    pub container_mixes: Vec<ContainerMixDataEntry>,
    /// `material_reducers` is a list of all material reducers which is used in education edition
    /// chemistry.
    pub material_reducers: Vec<MaterialReducerDataEntry>,
    /// `clear_recipes` indicates if all recipes currently active on the client should be cleaned. Doing
    /// this means that the client will have no recipes active by itself: Any CraftingData packets
    /// previously sent will also be discarded, and only the recipes in this CraftingData packet will be
    /// used.
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
/// BlockActorData is sent by the server to update data of a block entity client-side, for example
/// the data of a chest.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BlockActorData {
    pub block_position: BlockPos,
    pub actor_data_tags: wire::NetworkNbt,
}

impl BlockActorData {
    pub const ID: u32 = 56;
}
/// LevelChunk is sent by the server to provide the client with a chunk of a world data (16xYx16
/// blocks). Typically, a certain amount of chunks is sent to the client before sending it the spawn
/// PlayStatus packet, so that the client spawns in a loaded world.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelChunk {
    pub chunk_position: ChunkPos,
    pub dimension_id: DimensionType,
    pub sub_chunks_count: wire::VarUInt,
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_request_sub_chunk_limit: Option<wire::ZigZag32>,
    /// `cache_enabled` specifies if the client blob cache should be enabled. This system is based on
    /// hashes of blobs which are consistent and saved by the client in combination with that blob, so
    /// that the server does not have the same chunk multiple times. If the client does not yet have a
    /// blob with the hash sent, it will send a ClientCacheBlobStatus packet containing the hashes is
    /// does not have the data of.
    pub cache_enabled: bool,
    pub cache_metadata: Vec<SubChunkMetadata>,
    pub serialized_chunk_data: bytes::Bytes,
}

impl LevelChunk {
    pub const ID: u32 = 58;
}
/// SetCommandsEnabled is sent by the server to enable or disable the ability to execute commands
/// for the client. If disabled, the client itself will stop the execution of commands.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetCommandsEnabled {
    /// `commands_enabled` defines if the commands should be enabled, or if false, disabled.
    pub commands_enabled: bool,
}

impl SetCommandsEnabled {
    pub const ID: u32 = 59;
}
/// SetDifficulty is sent by the server to update the client-side difficulty of the client. The
/// actual effect of this packet on the client isn't very significant, as the difficulty is handled
/// server-side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetDifficulty {
    /// `difficulty` is the new difficulty that the world has.
    pub difficulty: wire::VarUInt,
}

impl SetDifficulty {
    pub const ID: u32 = 60;
}
/// ChangeDimension is sent by the server to the client to send a dimension change screen
/// client-side. Once the screen is cleared client-side, the client will send a PlayerAction packet
/// with PlayerActionDimensionChangeDone.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChangeDimension {
    pub dimension_id: DimensionType,
    /// `position` is the position in the new dimension that the player is spawned in.
    pub position: glam::Vec3,
    /// `respawn` specifies if the dimension change was respawn based, meaning that the player died in
    /// one dimension and got respawned into another. The client will send a PlayerAction packet with
    /// PlayerActionDimensionChangeRequest if it dies in another dimension, indicating that it needs a
    /// DimensionChange packet with Respawn set to true.
    pub respawn: bool,
    /// `loading_screen_id` is a unique ID for the loading screen that the player is currently in. The
    /// client will update the server on its state through the ServerBoundLoadingScreen packet, and it
    /// can be used to not send specific packets to the client if it is changing dimensions. This field
    /// should be unique for every ChangeDimension packet sent.
    /// Wire presence: optional value is preceded by a presence marker.
    pub loading_screen_id: Option<wire::U32LE>,
}

impl ChangeDimension {
    pub const ID: u32 = 61;
}
/// SetPlayerGameType is sent by the server to update the game type, which is otherwise known as the
/// game mode, of a player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetPlayerGameType {
    pub player_game_type: GameType,
}

impl SetPlayerGameType {
    pub const ID: u32 = 62;
}
/// PlayerList is sent by the server to update the client-side player list in the in-game menu
/// screen. It shows the icon of each player if the correct XUID is written in the packet. Sending
/// the PlayerList packet is obligatory when sending an AddPlayer packet. The added player will not
/// show up to a client if it has not been added to the player list, because several properties of
/// the player are obtained from the player list, such as the skin.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerList {
    /// `entries` is a list of all player list entries that should be added/removed from the player
    /// list, depending on the ActionType set.
    pub entries: Vec<PlayerListData>,
}

impl PlayerList {
    pub const ID: u32 = 63;
}
/// SimpleEvent is used for enabling or disabling commands and for unlocking world template settings
/// (both unlocking UI buttons on client and the actual setting on the server). This is fired from
/// the client to the server and a SetCommandsEnabled is sent back when enabling commands.
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
    pub event_data: EventData,
}

impl LegacyTelemetryEvent {
    pub const ID: u32 = 65;
}
/// SpawnExperienceOrb is sent by the server to spawn an experience orb entity client-side. Much
/// like the AddPainting packet, it is one of the few packets that spawn an entity without using the
/// AddActor packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SpawnExperienceOrb {
    /// `position` is the position to spawn the experience orb on. If the entity is on a distance that
    /// the player cannot see it, the entity will still show up if the player moves closer.
    pub position: glam::Vec3,
    /// `xp_value` is the amount of experience in experience points that the orb carries. The
    /// client-side size of the orb depends on the amount of experience in the orb: There are 11
    /// possible sizes for the orb, for 1–2, 3–6, 7–16, 17–36, 37–72, 73–148, 149–306,
    /// 307–616, 617–1236, 1237–2476, and 2477 and up.
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
/// MapInfoRequest is sent by the client to request the server to deliver information of a certain
/// map in the inventory of the player. The server should respond with a ClientBoundMapItemData
/// packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapInfoRequest {
    pub map_unique_id: ActorUniqueID,
    pub client_pixels_list: Vec<PixelRequest>,
}

impl MapInfoRequest {
    pub const ID: u32 = 68;
}
/// RequestChunkRadius is sent by the client to the server to update the server on the chunk view
/// radius that it has set in the settings. The server may respond with a ChunkRadiusUpdated packet
/// with either the chunk radius requested, or a different chunk radius if the server chooses so.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestChunkRadius {
    /// `chunk_radius` is the requested chunk radius. This value is always the value set in the settings
    /// of the player.
    pub chunk_radius: wire::ZigZag32,
    /// `max_chunk_radius` is the maximum chunk radius that the player wants to receive. The reason for
    /// the client sending this is currently unknown.
    pub max_chunk_radius: wire::U8,
}

impl RequestChunkRadius {
    pub const ID: u32 = 69;
}
/// ChunkRadiusUpdated is sent by the server in response to a RequestChunkRadius packet. It defines
/// the chunk radius that the server allows the client to have. This may be lower than the chunk
/// radius requested by the client in the RequestChunkRadius packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ChunkRadiusUpdated {
    /// `chunk_radius` is the final chunk radius that the client will adapt when it receives the packet.
    /// It does not have to be the same as the requested chunk radius.
    pub chunk_radius: wire::ZigZag32,
}

impl ChunkRadiusUpdated {
    pub const ID: u32 = 70;
}
/// GameRulesChanged is sent by the server to the client to update client-side game rules, such as
/// game rules like the 'showCoordinates' game rule.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameRulesChanged {
    pub rule_data: GameRulesChangedData,
}

impl GameRulesChanged {
    pub const ID: u32 = 72;
}
/// Camera is sent by the server to use an Education Edition camera on a player. It produces an
/// image client-side.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Camera {
    pub camera_id: ActorUniqueID,
    pub target_player_id: ActorUniqueID,
}

impl Camera {
    pub const ID: u32 = 73;
}
/// BossEvent is sent by the server to make a specific 'boss event' occur in the world. It includes
/// features such as showing a boss bar to the player and turning the sky dark.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BossEvent {
    pub target_actor_id: ActorUniqueID,
    pub player_id: ActorUniqueID,
    /// `event_type` is the type of the event. It is one of the BossEvent constants above.
    pub event_type: BossEventUpdateType,
    pub name: String,
    pub filtered_name: String,
    pub health_percent: wire::F32LE,
    pub color: BossBarColor,
    /// `overlay` is the overlay of the boss bar that is shown on top of the boss bar when a player is
    /// subscribed. It is one of the BossEventOverlay constants listed above.
    pub overlay: BossBarOverlay,
}

impl BossEvent {
    pub const ID: u32 = 74;
}
/// ShowCredits is sent by the server to show the Minecraft credits screen to the client. It is
/// typically sent when the player beats the ender dragon and leaves the End.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShowCredits {
    /// `player_runtime_id` is the entity runtime ID of the player to show the credits to. It's not
    /// clear why this field is actually here in the first place.
    pub player_runtime_id: ActorRuntimeID,
    pub credits_state: wire::ZigZag32,
}

impl ShowCredits {
    pub const ID: u32 = 75;
}
/// AvailableCommands is sent by the server to send a list of all commands that the player is able
/// to use on the server. This packet holds all the arguments of each commands as well, making it
/// possible for the client to provide auto-completion and command usages. AvailableCommands packets
/// can be resent, but the packet is often very big, so doing this very often should be avoided.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableCommands {
    /// `enum_values` is a slice of all enum values of any enum in the AvailableCommands packet.
    /// EnumValues generally should contain each possible value only once. Enums are built by pointing
    /// to entries in this slice.
    pub enum_values: Vec<String>,
    /// `chained_subcommand_values` is a slice of all chained subcommand names. ChainedSubcommandValues
    /// generally should contain each possible value only once. ChainedSubcommands are built by pointing
    /// to entries in this slice.
    pub chained_subcommand_values: Vec<String>,
    /// `post_fixes`, like EnumValues, is a slice of all suffix values of any command parameter in the
    /// AvailableCommands packet.
    pub post_fixes: Vec<String>,
    /// `enum_data` is a slice of all (fixed) command enums present in any of the commands.
    pub enum_data: Vec<CommandEnum>,
    /// `chained_subcommand_data` is a slice of all subcommands that are followed by a chained command.
    /// An example usage of this is /execute which allows you to run another command as another entity
    /// or at a different position etc.
    pub chained_subcommand_data: Vec<ChainedSubcommand>,
    /// `commands` is a list of all commands that the client should show client-side. The
    /// AvailableCommands packet replaces any commands sent before. It does not only add the commands
    /// that are sent in it.
    pub commands: Vec<Command>,
    /// `soft_enums` is a slice of dynamic command enums. These command enums can be changed during
    /// runtime without having to resend an AvailableCommands packet.
    pub soft_enums: Vec<DynamicEnum>,
    /// `constraints` is a list of constraints that should be applied to certain options of enums in the
    /// commands above.
    pub constraints: Vec<CommandEnumConstraint>,
}

impl AvailableCommands {
    pub const ID: u32 = 76;
}
/// CommandRequest is sent by the client to request the execution of a server-side command. Although
/// some servers support sending commands using the Text packet, this packet is guaranteed to have
/// the correct result.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandRequest {
    /// `command` is the raw entered command line. The client does no parsing of the command line by
    /// itself (unlike it did in the early stages), but lets the server do that.
    pub command: String,
    /// `origin` is the data specifying the origin of the command. In other words, the source that the
    /// command was from, such as the player itself or a websocket server.
    pub origin: CommandOriginData,
    /// `is_internal` specifies if the command request internal. Setting it to false seems to work and
    /// the usage of this field is not known.
    pub is_internal: bool,
    /// `version` is the version of the command that is being executed. This field currently has no
    /// purpose or functionality.
    pub version: String,
}

impl CommandRequest {
    pub const ID: u32 = 77;
}
/// CommandBlockUpdate is sent by the client to update a command block at a specific position. The
/// command block may be either a physical block or an entity.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandBlockUpdate {
    pub target: CommandBlockUpdateData,
    /// `command` is the command currently entered in the command block. This is the command that is
    /// executed when the command block is activated.
    pub command: String,
    /// `last_output` is the output of the last command executed by the command block. It may be left
    /// empty to show simply no output at all, in combination with setting ShouldTrackOutput to false.
    pub last_output: String,
    /// `name` is the name of the command block updated. If not empty, it will show this name hovering
    /// above the command block when hovering over the block with the cursor.
    pub name: String,
    /// `filtered_name` is a filtered version of Name with all the profanity removed. The client will
    /// use this over Name if this field is not empty and they have the "Filter Profanity" setting
    /// enabled.
    pub filtered_name: String,
    pub track_output: bool,
    /// `tick_delay` is the delay in ticks between executions of a command block, if it is a repeating
    /// command block.
    pub tick_delay: wire::I32LE,
    /// `execute_on_first_tick` specifies if the command block should execute on the first tick, AKA as
    /// soon as the command block is enabled.
    pub execute_on_first_tick: bool,
}

impl CommandBlockUpdate {
    pub const ID: u32 = 78;
}
/// CommandOutput is sent by the server to the client to send text as output of a command. Most
/// servers do not use this packet and instead simply send Text packets, but there is reason to send
/// it. If the origin of a CommandRequest packet is not the player itself, but, for example, a
/// websocket server, sending a Text packet will not do what is expected: The message should go to
/// the websocket server, not to the client's chat. The CommandOutput packet will make sure the
/// messages are relayed to the correct origin of the command request.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CommandOutput {
    pub origin_data: CommandOriginData,
    pub output: CommandOutputData,
}

impl CommandOutput {
    pub const ID: u32 = 79;
}
/// UpdateTrade is sent by the server to update the trades offered by a villager to a player. It is
/// sent at the moment that a player interacts with a villager.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateTrade {
    pub container_id: wire::U8,
    pub type_: wire::U8,
    /// `size` is the amount of trading options that the villager has.
    pub size: wire::ZigZag32,
    pub trader_tier: wire::ZigZag32,
    /// `entity_unique_id` is the unique ID of the entity (usually a player) for which the trades are
    /// updated. The updated trades may apply only to this entity.
    pub entity_unique_id: ActorUniqueID,
    pub last_trading_player: ActorUniqueID,
    /// `display_name` is the name displayed at the top of the trading UI. It is usually used to
    /// represent the profession of the villager in the UI.
    pub display_name: String,
    pub use_new_trade_screen: bool,
    pub using_economy_trade: bool,
    pub data: wire::NetworkNbt,
}

impl UpdateTrade {
    pub const ID: u32 = 80;
}
/// UpdateEquip is sent by the server to the client upon opening a horse inventory. It is used to
/// set the content of the inventory and specify additional properties, such as the items that are
/// allowed to be put in slots of the inventory.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateEquip {
    pub container_id: wire::U8,
    pub type_: wire::U8,
    /// `size` is the size of the horse inventory that should be opened. A bigger size does, in fact,
    /// change the amount of slots displayed.
    pub size: wire::ZigZag32,
    /// `entity_unique_id` is the unique ID of the entity whose equipment was 'updated' to the player.
    /// It is typically the horse entity that had its inventory opened.
    pub entity_unique_id: ActorUniqueID,
    pub data: wire::NetworkNbt,
}

impl UpdateEquip {
    pub const ID: u32 = 81;
}
/// ResourcePackDataInfo is sent by the server to the client to inform the client about the data
/// contained in one of the resource packs that are about to be sent.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackDataInfo {
    pub resource_name: String,
    pub chunk_size: wire::U32LE,
    pub number_of_chunks: wire::U32LE,
    pub file_size: wire::U64LE,
    pub file_hash: String,
    pub is_premium_pack: bool,
    /// `pack_type` is the type of the resource pack. It is one of the resource pack types that may be
    /// found in the constants above.
    pub pack_type: wire::U8,
}

impl ResourcePackDataInfo {
    pub const ID: u32 = 82;
}
/// ResourcePackChunkData is sent to the client so that the client can download the resource pack.
/// Each packet holds a chunk of the compressed resource pack, of which the size is defined in the
/// ResourcePackDataInfo packet sent before.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackChunkData {
    /// `resource_name` is the unique ID of the resource pack that the chunk of data is taken out of.
    pub resource_name: String,
    /// `chunk_id` is the current chunk index of the chunk. It is a number that starts at 0 and is
    /// incremented for each resource pack data chunk sent to the client.
    pub chunk_id: wire::U32LE,
    /// `byte_offset` is the current progress in bytes or offset in the data that the resource pack data
    /// chunk is taken from.
    pub byte_offset: wire::U64LE,
    /// RawPayload is a byte slice containing a chunk of data from the resource pack. It must be of the
    /// same size or less than the DataChunkSize set in the ResourcePackDataInfo packet.
    pub chunk_data: bytes::Bytes,
}

impl ResourcePackChunkData {
    pub const ID: u32 = 83;
}
/// ResourcePackChunkRequest is sent by the client to request a chunk of data from a particular
/// resource pack, that it has obtained information about in a ResourcePackDataInfo packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePackChunkRequest {
    /// `resource_name` is the unique ID of the resource pack that the chunk of data is requested from.
    pub resource_name: String,
    /// `chunk` is the requested chunk index of the chunk. It is a number that starts at 0 and is
    /// incremented for each resource pack data chunk requested.
    pub chunk: wire::I32LE,
}

impl ResourcePackChunkRequest {
    pub const ID: u32 = 84;
}
/// Transfer is sent by the server to transfer a player from the current server to another. Doing so
/// will fully disconnect the client, bring it back to the main menu and make it connect to the next
/// server.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Transfer {
    /// `server_address` is the address of the new server, which might be either a hostname or an actual
    /// IP address.
    pub server_address: String,
    /// `server_port` is the UDP port of the new server.
    pub server_port: wire::U16LE,
    /// `reload_world` currently has an unknown usage.
    pub reload_world: bool,
    /// `gatherings_configuration` optionally identifies the gathering being joined on the target
    /// server.
    /// Wire presence: optional value is preceded by a presence marker.
    pub gatherings_configuration: Option<ServerConfigurationGatheringsConfigurationJoinInfo>,
}

impl Transfer {
    pub const ID: u32 = 85;
}
/// PlaySound is sent by the server to play a sound to the client. Some of the sounds may only be
/// started using this packet and must be stopped using the StopSound packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlaySound {
    /// `name` is the name of the sound to play.
    pub name: String,
    /// `position` is the position at which the sound was played. Some sounds do not depend on a
    /// position, which will then ignore it, but most of them will play with the direction based on the
    /// position compared to the player's position.
    pub position: BlockPos,
    /// `volume` is the relative volume of the sound to play. It will be less loud for the player if it
    /// is farther away from the position of the sound.
    pub volume: wire::F32LE,
    /// `pitch` is the pitch of the sound to play. Some sounds completely ignore this field, whereas
    /// others use it to specify the pitch as the field is intended.
    pub pitch: wire::F32LE,
    /// `loop_count` is the number of times to loop the sound before stopping. -1 means no looping at
    /// all.
    pub loop_count: wire::ZigZag32,
    /// `server_sound_handle` is an optional sound handle ID. It is currently unknown what this is for,
    /// and is not required to be set by servers.
    /// Wire presence: optional value is preceded by a presence marker.
    pub server_sound_handle: Option<ServerSoundHandle>,
}

impl PlaySound {
    pub const ID: u32 = 86;
}
/// StopSound is sent by the server to stop a sound playing to the player, such as a playing music
/// disk track or other long-lasting sounds.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StopSound {
    /// `sound_name` is the name of the sound that should be stopped from playing. If no sound with this
    /// name is currently active, the packet is ignored.
    pub sound_name: String,
    /// `stop_all_sounds` specifies if all sounds currently playing to the player should be stopped. If
    /// set to true, the SoundName field may be left empty.
    pub stop_all_sounds: bool,
    /// `stop_music_legacy` is currently unknown.
    pub stop_music_legacy: bool,
}

impl StopSound {
    pub const ID: u32 = 87;
}
/// SetTitle is sent by the server to make a title, subtitle or action bar shown to a player. It has
/// several fields that allow setting the duration of the titles.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetTitle {
    pub title_type: TitleType,
    pub title_text: String,
    pub fade_in_time: wire::ZigZag32,
    pub stay_time: wire::ZigZag32,
    pub fade_out_time: wire::ZigZag32,
    /// `xuid` is the XBOX Live user ID of the player, which will remain consistent as long as the
    /// player is logged in with the XBOX Live account. It is empty if the user is not logged into its
    /// XBL account.
    pub xuid: String,
    /// `platform_online_id` is either a uint64 or an empty string.
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
/// StructureBlockUpdate is sent by the client when it updates a structure block using the in-game
/// UI. The data it contains depends on the type of structure block that it is. In Minecraft Bedrock
/// Edition v1.11, there is only the Export structure block type, but in v1.13 the ones present in
/// Java Edition will, according to the wiki, be added too.
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
/// ShowStoreOffer is sent by the server to show a Marketplace store offer to a player. It opens a
/// window client-side that displays the item. The ShowStoreOffer packet only works on the partnered
/// servers: Servers that are not partnered will not have a store buttons show up in the in-game
/// pause menu and will, as a result, not be able to open store offers on the client side. Sending
/// the packet does therefore not work when using a proxy that is not connected to with the domain
/// of one of the partnered servers.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShowStoreOffer {
    /// `offer_id` is a UUID that identifies the offer for which a window should be opened.
    pub offer_id: uuid::Uuid,
    pub redirect_type: ShowStoreOfferRedirectType,
}

impl ShowStoreOffer {
    pub const ID: u32 = 91;
}
/// PurchaseReceipt is sent by the client to the server to notify the server it purchased an item
/// from the Marketplace store that was offered by the server. The packet is only used for partnered
/// servers.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PurchaseReceipt {
    /// `purchase_receipts` is a list of receipts, or proofs of purchases, for the offers that have been
    /// purchased by the player.
    pub purchase_receipts: Vec<String>,
}

impl PurchaseReceipt {
    pub const ID: u32 = 92;
}
/// PlayerSkin is sent by the client to the server when it updates its own skin using the in-game
/// skin picker. It is relayed by the server, or sent if the server changes the skin of a player on
/// its own accord. Note that the packet can only be sent for players that are in the player list at
/// the time of sending.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerSkin {
    /// `uuid` is the UUID of the player as sent in the Login packet when the client joined the server.
    /// It must match this UUID exactly for the skin to show up on the player.
    pub uuid: uuid::Uuid,
    pub serialized_skin: SerializedSkinRef,
    pub localized_new_skin_name: String,
    pub localized_old_skin_name: String,
}

impl PlayerSkin {
    pub const ID: u32 = 93;
}
/// SubClientLogin is sent when a sub-client joins the server while another client is already
/// connected to it. The packet is sent as a result of split-screen game play, and allows up to four
/// players to play using the same network connection. After an initial Login packet from the 'main'
/// client, each sub-client that connects sends a SubClientLogin to request their own login.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubClientLogin {
    pub sub_client_connection_request: String,
}

impl SubClientLogin {
    pub const ID: u32 = 94;
}
/// AutomationClientConnect is used to make the client connect to a websocket server. This websocket
/// server has the ability to execute commands on the behalf of the client and it can listen for
/// certain events fired by the client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AutomationClientConnect {
    pub web_socket_data: WebSocketData,
}

impl AutomationClientConnect {
    pub const ID: u32 = 95;
}
/// SetLastHurtBy is sent by the server to let the client know what entity type it was last hurt by.
/// At this moment, the packet is useless and should not be used. There is no behaviour that depends
/// on if this packet is sent or not.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetLastHurtBy {
    pub last_hurt_by: ActorType,
}

impl SetLastHurtBy {
    pub const ID: u32 = 96;
}
/// BookEdit is sent by the client when it edits a book. It is sent each time a modification was
/// made and the player stops its typing 'session', rather than simply after closing the book.
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
/// PhotoTransfer is sent by the server to transfer a photo (image) file to the client. It is
/// typically used to transfer photos so that the client can display it in a portfolio in Education
/// Edition. While previously usable in the default Bedrock Edition, the displaying of photos in
/// books was disabled and the packet now has little use anymore.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PhotoTransfer {
    /// `photo_name` is the name of the photo to transfer. It is the exact file name that the client
    /// will download the photo as, including the extension of the file.
    pub photo_name: String,
    /// `photo_data` is the raw data of the photo image. The format of this data may vary: Formats such
    /// as JPEG or PNG work, as long as PhotoName has the correct extension.
    pub photo_data: bytes::Bytes,
    /// `book_id` is the ID of the book that the photo is associated with. If the PhotoName in a book
    /// with this ID is set to PhotoName, it will display the photo (provided Education Edition is
    /// used). The photo image is downloaded to a sub-folder with this book ID.
    pub book_id: String,
    /// `type_` is one of the three photo types above.
    pub type_: PhotoType,
    /// `source_type` is the source photo type. It is one of the three photo types above.
    pub source_type: PhotoType,
    /// `owner_id` is the entity unique ID of the photo's owner.
    pub owner_id: wire::I64LE,
    /// `new_photo_name` is the new name of the photo.
    pub new_photo_name: String,
}

impl PhotoTransfer {
    pub const ID: u32 = 99;
}
/// ModalFormRequest is sent by the server to make the client open a form. This form may be either a
/// modal form which has two options, a menu form for a selection of options and a custom form for
/// properties.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ModalFormRequest {
    /// `form_id` is an ID used to identify the form. The ID is saved by the client and sent back when
    /// the player submits the form, so that the server can identify which form was submitted.
    pub form_id: wire::VarUInt,
    pub form_ui_json: String,
}

impl ModalFormRequest {
    pub const ID: u32 = 100;
}
/// ModalFormResponse is sent by the client in response to a ModalFormRequest, after the player has
/// submitted the form sent. It contains the options/properties selected by the player, or a JSON
/// encoded 'null' if the form was closed by clicking the X at the top right corner of the form.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ModalFormResponse {
    /// `form_id` is the form ID of the form the client has responded to. It is the same as the ID sent
    /// in the ModalFormRequest, and may be used to identify which form was submitted.
    pub form_id: wire::VarUInt,
    /// `json_response` is a JSON encoded value representing the response of the player. For a modal
    /// form, the response is either true or false, for a menu form, the response is an integer
    /// specifying the index of the button clicked, and for a custom form, the response is an array
    /// containing a value for each element.
    /// Wire presence: optional value is preceded by a presence marker.
    pub json_response: Option<String>,
    /// `form_cancel_reason` represents the reason why the form was cancelled. It is one of the
    /// constants above.
    /// Wire presence: optional value is preceded by a presence marker.
    pub form_cancel_reason: Option<ModalFormCancelReason>,
}

impl ModalFormResponse {
    pub const ID: u32 = 101;
}
/// ServerSettingsRequest is sent by the client to request the settings specific to the server.
/// These settings are shown in a separate tab client-side, and have the same structure as a custom
/// form.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerSettingsRequest {
}

impl ServerSettingsRequest {
    pub const ID: u32 = 102;
}
/// ServerSettingsResponse is optionally sent by the server in response to a ServerSettingsRequest
/// from the client. It is structured the same as a ModalFormRequest packet, and if filled out
/// correctly, will show a specific tab for the server in the settings of the client. A
/// ModalFormResponse packet is sent by the client in response to a ServerSettingsResponse, when the
/// client fills out the settings and closes the settings again.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerSettingsResponse {
    /// `form_id` is an ID used to identify the form. The ID is saved by the client and sent back when
    /// the player submits the form, so that the server can identify which form was submitted.
    pub form_id: wire::VarUInt,
    pub form_ui_json: String,
}

impl ServerSettingsResponse {
    pub const ID: u32 = 103;
}
/// ShowProfile is sent by the server to show the XBOX Live profile of one player to another.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ShowProfile {
    /// `player_xuid` is the XBOX Live User ID of the player whose profile should be shown to the
    /// player. If it is not a valid XUID, the client ignores the packet.
    pub player_xuid: String,
}

impl ShowProfile {
    pub const ID: u32 = 104;
}
/// SetDefaultGameType is sent by the client when it toggles the default game type in the settings
/// UI, and is sent by the server when it actually changes the default game type, resulting in the
/// toggle being changed in the settings UI.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetDefaultGameType {
    pub default_game_type: GameType,
}

impl SetDefaultGameType {
    pub const ID: u32 = 105;
}
/// RemoveObjective is sent by the server to remove a scoreboard objective. It is used to stop
/// showing a scoreboard to a player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RemoveObjective {
    /// `objective_name` is the name of the objective that the scoreboard currently active has. This
    /// name must be identical to the one sent in the SetDisplayObjective packet.
    pub objective_name: String,
}

impl RemoveObjective {
    pub const ID: u32 = 106;
}
/// SetDisplayObjective is sent by the server to display an object as a scoreboard to the player.
/// Once sent, it should be followed up by a SetScore packet to set the lines of the packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetDisplayObjective {
    /// `display_slot_name` is the slot in which the scoreboard should be displayed. Available options
    /// can be found in the constants above.
    pub display_slot_name: String,
    /// `objective_name` is the name of the objective that the scoreboard displays. Filling out a random
    /// unique value for this field works: It is not displayed in the scoreboard.
    pub objective_name: String,
    /// `objective_display_name` is the name, or title, that is displayed at the top of the scoreboard.
    pub objective_display_name: String,
    /// `criteria_name` is the name of the criteria that need to be fulfilled in order for the score to
    /// be increased. This can be any kind of string and does not show up client-side.
    pub criteria_name: String,
    /// `sort_order` is the order in which entries on the scoreboard should be sorted. It is one of the
    /// constants that may be found above.
    pub sort_order: wire::ZigZag32,
}

impl SetDisplayObjective {
    pub const ID: u32 = 107;
}
/// SetScore is sent by the server to send the contents of a scoreboard to the player. It may be
/// used to either add, remove or edit entries on the scoreboard.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetScore {
    /// `score_info` is a list of all entries that the client should operate on. Each entry's
    /// IdentityType specifies whether it is added, modified or removed.
    pub score_info: Vec<SetScoreInfoItem>,
}

impl SetScore {
    pub const ID: u32 = 108;
}
/// LabTable is sent by the client to let the server know it started a chemical reaction in
/// Education Edition, and is sent by the server to other clients to show the effects. The packet is
/// only functional if Education features are enabled.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LabTable {
    /// `type_` is the type of the action that was executed. It is one of the constants above.
    /// Typically, only LabTableActionCombine is sent by the client, whereas LabTableActionReact is sent
    /// by the server.
    pub type_: LabTableType,
    /// `position` is the position at which the lab table used was located.
    pub position: BlockPos,
    /// `reaction` is the type of the reaction that took place as a result of the items put into the lab
    /// table. The reaction type can be either that of an item or a particle, depending on whatever the
    /// result was of the reaction.
    pub reaction: LabTableReactionType,
}

impl LabTable {
    pub const ID: u32 = 109;
}
/// UpdateBlockSynced is sent by the server to synchronise the falling of a falling block entity
/// with the transitioning back and forth from and to a solid block. It is used to prevent the
/// entity from flickering, and is used in places such as the pushing of blocks with pistons.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateBlockSynced {
    /// `block_position` is the block position at which a block is updated.
    pub block_position: BlockPos,
    /// `block_runtime_id` is the runtime ID of the block that is placed at Position after sending the
    /// packet to the client.
    pub block_runtime_id: wire::VarUInt,
    /// `flags` is a combination of flags that specify the way the block is updated client-side. It is a
    /// combination of the flags above, but typically sending only the BlockUpdateNetwork flag is
    /// sufficient.
    pub flags: wire::VarUInt,
    /// `layer` is the world layer on which the block is updated. For most blocks, this is the first
    /// layer, as that layer is the default layer to place blocks on, but for blocks inside of each
    /// other, this differs.
    pub layer: wire::VarUInt,
    /// `unique_actor_id` is the unique ID of the falling block entity that the block transitions to or
    /// that the entity transitions from. Note that for both possible values for TransitionType, the
    /// EntityUniqueID should point to the falling block entity involved.
    pub unique_actor_id: wire::VarULong,
    /// `actor_sync_message` is the type of the transition that happened. It is either
    /// BlockToEntityTransition, when a block placed becomes a falling entity, or
    /// EntityToBlockTransition, when a falling entity hits the ground and becomes a solid block again.
    pub actor_sync_message: wire::VarULong,
}

impl UpdateBlockSynced {
    pub const ID: u32 = 110;
}
/// MoveActorDelta is sent by the server to move an entity. The packet is specifically optimised to
/// save as much space as possible, by only writing non-zero fields. As of 1.16.100, this packet no
/// longer actually contains any deltas.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MoveActorDelta {
    pub move_data: MoveActorDeltaData,
}

impl MoveActorDelta {
    pub const ID: u32 = 111;
}
/// SetScoreboardIdentity is sent by the server to change the identity type of one of the entries on
/// a scoreboard. This is used to change, for example, an entry pointing to a player, to a fake
/// player when it leaves the server, and to change it back to a real player when it joins again. In
/// non-vanilla situations, the packet is quite useless.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetScoreboardIdentity {
    /// `scoreboard_identity_packet_type` is the type of the action to execute. The action is either
    /// ScoreboardIdentityActionRegister to associate an identity with the entry, or
    /// ScoreboardIdentityActionClear to remove associations with an entity.
    pub scoreboard_identity_packet_type: ScoreboardIdentityPacketType,
    /// `scoreboard_identity_info` is a list of all entries in the packet. Each of these entries points
    /// to one of the entries on a scoreboard. Depending on ActionType, their identity will either be
    /// registered or cleared.
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
/// UpdateSoftEnum is sent by the server to update a soft enum, also known as a dynamic enum,
/// previously sent in the AvailableCommands packet. It is sent whenever the enum should get new
/// options or when some of its options should be removed. The UpdateSoftEnum packet will apply for
/// enums that have been set in the AvailableCommands packet with the 'Dynamic' field of the
/// CommandEnum set to true.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateSoftEnum {
    /// `enum_name` is the type of the enum. This type must be identical to the one set in the
    /// AvailableCommands packet, because the client uses this to recognise which enum to update.
    pub enum_name: String,
    /// `values` is a list of options that should be updated. Depending on the ActionType field, either
    /// these options will be added to the enum, the enum options will be set to these options or all of
    /// these options will be removed from the enum.
    pub values: Vec<String>,
    /// `update_type` is the type of the action to execute on the enum. The Options field has a
    /// different result, depending on what ActionType is used.
    pub update_type: SoftEnumUpdateType,
}

impl UpdateSoftEnum {
    pub const ID: u32 = 114;
}
/// NetworkStackLatency is sent by the server (and the client, on development builds) to measure the
/// latency over the entire Minecraft stack, rather than the RakNet latency. It has other usages
/// too, such as the ability to be used as some kind of acknowledgement packet, to know when the
/// client has received a certain other packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkStackLatency {
    pub creation_time: wire::U64LE,
    pub is_from_server: bool,
}

impl NetworkStackLatency {
    pub const ID: u32 = 115;
}
/// SpawnParticleEffect is sent by the server to spawn a particle effect client-side. Unlike other
/// packets that result in the appearing of particles, this packet can show particles that are not
/// hardcoded in the client. They can be added and changed through behaviour packs to implement
/// custom particles.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SpawnParticleEffect {
    pub dimension_id: wire::U8,
    pub actor_id: ActorUniqueID,
    /// `position` is the position that the particle should be spawned at. If the position is too far
    /// away from the player, it will not show up. If EntityUniqueID is not -1, the position will be
    /// relative to the position of the entity.
    pub position: glam::Vec3,
    pub effect_name: String,
    /// `molang_variables` is an encoded JSON map of MoLang variables that may be applicable to the
    /// particle spawn. This can just be left empty in most cases.
    /// Wire presence: optional value is preceded by a presence marker.
    pub molang_variables: Option<String>,
}

impl SpawnParticleEffect {
    pub const ID: u32 = 118;
}
/// AvailableActorIdentifiers is sent by the server at the start of the game to let the client know
/// all entities that are available on the server.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AvailableActorIdentifiers {
    pub identifier_list: wire::NetworkNbt,
}

impl AvailableActorIdentifiers {
    pub const ID: u32 = 119;
}
/// NetworkChunkPublisherUpdate is sent by the server to change the point around which chunks are
/// and remain loaded. This is useful for mini-game servers, where only one area is ever loaded, in
/// which case the NetworkChunkPublisherUpdate packet can be sent in the middle of it, so that no
/// chunks ever need to be additionally sent during the course of the game. In reality, the packet
/// is not extraordinarily useful, and most servers just send it constantly at the position of the
/// player. If the packet is not sent at all, no chunks will be shown to the player, regardless of
/// where they are sent.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkChunkPublisherUpdate {
    /// `new_position_for_view` is the block position around which chunks loaded will remain shown to
    /// the client. Most servers set this position to the position of the player itself.
    pub new_position_for_view: BlockPos,
    /// `new_radius_for_view` is the radius in blocks around Position that chunks sent show up in and
    /// will remain loaded in. Unlike the RequestChunkRadius and ChunkRadiusUpdated packets, this radius
    /// is in blocks rather than chunks, so the chunk radius needs to be multiplied by 16. (Or shifted
    /// to the left by 4.)
    pub new_radius_for_view: wire::VarUInt,
    /// `server_built_chunks_list` ... TODO: Figure out what this field is used for.
    pub server_built_chunks_list: Vec<ChunkPos>,
}

impl NetworkChunkPublisherUpdate {
    pub const ID: u32 = 121;
}
/// BiomeDefinitionList is sent by the server to let the client know all biomes that are available
/// and implemented on the server side. When enabled, it also includes information for the client to
/// accurately recreate the server-side generation in vanilla worlds/servers for increased
/// performance.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct BiomeDefinitionList {
    pub map_of_biome_names_to_data: Vec<(wire::U16LE, BiomeDefinitionData)>,
    /// `string_list` is a makeshift dictionary implementation Mojang created to try and reduce the size
    /// of the overall packet. It is a list of common strings that are used in the biome definitions,
    /// such as biome names, float values or query expressions.
    pub string_list: BiomeStringList,
}

impl BiomeDefinitionList {
    pub const ID: u32 = 122;
}
/// LevelSoundEvent is sent by the server to make any kind of built-in sound heard to a player. It
/// is sent to, for example, play a stepping sound or a shear sound. The packet is also sent by the
/// client, in which case it could be forwarded by the server to the other players online. If
/// possible, the packets from the client should be ignored however, and the server should play them
/// on its own accord.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelSoundEvent {
    /// `sound_event` is the type of the sound to play. It is one of the constants above. Some of the
    /// sound types require additional data, which is set in the ExtraData field.
    pub sound_event: String,
    /// `position` is the position of the sound event. The player will be able to hear the direction of
    /// the sound based on what position is sent here.
    pub position: glam::Vec3,
    /// `data` is a packed integer that some sound types use to provide extra data. An example of this
    /// is the note sound, which is composed of a pitch and an instrument type.
    pub data: wire::ZigZag32,
    /// `actor_identifier` is the string entity type of the entity that emitted the sound, for example
    /// 'minecraft:skeleton'. Some sound types use this entity type for additional data.
    pub actor_identifier: String,
    /// `is_baby` specifies if the sound should be that of a baby mob. It is most notably used for
    /// parrot imitations, which will change based on if this field is set to true or not.
    pub is_baby: bool,
    /// `is_global` specifies if the sound should be played relatively or not. If set to true, the sound
    /// will have full volume, regardless of where the Position is, whereas if set to false, the sound's
    /// volume will be based on the distance to Position.
    pub is_global: bool,
    /// `actor_unique_id` is the unique ID of a source entity. The unique ID is a value that remains
    /// consistent across different sessions of the same world, but most servers simply fill the runtime
    /// ID of the entity out for this field.
    pub actor_unique_id: wire::I64LE,
    /// `fire_at_position` is the position in the same world at which the event should fire. If this is
    /// not present, the position entity will be used instead.
    /// Wire presence: optional value is preceded by a presence marker.
    pub fire_at_position: Option<glam::Vec3>,
}

impl LevelSoundEvent {
    pub const ID: u32 = 123;
}
/// LevelEventGeneric is sent by the server to send a 'generic' level event to the client. This
/// packet sends an NBT serialised object and may for that reason be used for any event holding
/// additional data.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LevelEventGeneric {
    /// `event_id` is a unique identifier that identifies the event called. The data that follows has
    /// fields in the NBT depending on what event it is.
    pub event_id: wire::ZigZag32,
    pub ctd: wire::NetworkNbt,
}

impl LevelEventGeneric {
    pub const ID: u32 = 124;
}
/// LecternUpdate is sent by the client to update the server on which page was opened in a book on a
/// lectern, or if the book should be removed from it.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LecternUpdate {
    pub new_page_to_show: wire::U8,
    pub total_pages: wire::U8,
    pub position_of_lectern_to_update: BlockPos,
}

impl LecternUpdate {
    pub const ID: u32 = 125;
}
/// ClientCacheStatus is sent by the client to the server at the start of the game. It is sent to
/// let the server know if it supports the client-side blob cache. Clients such as Nintendo Switch
/// do not support the cache, and attempting to use it anyway will fail.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCacheStatus {
    /// `is_cache_supported` specifies if the blob cache is enabled. If false, the server should not
    /// attempt to use the blob cache. If true, it may do so, but it may also choose not to use it.
    pub is_cache_supported: bool,
}

impl ClientCacheStatus {
    pub const ID: u32 = 129;
}
/// OnScreenTextureAnimation is sent by the server to show a certain animation on the screen of the
/// player. The packet is used, as an example, for when a raid is triggered and when a raid is
/// defeated.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct OnScreenTextureAnimation {
    /// `effect_id` is the type of the animation to show. The packet provides no further extra data to
    /// allow modifying the duration or other properties of the animation.
    pub effect_id: wire::U32LE,
}

impl OnScreenTextureAnimation {
    pub const ID: u32 = 130;
}
/// MapCreateLockedCopy is sent by the client to create a locked copy of one map into another map.
/// In vanilla, it is used in the cartography table to create a map that is locked and cannot be
/// modified.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MapCreateLockedCopy {
    /// `original_map_id` is the ID of the map that is being copied. The locked copy will obtain all
    /// content that is visible on this map, except the content will not change.
    pub original_map_id: ActorUniqueID,
    /// `new_map_id` is the ID of the map that holds the locked copy of the map that OriginalMapID
    /// points to. Its contents will be impossible to change.
    pub new_map_id: ActorUniqueID,
}

impl MapCreateLockedCopy {
    pub const ID: u32 = 131;
}
/// StructureTemplateDataRequest is sent by the client to request data of a structure.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureTemplateDataRequest {
    /// `structure_name` is the name of the structure that was set in the structure block's UI. This is
    /// the name used to export the structure to a file.
    pub structure_name: String,
    /// `structure_position` is the position of the structure block that has its template data
    /// requested.
    pub structure_position: BlockPos,
    /// `structure_settings` is a struct of settings that should be used for exporting the structure.
    /// These settings are identical to the last sent in the StructureBlockUpdate packet by the client.
    pub structure_settings: StructureSettings,
    /// `requested_operation` specifies the type of template data request that the player sent. It is
    /// one of the constants found above.
    pub requested_operation: StructureTemplateRequestOperation,
}

impl StructureTemplateDataRequest {
    pub const ID: u32 = 132;
}
/// StructureTemplateDataResponse is sent by the server to send data of a structure to the client in
/// response to a StructureTemplateDataRequest packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct StructureTemplateDataResponse {
    /// `structure_name` is the name of the structure that was requested. This is the name used to
    /// export the structure to a file.
    pub structure_name: String,
    pub structure_nbt: wire::NetworkNbt,
    /// `response_type` specifies the response type of the packet. This depends on the RequestType field
    /// sent in the StructureTemplateDataRequest packet and is one of the constants above.
    pub response_type: StructureTemplateResponseType,
}

impl StructureTemplateDataResponse {
    pub const ID: u32 = 133;
}
/// ClientCacheBlobStatus is part of the blob cache protocol. It is sent by the client to let the
/// server know what blobs it needs and which blobs it already has, in an ACK type system.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCacheBlobStatus {
    /// `missing_ids` is a list of blob hashes that the client does not have a blob available for. The
    /// server should send the blobs matching these hashes as soon as possible.
    pub missing_ids: Vec<wire::U64LE>,
    /// `found_ids` is a list of blob hashes that the client has a blob available for. The blobs hashes
    /// here mean that the client already has them: The server does not need to send the blobs anymore.
    pub found_ids: Vec<wire::U64LE>,
}

impl ClientCacheBlobStatus {
    pub const ID: u32 = 135;
}
/// ClientCacheMissResponse is part of the blob cache protocol. It is sent by the server in response
/// to a ClientCacheBlobStatus packet and contains the blob data of all blobs that the client
/// acknowledged not to have yet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCacheMissResponse {
    /// `missing_blobs` is a list of all blobs that the client sent misses for in the
    /// ClientCacheBlobStatus. These blobs hold the data of the blobs with the hashes they are matched
    /// with.
    pub missing_blobs: Vec<MissingBlobData>,
}

impl ClientCacheMissResponse {
    pub const ID: u32 = 136;
}
/// EducationSettings is a packet sent by the server to update Minecraft: Education Edition related
/// settings. It is unused by the normal base game.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EducationSettings {
    pub education_level_settings: EducationLevelSettings,
}

impl EducationSettings {
    pub const ID: u32 = 137;
}
/// Emote is sent by both the server and the client. When the client sends an emote, it sends this
/// packet to the server, after which the server will broadcast the packet to other players online.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct Emote {
    pub actor_runtime_id: ActorRuntimeID,
    /// `emote_id` is the ID of the emote to send.
    pub emote_id: String,
    pub emote_length_ticks: wire::VarUInt,
    /// `xuid` is the Xbox User ID of the player that sent the emote. It is only set when the emote is
    /// used by a player that is authenticated with Xbox Live.
    pub xuid: String,
    /// `platform_id` is an identifier only set for particular platforms when using an emote (presumably
    /// only for Nintendo Switch). It is otherwise an empty string, and is used to decide which players
    /// are able to emote with each other.
    pub platform_id: String,
    /// `flags` is a combination of flags that change the way the Emote packet operates. When the server
    /// sends this packet to other players, EmoteFlagServerSide must be present.
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
/// SettingsCommand is sent by the client when it changes a setting in the settings that results in
/// the issuing of a command to the server, such as when Show Coordinates is enabled.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SettingsCommand {
    /// `command` is the full command line that was sent to the server as a result of the setting that
    /// the client changed.
    pub command: String,
    /// `suppress_output` specifies if the client requests the suppressing of the output of the command
    /// that was executed. Generally this is set to true, as the client won't need a message to confirm
    /// the output of the change.
    pub suppress_output: bool,
}

impl SettingsCommand {
    pub const ID: u32 = 140;
}
/// AnvilDamage is sent by the client to request the dealing damage to an anvil. This packet is
/// completely pointless and the server should never listen to it.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AnvilDamage {
    /// `block_position` is the position in the world that the anvil can be found at.
    pub block_position: BlockPos,
}

impl AnvilDamage {
    pub const ID: u32 = 141;
}
/// CompletedUsingItem is sent by the server to tell the client that it should be done using the
/// item it is currently using.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CompletedUsingItem {
    /// `item_id` is the item ID of the item that the client completed using. This should typically be
    /// the ID of the item held in the hand.
    pub item_id: wire::I16LE,
    /// `item_use_method` is the method of the using of the item that was completed. It is one of the
    /// constants that may be found above.
    pub item_use_method: wire::I32LE,
}

impl CompletedUsingItem {
    pub const ID: u32 = 142;
}
/// NetworkSettings is sent by the server to update a variety of network settings. These settings
/// modify the way packets are sent over the network stack.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct NetworkSettings {
    /// `compression_threshold` is the minimum size of a packet that is compressed when sent. If the
    /// size of a packet is under this value, it is not compressed. When set to 0, all packets will be
    /// left uncompressed.
    pub compression_threshold: wire::U16LE,
    /// `compression_algorithm` is the algorithm that is used to compress packets.
    pub compression_algorithm: PacketCompressionAlgorithm,
    pub client_throttle_enabled: bool,
    /// `client_throttle_threshold` is the threshold for client throttling. If the number of players
    /// exceeds this value, the client will throttle players.
    pub client_throttle_threshold: wire::U8,
    /// `client_throttle_scalar` is the scalar for client throttling. The scalar is the amount of
    /// players that are ticked when throttling is enabled.
    pub client_throttle_scalar: wire::F32LE,
}

impl NetworkSettings {
    pub const ID: u32 = 143;
}
/// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to
/// synchronise the player input with the position server-side. The client sends this packet when
/// the ServerAuthoritativeMovementMode field in the StartGame packet is set to true, instead of the
/// MovePlayer packet. The client will send this packet once every tick.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerAuthInput {
    pub player_rotation: glam::Vec2,
    /// `position` holds the position that the player reports it has.
    pub position: glam::Vec3,
    /// `move_vector` is a Vec2 that specifies the direction in which the player moved, as a combination
    /// of X/Z values which are created using the WASD/controller stick state.
    pub move_vector: glam::Vec2,
    pub player_head_rotation: wire::F32LE,
    /// `input_data` is the set of input flags that together specify the way the player moved last tick.
    /// It holds the flags above.
    /// Wire presence: optional value is preceded by a presence marker.
    pub input_data: Option<Vec<InputData>>,
    /// `input_mode` specifies the way that the client inputs data to the screen. It is one of the
    /// constants that may be found above.
    pub input_mode: InputMode,
    /// `play_mode` specifies the way that the player is playing. The values it holds, which are rather
    /// random, may be found above.
    pub play_mode: ClientPlayMode,
    pub new_interaction_model: NewInteractionModel,
    pub interact_rotation: glam::Vec2,
    pub client_tick: PlayerInputTick,
    pub pos_delta: glam::Vec3,
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_use_transaction: Option<PackedItemUseLegacyInventoryTransaction>,
    /// `item_stack_request` is sent by the client to change an item in their inventory.
    /// Wire presence: optional value is preceded by a presence marker.
    pub item_stack_request: Option<ItemStackRequestData>,
    /// Wire presence: optional value is preceded by a presence marker.
    pub player_block_actions: Option<Vec<PlayerBlockActionData>>,
    /// `vehicle_rotation` is the rotation of the vehicle that the player is in, if any.
    /// Wire presence: optional value is preceded by a presence marker.
    pub vehicle_rotation: Option<glam::Vec2>,
    /// `client_predicted_vehicle` is the unique ID of the vehicle that the client predicts the player
    /// to be in.
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_predicted_vehicle: Option<ActorUniqueID>,
    pub analog_move_vector: glam::Vec2,
    /// `camera_orientation` is the vector that represents the camera's forward direction which can be
    /// used to transform movement to be camera relative.
    pub camera_orientation: glam::Vec3,
    /// `raw_move_vector` is the value of MoveVector before it is affected by input permissions,
    /// sneaking/fly speeds and isn't normalised for analogue inputs.
    pub raw_move_vector: glam::Vec2,
}

impl PlayerAuthInput {
    pub const ID: u32 = 144;
}
/// CreativeContent is a packet sent by the server to set the creative inventory's content for a
/// player. Introduced in 1.16, this packet replaces the previous method - sending an
/// InventoryContent packet with creative inventory window ID. As of v1.21.60, this packet is no
/// longer required to be sent as part of the login sequence however the client will crash if they
/// try to open their creative inventory before receiving this packet. Every item must be part of a
/// group, any items that are not part of a group will need to reference an "anonymous group" which
/// has an empty name OR no icon. The order of Groups and Items is how the client will render items
/// in the creative inventory compared to the previous, hard coded order.
/// Below is an example of defining 2 ungrouped items, 2 grouped items and then another 2 ungrouped
/// items, all in the nature category.
/// CreativeContent{ Groups: []protocol.CreativeGroup{ {Category: 1}, // No name or icon, this is
/// the "anonymous group" {Category: 1, Name: "itemGroup.name.planks", Icon:
/// protocol.ItemStack{...}}, // A "planks" group {Category: 1}, // Another "anonymous group" },
/// Items: []protocol.CreativeItem{ {CreativeItemNetworkID: 0, Item: protocol.ItemStack{...},
/// GroupIndex: 0}, // Ungrouped before "planks" {CreativeItemNetworkID: 1, Item:
/// protocol.ItemStack{...}, GroupIndex: 0}, // Ungrouped before "planks" {CreativeItemNetworkID: 2,
/// Item: protocol.ItemStack{...}, GroupIndex: 1}, // Grouped under the "planks" group
/// {CreativeItemNetworkID: 3, Item: protocol.ItemStack{...}, GroupIndex: 1}, // Grouped under the
/// "planks" group {CreativeItemNetworkID: 4, Item: protocol.ItemStack{...}, GroupIndex: 2}, //
/// Ungrouped after "planks" {CreativeItemNetworkID: 5, Item: protocol.ItemStack{...}, GroupIndex:
/// 2}, // Ungrouped after "planks" } }
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreativeContent {
    /// `groups` is a list of the groups that should be added to the creative inventory.
    pub groups: Vec<CreativeGroupInfo>,
    /// `entries` is a list of the items that should be added to the creative inventory.
    pub entries: Vec<CreativeItemEntry>,
}

impl CreativeContent {
    pub const ID: u32 = 145;
}
/// PlayerEnchantOptions is sent by the server to update the enchantment options displayed when the
/// user opens the enchantment table and puts an item in. This packet was added in 1.16 and allows
/// the server to decide on the enchantments that can be selected by the player. The
/// PlayerEnchantOptions packet should be sent once for every slot update of the enchantment table.
/// The vanilla server sends an empty PlayerEnchantOptions packet when the player opens the
/// enchantment table (air is present in the enchantment table slot) and sends the packet with
/// actual enchantments in it when items are put in that can have enchantments.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerEnchantOptions {
    /// `options` is a list of possible enchantment options for the item that was put into the
    /// enchantment table.
    pub options: Vec<ItemEnchantOption>,
}

impl PlayerEnchantOptions {
    pub const ID: u32 = 146;
}
/// ItemStackRequest is sent by the client to change item stacks in an inventory. It is essentially
/// a replacement of the InventoryTransaction packet added in 1.16 for inventory specific actions,
/// such as moving items around or crafting. The InventoryTransaction packet is still used for
/// actions such as placing blocks and interacting with entities.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackRequest {
    /// `requests` holds a list of item stack requests. These requests are all separate, but the client
    /// buffers the requests, so you might find multiple unrelated requests in this packet.
    pub requests: Vec<ItemStackRequestPacketData>,
}

impl ItemStackRequest {
    pub const ID: u32 = 147;
}
/// ItemStackResponse is sent by the server in response to an ItemStackRequest packet from the
/// client. This packet is used to either approve or reject ItemStackRequests from the client. If a
/// request is approved, the client will simply continue as normal. If rejected, the client will
/// undo the actions so that the inventory should be in sync with the server again.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemStackResponse {
    /// `responses` is a list of responses to ItemStackRequests sent by the client before. Responses
    /// either approve or reject a request from the client. Vanilla limits the size of this slice to
    /// 4096.
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
/// CodeBuilder is an Education Edition packet sent by the server to the client to open the URL to a
/// Code Builder (websocket) server.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CodeBuilder {
    /// `url` is the url to the Code Builder (websocket) server.
    pub url: String,
    /// `should_open_code_builder` specifies if the client should automatically open the Code Builder
    /// app. If set to true, the client will attempt to use the Code Builder app to connect to and
    /// interface with the server running at the URL above.
    pub should_open_code_builder: bool,
}

impl CodeBuilder {
    pub const ID: u32 = 150;
}
/// UpdatePlayerGameType is sent by the server to change the game mode of a player. It is
/// functionally identical to the SetPlayerGameType packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdatePlayerGameType {
    pub player_game_type: GameType,
    pub target_player: ActorUniqueID,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
}

impl UpdatePlayerGameType {
    pub const ID: u32 = 151;
}
/// EmoteList is sent by the client every time it joins the server and when it equips new emotes. It
/// may be used by the server to find out which emotes the client has available. If the player has
/// no emotes equipped, this packet is not sent. Under certain circumstances, this packet is also
/// sent from the server to the client, but I was unable to find when this is done.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EmoteList {
    pub runtime_id: ActorRuntimeID,
    pub emote_piece_ids: Vec<uuid::Uuid>,
}

impl EmoteList {
    pub const ID: u32 = 152;
}
/// PositionTrackingDBServerBroadcast is sent by the server in response to the
/// PositionTrackingDBClientRequest packet. This packet is, as of 1.16, currently only used for
/// lodestones. The server maintains a database with tracking IDs and their position and dimension.
/// The client will request these tracking IDs, (NBT tag set on the lodestone compass with the
/// tracking ID?) and the server will respond with the status of those tracking IDs. What is
/// actually done with the data sent depends on what the client chooses to do with it. For the
/// lodestone compass, it is used to make the compass point towards lodestones and to make it spin
/// if the lodestone at a position is no longer there.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PositionTrackingDBServerBroadcast {
    pub action: PositionTrackingDBServerBroadcastAction,
    pub id: PositionTrackingId,
    pub position_tracking_data: wire::NetworkNbt,
}

impl PositionTrackingDBServerBroadcast {
    pub const ID: u32 = 153;
}
/// PositionTrackingDBClientRequest is a packet sent by the client to request the position and
/// dimension of a 'tracking ID'. These IDs are tracked in a database by the server. In 1.16, this
/// is used for lodestones. The client will send this request to find the position a lodestone
/// compass needs to point to. If found, it will point to the lodestone. If not, it will start
/// spinning around. A PositionTrackingDBServerBroadcast packet should be sent in response to this
/// packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PositionTrackingDBClientRequest {
    pub action: PositionTrackingDBClientRequestAction,
    pub id: PositionTrackingId,
}

impl PositionTrackingDBClientRequest {
    pub const ID: u32 = 154;
}
/// DebugInfo is a packet sent by the server to the client. It does not seem to do anything when
/// sent to the normal client in 1.16.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DebugInfo {
    pub actor_id: ActorUniqueID,
    /// `data` is the debug data.
    pub data: String,
}

impl DebugInfo {
    pub const ID: u32 = 155;
}
/// PacketViolationWarning is sent by the client when it receives an invalid packet from the server.
/// It holds some information on the error that occurred. noinspection GoNameStartsWithPackageName
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PacketViolationWarning {
    pub violation_type: PacketViolationType,
    pub violation_severity: PacketViolationSeverity,
    pub violation_packet_id: wire::ZigZag32,
    /// `violation_context` holds a description on the violation of the packet.
    pub violation_context: String,
}

impl PacketViolationWarning {
    pub const ID: u32 = 156;
}
/// MotionPredictionHints is sent by the server to the client. There is a predictive movement
/// component for entities. This packet fills the "history" of that component and entity movement is
/// computed based on the points. Vanilla sends this packet instead of the SetActorMotion packet
/// when 'spatial optimisations' are enabled.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MotionPredictionHints {
    pub m_runtime_id: ActorRuntimeID,
    pub m_motion: glam::Vec3,
    pub m_on_ground: bool,
}

impl MotionPredictionHints {
    pub const ID: u32 = 157;
}
/// AnimateEntity is sent by the server to animate an entity client-side. It may be used to play a
/// single animation, or to activate a controller which can start a sequence of animations based on
/// different conditions specified in an animation controller. Much of the documentation of this
/// packet can be found at
/// https://learn.microsoft.com/en-us/minecraft/creator/reference/content/animationsreference
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AnimateEntity {
    /// `m_animation` is the name of a single animation to start playing.
    pub m_animation: String,
    /// `m_next_state` is the first state to start with. These states are declared in animation
    /// controllers (which, in themselves, are animations too). These states in turn may have animations
    /// and transitions to move to a next state.
    pub m_next_state: String,
    /// `m_stop_expression` is a MoLang expression that specifies when the animation should be stopped.
    pub m_stop_expression: String,
    /// `m_stop_expression_version` is the MoLang stop condition version.
    pub m_stop_expression_version: wire::I32LE,
    /// `m_controller` is the animation controller that is used to manage animations. These controllers
    /// decide when to play which animation.
    pub m_controller: String,
    /// `m_blend_out_time` does not currently seem to be used.
    pub m_blend_out_time: wire::F32LE,
    /// `m_runtime_ids` is list of runtime IDs of entities that the animation should be applied to.
    pub m_runtime_ids: Vec<ActorRuntimeID>,
}

impl AnimateEntity {
    pub const ID: u32 = 158;
}
/// CameraShake is sent by the server to make the camera shake client-side. This feature was added
/// for map- making partners.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraShake {
    /// `intensity` is the intensity of the shaking. The client limits this value to 4, so anything
    /// higher may not work.
    pub intensity: wire::F32LE,
    pub seconds: wire::F32LE,
    pub shake_type: CameraShakeType,
    pub shake_action: CameraShakeAction,
}

impl CameraShake {
    pub const ID: u32 = 159;
}
/// PlayerFog is sent by the server to render the different fogs in the Stack. The types of fog are
/// controlled by resource packs to change how they are rendered, and the ability to create custom
/// fog.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerFog {
    /// `fog_stack` is a list of fog identifiers to be sent to the client. Examples of fog identifiers
    /// are "minecraft:fog_ocean" and "minecraft:fog_hell".
    pub fog_stack: Vec<String>,
}

impl PlayerFog {
    pub const ID: u32 = 160;
}
/// CorrectPlayerMovePrediction is sent by the server if and only if
/// StartGame.ServerAuthoritativeMovementMode is set to AuthoritativeMovementModeServerWithRewind.
/// The packet is used to correct movement at a specific point in time.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CorrectPlayerMovePrediction {
    /// `prediction_type` is the type of prediction that was corrected. It is one of the constants
    /// above.
    pub prediction_type: RewindType,
    pub pos: glam::Vec3,
    pub pos_delta: glam::Vec3,
    /// `rotation` is the rotation of the player at the tick written in the field below.
    pub rotation: glam::Vec2,
    /// `vehicle_angular_velocity` is the angular velocity of the vehicle that the rider is riding.
    /// Wire presence: optional value is preceded by a presence marker.
    pub vehicle_angular_velocity: Option<wire::F32LE>,
    /// `on_ground` specifies if the player was on the ground at the time of the tick below.
    pub on_ground: bool,
    /// `tick` is the tick of the movement which was corrected by this packet.
    pub tick: PlayerInputTick,
}

impl CorrectPlayerMovePrediction {
    pub const ID: u32 = 161;
}
/// ItemRegistry is sent by the server to send the client a list of available items and attach
/// client-side components to a custom item. This packet was formerly known as the ItemComponent
/// packet before 1.21.60, which did not include item definitions but only the components.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ItemRegistry {
    /// `item_data` is a list of all items with their legacy IDs which are available in the game.
    /// Failing to send any of the items that are in the game will crash mobile clients. Any custom
    /// components are also attached to the items in this list.
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
/// SyncActorProperty is an alternative to synced actor data.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SyncActorProperty {
    /// `property_data` ...
    pub property_data: wire::NetworkNbt,
}

impl SyncActorProperty {
    pub const ID: u32 = 165;
}
/// AddVolumeEntity sends a volume entity's definition and metadata from server to client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AddVolumeEntity {
    pub entity_network_id: EntityNetId,
    pub components: wire::NetworkNbt,
    pub json_identifier: String,
    pub instance_name: String,
    pub min_bounds: BlockPos,
    pub max_bounds: BlockPos,
    pub dimension_type: DimensionType,
    /// `engine_version` is the engine version the entity is using, for example, '1.17.0'.
    pub engine_version: String,
}

impl AddVolumeEntity {
    pub const ID: u32 = 166;
}
/// RemoveVolumeEntity indicates a volume entity to be removed from server to client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RemoveVolumeEntity {
    pub entity_network_id: EntityNetId,
    pub dimension_type: DimensionType,
}

impl RemoveVolumeEntity {
    pub const ID: u32 = 167;
}
/// SimulationType is an in-progress packet. We currently do not know the use case.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SimulationType {
    /// `sim_type` is the simulation type selected.
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
/// CreatePhoto is a packet that allows players to export photos from their portfolios into items in
/// their inventory. This packet only works on the Education Edition version of Minecraft.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CreatePhoto {
    pub raw_id: wire::U64LE,
    /// `photo_name` is the name of the photo.
    pub photo_name: String,
    pub photo_item_name: String,
}

impl CreatePhoto {
    pub const ID: u32 = 171;
}
/// UpdateSubChunkBlocks is essentially just UpdateBlock packet, however for a set of blocks in a
/// sub-chunk.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateSubChunkBlocks {
    pub sub_chunk_block_position: BlockPos,
    pub blocks_changed: UpdateSubChunkBlocksChangedInfo,
}

impl UpdateSubChunkBlocks {
    pub const ID: u32 = 172;
}
/// SubChunk sends data about multiple sub-chunks around a center point.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SubChunk {
    /// `cache_enabled` is whether the sub-chunk caching is enabled or not.
    pub cache_enabled: bool,
    pub dimension_type: DimensionType,
    pub center_pos: SubChunkPos,
    pub sub_chunk_data: Vec<SubChunkData>,
}

impl SubChunk {
    pub const ID: u32 = 174;
}
/// SubChunkRequest requests specific sub-chunks from the server using a center point.
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
/// ScriptMessage is used to communicate custom messages from the client to the server, or from the
/// server to the client. While the name may suggest this packet is used for the discontinued
/// scripting API, it is likely instead for the GameTest framework.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ScriptMessage {
    pub message_id: String,
    pub message_value: String,
}

impl ScriptMessage {
    pub const ID: u32 = 177;
}
/// CodeBuilderSource is an Education Edition packet sent by the client to the server to run an
/// operation with a code builder.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CodeBuilderSource {
    /// `operation` is used to distinguish the operation performed. It is always one of the constants
    /// listed above.
    pub operation: CodeBuilderStorageQueryOptionsOperation,
    /// `category` is used to distinguish the category of the operation performed. It is always one of
    /// the constants listed above.
    pub category: CodeBuilderStorageQueryOptionsCategory,
    /// `code_status` is the status of the code builder. It is always one of the constants listed above.
    pub code_status: CodeBuilderExecutionStateCodeStatus,
}

impl CodeBuilderSource {
    pub const ID: u32 = 178;
}
/// TickingAreasLoadStatus is sent by the server to the client to notify the client of a ticking
/// area's loading status.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TickingAreasLoadStatus {
    /// `waiting_for_preload` is true if the server is waiting for the area's preload.
    pub waiting_for_preload: bool,
}

impl TickingAreasLoadStatus {
    pub const ID: u32 = 179;
}
/// DimensionData is a packet sent from the server to the client containing information about
/// data-driven dimensions that the server may have registered. This packet does not seem to be sent
/// by default, rather only being sent when any data-driven dimensions are registered.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DimensionData {
    /// `definitions` contain a list of data-driven dimension definitions registered on the server.
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
/// ChangeMobProperty is a packet sent from the server to the client to change one of the properties
/// of a mob client-side.
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
/// LessonProgress is a packet sent by the server to the client to inform the client of updated
/// progress on a lesson. This packet only functions on the Minecraft: Education Edition version of
/// the game.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LessonProgress {
    pub lesson_action: wire::ZigZag32,
    /// `score` is the score the client should use when displaying the progress.
    pub score: wire::ZigZag32,
    pub activity_id: String,
}

impl LessonProgress {
    pub const ID: u32 = 183;
}
/// RequestAbility is a packet sent by the client to the server to request permission for a specific
/// ability from the server. These abilities are defined above.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestAbility {
    /// `ability` is the ability that the client is requesting. This is one of the constants defined in
    /// the protocol/ability.go file.
    pub ability: wire::ZigZag32,
    pub value_type: RequestAbilityType,
    pub bool: bool,
    pub float: wire::F32LE,
}

impl RequestAbility {
    pub const ID: u32 = 184;
}
/// RequestPermissions is a packet sent from the client to the server to request permissions that
/// the client does not currently have. It can only be sent by operators and host in vanilla
/// Minecraft.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestPermissions {
    /// `target_player_id_raw_id` is the unique ID of the player. The unique ID is unique for the entire
    /// world and is often used in packets. Most servers send an EntityUniqueID equal to the
    /// EntityRuntimeID.
    pub target_player_id_raw_id: wire::I64LE,
    /// `player_permission_level` is the current permission level of the player. This is one of the
    /// constants that may be found in the AdventureSettings packet.
    pub player_permission_level: wire::ZigZag32,
    /// `custom_permission_flags` contains the requested permission flags.
    pub custom_permission_flags: wire::U16LE,
}

impl RequestPermissions {
    pub const ID: u32 = 185;
}
/// ToastRequest is a packet sent from the server to the client to display a toast to the top of the
/// screen. These toasts are the same as the ones seen when, for example, loading a new resource
/// pack or obtaining an achievement.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ToastRequest {
    /// `title` is the title of the toast.
    pub title: String,
    /// `content` is the message that the toast may contain alongside the title.
    pub content: String,
}

impl ToastRequest {
    pub const ID: u32 = 186;
}
/// UpdateAbilities is a packet sent from the server to the client to update the abilities of the
/// player. It, along with the UpdateAdventureSettings packet, are replacements of the
/// AdventureSettings packet since v1.19.10.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateAbilities {
    /// `data` represents various data about the abilities of a player, such as ability layers or
    /// permissions.
    pub data: SerializedAbilitiesData,
}

impl UpdateAbilities {
    pub const ID: u32 = 187;
}
/// UpdateAdventureSettings is a packet sent from the server to the client to update the adventure
/// settings of the player. It, along with the UpdateAbilities packet, are replacements of the
/// AdventureSettings packet since v1.19.10.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateAdventureSettings {
    pub adventure_settings: AdventureSettings,
}

impl UpdateAdventureSettings {
    pub const ID: u32 = 188;
}
/// DeathInfo is a packet sent from the server to the client expected to be sent when a player dies.
/// It contains messages related to the player's death, which are shown on the death screen as of
/// v1.19.10.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct DeathInfo {
    /// `death_cause_attack_name` is the cause of the player's death, such as "suffocation" or
    /// "suicide".
    pub death_cause_attack_name: String,
    /// `death_cause_message_list` is a list of death messages to be shown on the death screen.
    pub death_cause_message_list: Vec<String>,
}

impl DeathInfo {
    pub const ID: u32 = 189;
}
/// EditorNetwork is a packet sent from the server to the client and vise-versa to communicate
/// editor-mode related information. It carries a single compound tag containing the relevant
/// information.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct EditorNetwork {
    /// `route_to_manager` ...
    pub route_to_manager: bool,
    /// `payload` is a network little endian compound tag holding data relevant to the editor.
    pub payload: wire::NetworkNbt,
}

impl EditorNetwork {
    pub const ID: u32 = 190;
}
/// FeatureRegistry is a packet used to notify the client about the world generation features the
/// server is currently using. This is used in combination with the client-side world generation
/// system introduced in v1.19.20, allowing the client to completely generate the chunks of the
/// world without having to rely on the server.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct FeatureRegistry {
    /// `features_data_list` is a slice of all registered world generation features.
    pub features_data_list: Vec<FeatureRegistryFeatureBinaryJsonFormat>,
}

impl FeatureRegistry {
    pub const ID: u32 = 191;
}
/// ServerStats is a packet sent from the server to the client to update the client on server
/// statistics. It is purely used for telemetry.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerStats {
    /// `server_time` ...
    pub server_time: wire::F32LE,
    /// `network_time` ...
    pub network_time: wire::F32LE,
}

impl ServerStats {
    pub const ID: u32 = 192;
}
/// RequestNetworkSettings is sent by the client to request network settings, such as compression,
/// from the server.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RequestNetworkSettings {
    /// `client_network_version` is the protocol version of the player. The player is disconnected if
    /// the protocol is incompatible with the protocol of the server.
    pub client_network_version: wire::I32BE,
}

impl RequestNetworkSettings {
    pub const ID: u32 = 193;
}
/// GameTestRequest ...
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameTestRequest {
    /// `max_tests_per_batch` ...
    pub max_tests_per_batch: wire::ZigZag32,
    pub repeat_count: wire::ZigZag32,
    /// `rotation` represents the rotation of the test. It is one of the constants above.
    pub rotation: Rotation,
    pub stop_on_failure: bool,
    pub test_pos: BlockPos,
    /// `tests_per_row` ...
    pub tests_per_row: wire::ZigZag32,
    pub test_name: String,
}

impl GameTestRequest {
    pub const ID: u32 = 194;
}
/// GameTestResults is a packet sent in response to the GameTestRequest packet, with a boolean
/// indicating whether the test was successful or not, and an error string if the test failed.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GameTestResults {
    /// `succeeded` indicates whether the test succeeded or not.
    pub succeeded: bool,
    /// `error` is the error that occurred. If Succeeded is true, this field is empty.
    pub error: String,
    pub test_name: String,
}

impl GameTestResults {
    pub const ID: u32 = 195;
}
/// UpdateClientInputLocks is sent by the server to the client to lock specific player inputs such
/// as camera rotation, movement, jumping, sneaking, mounting or individual directional movement.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateClientInputLocks {
    /// `input_lock_component_data` is a set of flags that specify which client inputs are disabled,
    /// such as whether the player can move, rotate the camera, jump, sneak or mount/dismount entities.
    /// It is a combination of the ClientInputLock constants above.
    pub input_lock_component_data: wire::VarUInt,
}

impl UpdateClientInputLocks {
    pub const ID: u32 = 196;
}
/// CameraPresets gives the client a list of custom camera presets.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraPresets {
    pub camera_presets: CameraPresetList,
}

impl CameraPresets {
    pub const ID: u32 = 198;
}
/// UnlockedRecipes gives the client a list of recipes that have been unlocked, restricting the
/// recipes that appear in the recipe book.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UnlockedRecipes {
    pub packet_type: PacketType,
    pub unlocked_recipes_list: Vec<String>,
}

impl UnlockedRecipes {
    pub const ID: u32 = 199;
}
/// CameraInstruction gives a custom camera specific instructions to operate.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraInstruction {
    pub camera_instruction: CameraInstructionData,
}

impl CameraInstruction {
    pub const ID: u32 = 300;
}
/// TrimData is sent by the server to the client when they first join the server. It contains a list
/// of all the patterns and materials that can be applied via armour trims.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct TrimData {
    /// `trim_pattern_list` is a list of patterns that can be applied to armour. Each pattern has its
    /// own style and texture that is defined through resource packs.
    pub trim_pattern_list: Vec<TrimPattern>,
    /// `trim_material_list` is a list of materials that can be applied to armour. These are mostly
    /// different ores that have different colours for more customization.
    pub trim_material_list: Vec<TrimMaterial>,
}

impl TrimData {
    pub const ID: u32 = 302;
}
/// OpenSign is sent by the server to open a sign for editing. As of 1.19.80, the player can
/// interact with a sign to edit the text on both sides instead of just the front.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct OpenSign {
    /// `pos` is the position of the sign to edit. The client uses this position to get the data of the
    /// sign, including the existing text and formatting etc.
    pub pos: BlockPos,
    /// `is_front_side` dictates whether the front side of the sign should be opened for editing. If
    /// false, the back side is assumed to be edited.
    pub is_front_side: bool,
}

impl OpenSign {
    pub const ID: u32 = 303;
}
/// AgentAnimation is an Education Edition packet sent from the server to the client to make an
/// agent perform an animation.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AgentAnimation {
    pub agent_animation: AgentAnimationType,
    pub runtime_id: ActorRuntimeID,
}

impl AgentAnimation {
    pub const ID: u32 = 304;
}
/// RefreshEntitlements is sent by the client to the server to refresh the entitlements of the
/// player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct RefreshEntitlements {
}

impl RefreshEntitlements {
    pub const ID: u32 = 305;
}
/// PlayerToggleCrafterSlotRequest is sent by the client when it tries to toggle the state of a slot
/// within a Crafter.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerToggleCrafterSlotRequest {
    /// `pos_x` is the X position of the Crafter that is being modified.
    pub pos_x: wire::I32LE,
    /// `pos_y` is the Y position of the Crafter that is being modified.
    pub pos_y: wire::I32LE,
    /// `pos_z` is the Z position of the Crafter that is being modified.
    pub pos_z: wire::I32LE,
    pub slot_index: wire::U8,
    pub is_disabled: bool,
}

impl PlayerToggleCrafterSlotRequest {
    pub const ID: u32 = 306;
}
/// SetPlayerInventoryOptions is a bidirectional packet that can be used to update the inventory
/// options of a player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetPlayerInventoryOptions {
    pub inventory_options: InventoryOptions,
}

impl SetPlayerInventoryOptions {
    pub const ID: u32 = 307;
}
/// SetHud is sent by the server to set the visibility of individual HUD elements on the client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SetHud {
    pub hud_element: Vec<HudElement>,
    pub hud_visible: HudVisibility,
}

impl SetHud {
    pub const ID: u32 = 308;
}
/// AwardAchievement is sent by the server to award an achievement to a player.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct AwardAchievement {
    /// `achievement_id` is the ID of the achievement that should be awarded to the player. The values
    /// for these IDs are currently unknown.
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
/// JigsawStructureData is sent by the server to let the client know all the rules for jigsaw
/// structures.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct JigsawStructureData {
    pub jigsaw_structure_data_tag: wire::NetworkNbt,
}

impl JigsawStructureData {
    pub const ID: u32 = 313;
}
/// CurrentStructureFeature is sent by the server to let the client know the name of the structure
/// feature that the player is currently occupying.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CurrentStructureFeature {
    /// `current_structure_feature` is the identifier of the structure feature that the player is
    /// currently occupying. If the player is not occupying any structure feature, this field is empty.
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
/// CameraAimAssist is sent by the server to the client to set up aim assist for the client's
/// camera.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssist {
    /// `preset_id` is the ID of the preset that has previously been defined in the
    /// CameraAimAssistPresets packet.
    pub preset_id: String,
    /// `view_angle` is the maximum angle around the playes's cursor that the aim assist should check
    /// for a target, if TargetMode is set to protocol.AimAssistTargetModeAngle.
    pub view_angle: glam::Vec2,
    /// `distance` is the maximum distance from the player's cursor should check for a target, if
    /// TargetMode is set to protocol.AimAssistTargetModeDistance.
    pub distance: wire::F32LE,
    /// `target_mode` is the mode that the camera should use for detecting targets. This is currently
    /// one of protocol.AimAssistTargetModeAngle or protocol.AimAssistTargetModeDistance.
    pub target_mode: TargetMode,
    /// `action` is the action that should be performed with the aim assist. This is one of the
    /// constants above.
    pub action: CameraAimAssistAction,
    /// `show_debug_render` specifies if debug render should be shown.
    pub show_debug_render: bool,
}

impl CameraAimAssist {
    pub const ID: u32 = 316;
}
/// ContainerRegistryCleanup is sent by the server to trigger a client-side cleanup of the dynamic
/// container registry.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ContainerRegistryCleanup {
    /// `removed_containers` is a list of protocol.FullContainerName's that should be removed from the
    /// client-side container registry.
    pub removed_containers: Vec<FullContainerName>,
}

impl ContainerRegistryCleanup {
    pub const ID: u32 = 317;
}
/// MovementEffect is sent by the server to the client to update specific movement effects to allow
/// the client to predict its movement. For example, fireworks used during gliding will send this
/// packet to tell the client the exact duration of the boost.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct MovementEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub effect_id: MovementEffectType,
    pub effect_duration: wire::ZigZag32,
    /// `tick` is the server tick at which the packet was sent. It is used in relation to
    /// CorrectPlayerMovePrediction.
    pub tick: PlayerInputTick,
}

impl MovementEffect {
    pub const ID: u32 = 318;
}
/// CameraAimAssistPresets is sent by the server to the client to provide a list of categories and
/// presets that can be used when sending a CameraAimAssist packet or a CameraInstruction including
/// aim assist.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistPresets {
    /// `camera_aim_assist_presets` is a list of categories which can be referenced by one of the
    /// Presets.
    pub camera_aim_assist_presets: Vec<CameraAimAssistCategoryDefinition>,
    /// `camera_aim_assist_categories` is a list of presets which define a base for how aim assist
    /// should behave
    pub camera_aim_assist_categories: Vec<CameraAimAssistPresetDefinition>,
    /// `operation` is the operation to perform with the presets. It is one of the constants above.
    pub operation: CameraAimAssistPresetOperation,
}

impl CameraAimAssistPresets {
    pub const ID: u32 = 320;
}
/// ClientCameraAimAssist is sent by the server to send a player animation from one player to all
/// viewers of that player. It is used for a couple of actions, such as arm swimming and critical
/// hits.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientCameraAimAssist {
    /// `camera_preset_id` is the identifier of the preset to use which was previously defined in the
    /// CameraAimAssistPresets packet.
    pub camera_preset_id: String,
    /// `action` is the action to perform with the aim assist. It is one of the constants above.
    pub action: ClientCameraAimAssistAction,
    /// `allow_aim_assist` specifies the client can use aim assist or not.
    pub allow_aim_assist: bool,
}

impl ClientCameraAimAssist {
    pub const ID: u32 = 321;
}
/// ClientMovementPredictionSync is sent by the client to the server periodically if the client has
/// received movement corrections from the server, containing information about client-predictions
/// that are relevant to movement.
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
/// UpdateClientOptions is sent by the client when some of the client's options are updated, such as
/// the graphics mode.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct UpdateClientOptions {
    /// `graphics_mode_change` is the graphics mode that the client is using. It is one of the constants
    /// above.
    /// Wire presence: optional value is preceded by a presence marker.
    pub graphics_mode_change: Option<GraphicsMode>,
    /// `filter_profanity_change` is if the client only uses filtered messages or not.
    /// Wire presence: optional value is preceded by a presence marker.
    pub filter_profanity_change: Option<bool>,
}

impl UpdateClientOptions {
    pub const ID: u32 = 323;
}
/// PlayerVideoCapture packet is sent by the server to start or stop video recording for a player.
/// This packet only works on development builds and has no effect on retail builds. When recording,
/// the client will save individual frames to '/LocalCache/minecraftpe' in the format specified
/// below.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerVideoCapture {
    /// `action` is the action to perform with the video capture. It is one of the constants above.
    pub action: PlayerVideoCaptureData,
}

impl PlayerVideoCapture {
    pub const ID: u32 = 324;
}
/// PlayerUpdateEntityOverrides is sent by the server to modify an entity's properties individually.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PlayerUpdateEntityOverrides {
    pub target_id: ActorUniqueID,
    /// `property_index` is the index of the property to modify. The index is unique for each property
    /// of an entity.
    pub property_index: wire::VarUInt,
    pub update: PlayerUpdateEntityOverridesData,
}

impl PlayerUpdateEntityOverrides {
    pub const ID: u32 = 325;
}
/// PlayerLocation is sent by the server to the client to either update a player's position on the
/// locator bar, or remove them completely. The client will determine how to render the player on
/// the locator bar based on their own distance to Position.
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
/// PrimitiveShapes is a packet sent by the server to instruct the client to render one or more
/// shapes in the world. Shapes can be added, removed or updated based on the data provided
/// individually.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PrimitiveShapes {
    /// `array_of_primitive_shapes_can_be_a_mix_of_new_updated_or_removed` is a list of shapes to draw
    /// on the client-side.
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
/// GraphicsOverrideParameter is sent by the server to override graphics parameters.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct GraphicsOverrideParameter {
    pub parameter_keyframe_values: Vec<(wire::F32LE, glam::Vec3)>,
    /// `float_value` is an optional single float graphics parameter to be overridden.
    /// Wire presence: optional value is preceded by a presence marker.
    pub float_value: Option<wire::F32LE>,
    /// `vec3_value` is an optional single Vec3 graphics parameter to be overridden.
    /// Wire presence: optional value is preceded by a presence marker.
    pub vec3_value: Option<glam::Vec3>,
    /// `biome_identifier` is the identifier of the biome for which the parameters apply.
    pub biome_identifier: String,
    /// `player_identifier` is the optional identifier of the player for which the override parameter
    /// applies.
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
/// VoxelShapes is sent by the server to send voxel shape data to the client.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct VoxelShapes {
    /// `shapes` is a list of voxel shapes.
    pub shapes: Vec<VoxelShapesSerializableVoxelShape>,
    /// `name_map` is a map of shape names to IDs.
    pub name_map: Vec<(String, VoxelShapesRegistryHandle)>,
    /// `custom_shape_count` is the number of custom shapes.
    pub custom_shape_count: wire::U16LE,
}

impl VoxelShapes {
    pub const ID: u32 = 337;
}
/// CameraSpline is sent by the server to define camera spline paths.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraSpline {
    /// `camera_data_splines` is a list of camera spline definitions.
    pub camera_data_splines: Vec<CameraSplineDefinition>,
}

impl CameraSpline {
    pub const ID: u32 = 338;
}
/// CameraAimAssistActorPriority is sent by the server to define actor-specific aim assist
/// priorities.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct CameraAimAssistActorPriority {
    /// `camera_aim_assist_actor_priority_list` is a list of aim assist actor priority entries.
    pub camera_aim_assist_actor_priority_list: Vec<CameraAimAssistActorPriorityData>,
}

impl CameraAimAssistActorPriority {
    pub const ID: u32 = 339;
}
/// ResourcePacksReadyForValidation is sent by the client to inform the server that the client has
/// finished loading resource packs and is ready for validation.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ResourcePacksReadyForValidation {
}

impl ResourcePacksReadyForValidation {
    pub const ID: u32 = 340;
}
/// LocatorBar is sent by the server to add, remove or update waypoints on the client's locator bar.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct LocatorBar {
    /// `waypoints` is a slice of waypoints to add, remove or update.
    pub waypoints: Vec<LocatorBarWaypoint>,
}

impl LocatorBar {
    pub const ID: u32 = 341;
}
/// PartyChanged is sent by the client to the server to indicate that the player's party ID has
/// changed.
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
/// SyncWorldClocks is sent by the server to initialise and synchronise world clocks with the
/// client.
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
/// ServerStoreInfo is sent by the server to provide the client with a store entry point. Like the
/// ShowStoreOffer packet, this only has an effect on partnered servers.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerStoreInfo {
    /// `client_store_entry_point_configuration` is the store info to set, or nothing to fall back to
    /// the default.
    /// Wire presence: optional value is preceded by a presence marker.
    pub client_store_entry_point_configuration: Option<ServerConfigurationClientStoreEntryPointConfiguration>,
}

impl ServerStoreInfo {
    pub const ID: u32 = 346;
}
/// ServerPresenceInfo is sent by the server to provide the client with presence info.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ServerPresenceInfo {
    /// `presence_configuration` is the presence info to set, or nothing to fall back to the default.
    /// Wire presence: optional value is preceded by a presence marker.
    pub presence_configuration: Option<ServerConfigurationPresenceConfiguration>,
}

impl ServerPresenceInfo {
    pub const ID: u32 = 347;
}
/// ClientboundUpdateSoundData is sent by the server to update a sound that is currently playing,
/// identified by the handle that the server sent in the PlaySound packet that started it. Each
/// optional field is a Cereal union slot that may hold any SoundDataUpdate variant; its name does
/// not constrain the variant on the wire.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct ClientboundUpdateSoundData {
    /// `server_sound_handle` is the server-side handle of the sound to update.
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
/// SendPartyDestinationCookie is sent by the server to a client with a party destination cookie.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct SendPartyDestinationCookie {
    /// `cookie` is the opaque party destination cookie.
    pub cookie: String,
    /// `intent` is the intent of the cookie. It is one of the PartyDestinationCookieIntent constants.
    pub intent: String,
    /// `destination_name` is the name of the destination the cookie refers to.
    pub destination_name: String,
}

impl SendPartyDestinationCookie {
    pub const ID: u32 = 349;
}
/// PartyDestinationCookieResponse is sent by the client to the server in response to a
/// SendPartyDestinationCookie packet.
#[derive(Clone, Debug, Default, PartialEq)]
pub struct PartyDestinationCookieResponse {
    /// `cookie` is the opaque party destination cookie echoed back from the SendPartyDestinationCookie
    /// packet.
    pub cookie: String,
    /// `accepted` is true if the client accepted the party destination.
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
