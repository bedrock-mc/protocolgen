// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunkRequest {
    pub dimension_type: DimensionType,
    pub sub_chunk_position_offset_list: Vec<SubChunkSubChunkPosOffset>,
    pub center_pos: SubChunkPos,
}

impl SubChunkRequest {
    pub const ID: u32 = 175;
}
