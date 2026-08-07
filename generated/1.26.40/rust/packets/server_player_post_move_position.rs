// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ServerPlayerPostMovePosition {
    pub pos: glam::Vec3,
}

pub const SERVERPLAYERPOSTMOVEPOSITION_POS_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl ServerPlayerPostMovePosition {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ServerPlayerPostMovePositionPacket.Pos", SERVERPLAYERPOSTMOVEPOSITION_POS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ServerPlayerPostMovePositionPacket.Pos", SERVERPLAYERPOSTMOVEPOSITION_POS_SHAPE);
    }
}
