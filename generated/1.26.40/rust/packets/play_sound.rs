// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct PlaySound {
    pub name: String,
    pub position: BlockPos,
    pub volume: f32,
    pub pitch: f32,
    pub loop_count: i32,
    pub server_sound_handle: Option<ServerSoundHandle>,
}

pub const PLAYSOUND_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const PLAYSOUND_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"BlockPos","type_id":"BlockPos","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const PLAYSOUND_VOLUME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const PLAYSOUND_PITCH_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const PLAYSOUND_LOOP_COUNT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const PLAYSOUND_SERVER_SOUND_HANDLE_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"ServerSoundHandle","type_id":"ServerSoundHandle","fields":[{"ordinal":0,"name":"Server Sound Handle","semantic":"Server Sound Handle","encode":{"kind":"primitive","primitive":{"code":"u64le","width":64,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl PlaySound {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("PlaySoundPacket.Name", PLAYSOUND_NAME_SHAPE);
        encoder.field("PlaySoundPacket.Position", PLAYSOUND_POSITION_SHAPE);
        encoder.field("PlaySoundPacket.Volume", PLAYSOUND_VOLUME_SHAPE);
        encoder.field("PlaySoundPacket.Pitch", PLAYSOUND_PITCH_SHAPE);
        encoder.field("PlaySoundPacket.Loop Count", PLAYSOUND_LOOP_COUNT_SHAPE);
        encoder.field("PlaySoundPacket.Server Sound Handle", PLAYSOUND_SERVER_SOUND_HANDLE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("PlaySoundPacket.Name", PLAYSOUND_NAME_SHAPE);
        decoder.field("PlaySoundPacket.Position", PLAYSOUND_POSITION_SHAPE);
        decoder.field("PlaySoundPacket.Volume", PLAYSOUND_VOLUME_SHAPE);
        decoder.field("PlaySoundPacket.Pitch", PLAYSOUND_PITCH_SHAPE);
        decoder.field("PlaySoundPacket.Loop Count", PLAYSOUND_LOOP_COUNT_SHAPE);
        decoder.field("PlaySoundPacket.Server Sound Handle", PLAYSOUND_SERVER_SOUND_HANDLE_SHAPE);
    }
}
