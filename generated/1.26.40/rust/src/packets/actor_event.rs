// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ActorEvent {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: ActorEventType,
    pub data: i32,
    pub fire_at_position: Option<glam::Vec3>,
}

impl ActorEvent {
    pub const ID: u32 = 27;
}
