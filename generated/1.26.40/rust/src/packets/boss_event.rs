// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BossEvent {
    pub target_actor_id: ActorUniqueID,
    pub player_id: ActorUniqueID,
    pub event_type: BossEventUpdateType,
    pub name: String,
    pub filtered_name: String,
    pub health_percent: f32,
    pub color: BossBarColor,
    pub overlay: BossBarOverlay,
}

impl BossEvent {
    pub const ID: u32 = 74;
}
