// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CameraShake {
    pub intensity: f32,
    pub seconds: f32,
    pub shake_type: CameraShakeType,
    pub shake_action: CameraShakeAction,
}

pub const CAMERASHAKE_INTENSITY_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const CAMERASHAKE_SECONDS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"f32le","width":32,"signed":false,"zigzag":false,"endianness":"little"}}"#;
pub const CAMERASHAKE_SHAKE_TYPE_SHAPE: &str = r#"{"kind":"enum","semantic":"CameraShakeType","type_id":"enums/CameraShakeType","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Positional","encode":{"kind":"void"}},{"value":1,"name":"Rotational","encode":{"kind":"void"}}]}"#;
pub const CAMERASHAKE_SHAKE_ACTION_SHAPE: &str = r#"{"kind":"enum","semantic":"CameraShakeAction","type_id":"enums/CameraShakeAction","primitive":{"code":"u8","width":8,"signed":false,"zigzag":false,"endianness":"none"},"variants":[{"value":0,"name":"Add","encode":{"kind":"void"}},{"value":1,"name":"Stop","encode":{"kind":"void"}}]}"#;

impl CameraShake {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CameraShakePacket.Intensity", CAMERASHAKE_INTENSITY_SHAPE);
        encoder.field("CameraShakePacket.Seconds", CAMERASHAKE_SECONDS_SHAPE);
        encoder.field("CameraShakePacket.Shake Type", CAMERASHAKE_SHAKE_TYPE_SHAPE);
        encoder.field("CameraShakePacket.Shake Action", CAMERASHAKE_SHAKE_ACTION_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CameraShakePacket.Intensity", CAMERASHAKE_INTENSITY_SHAPE);
        decoder.field("CameraShakePacket.Seconds", CAMERASHAKE_SECONDS_SHAPE);
        decoder.field("CameraShakePacket.Shake Type", CAMERASHAKE_SHAKE_TYPE_SHAPE);
        decoder.field("CameraShakePacket.Shake Action", CAMERASHAKE_SHAKE_ACTION_SHAPE);
    }
}
