// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackChunkData {
    pub resource_name: String,
    pub chunk_id: u32,
    pub byte_offset: u64,
    pub chunk_data: String,
}

pub const RESOURCEPACKCHUNKDATA_RESOURCE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const RESOURCEPACKCHUNKDATA_CHUNK_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const RESOURCEPACKCHUNKDATA_BYTE_OFFSET_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const RESOURCEPACKCHUNKDATA_CHUNK_DATA_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ResourcePackChunkData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ResourcePackChunkDataPacket.Resource Name", RESOURCEPACKCHUNKDATA_RESOURCE_NAME_SHAPE);
        encoder.field("ResourcePackChunkDataPacket.Chunk ID", RESOURCEPACKCHUNKDATA_CHUNK_ID_SHAPE);
        encoder.field("ResourcePackChunkDataPacket.Byte Offset", RESOURCEPACKCHUNKDATA_BYTE_OFFSET_SHAPE);
        encoder.field("ResourcePackChunkDataPacket.Chunk Data", RESOURCEPACKCHUNKDATA_CHUNK_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ResourcePackChunkDataPacket.Resource Name", RESOURCEPACKCHUNKDATA_RESOURCE_NAME_SHAPE);
        decoder.field("ResourcePackChunkDataPacket.Chunk ID", RESOURCEPACKCHUNKDATA_CHUNK_ID_SHAPE);
        decoder.field("ResourcePackChunkDataPacket.Byte Offset", RESOURCEPACKCHUNKDATA_BYTE_OFFSET_SHAPE);
        decoder.field("ResourcePackChunkDataPacket.Chunk Data", RESOURCEPACKCHUNKDATA_CHUNK_DATA_SHAPE);
    }
}
