// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct AnimateEntity {
    pub m_animation: String,
    pub m_next_state: String,
    pub m_stop_expression: String,
    pub m_stop_expression_version: i32,
    pub m_controller: String,
    pub m_blend_out_time: f32,
    pub m_runtime_ids: Vec<ActorRuntimeID>,
}

pub const ANIMATEENTITY_M_ANIMATION_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const ANIMATEENTITY_M_NEXT_STATE_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const ANIMATEENTITY_M_STOP_EXPRESSION_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const ANIMATEENTITY_M_STOP_EXPRESSION_VERSION_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i32le","width":32,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const ANIMATEENTITY_M_CONTROLLER_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const ANIMATEENTITY_M_BLEND_OUT_TIME_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const ANIMATEENTITY_M_RUNTIME_IDS_SHAPE: &str = r#"{"kind":"array","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"element":{"kind":"struct","semantic":"ActorRuntimeID","type_id":"ActorRuntimeID","fields":[{"ordinal":0,"name":"Actor Runtime ID","semantic":"Actor Runtime ID","encode":{"kind":"primitive","primitive":{"code":"var_u64","width":64,"signed":false,"zigzag":false,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}}"#;

impl AnimateEntity {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("AnimateEntityPacket.mAnimation", ANIMATEENTITY_M_ANIMATION_SHAPE);
        encoder.field("AnimateEntityPacket.mNextState", ANIMATEENTITY_M_NEXT_STATE_SHAPE);
        encoder.field("AnimateEntityPacket.mStopExpression", ANIMATEENTITY_M_STOP_EXPRESSION_SHAPE);
        encoder.field("AnimateEntityPacket.mStopExpressionVersion", ANIMATEENTITY_M_STOP_EXPRESSION_VERSION_SHAPE);
        encoder.field("AnimateEntityPacket.mController", ANIMATEENTITY_M_CONTROLLER_SHAPE);
        encoder.field("AnimateEntityPacket.mBlendOutTime", ANIMATEENTITY_M_BLEND_OUT_TIME_SHAPE);
        encoder.field("AnimateEntityPacket.mRuntimeIds", ANIMATEENTITY_M_RUNTIME_IDS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("AnimateEntityPacket.mAnimation", ANIMATEENTITY_M_ANIMATION_SHAPE);
        decoder.field("AnimateEntityPacket.mNextState", ANIMATEENTITY_M_NEXT_STATE_SHAPE);
        decoder.field("AnimateEntityPacket.mStopExpression", ANIMATEENTITY_M_STOP_EXPRESSION_SHAPE);
        decoder.field("AnimateEntityPacket.mStopExpressionVersion", ANIMATEENTITY_M_STOP_EXPRESSION_VERSION_SHAPE);
        decoder.field("AnimateEntityPacket.mController", ANIMATEENTITY_M_CONTROLLER_SHAPE);
        decoder.field("AnimateEntityPacket.mBlendOutTime", ANIMATEENTITY_M_BLEND_OUT_TIME_SHAPE);
        decoder.field("AnimateEntityPacket.mRuntimeIds", ANIMATEENTITY_M_RUNTIME_IDS_SHAPE);
    }
}
