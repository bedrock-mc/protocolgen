// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct NetworkChunkPublisherUpdate {
    pub new_position_for_view: BlockPos,
    pub new_radius_for_view: u32,
    pub server_built_chunks_list: Vec<ChunkPos>,
}

pub const NETWORKCHUNKPUBLISHERUPDATE_NEW_POSITION_FOR_VIEW_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const NETWORKCHUNKPUBLISHERUPDATE_NEW_RADIUS_FOR_VIEW_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const NETWORKCHUNKPUBLISHERUPDATE_SERVER_BUILT_CHUNKS_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"element":{"kind":"struct","semantic":"ChunkPos","type_id":"ChunkPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl NetworkChunkPublisherUpdate {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("NetworkChunkPublisherUpdatePacket.New position for view", NETWORKCHUNKPUBLISHERUPDATE_NEW_POSITION_FOR_VIEW_SHAPE);
        encoder.field("NetworkChunkPublisherUpdatePacket.New radius for view", NETWORKCHUNKPUBLISHERUPDATE_NEW_RADIUS_FOR_VIEW_SHAPE);
        encoder.field("NetworkChunkPublisherUpdatePacket.Server Built Chunks List", NETWORKCHUNKPUBLISHERUPDATE_SERVER_BUILT_CHUNKS_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("NetworkChunkPublisherUpdatePacket.New position for view", NETWORKCHUNKPUBLISHERUPDATE_NEW_POSITION_FOR_VIEW_SHAPE);
        decoder.field("NetworkChunkPublisherUpdatePacket.New radius for view", NETWORKCHUNKPUBLISHERUPDATE_NEW_RADIUS_FOR_VIEW_SHAPE);
        decoder.field("NetworkChunkPublisherUpdatePacket.Server Built Chunks List", NETWORKCHUNKPUBLISHERUPDATE_SERVER_BUILT_CHUNKS_LIST_SHAPE);
    }
}
