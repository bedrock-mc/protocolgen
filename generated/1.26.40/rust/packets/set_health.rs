// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetHealth {
    pub health: i32,
}

pub const SETHEALTH_HEALTH_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl SetHealth {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetHealthPacket.Health", SETHEALTH_HEALTH_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetHealthPacket.Health", SETHEALTH_HEALTH_SHAPE);
    }
}
