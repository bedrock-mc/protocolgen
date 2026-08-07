// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ContainerSetData {
    pub container_id: u8,
    pub id: i32,
    pub value: i32,
}

pub const CONTAINERSETDATA_CONTAINER_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CONTAINERSETDATA_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const CONTAINERSETDATA_VALUE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl ContainerSetData {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ContainerSetDataPacket.Container ID", CONTAINERSETDATA_CONTAINER_ID_SHAPE);
        encoder.field("ContainerSetDataPacket.ID", CONTAINERSETDATA_ID_SHAPE);
        encoder.field("ContainerSetDataPacket.Value", CONTAINERSETDATA_VALUE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ContainerSetDataPacket.Container ID", CONTAINERSETDATA_CONTAINER_ID_SHAPE);
        decoder.field("ContainerSetDataPacket.ID", CONTAINERSETDATA_ID_SHAPE);
        decoder.field("ContainerSetDataPacket.Value", CONTAINERSETDATA_VALUE_SHAPE);
    }
}
