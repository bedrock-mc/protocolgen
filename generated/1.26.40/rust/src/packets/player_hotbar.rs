// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerHotbar {
    pub selected_slot: u32,
    pub container_id: u8,
    pub should_select_slot: bool,
}

impl PlayerHotbar {
    pub const ID: u32 = 48;
}
