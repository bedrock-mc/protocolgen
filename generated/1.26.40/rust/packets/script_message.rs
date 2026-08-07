// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ScriptMessage {
    pub message_id: String,
    pub message_value: String,
}

pub const SCRIPTMESSAGE_MESSAGE_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SCRIPTMESSAGE_MESSAGE_VALUE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ScriptMessage {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ScriptMessagePacket.Message Id", SCRIPTMESSAGE_MESSAGE_ID_SHAPE);
        encoder.field("ScriptMessagePacket.Message Value", SCRIPTMESSAGE_MESSAGE_VALUE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ScriptMessagePacket.Message Id", SCRIPTMESSAGE_MESSAGE_ID_SHAPE);
        decoder.field("ScriptMessagePacket.Message Value", SCRIPTMESSAGE_MESSAGE_VALUE_SHAPE);
    }
}
