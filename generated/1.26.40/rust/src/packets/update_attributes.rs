// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateAttributes {
    pub target_runtime_id: ActorRuntimeID,
    pub attribute_list: Vec<AttributeData>,
    pub tick: PlayerInputTick,
}

impl UpdateAttributes {
    pub const ID: u32 = 29;
}
