// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Animate {
    pub action: AnimateAction,
    pub target_actor_runtime_id: ActorRuntimeID,
    pub data: f32,
    pub swing_source: Option<String>,
}

impl Animate {
    pub const ID: u32 = 44;
}
