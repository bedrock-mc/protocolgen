// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SpawnParticleEffect {
    pub dimension_id: u8,
    pub actor_id: ActorUniqueID,
    pub position: glam::Vec3,
    pub effect_name: String,
    pub molang_variables: Option<String>,
}

pub const SPAWNPARTICLEEFFECT_DIMENSION_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const SPAWNPARTICLEEFFECT_ACTOR_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SPAWNPARTICLEEFFECT_POSITION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SPAWNPARTICLEEFFECT_EFFECT_NAME_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const SPAWNPARTICLEEFFECT_MOLANG_VARIABLES_SHAPE: &str = r##"{"kind":"optional","value":{"kind":"string","semantic":"Json::Value","type_id":"MolangVariableMap.json#","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}}"##;

impl SpawnParticleEffect {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SpawnParticleEffectPacket.Dimension Id", SPAWNPARTICLEEFFECT_DIMENSION_ID_SHAPE);
        encoder.field("SpawnParticleEffectPacket.Actor Id", SPAWNPARTICLEEFFECT_ACTOR_ID_SHAPE);
        encoder.field("SpawnParticleEffectPacket.Position", SPAWNPARTICLEEFFECT_POSITION_SHAPE);
        encoder.field("SpawnParticleEffectPacket.Effect Name", SPAWNPARTICLEEFFECT_EFFECT_NAME_SHAPE);
        encoder.field("SpawnParticleEffectPacket.Molang Variables", SPAWNPARTICLEEFFECT_MOLANG_VARIABLES_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SpawnParticleEffectPacket.Dimension Id", SPAWNPARTICLEEFFECT_DIMENSION_ID_SHAPE);
        decoder.field("SpawnParticleEffectPacket.Actor Id", SPAWNPARTICLEEFFECT_ACTOR_ID_SHAPE);
        decoder.field("SpawnParticleEffectPacket.Position", SPAWNPARTICLEEFFECT_POSITION_SHAPE);
        decoder.field("SpawnParticleEffectPacket.Effect Name", SPAWNPARTICLEEFFECT_EFFECT_NAME_SHAPE);
        decoder.field("SpawnParticleEffectPacket.Molang Variables", SPAWNPARTICLEEFFECT_MOLANG_VARIABLES_SHAPE);
    }
}
