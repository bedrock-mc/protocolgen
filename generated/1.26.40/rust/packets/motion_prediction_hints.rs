// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct MotionPredictionHints {
    pub m_runtime_id: ActorRuntimeID,
    pub m_motion: Vec3,
    pub m_on_ground: bool,
}

pub const MOTIONPREDICTIONHINTS_M_RUNTIME_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOTIONPREDICTIONHINTS_M_MOTION_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec3","type_id":"Vec3","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":2,"name":"Z","semantic":"Z","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const MOTIONPREDICTIONHINTS_M_ON_GROUND_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl MotionPredictionHints {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("MotionPredictionHintsPacket.mRuntimeId", MOTIONPREDICTIONHINTS_M_RUNTIME_ID_SHAPE);
        encoder.field("MotionPredictionHintsPacket.mMotion", MOTIONPREDICTIONHINTS_M_MOTION_SHAPE);
        encoder.field("MotionPredictionHintsPacket.mOnGround", MOTIONPREDICTIONHINTS_M_ON_GROUND_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("MotionPredictionHintsPacket.mRuntimeId", MOTIONPREDICTIONHINTS_M_RUNTIME_ID_SHAPE);
        decoder.field("MotionPredictionHintsPacket.mMotion", MOTIONPREDICTIONHINTS_M_MOTION_SHAPE);
        decoder.field("MotionPredictionHintsPacket.mOnGround", MOTIONPREDICTIONHINTS_M_ON_GROUND_SHAPE);
    }
}
