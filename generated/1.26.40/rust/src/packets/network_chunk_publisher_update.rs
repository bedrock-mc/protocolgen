// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct NetworkChunkPublisherUpdate {
    pub new_position_for_view: BlockPos,
    pub new_radius_for_view: u32,
    pub server_built_chunks_list: Vec<ChunkPos>,
}

impl NetworkChunkPublisherUpdate {
    pub const ID: u32 = 121;
}
