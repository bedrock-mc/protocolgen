// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LevelEventGeneric {
    pub event_id: i32,
    pub ctd: Vec<u8>,
}

pub const LEVELEVENTGENERIC_EVENT_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const LEVELEVENTGENERIC_CTD_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl LevelEventGeneric {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LevelEventGenericPacket.Event Id", LEVELEVENTGENERIC_EVENT_ID_SHAPE);
        encoder.field("LevelEventGenericPacket.__[[CTD]]__", LEVELEVENTGENERIC_CTD_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LevelEventGenericPacket.Event Id", LEVELEVENTGENERIC_EVENT_ID_SHAPE);
        decoder.field("LevelEventGenericPacket.__[[CTD]]__", LEVELEVENTGENERIC_CTD_SHAPE);
    }
}
