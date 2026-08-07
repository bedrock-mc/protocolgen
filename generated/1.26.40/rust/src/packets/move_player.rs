// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MovePlayer {
    pub player_runtime_id: ActorRuntimeID,
    pub position: glam::Vec3,
    pub rotation: glam::Vec2,
    pub y_head_rotation: f32,
    pub position_mode: PlayerPositionModeComponentPositionMode,
    pub on_ground: bool,
    pub riding_runtime_id: ActorRuntimeID,
    pub teleport_data: Option<MovePlayerTeleportData>,
    pub tick: PlayerInputTick,
}

impl MovePlayer {
    pub const ID: u32 = 19;
}
