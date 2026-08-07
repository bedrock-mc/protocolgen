// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct LevelSoundEvent {
    pub sound_event: String,
    pub position: Vec3,
    pub data: i32,
    pub actor_identifier: String,
    pub is_baby: bool,
    pub is_global: bool,
    pub actor_unique_id: i64,
    pub fire_at_position: Option<Vec3>,
}

pub const LEVELSOUNDEVENT_SOUND_EVENT_SHAPE: &str = r##"{"kind":"string","semantic":"SoundEventIdentifier","type_id":"SoundEventIdentifier.json#","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"##;
pub const LEVELSOUNDEVENT_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const LEVELSOUNDEVENT_DATA_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const LEVELSOUNDEVENT_ACTOR_IDENTIFIER_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const LEVELSOUNDEVENT_IS_BABY_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const LEVELSOUNDEVENT_IS_GLOBAL_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const LEVELSOUNDEVENT_ACTOR_UNIQUE_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i64le","width":64,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const LEVELSOUNDEVENT_FIRE_AT_POSITION_SHAPE: &str = r#"{"kind":"optional","value":{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl LevelSoundEvent {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("LevelSoundEventPacket.Sound Event", LEVELSOUNDEVENT_SOUND_EVENT_SHAPE);
        encoder.field("LevelSoundEventPacket.Position", LEVELSOUNDEVENT_POSITION_SHAPE);
        encoder.field("LevelSoundEventPacket.Data", LEVELSOUNDEVENT_DATA_SHAPE);
        encoder.field("LevelSoundEventPacket.Actor Identifier", LEVELSOUNDEVENT_ACTOR_IDENTIFIER_SHAPE);
        encoder.field("LevelSoundEventPacket.Is Baby", LEVELSOUNDEVENT_IS_BABY_SHAPE);
        encoder.field("LevelSoundEventPacket.Is Global", LEVELSOUNDEVENT_IS_GLOBAL_SHAPE);
        encoder.field("LevelSoundEventPacket.Actor Unique Id", LEVELSOUNDEVENT_ACTOR_UNIQUE_ID_SHAPE);
        encoder.field("LevelSoundEventPacket.Fire At Position", LEVELSOUNDEVENT_FIRE_AT_POSITION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("LevelSoundEventPacket.Sound Event", LEVELSOUNDEVENT_SOUND_EVENT_SHAPE);
        decoder.field("LevelSoundEventPacket.Position", LEVELSOUNDEVENT_POSITION_SHAPE);
        decoder.field("LevelSoundEventPacket.Data", LEVELSOUNDEVENT_DATA_SHAPE);
        decoder.field("LevelSoundEventPacket.Actor Identifier", LEVELSOUNDEVENT_ACTOR_IDENTIFIER_SHAPE);
        decoder.field("LevelSoundEventPacket.Is Baby", LEVELSOUNDEVENT_IS_BABY_SHAPE);
        decoder.field("LevelSoundEventPacket.Is Global", LEVELSOUNDEVENT_IS_GLOBAL_SHAPE);
        decoder.field("LevelSoundEventPacket.Actor Unique Id", LEVELSOUNDEVENT_ACTOR_UNIQUE_ID_SHAPE);
        decoder.field("LevelSoundEventPacket.Fire At Position", LEVELSOUNDEVENT_FIRE_AT_POSITION_SHAPE);
    }
}
