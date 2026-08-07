// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct SetActorMotion {
    pub target_runtime_id: ActorRuntimeID,
    pub motion: glam::Vec3,
    pub tick: PlayerInputTick,
}

pub const SETACTORMOTION_TARGET_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SETACTORMOTION_MOTION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const SETACTORMOTION_TICK_SHAPE: &str = r#"{"kind":"struct","semantic":"PlayerInputTick","type_id":"PlayerInputTick","fields":[{"ordinal":0,"name":"Input tick","semantic":"Input tick","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl SetActorMotion {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("SetActorMotionPacket.Target Runtime ID", SETACTORMOTION_TARGET_RUNTIME_ID_SHAPE);
        encoder.field("SetActorMotionPacket.Motion", SETACTORMOTION_MOTION_SHAPE);
        encoder.field("SetActorMotionPacket.Tick", SETACTORMOTION_TICK_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("SetActorMotionPacket.Target Runtime ID", SETACTORMOTION_TARGET_RUNTIME_ID_SHAPE);
        decoder.field("SetActorMotionPacket.Motion", SETACTORMOTION_MOTION_SHAPE);
        decoder.field("SetActorMotionPacket.Tick", SETACTORMOTION_TICK_SHAPE);
    }
}
