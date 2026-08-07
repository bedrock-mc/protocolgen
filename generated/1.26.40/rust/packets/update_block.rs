// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateBlock {
    pub block_position: BlockPos,
    pub block_runtime_id: u32,
    pub flags: u32,
    pub layer: u32,
}

pub const UPDATEBLOCK_BLOCK_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const UPDATEBLOCK_BLOCK_RUNTIME_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEBLOCK_FLAGS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEBLOCK_LAYER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl UpdateBlock {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateBlockPacket.Block Position", UPDATEBLOCK_BLOCK_POSITION_SHAPE);
        encoder.field("UpdateBlockPacket.Block Runtime ID", UPDATEBLOCK_BLOCK_RUNTIME_ID_SHAPE);
        encoder.field("UpdateBlockPacket.Flags", UPDATEBLOCK_FLAGS_SHAPE);
        encoder.field("UpdateBlockPacket.Layer", UPDATEBLOCK_LAYER_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateBlockPacket.Block Position", UPDATEBLOCK_BLOCK_POSITION_SHAPE);
        decoder.field("UpdateBlockPacket.Block Runtime ID", UPDATEBLOCK_BLOCK_RUNTIME_ID_SHAPE);
        decoder.field("UpdateBlockPacket.Flags", UPDATEBLOCK_FLAGS_SHAPE);
        decoder.field("UpdateBlockPacket.Layer", UPDATEBLOCK_LAYER_SHAPE);
    }
}
