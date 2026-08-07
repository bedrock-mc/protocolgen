// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct HurtArmor {
    pub cause: i32,
    pub damage: i32,
    pub armor_slots: u64,
}

impl HurtArmor {
    pub const ID: u32 = 38;
}
