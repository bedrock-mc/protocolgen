// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MotionPredictionHints {
    pub m_runtime_id: ActorRuntimeID,
    pub m_motion: glam::Vec3,
    pub m_on_ground: bool,
}

impl MotionPredictionHints {
    pub const ID: u32 = 157;
}
