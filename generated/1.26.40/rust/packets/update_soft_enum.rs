// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct UpdateSoftEnum {
    pub enum_name: String,
    pub values: Vec<String>,
    pub update_type: SoftEnumUpdateType,
}

pub const UPDATESOFTENUM_ENUM_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const UPDATESOFTENUM_VALUES_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;
pub const UPDATESOFTENUM_UPDATE_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"SoftEnumUpdateType","type_id":"enums/SoftEnumUpdateType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Add","encode":{"kind":"void"}},{"value":1,"name":"Remove","encode":{"kind":"void"}},{"value":2,"name":"Replace","encode":{"kind":"void"}}]}"#;

impl UpdateSoftEnum {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("UpdateSoftEnumPacket.Enum Name", UPDATESOFTENUM_ENUM_NAME_SHAPE);
        encoder.field("UpdateSoftEnumPacket.Values", UPDATESOFTENUM_VALUES_SHAPE);
        encoder.field("UpdateSoftEnumPacket.Update Type", UPDATESOFTENUM_UPDATE_TYPE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("UpdateSoftEnumPacket.Enum Name", UPDATESOFTENUM_ENUM_NAME_SHAPE);
        decoder.field("UpdateSoftEnumPacket.Values", UPDATESOFTENUM_VALUES_SHAPE);
        decoder.field("UpdateSoftEnumPacket.Update Type", UPDATESOFTENUM_UPDATE_TYPE_SHAPE);
    }
}
