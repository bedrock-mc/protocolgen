// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct ClientCameraAimAssist {
    pub camera_preset_id: String,
    pub action: ClientCameraAimAssistPacketAction,
    pub allow_aim_assist: bool,
}

impl ClientCameraAimAssist {
    pub const ID: u32 = 321;
}
