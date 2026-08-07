// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SpawnExperienceOrb {
    pub position: Vec3,
    pub xp_value: i32,
}

pub const SPAWNEXPERIENCEORB_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SPAWNEXPERIENCEORB_XP_VALUE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;

impl SpawnExperienceOrb {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SpawnExperienceOrbPacket.Position", SPAWNEXPERIENCEORB_POSITION_SHAPE);
        encoder.field("SpawnExperienceOrbPacket.XP Value", SPAWNEXPERIENCEORB_XP_VALUE_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SpawnExperienceOrbPacket.Position", SPAWNEXPERIENCEORB_POSITION_SHAPE);
        decoder.field("SpawnExperienceOrbPacket.XP Value", SPAWNEXPERIENCEORB_XP_VALUE_SHAPE);
    }
}
