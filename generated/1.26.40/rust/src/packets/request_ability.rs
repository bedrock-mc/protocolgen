// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestAbility {
    pub ability: i32,
    pub value_type: RequestAbilityType,
    pub bool: bool,
    pub float: f32,
}

impl RequestAbility {
    pub const ID: u32 = 184;
}
