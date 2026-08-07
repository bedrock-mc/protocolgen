// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ShowCredits {
    pub player_runtime_id: ActorRuntimeID,
    pub credits_state: i32,
}

pub const SHOWCREDITS_PLAYER_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SHOWCREDITS_CREDITS_STATE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl ShowCredits {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ShowCreditsPacket.Player Runtime ID", SHOWCREDITS_PLAYER_RUNTIME_ID_SHAPE);
        encoder.field("ShowCreditsPacket.Credits State", SHOWCREDITS_CREDITS_STATE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ShowCreditsPacket.Player Runtime ID", SHOWCREDITS_PLAYER_RUNTIME_ID_SHAPE);
        decoder.field("ShowCreditsPacket.Credits State", SHOWCREDITS_CREDITS_STATE_SHAPE);
    }
}
