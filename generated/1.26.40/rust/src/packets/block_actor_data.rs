// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BlockActorData {
    pub block_position: BlockPos,
    pub actor_data_tags: Nbt,
}

impl BlockActorData {
    pub const ID: u32 = 56;
}
