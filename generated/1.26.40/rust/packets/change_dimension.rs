// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ChangeDimension {
    pub dimension_id: DimensionType,
    pub position: Vec3,
    pub respawn: bool,
    pub loading_screen_id: Option<u32>,
}

pub const CHANGEDIMENSION_DIMENSION_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"DimensionType","type_id":"DimensionType","fields":[{"ordinal":0,"name":"value","semantic":"value","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CHANGEDIMENSION_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CHANGEDIMENSION_RESPAWN_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CHANGEDIMENSION_LOADING_SCREEN_ID_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"primitive","primitive":{"code":"u32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}}"#;

impl ChangeDimension {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ChangeDimensionPacket.Dimension ID", CHANGEDIMENSION_DIMENSION_ID_SHAPE);
        encoder.field("ChangeDimensionPacket.Position", CHANGEDIMENSION_POSITION_SHAPE);
        encoder.field("ChangeDimensionPacket.Respawn", CHANGEDIMENSION_RESPAWN_SHAPE);
        encoder.field("ChangeDimensionPacket.Loading Screen Id", CHANGEDIMENSION_LOADING_SCREEN_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ChangeDimensionPacket.Dimension ID", CHANGEDIMENSION_DIMENSION_ID_SHAPE);
        decoder.field("ChangeDimensionPacket.Position", CHANGEDIMENSION_POSITION_SHAPE);
        decoder.field("ChangeDimensionPacket.Respawn", CHANGEDIMENSION_RESPAWN_SHAPE);
        decoder.field("ChangeDimensionPacket.Loading Screen Id", CHANGEDIMENSION_LOADING_SCREEN_ID_SHAPE);
    }
}
