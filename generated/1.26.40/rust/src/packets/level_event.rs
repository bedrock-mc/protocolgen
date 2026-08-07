// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LevelEvent {
    pub event_id: i32,
    pub position: glam::Vec3,
    pub data: i32,
}

impl LevelEvent {
    pub const ID: u32 = 25;
}
