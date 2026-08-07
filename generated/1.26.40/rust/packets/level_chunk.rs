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

pub const LEVELCHUNK_CHUNK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"ChunkPos","type_id":"ChunkPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const LEVELCHUNK_DIMENSION_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const LEVELCHUNK_SUB_CHUNKS_COUNT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const LEVELCHUNK_CLIENT_REQUEST_SUB_CHUNK_LIMIT_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}}"#;
pub const LEVELCHUNK_CACHE_ENABLED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const LEVELCHUNK_CACHE_METADATA_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"LevelChunkPacketPayload::SubChunkMetadata","type_id":"LevelChunkPacketPayload::SubChunkMetadata","fields":[{"ordinal":0,"name":"Blob Id","semantic":"Blob Id","encode":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const LEVELCHUNK_SERIALIZED_CHUNK_DATA_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl LevelChunk {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LevelChunkPacket.Chunk Position", LEVELCHUNK_CHUNK_POSITION_SHAPE);
        encoder.field("LevelChunkPacket.Dimension Id", LEVELCHUNK_DIMENSION_ID_SHAPE);
        encoder.field("LevelChunkPacket.Sub-chunks Count", LEVELCHUNK_SUB_CHUNKS_COUNT_SHAPE);
        encoder.field("LevelChunkPacket.Client Request SubChunk Limit", LEVELCHUNK_CLIENT_REQUEST_SUB_CHUNK_LIMIT_SHAPE);
        encoder.field("LevelChunkPacket.Cache Enabled", LEVELCHUNK_CACHE_ENABLED_SHAPE);
        encoder.field("LevelChunkPacket.Cache Metadata", LEVELCHUNK_CACHE_METADATA_SHAPE);
        encoder.field("LevelChunkPacket.Serialized Chunk Data", LEVELCHUNK_SERIALIZED_CHUNK_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LevelChunkPacket.Chunk Position", LEVELCHUNK_CHUNK_POSITION_SHAPE);
        decoder.field("LevelChunkPacket.Dimension Id", LEVELCHUNK_DIMENSION_ID_SHAPE);
        decoder.field("LevelChunkPacket.Sub-chunks Count", LEVELCHUNK_SUB_CHUNKS_COUNT_SHAPE);
        decoder.field("LevelChunkPacket.Client Request SubChunk Limit", LEVELCHUNK_CLIENT_REQUEST_SUB_CHUNK_LIMIT_SHAPE);
        decoder.field("LevelChunkPacket.Cache Enabled", LEVELCHUNK_CACHE_ENABLED_SHAPE);
        decoder.field("LevelChunkPacket.Cache Metadata", LEVELCHUNK_CACHE_METADATA_SHAPE);
        decoder.field("LevelChunkPacket.Serialized Chunk Data", LEVELCHUNK_SERIALIZED_CHUNK_DATA_SHAPE);
    }
}
