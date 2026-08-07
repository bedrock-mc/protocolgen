// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateBlock {
    pub block_position: BlockPos,
    pub block_runtime_id: u32,
    pub flags: u32,
    pub layer: u32,
}

impl UpdateBlock {
    pub const ID: u32 = 21;
}
