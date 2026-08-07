// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackChunkRequest {
    pub resource_name: String,
    pub chunk: i32,
}

pub const RESOURCEPACKCHUNKREQUEST_RESOURCE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const RESOURCEPACKCHUNKREQUEST_CHUNK_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;

impl ResourcePackChunkRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ResourcePackChunkRequestPacket.Resource Name", RESOURCEPACKCHUNKREQUEST_RESOURCE_NAME_SHAPE);
        encoder.field("ResourcePackChunkRequestPacket.Chunk", RESOURCEPACKCHUNKREQUEST_CHUNK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ResourcePackChunkRequestPacket.Resource Name", RESOURCEPACKCHUNKREQUEST_RESOURCE_NAME_SHAPE);
        decoder.field("ResourcePackChunkRequestPacket.Chunk", RESOURCEPACKCHUNKREQUEST_CHUNK_SHAPE);
    }
}
