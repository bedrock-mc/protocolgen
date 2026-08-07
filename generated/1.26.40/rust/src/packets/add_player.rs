// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddPlayer {
    pub uuid: uuid::Uuid,
    pub player_name: String,
    pub target_runtime_id: ActorRuntimeID,
    pub platform_chat_id: String,
    pub position: glam::Vec3,
    pub velocity: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: f32,
    pub carried_item: CerealizerNetworkItemStackDescriptorSerializedData,
    pub player_game_type: GameType,
    pub entity_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub abilities_data: SerializedAbilitiesData,
    pub actor_links: Vec<ActorLink>,
    pub device_id: String,
    pub build_platform: BuildPlatform,
}

impl AddPlayer {
    pub const ID: u32 = 12;
}
