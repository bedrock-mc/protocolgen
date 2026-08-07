// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct RequestPermissions {
    pub target_player_id_s_raw_id: i64,
    pub player_permission_level: i32,
    pub custom_permission_flags: u16,
}

impl RequestPermissions {
    pub const ID: u32 = 185;
}
