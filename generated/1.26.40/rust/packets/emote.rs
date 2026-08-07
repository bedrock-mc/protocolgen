// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Emote {
    pub actor_runtime_id: ActorRuntimeID,
    pub emote_id: String,
    pub emote_length_ticks: u32,
    pub xuid: String,
    pub platform_id: String,
    pub flags: u8,
}

pub const EMOTE_ACTOR_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const EMOTE_EMOTE_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const EMOTE_EMOTE_LENGTH_TICKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const EMOTE_XUID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const EMOTE_PLATFORM_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const EMOTE_FLAGS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl Emote {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("EmotePacket.Actor Runtime Id", EMOTE_ACTOR_RUNTIME_ID_SHAPE);
        encoder.field("EmotePacket.Emote Id", EMOTE_EMOTE_ID_SHAPE);
        encoder.field("EmotePacket.Emote Length Ticks", EMOTE_EMOTE_LENGTH_TICKS_SHAPE);
        encoder.field("EmotePacket.Xuid", EMOTE_XUID_SHAPE);
        encoder.field("EmotePacket.PlatformId", EMOTE_PLATFORM_ID_SHAPE);
        encoder.field("EmotePacket.Flags", EMOTE_FLAGS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("EmotePacket.Actor Runtime Id", EMOTE_ACTOR_RUNTIME_ID_SHAPE);
        decoder.field("EmotePacket.Emote Id", EMOTE_EMOTE_ID_SHAPE);
        decoder.field("EmotePacket.Emote Length Ticks", EMOTE_EMOTE_LENGTH_TICKS_SHAPE);
        decoder.field("EmotePacket.Xuid", EMOTE_XUID_SHAPE);
        decoder.field("EmotePacket.PlatformId", EMOTE_PLATFORM_ID_SHAPE);
        decoder.field("EmotePacket.Flags", EMOTE_FLAGS_SHAPE);
    }
}
