// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerToggleCrafterSlotRequest {
    pub pos_x: i32,
    pub pos_y: i32,
    pub pos_z: i32,
    pub slot_index: u8,
    pub is_disabled: bool,
}

impl PlayerToggleCrafterSlotRequest {
    pub const ID: u32 = 306;
}
