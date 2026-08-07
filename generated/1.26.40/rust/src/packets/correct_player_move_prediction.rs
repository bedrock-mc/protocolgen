// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CorrectPlayerMovePrediction {
    pub prediction_type: RewindType,
    pub pos: glam::Vec3,
    pub pos_delta: glam::Vec3,
    pub rotation: glam::Vec2,
    pub vehicle_angular_velocity: Option<f32>,
    pub on_ground: bool,
    pub tick: PlayerInputTick,
}

impl CorrectPlayerMovePrediction {
    pub const ID: u32 = 161;
}
