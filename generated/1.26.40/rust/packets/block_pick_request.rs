// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BlockPickRequest {
    pub position: BlockPos,
    pub with_data: bool,
    pub max_slots: u8,
}

pub const BLOCKPICKREQUEST_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const BLOCKPICKREQUEST_WITH_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const BLOCKPICKREQUEST_MAX_SLOTS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl BlockPickRequest {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("BlockPickRequestPacket.Position", BLOCKPICKREQUEST_POSITION_SHAPE);
        encoder.field("BlockPickRequestPacket.With Data?", BLOCKPICKREQUEST_WITH_DATA_SHAPE);
        encoder.field("BlockPickRequestPacket.Max Slots", BLOCKPICKREQUEST_MAX_SLOTS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("BlockPickRequestPacket.Position", BLOCKPICKREQUEST_POSITION_SHAPE);
        decoder.field("BlockPickRequestPacket.With Data?", BLOCKPICKREQUEST_WITH_DATA_SHAPE);
        decoder.field("BlockPickRequestPacket.Max Slots", BLOCKPICKREQUEST_MAX_SLOTS_SHAPE);
    }
}
