// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LevelEvent {
    pub event_id: i32,
    pub position: glam::Vec3,
    pub data: i32,
}

pub const LEVELEVENT_EVENT_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const LEVELEVENT_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const LEVELEVENT_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl LevelEvent {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LevelEventPacket.Event Id", LEVELEVENT_EVENT_ID_SHAPE);
        encoder.field("LevelEventPacket.Position", LEVELEVENT_POSITION_SHAPE);
        encoder.field("LevelEventPacket.Data", LEVELEVENT_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LevelEventPacket.Event Id", LEVELEVENT_EVENT_ID_SHAPE);
        decoder.field("LevelEventPacket.Position", LEVELEVENT_POSITION_SHAPE);
        decoder.field("LevelEventPacket.Data", LEVELEVENT_DATA_SHAPE);
    }
}
