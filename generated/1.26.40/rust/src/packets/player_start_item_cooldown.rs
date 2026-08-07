// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlayerStartItemCooldown {
    pub item_category: String,
    pub duration_ticks: i32,
}

impl PlayerStartItemCooldown {
    pub const ID: u32 = 176;
}
