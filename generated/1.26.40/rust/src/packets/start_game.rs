// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
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
    pub level_current_time: u64,
    pub enchantment_seed: i32,
    pub block_properties: Vec<ServerBlockProperty>,
    pub multiplayer_correlation_id: String,
    pub enable_item_stack_net_manager: bool,
    pub server_version: String,
    pub player_property_data: Nbt,
    pub server_block_type_registry_checksum: u64,
    pub world_template_id: uuid::Uuid,
    pub server_enabled_client_side_generation: bool,
    pub block_network_ids_are_hashes: bool,
    pub network_permissions: NetworkPermissions,
    pub server_configuration_join_info: Option<ServerConfigurationServerConfigurationJoinInfo>,
    pub server_telemetry_data: SocialEventsServerTelemetryData,
}

impl StartGame {
    pub const ID: u32 = 11;
}
