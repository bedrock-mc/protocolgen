// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientboundTextureShift {
    pub action_id: ClientboundTextureShiftAction,
    pub collection_name: String,
    pub from_step: String,
    pub to_step: String,
    pub all_steps: Vec<String>,
    pub current_length_in_ticks: u64,
    pub total_length_in_ticks: u64,
    pub enabled: bool,
}

pub const CLIENTBOUNDTEXTURESHIFT_ACTION_ID_SHAPE: &str = r#"{"kind":"enum","semantic":"ClientboundTextureShiftPacketPayload::Action","type_id":"enums/ClientboundTextureShiftPacketPayload::Action","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Invalid","encode":{"kind":"void"}},{"value":1,"name":"Initialize","encode":{"kind":"void"}},{"value":2,"name":"Start","encode":{"kind":"void"}},{"value":3,"name":"SetEnabled","encode":{"kind":"void"}},{"value":4,"name":"Sync","encode":{"kind":"void"}}]}"#;
pub const CLIENTBOUNDTEXTURESHIFT_COLLECTION_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CLIENTBOUNDTEXTURESHIFT_FROM_STEP_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CLIENTBOUNDTEXTURESHIFT_TO_STEP_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CLIENTBOUNDTEXTURESHIFT_ALL_STEPS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;
pub const CLIENTBOUNDTEXTURESHIFT_CURRENT_LENGTH_IN_TICKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CLIENTBOUNDTEXTURESHIFT_TOTAL_LENGTH_IN_TICKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const CLIENTBOUNDTEXTURESHIFT_ENABLED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ClientboundTextureShift {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientboundTextureShiftPacket.Action ID", CLIENTBOUNDTEXTURESHIFT_ACTION_ID_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.Collection Name", CLIENTBOUNDTEXTURESHIFT_COLLECTION_NAME_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.From Step", CLIENTBOUNDTEXTURESHIFT_FROM_STEP_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.To Step", CLIENTBOUNDTEXTURESHIFT_TO_STEP_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.All Steps", CLIENTBOUNDTEXTURESHIFT_ALL_STEPS_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.Current Length In Ticks", CLIENTBOUNDTEXTURESHIFT_CURRENT_LENGTH_IN_TICKS_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.Total Length In Ticks", CLIENTBOUNDTEXTURESHIFT_TOTAL_LENGTH_IN_TICKS_SHAPE);
        encoder.field("ClientboundTextureShiftPacket.Enabled", CLIENTBOUNDTEXTURESHIFT_ENABLED_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientboundTextureShiftPacket.Action ID", CLIENTBOUNDTEXTURESHIFT_ACTION_ID_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.Collection Name", CLIENTBOUNDTEXTURESHIFT_COLLECTION_NAME_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.From Step", CLIENTBOUNDTEXTURESHIFT_FROM_STEP_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.To Step", CLIENTBOUNDTEXTURESHIFT_TO_STEP_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.All Steps", CLIENTBOUNDTEXTURESHIFT_ALL_STEPS_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.Current Length In Ticks", CLIENTBOUNDTEXTURESHIFT_CURRENT_LENGTH_IN_TICKS_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.Total Length In Ticks", CLIENTBOUNDTEXTURESHIFT_TOTAL_LENGTH_IN_TICKS_SHAPE);
        decoder.field("ClientboundTextureShiftPacket.Enabled", CLIENTBOUNDTEXTURESHIFT_ENABLED_SHAPE);
    }
}
