// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct BlockEvent {
    pub block_position: BlockPos,
    pub event_type: i32,
    pub event_value: i32,
}

pub const BLOCKEVENT_BLOCK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const BLOCKEVENT_EVENT_TYPE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const BLOCKEVENT_EVENT_VALUE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl BlockEvent {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("BlockEventPacket.Block Position", BLOCKEVENT_BLOCK_POSITION_SHAPE);
        encoder.field("BlockEventPacket.Event Type", BLOCKEVENT_EVENT_TYPE_SHAPE);
        encoder.field("BlockEventPacket.Event Value", BLOCKEVENT_EVENT_VALUE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("BlockEventPacket.Block Position", BLOCKEVENT_BLOCK_POSITION_SHAPE);
        decoder.field("BlockEventPacket.Event Type", BLOCKEVENT_EVENT_TYPE_SHAPE);
        decoder.field("BlockEventPacket.Event Value", BLOCKEVENT_EVENT_VALUE_SHAPE);
    }
}
