// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BlockActorData {
    pub block_position: BlockPos,
    pub actor_data_tags: Vec<u8>,
}

pub const BLOCKACTORDATA_BLOCK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const BLOCKACTORDATA_ACTOR_DATA_TAGS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl BlockActorData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("BlockActorDataPacket.Block Position", BLOCKACTORDATA_BLOCK_POSITION_SHAPE);
        encoder.field("BlockActorDataPacket.Actor Data Tags", BLOCKACTORDATA_ACTOR_DATA_TAGS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("BlockActorDataPacket.Block Position", BLOCKACTORDATA_BLOCK_POSITION_SHAPE);
        decoder.field("BlockActorDataPacket.Actor Data Tags", BLOCKACTORDATA_ACTOR_DATA_TAGS_SHAPE);
    }
}
