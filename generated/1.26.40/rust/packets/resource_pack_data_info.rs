// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ResourcePackDataInfo {
    pub resource_name: String,
    pub chunk_size: u32,
    pub number_of_chunks: u32,
    pub file_size: u64,
    pub file_hash: String,
    pub is_premium_pack: bool,
    pub pack_type: u8,
}

pub const RESOURCEPACKDATAINFO_RESOURCE_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const RESOURCEPACKDATAINFO_CHUNK_SIZE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const RESOURCEPACKDATAINFO_NUMBER_OF_CHUNKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const RESOURCEPACKDATAINFO_FILE_SIZE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const RESOURCEPACKDATAINFO_FILE_HASH_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const RESOURCEPACKDATAINFO_IS_PREMIUM_PACK_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const RESOURCEPACKDATAINFO_PACK_TYPE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ResourcePackDataInfo {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ResourcePackDataInfoPacket.Resource Name", RESOURCEPACKDATAINFO_RESOURCE_NAME_SHAPE);
        encoder.field("ResourcePackDataInfoPacket.Chunk Size", RESOURCEPACKDATAINFO_CHUNK_SIZE_SHAPE);
        encoder.field("ResourcePackDataInfoPacket.Number of Chunks", RESOURCEPACKDATAINFO_NUMBER_OF_CHUNKS_SHAPE);
        encoder.field("ResourcePackDataInfoPacket.File Size", RESOURCEPACKDATAINFO_FILE_SIZE_SHAPE);
        encoder.field("ResourcePackDataInfoPacket.File Hash", RESOURCEPACKDATAINFO_FILE_HASH_SHAPE);
        encoder.field("ResourcePackDataInfoPacket.Is Premium Pack", RESOURCEPACKDATAINFO_IS_PREMIUM_PACK_SHAPE);
        encoder.field("ResourcePackDataInfoPacket.Pack Type", RESOURCEPACKDATAINFO_PACK_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ResourcePackDataInfoPacket.Resource Name", RESOURCEPACKDATAINFO_RESOURCE_NAME_SHAPE);
        decoder.field("ResourcePackDataInfoPacket.Chunk Size", RESOURCEPACKDATAINFO_CHUNK_SIZE_SHAPE);
        decoder.field("ResourcePackDataInfoPacket.Number of Chunks", RESOURCEPACKDATAINFO_NUMBER_OF_CHUNKS_SHAPE);
        decoder.field("ResourcePackDataInfoPacket.File Size", RESOURCEPACKDATAINFO_FILE_SIZE_SHAPE);
        decoder.field("ResourcePackDataInfoPacket.File Hash", RESOURCEPACKDATAINFO_FILE_HASH_SHAPE);
        decoder.field("ResourcePackDataInfoPacket.Is Premium Pack", RESOURCEPACKDATAINFO_IS_PREMIUM_PACK_SHAPE);
        decoder.field("ResourcePackDataInfoPacket.Pack Type", RESOURCEPACKDATAINFO_PACK_TYPE_SHAPE);
    }
}
