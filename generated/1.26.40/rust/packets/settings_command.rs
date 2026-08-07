// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SettingsCommand {
    pub command: String,
    pub suppress_output: bool,
}

pub const SETTINGSCOMMAND_COMMAND_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SETTINGSCOMMAND_SUPPRESS_OUTPUT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl SettingsCommand {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SettingsCommandPacket.Command", SETTINGSCOMMAND_COMMAND_SHAPE);
        encoder.field("SettingsCommandPacket.Suppress Output?", SETTINGSCOMMAND_SUPPRESS_OUTPUT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SettingsCommandPacket.Command", SETTINGSCOMMAND_COMMAND_SHAPE);
        decoder.field("SettingsCommandPacket.Suppress Output?", SETTINGSCOMMAND_SUPPRESS_OUTPUT_SHAPE);
    }
}
