// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestChunkRadius {
    pub chunk_radius: i32,
    pub max_chunk_radius: u8,
}

pub const REQUESTCHUNKRADIUS_CHUNK_RADIUS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const REQUESTCHUNKRADIUS_MAX_CHUNK_RADIUS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl RequestChunkRadius {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RequestChunkRadiusPacket.Chunk Radius", REQUESTCHUNKRADIUS_CHUNK_RADIUS_SHAPE);
        encoder.field("RequestChunkRadiusPacket.Max ChunkRadius", REQUESTCHUNKRADIUS_MAX_CHUNK_RADIUS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RequestChunkRadiusPacket.Chunk Radius", REQUESTCHUNKRADIUS_CHUNK_RADIUS_SHAPE);
        decoder.field("RequestChunkRadiusPacket.Max ChunkRadius", REQUESTCHUNKRADIUS_MAX_CHUNK_RADIUS_SHAPE);
    }
}
