// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct OpenSign {
    pub pos: BlockPos,
    pub is_front_side: bool,
}

pub const OPENSIGN_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const OPENSIGN_IS_FRONT_SIDE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl OpenSign {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("OpenSignPacket.Pos", OPENSIGN_POS_SHAPE);
        encoder.field("OpenSignPacket.Is Front Side", OPENSIGN_IS_FRONT_SIDE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("OpenSignPacket.Pos", OPENSIGN_POS_SHAPE);
        decoder.field("OpenSignPacket.Is Front Side", OPENSIGN_IS_FRONT_SIDE_SHAPE);
    }
}
