// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ChangeMobProperty {
    pub actor_id: ActorUniqueID,
    pub property_name: String,
    pub bool_component_value: bool,
    pub string_component_value: String,
    pub int_component_value: i32,
    pub float_component_value: f32,
}

impl ChangeMobProperty {
    pub const ID: u32 = 182;
}
