// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct StopSound {
    pub sound_name: String,
    pub stop_all_sounds: bool,
    pub stop_music_legacy: bool,
}

pub const STOPSOUND_SOUND_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const STOPSOUND_STOP_ALL_SOUNDS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const STOPSOUND_STOP_MUSIC_LEGACY_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl StopSound {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("StopSoundPacket.Sound Name", STOPSOUND_SOUND_NAME_SHAPE);
        encoder.field("StopSoundPacket.Stop All Sounds?", STOPSOUND_STOP_ALL_SOUNDS_SHAPE);
        encoder.field("StopSoundPacket.Stop Music (Legacy)", STOPSOUND_STOP_MUSIC_LEGACY_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("StopSoundPacket.Sound Name", STOPSOUND_SOUND_NAME_SHAPE);
        decoder.field("StopSoundPacket.Stop All Sounds?", STOPSOUND_STOP_ALL_SOUNDS_SHAPE);
        decoder.field("StopSoundPacket.Stop Music (Legacy)", STOPSOUND_STOP_MUSIC_LEGACY_SHAPE);
    }
}
