// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ChangeMobProperty {
    pub actor_id: ActorUniqueID,
    pub property_name: String,
    pub bool_component_value: bool,
    pub string_component_value: String,
    pub int_component_value: i32,
    pub float_component_value: f32,
}

pub const CHANGEMOBPROPERTY_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CHANGEMOBPROPERTY_PROPERTY_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CHANGEMOBPROPERTY_BOOL_COMPONENT_VALUE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CHANGEMOBPROPERTY_STRING_COMPONENT_VALUE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CHANGEMOBPROPERTY_INT_COMPONENT_VALUE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const CHANGEMOBPROPERTY_FLOAT_COMPONENT_VALUE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl ChangeMobProperty {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ChangeMobPropertyPacket.Actor Id", CHANGEMOBPROPERTY_ACTOR_ID_SHAPE);
        encoder.field("ChangeMobPropertyPacket.Property Name", CHANGEMOBPROPERTY_PROPERTY_NAME_SHAPE);
        encoder.field("ChangeMobPropertyPacket.BoolComponent Value", CHANGEMOBPROPERTY_BOOL_COMPONENT_VALUE_SHAPE);
        encoder.field("ChangeMobPropertyPacket.StringComponent Value", CHANGEMOBPROPERTY_STRING_COMPONENT_VALUE_SHAPE);
        encoder.field("ChangeMobPropertyPacket.IntComponent Value", CHANGEMOBPROPERTY_INT_COMPONENT_VALUE_SHAPE);
        encoder.field("ChangeMobPropertyPacket.FloatComponent Value", CHANGEMOBPROPERTY_FLOAT_COMPONENT_VALUE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ChangeMobPropertyPacket.Actor Id", CHANGEMOBPROPERTY_ACTOR_ID_SHAPE);
        decoder.field("ChangeMobPropertyPacket.Property Name", CHANGEMOBPROPERTY_PROPERTY_NAME_SHAPE);
        decoder.field("ChangeMobPropertyPacket.BoolComponent Value", CHANGEMOBPROPERTY_BOOL_COMPONENT_VALUE_SHAPE);
        decoder.field("ChangeMobPropertyPacket.StringComponent Value", CHANGEMOBPROPERTY_STRING_COMPONENT_VALUE_SHAPE);
        decoder.field("ChangeMobPropertyPacket.IntComponent Value", CHANGEMOBPROPERTY_INT_COMPONENT_VALUE_SHAPE);
        decoder.field("ChangeMobPropertyPacket.FloatComponent Value", CHANGEMOBPROPERTY_FLOAT_COMPONENT_VALUE_SHAPE);
    }
}
