// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LevelChunk {
    pub chunk_position: ChunkPos,
    pub dimension_id: DimensionType,
    pub sub_chunks_count: u32,
    pub client_request_sub_chunk_limit: Option<i32>,
    pub cache_enabled: bool,
    pub cache_metadata: Vec<LevelChunkSubChunkMetadata>,
    pub serialized_chunk_data: String,
}

impl LevelChunk {
    pub const ID: u32 = 58;
}
