// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateSubChunkBlocks {
    pub sub_chunk_block_position: BlockPos,
    pub blocks_changed: UpdateSubChunkBlocksChangedInfo,
}

impl UpdateSubChunkBlocks {
    pub const ID: u32 = 172;
}
