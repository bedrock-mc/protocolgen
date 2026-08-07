// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundTextureShift {
    pub action_id: ClientboundTextureShiftAction,
    pub collection_name: String,
    pub from_step: String,
    pub to_step: String,
    pub all_steps: Vec<String>,
    pub current_length_in_ticks: u64,
    pub total_length_in_ticks: u64,
    pub enabled: bool,
}

impl ClientboundTextureShift {
    pub const ID: u32 = 336;
}
