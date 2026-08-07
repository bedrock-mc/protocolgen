// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SubChunkRequest {
    pub dimension_type: DimensionType,
    pub sub_chunk_position_offset_list: Vec<SubChunkSubChunkPosOffset>,
    pub center_pos: SubChunkPos,
}

pub const SUBCHUNKREQUEST_DIMENSION_TYPE_SHAPE: &str = r#"{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SUBCHUNKREQUEST_SUB_CHUNK_POSITION_OFFSET_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"SubChunkPacketPayload::SubChunkPosOffset","type_id":"SubChunkPacketPayload::SubChunkPosOffset","fields":[{"ordinal":0,"name":"Subchunk Offset X","semantic":"Subchunk Offset X","encode":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Subchunk Offset Y","semantic":"Subchunk Offset Y","encode":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Subchunk Offset Z","semantic":"Subchunk Offset Z","encode":{"kind":"primitive","primitive":{"code":"i8","width":8,"signed":true,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;
pub const SUBCHUNKREQUEST_CENTER_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"SubChunkPos","type_id":"SubChunkPos","fields":[{"ordinal":0,"name":"Subchunk Position X","semantic":"Subchunk Position X","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Subchunk Position Y","semantic":"Subchunk Position Y","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Subchunk Position Z","semantic":"Subchunk Position Z","encode":{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl SubChunkRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SubChunkRequestPacket.Dimension Type", SUBCHUNKREQUEST_DIMENSION_TYPE_SHAPE);
        encoder.field("SubChunkRequestPacket.SubChunk Position Offset List", SUBCHUNKREQUEST_SUB_CHUNK_POSITION_OFFSET_LIST_SHAPE);
        encoder.field("SubChunkRequestPacket.Center Pos", SUBCHUNKREQUEST_CENTER_POS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SubChunkRequestPacket.Dimension Type", SUBCHUNKREQUEST_DIMENSION_TYPE_SHAPE);
        decoder.field("SubChunkRequestPacket.SubChunk Position Offset List", SUBCHUNKREQUEST_SUB_CHUNK_POSITION_OFFSET_LIST_SHAPE);
        decoder.field("SubChunkRequestPacket.Center Pos", SUBCHUNKREQUEST_CENTER_POS_SHAPE);
    }
}
