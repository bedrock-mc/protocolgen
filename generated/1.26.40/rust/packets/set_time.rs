// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetTime {
    pub time: i32,
}

pub const SETTIME_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl SetTime {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetTimePacket.Time", SETTIME_TIME_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetTimePacket.Time", SETTIME_TIME_SHAPE);
    }
}
