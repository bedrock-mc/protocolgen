// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateEquip {
    pub container_id: u8,
    pub r#type: u8,
    pub size: i32,
    pub entity_unique_id: ActorUniqueID,
    pub data: Vec<u8>,
}

pub const UPDATEEQUIP_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEEQUIP_R_TYPE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const UPDATEEQUIP_SIZE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const UPDATEEQUIP_ENTITY_UNIQUE_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const UPDATEEQUIP_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"nbt_le","width":0,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl UpdateEquip {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateEquipPacket.Container Id", UPDATEEQUIP_CONTAINER_ID_SHAPE);
        encoder.field("UpdateEquipPacket.Type", UPDATEEQUIP_R_TYPE_SHAPE);
        encoder.field("UpdateEquipPacket.Size", UPDATEEQUIP_SIZE_SHAPE);
        encoder.field("UpdateEquipPacket.Entity Unique Id", UPDATEEQUIP_ENTITY_UNIQUE_ID_SHAPE);
        encoder.field("UpdateEquipPacket.Data", UPDATEEQUIP_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateEquipPacket.Container Id", UPDATEEQUIP_CONTAINER_ID_SHAPE);
        decoder.field("UpdateEquipPacket.Type", UPDATEEQUIP_R_TYPE_SHAPE);
        decoder.field("UpdateEquipPacket.Size", UPDATEEQUIP_SIZE_SHAPE);
        decoder.field("UpdateEquipPacket.Entity Unique Id", UPDATEEQUIP_ENTITY_UNIQUE_ID_SHAPE);
        decoder.field("UpdateEquipPacket.Data", UPDATEEQUIP_DATA_SHAPE);
    }
}
