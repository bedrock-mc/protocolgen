// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestPermissions {
    pub target_player_id_s_raw_id: i64,
    pub player_permission_level: i32,
    pub custom_permission_flags: u16,
}

pub const REQUESTPERMISSIONS_TARGET_PLAYER_ID_S_RAW_ID_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"i64le","width":64,"signed":true,"zigzag":false,"endianness":"little"}}"#;
pub const REQUESTPERMISSIONS_PLAYER_PERMISSION_LEVEL_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"zigzag_i32","width":32,"signed":true,"zigzag":true,"endianness":"none"}}"#;
pub const REQUESTPERMISSIONS_CUSTOM_PERMISSION_FLAGS_SHAPE: &str = r#"{"kind":"primitive","primitive":{"code":"u16le","width":16,"signed":false,"zigzag":false,"endianness":"little"}}"#;

impl RequestPermissions {
    pub fn encode<E: WireEncoder>(&self, encoder: &mut E) {
        encoder.field("RequestPermissionsPacket.Target Player Id's Raw ID", REQUESTPERMISSIONS_TARGET_PLAYER_ID_S_RAW_ID_SHAPE);
        encoder.field("RequestPermissionsPacket.Player Permission Level", REQUESTPERMISSIONS_PLAYER_PERMISSION_LEVEL_SHAPE);
        encoder.field("RequestPermissionsPacket.Custom Permission Flags", REQUESTPERMISSIONS_CUSTOM_PERMISSION_FLAGS_SHAPE);
    }
    pub fn decode<D: WireDecoder>(decoder: &mut D) {
        decoder.field("RequestPermissionsPacket.Target Player Id's Raw ID", REQUESTPERMISSIONS_TARGET_PLAYER_ID_S_RAW_ID_SHAPE);
        decoder.field("RequestPermissionsPacket.Player Permission Level", REQUESTPERMISSIONS_PLAYER_PERMISSION_LEVEL_SHAPE);
        decoder.field("RequestPermissionsPacket.Custom Permission Flags", REQUESTPERMISSIONS_CUSTOM_PERMISSION_FLAGS_SHAPE);
    }
}
