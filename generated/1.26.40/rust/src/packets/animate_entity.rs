// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AnimateEntity {
    pub m_animation: String,
    pub m_next_state: String,
    pub m_stop_expression: String,
    pub m_stop_expression_version: i32,
    pub m_controller: String,
    pub m_blend_out_time: f32,
    pub m_runtime_ids: Vec<ActorRuntimeID>,
}

impl AnimateEntity {
    pub const ID: u32 = 158;
}
