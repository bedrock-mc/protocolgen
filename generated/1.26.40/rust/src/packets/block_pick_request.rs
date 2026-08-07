// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BlockPickRequest {
    pub position: BlockPos,
    pub with_data: bool,
    pub max_slots: u8,
}

impl BlockPickRequest {
    pub const ID: u32 = 34;
}
