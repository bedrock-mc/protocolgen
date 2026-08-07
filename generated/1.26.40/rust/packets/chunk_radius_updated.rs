// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ChunkRadiusUpdated {
    pub chunk_radius: i32,
}

pub const CHUNKRADIUSUPDATED_CHUNK_RADIUS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl ChunkRadiusUpdated {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ChunkRadiusUpdatedPacket.Chunk Radius", CHUNKRADIUSUPDATED_CHUNK_RADIUS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ChunkRadiusUpdatedPacket.Chunk Radius", CHUNKRADIUSUPDATED_CHUNK_RADIUS_SHAPE);
    }
}
