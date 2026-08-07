// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct DebugInfo {
    pub actor_id: ActorUniqueID,
    pub data: String,
}

pub const DEBUGINFO_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const DEBUGINFO_DATA_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;

impl DebugInfo {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("DebugInfoPacket.Actor Id", DEBUGINFO_ACTOR_ID_SHAPE);
        encoder.field("DebugInfoPacket.Data", DEBUGINFO_DATA_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("DebugInfoPacket.Actor Id", DEBUGINFO_ACTOR_ID_SHAPE);
        decoder.field("DebugInfoPacket.Data", DEBUGINFO_DATA_SHAPE);
    }
}
