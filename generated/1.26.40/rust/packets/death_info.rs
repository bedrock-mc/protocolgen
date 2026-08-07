// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct DeathInfo {
    pub death_cause_attack_name: String,
    pub death_cause_message_list: Vec<String>,
}

pub const DEATHINFO_DEATH_CAUSE_ATTACK_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const DEATHINFO_DEATH_CAUSE_MESSAGE_LIST_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"#;

impl DeathInfo {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("DeathInfoPacket.Death Cause Attack Name", DEATHINFO_DEATH_CAUSE_ATTACK_NAME_SHAPE);
        encoder.field("DeathInfoPacket.Death Cause Message List", DEATHINFO_DEATH_CAUSE_MESSAGE_LIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("DeathInfoPacket.Death Cause Attack Name", DEATHINFO_DEATH_CAUSE_ATTACK_NAME_SHAPE);
        decoder.field("DeathInfoPacket.Death Cause Message List", DEATHINFO_DEATH_CAUSE_MESSAGE_LIST_SHAPE);
    }
}
