// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BlockEvent {
    pub block_position: BlockPos,
    pub event_type: i32,
    pub event_value: i32,
}

impl BlockEvent {
    pub const ID: u32 = 26;
}
