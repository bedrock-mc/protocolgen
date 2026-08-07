// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct Camera {
    pub camera_id: ActorUniqueID,
    pub target_player_id: ActorUniqueID,
}

pub const CAMERA_CAMERA_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;
pub const CAMERA_TARGET_PLAYER_ID_SHAPE: &str = r#"{"kind":"struct","semantic":"ActorUniqueID","type_id":"ActorUniqueID","fields":[{"ordinal":0,"name":"Actor Unique ID","semantic":"Actor Unique ID","encode":{"kind":"primitive","primitive":{"code":"zigzag_i64","width":64,"signed":true,"zigzag":true,"endianness":"none"}},"symmetry":"symmetric","provenance":{"pins":["endstone"]}}]}"#;

impl Camera {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("CameraPacket.Camera ID", CAMERA_CAMERA_ID_SHAPE);
        encoder.field("CameraPacket.Target Player ID", CAMERA_TARGET_PLAYER_ID_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("CameraPacket.Camera ID", CAMERA_CAMERA_ID_SHAPE);
        decoder.field("CameraPacket.Target Player ID", CAMERA_TARGET_PLAYER_ID_SHAPE);
    }
}
