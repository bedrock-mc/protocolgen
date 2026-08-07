// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MovementEffect {
    pub target_runtime_id: ActorRuntimeID,
    pub effect_id: MovementEffectType,
    pub effect_duration: i32,
    pub tick: PlayerInputTick,
}

pub const MOVEMENTEFFECT_TARGET_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOVEMENTEFFECT_EFFECT_ID_SHAPE: &str = r#"{"kind":"enum","semantic":"MovementEffectType","type_id":"enums/MovementEffectType","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"},"variants":[{"value":0,"name":"GLIDE_BOOST","encode":{"kind":"void"}},{"value":1,"name":"DOLPHIN_BOOST","encode":{"kind":"void"}},{"value":2,"name":"GEYSER_BOOST","encode":{"kind":"void"}}]}"#;
pub const MOVEMENTEFFECT_EFFECT_DURATION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const MOVEMENTEFFECT_TICK_SHAPE: &str = r#"{"kind":"struct","semantic":"PlayerInputTick","type_id":"PlayerInputTick","fields":[{"ordinal":0,"name":"Input tick","semantic":"Input tick","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl MovementEffect {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MovementEffectPacket.Target Runtime ID", MOVEMENTEFFECT_TARGET_RUNTIME_ID_SHAPE);
        encoder.field("MovementEffectPacket.Effect ID", MOVEMENTEFFECT_EFFECT_ID_SHAPE);
        encoder.field("MovementEffectPacket.Effect Duration", MOVEMENTEFFECT_EFFECT_DURATION_SHAPE);
        encoder.field("MovementEffectPacket.Tick", MOVEMENTEFFECT_TICK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MovementEffectPacket.Target Runtime ID", MOVEMENTEFFECT_TARGET_RUNTIME_ID_SHAPE);
        decoder.field("MovementEffectPacket.Effect ID", MOVEMENTEFFECT_EFFECT_ID_SHAPE);
        decoder.field("MovementEffectPacket.Effect Duration", MOVEMENTEFFECT_EFFECT_DURATION_SHAPE);
        decoder.field("MovementEffectPacket.Tick", MOVEMENTEFFECT_TICK_SHAPE);
    }
}
