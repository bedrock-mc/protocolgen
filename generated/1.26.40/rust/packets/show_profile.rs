// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ShowProfile {
    pub player_xuid: String,
}

pub const SHOWPROFILE_PLAYER_XUID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl ShowProfile {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ShowProfilePacket.Player XUID", SHOWPROFILE_PLAYER_XUID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ShowProfilePacket.Player XUID", SHOWPROFILE_PLAYER_XUID_SHAPE);
    }
}
