// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssist {
    pub preset_id: String,
    pub view_angle: glam::Vec2,
    pub distance: f32,
    pub target_mode: CameraAimAssistTargetModeType,
    pub action: CameraAimAssistAction,
    pub show_debug_render: bool,
}

pub const CAMERAAIMASSIST_PRESET_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CAMERAAIMASSIST_VIEW_ANGLE_SHAPE: &str = r#"{"kind":"struct","semantic":"Vec2","type_id":"Vec2","fields":[{"ordinal":0,"name":"X","semantic":"X","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}},{"ordinal":1,"name":"Y","semantic":"Y","encode":{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CAMERAAIMASSIST_DISTANCE_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const CAMERAAIMASSIST_TARGET_MODE_SHAPE: &str = r#"{"kind":"enum","semantic":"CameraAimAssistPacketPayload::TargetMode","type_id":"enums/CameraAimAssistPacketPayload::TargetMode","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Angle","encode":{"kind":"void"}},{"value":1,"name":"Distance","encode":{"kind":"void"}}]}"#;
pub const CAMERAAIMASSIST_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"CameraAimAssistPacketPayload::Action","type_id":"enums/CameraAimAssistPacketPayload::Action","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Set","encode":{"kind":"void"}},{"value":1,"name":"Clear","encode":{"kind":"void"}}]}"#;
pub const CAMERAAIMASSIST_SHOW_DEBUG_RENDER_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl CameraAimAssist {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CameraAimAssistPacket.Preset Id", CAMERAAIMASSIST_PRESET_ID_SHAPE);
        encoder.field("CameraAimAssistPacket.View Angle", CAMERAAIMASSIST_VIEW_ANGLE_SHAPE);
        encoder.field("CameraAimAssistPacket.Distance", CAMERAAIMASSIST_DISTANCE_SHAPE);
        encoder.field("CameraAimAssistPacket.Target Mode", CAMERAAIMASSIST_TARGET_MODE_SHAPE);
        encoder.field("CameraAimAssistPacket.Action", CAMERAAIMASSIST_ACTION_SHAPE);
        encoder.field("CameraAimAssistPacket.Show Debug Render", CAMERAAIMASSIST_SHOW_DEBUG_RENDER_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CameraAimAssistPacket.Preset Id", CAMERAAIMASSIST_PRESET_ID_SHAPE);
        decoder.field("CameraAimAssistPacket.View Angle", CAMERAAIMASSIST_VIEW_ANGLE_SHAPE);
        decoder.field("CameraAimAssistPacket.Distance", CAMERAAIMASSIST_DISTANCE_SHAPE);
        decoder.field("CameraAimAssistPacket.Target Mode", CAMERAAIMASSIST_TARGET_MODE_SHAPE);
        decoder.field("CameraAimAssistPacket.Action", CAMERAAIMASSIST_ACTION_SHAPE);
        decoder.field("CameraAimAssistPacket.Show Debug Render", CAMERAAIMASSIST_SHOW_DEBUG_RENDER_SHAPE);
    }
}
