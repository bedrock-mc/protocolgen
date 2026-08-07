// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunk {
    pub cache_enabled: bool,
    pub dimension_type: DimensionType,
    pub center_pos: SubChunkPos,
    pub sub_chunk_data: Vec<SubChunkSubChunkPacketData>,
}

impl SubChunk {
    pub const ID: u32 = 174;
}
