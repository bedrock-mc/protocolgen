// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestAbility {
    pub ability: i32,
    pub value_type: RequestAbilityType,
    pub bool: bool,
    pub float: f32,
}

pub const REQUESTABILITY_ABILITY_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const REQUESTABILITY_VALUE_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"RequestAbilityPacketPayload::Type","type_id":"enums/RequestAbilityPacketPayload::Type","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Unset","encode":{"kind":"void"}},{"value":1,"name":"Bool","encode":{"kind":"void"}},{"value":2,"name":"Float","encode":{"kind":"void"}}]}"#;
pub const REQUESTABILITY_BOOL_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const REQUESTABILITY_FLOAT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl RequestAbility {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RequestAbilityPacket.Ability", REQUESTABILITY_ABILITY_SHAPE);
        encoder.field("RequestAbilityPacket.Value Type", REQUESTABILITY_VALUE_TYPE_SHAPE);
        encoder.field("RequestAbilityPacket.Bool", REQUESTABILITY_BOOL_SHAPE);
        encoder.field("RequestAbilityPacket.Float", REQUESTABILITY_FLOAT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RequestAbilityPacket.Ability", REQUESTABILITY_ABILITY_SHAPE);
        decoder.field("RequestAbilityPacket.Value Type", REQUESTABILITY_VALUE_TYPE_SHAPE);
        decoder.field("RequestAbilityPacket.Bool", REQUESTABILITY_BOOL_SHAPE);
        decoder.field("RequestAbilityPacket.Float", REQUESTABILITY_FLOAT_SHAPE);
    }
}
