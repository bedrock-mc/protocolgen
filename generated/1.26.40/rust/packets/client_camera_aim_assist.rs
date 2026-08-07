// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCameraAimAssist {
    pub camera_preset_id: String,
    pub action: ClientCameraAimAssistPacketAction,
    pub allow_aim_assist: bool,
}

pub const CLIENTCAMERAAIMASSIST_CAMERA_PRESET_ID_SHAPE: &str = r#"{"kind":"string","prefix":{"kind":"primitive","primitive":{"code":"var_u32","width":32,"signed":false,"zigzag":false,"endianness":"none"}},"encoding":"utf8","representation":"text"}"#;
pub const CLIENTCAMERAAIMASSIST_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"ClientCameraAimAssistPacketAction","type_id":"enums/ClientCameraAimAssistPacketAction","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"SetFromCameraPreset","encode":{"kind":"void"}},{"value":1,"name":"Clear","encode":{"kind":"void"}}]}"#;
pub const CLIENTCAMERAAIMASSIST_ALLOW_AIM_ASSIST_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"bool","width":1,"signed":false,"zigzag":false,"endianness":"none"}}"#;

impl ClientCameraAimAssist {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("ClientCameraAimAssistPacket.Camera Preset Id", CLIENTCAMERAAIMASSIST_CAMERA_PRESET_ID_SHAPE);
        encoder.field("ClientCameraAimAssistPacket.Action", CLIENTCAMERAAIMASSIST_ACTION_SHAPE);
        encoder.field("ClientCameraAimAssistPacket.Allow aim assist", CLIENTCAMERAAIMASSIST_ALLOW_AIM_ASSIST_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("ClientCameraAimAssistPacket.Camera Preset Id", CLIENTCAMERAAIMASSIST_CAMERA_PRESET_ID_SHAPE);
        decoder.field("ClientCameraAimAssistPacket.Action", CLIENTCAMERAAIMASSIST_ACTION_SHAPE);
        decoder.field("ClientCameraAimAssistPacket.Allow aim assist", CLIENTCAMERAAIMASSIST_ALLOW_AIM_ASSIST_SHAPE);
    }
}
