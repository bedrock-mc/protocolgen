// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LabTable {
    pub r#type: LabTableType,
    pub position: BlockPos,
    pub reaction: LabTableReactionType,
}

pub const LABTABLE_R_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"LabTablePacketPayload::Type","type_id":"enums/LabTablePacketPayload::Type","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"StartCombine","encode":{"kind":"void"}},{"value":1,"name":"StartReaction","encode":{"kind":"void"}},{"value":2,"name":"Reset","encode":{"kind":"void"}}]}"#;
pub const LABTABLE_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const LABTABLE_REACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"LabTableReactionType","type_id":"enums/LabTableReactionType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"None","encode":{"kind":"void"}},{"value":1,"name":"IceBomb","encode":{"kind":"void"}},{"value":2,"name":"Bleach","encode":{"kind":"void"}},{"value":3,"name":"ElephantToothpaste","encode":{"kind":"void"}},{"value":4,"name":"Fertilizer","encode":{"kind":"void"}},{"value":5,"name":"HeatBlock","encode":{"kind":"void"}},{"value":6,"name":"MagnesiumSalts","encode":{"kind":"void"}},{"value":7,"name":"MiscFire","encode":{"kind":"void"}},{"value":8,"name":"MiscExplosion","encode":{"kind":"void"}},{"value":9,"name":"MiscLava","encode":{"kind":"void"}},{"value":10,"name":"MiscMystical","encode":{"kind":"void"}},{"value":11,"name":"MiscSmoke","encode":{"kind":"void"}},{"value":12,"name":"MiscLargeSmoke","encode":{"kind":"void"}}]}"#;

impl LabTable {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LabTablePacket.Type", LABTABLE_R_TYPE_SHAPE);
        encoder.field("LabTablePacket.Position", LABTABLE_POSITION_SHAPE);
        encoder.field("LabTablePacket.Reaction", LABTABLE_REACTION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LabTablePacket.Type", LABTABLE_R_TYPE_SHAPE);
        decoder.field("LabTablePacket.Position", LABTABLE_POSITION_SHAPE);
        decoder.field("LabTablePacket.Reaction", LABTABLE_REACTION_SHAPE);
    }
}
