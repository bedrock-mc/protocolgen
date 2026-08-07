// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetSpawnPosition {
    pub spawn_position_type: SpawnPositionType,
    pub block_position: BlockPos,
    pub dimension_type: DimensionType,
    pub spawn_block_pos: BlockPos,
}

impl SetSpawnPosition {
    pub const ID: u32 = 43;
}
