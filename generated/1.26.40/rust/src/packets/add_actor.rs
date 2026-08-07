// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AddActor {
    pub target_actor_id: ActorUniqueID,
    pub target_runtime_id: ActorRuntimeID,
    pub actor_type: String,
    pub position: glam::Vec3,
    pub velocity: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: f32,
    pub y_body_rotation: f32,
    pub attributes_list: Vec<SyncedAttribute>,
    pub actor_data: SynchedActorDataCopyableDataList,
    pub synched_properties: PropertySyncData,
    pub actor_links: Vec<ActorLink>,
}

impl AddActor {
    pub const ID: u32 = 13;
}
