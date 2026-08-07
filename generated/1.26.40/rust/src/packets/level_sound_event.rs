// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LevelSoundEvent {
    pub sound_event: String,
    pub position: glam::Vec3,
    pub data: i32,
    pub actor_identifier: String,
    pub is_baby: bool,
    pub is_global: bool,
    pub actor_unique_id: i64,
    pub fire_at_position: Option<glam::Vec3>,
}

impl LevelSoundEvent {
    pub const ID: u32 = 123;
}
