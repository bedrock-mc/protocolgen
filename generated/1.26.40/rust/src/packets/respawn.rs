// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Respawn {
    pub position: glam::Vec3,
    pub state: PlayerRespawnState,
    pub player_runtime_id: ActorRuntimeID,
}

impl Respawn {
    pub const ID: u32 = 45;
}
