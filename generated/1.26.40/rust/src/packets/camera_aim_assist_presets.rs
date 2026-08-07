// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[allow(unused_imports)]
use crate::*;

#[derive(Clone, Debug, PartialEq)]
pub struct CameraAimAssistPresets {
    pub camera_aim_assist_presets: Vec<CameraAimAssistCategoryDefinition>,
    pub camera_aim_assist_categories: Vec<CameraAimAssistPresetDefinition>,
    pub operation: CameraAimAssistPresetsPacketOperation,
}

impl CameraAimAssistPresets {
    pub const ID: u32 = 320;
}
