// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MobEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub event_id: MobEffectEvent,
    pub effect_id: i32,
    pub effect_amplifier: i32,
    pub show_particles: bool,
    pub effect_duration_ticks: i32,
    pub tick: PlayerInputTick,
    pub ambient: bool,
}

pub const MOBEFFECT_TARGET_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOBEFFECT_EVENT_ID_SHAPE: &str = r#"{"kind":"enum","semantic":"MobEffectPacketPayload::Event","type_id":"enums/MobEffectPacketPayload::Event","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Invalid","encode":{"kind":"void"}},{"value":1,"name":"Add","encode":{"kind":"void"}},{"value":2,"name":"Update","encode":{"kind":"void"}},{"value":3,"name":"Remove","encode":{"kind":"void"}}]}"#;
pub const MOBEFFECT_EFFECT_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const MOBEFFECT_EFFECT_AMPLIFIER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const MOBEFFECT_SHOW_PARTICLES_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;
pub const MOBEFFECT_EFFECT_DURATION_TICKS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const MOBEFFECT_TICK_SHAPE: &str = r#"{"kind":"struct","semantic":"PlayerInputTick","type_id":"PlayerInputTick","fields":[{"ordinal":0,"name":"Input tick","semantic":"Input tick","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOBEFFECT_AMBIENT_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl MobEffect {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MobEffectPacket.Target Runtime ID", MOBEFFECT_TARGET_RUNTIME_ID_SHAPE);
        encoder.field("MobEffectPacket.Event ID", MOBEFFECT_EVENT_ID_SHAPE);
        encoder.field("MobEffectPacket.Effect ID", MOBEFFECT_EFFECT_ID_SHAPE);
        encoder.field("MobEffectPacket.Effect Amplifier", MOBEFFECT_EFFECT_AMPLIFIER_SHAPE);
        encoder.field("MobEffectPacket.Show Particles", MOBEFFECT_SHOW_PARTICLES_SHAPE);
        encoder.field("MobEffectPacket.Effect Duration Ticks", MOBEFFECT_EFFECT_DURATION_TICKS_SHAPE);
        encoder.field("MobEffectPacket.Tick", MOBEFFECT_TICK_SHAPE);
        encoder.field("MobEffectPacket.Ambient", MOBEFFECT_AMBIENT_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MobEffectPacket.Target Runtime ID", MOBEFFECT_TARGET_RUNTIME_ID_SHAPE);
        decoder.field("MobEffectPacket.Event ID", MOBEFFECT_EVENT_ID_SHAPE);
        decoder.field("MobEffectPacket.Effect ID", MOBEFFECT_EFFECT_ID_SHAPE);
        decoder.field("MobEffectPacket.Effect Amplifier", MOBEFFECT_EFFECT_AMPLIFIER_SHAPE);
        decoder.field("MobEffectPacket.Show Particles", MOBEFFECT_SHOW_PARTICLES_SHAPE);
        decoder.field("MobEffectPacket.Effect Duration Ticks", MOBEFFECT_EFFECT_DURATION_TICKS_SHAPE);
        decoder.field("MobEffectPacket.Tick", MOBEFFECT_TICK_SHAPE);
        decoder.field("MobEffectPacket.Ambient", MOBEFFECT_AMBIENT_SHAPE);
    }
}
