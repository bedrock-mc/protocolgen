// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateBlockSynced {
    pub block_position: BlockPos,
    pub block_runtime_id: u32,
    pub flags: u32,
    pub layer: u32,
    pub unique_actor_id: u64,
    pub actor_sync_message: u64,
}

impl UpdateBlockSynced {
    pub const ID: u32 = 110;
}
