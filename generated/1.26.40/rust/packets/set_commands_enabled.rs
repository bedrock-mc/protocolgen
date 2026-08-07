// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetCommandsEnabled {
    pub commands_enabled: bool,
}

pub const SETCOMMANDSENABLED_COMMANDS_ENABLED_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl SetCommandsEnabled {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetCommandsEnabledPacket.Commands Enabled", SETCOMMANDSENABLED_COMMANDS_ENABLED_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetCommandsEnabledPacket.Commands Enabled", SETCOMMANDSENABLED_COMMANDS_ENABLED_SHAPE);
    }
}
