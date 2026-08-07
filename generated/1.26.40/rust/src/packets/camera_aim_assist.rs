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

impl CameraAimAssist {
    pub const ID: u32 = 316;
}
